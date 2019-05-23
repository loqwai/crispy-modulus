using Unity.Entities;
using Unity.Mathematics;
using UnityEngine;
using Unity.Jobs;
using Unity.Collections;
using System;
using System.Collections.Generic;
using System.Linq;
using OurECS;
namespace OurECS {
  public class GameSystem : ComponentSystem {
    protected EntityArchetype undrawnCardArchetype;
    System.Random random;
    
    protected struct CardEntity {
      public Entity entity;
      public Card card;
    }

    protected struct PlayerEntity {
      public Entity entity;
      public Player player;
    }

    protected override void OnCreateManager(){
      RequireSingletonForUpdate<Game>();      
      random = new System.Random((int)DateTime.Now.Ticks);
      
      undrawnCardArchetype =
        EntityManager.CreateArchetype(
          typeof(Card),
          typeof(CardFacedDown)
      );  

    }

    private void CreatePlayers(Game game) {
      var query = GetEntityQuery(typeof(Player));
      EntityManager.DestroyEntity(query);
      
      for(int i = 0; i < game.playerCount; i++) {
        var e = PostUpdateCommands.CreateEntity();
        PostUpdateCommands.AddComponent<Player>(e, new Player());
      }      
    }

    private void DealCards(Game game) {
      var query = GetEntityQuery(typeof(Card));
      EntityManager.DestroyEntity(query);
      
      Entities.ForEach((Entity pe, ref Player p) => {
        for (int i = 1; i < game.cardCount+1; i++) {
          var e = PostUpdateCommands.CreateEntity(undrawnCardArchetype);
          PostUpdateCommands.SetComponent(e, new Card { Value = i, OriginalPlayer=pe});
        }
      });
    }

    private void DrawUntilUnequalMods(Game game) {
      var mods = DrawAll(game);
      if(mods.Distinct().Count() == 1)
      DrawUntilUnequalMods(game);      
    }

    private List<int> DrawAll(Game game) {
      var mods = new List<int>();
      Entities.ForEach((Entity e, ref Player p) => {         
          Draw(e, ref p, game);
          mods.Add(p.mod);
      });
      return mods;
    }

    protected void Draw(Entity pe, ref Player player, Game game) {                                  
        var ourCards = GetDrawableCardsForPlayer(pe);
        if(ourCards.Count == 0) return;

        var cardToDraw = random.Next(0, ourCards.Count);
        var drawnCard = ourCards[cardToDraw];
        player.cardCount++;
        player.cardSum += drawnCard.card.Value;
        player.mod = player.cardSum % game.mod;        
        PostUpdateCommands.RemoveComponent<CardFacedDown>(drawnCard.entity);
    } 

    protected List<CardEntity> GetDrawableCardsForPlayer(Entity pe) {
      var ourCards = new List<CardEntity>();

      Entities.WithAll<Card, CardFacedDown>().
        ForEach((Entity e, ref Card c)=>{
          if(c.OriginalPlayer != pe) return;
          ourCards.Add(new CardEntity{entity=e, card=c});
      });

      return ourCards;
    }

    protected void DoActivePlayerAction(Game game) {
        var pe = findActivePlayer();        
        switch(pe.player.action) {
          case Player.Actions.Nothing:
          break;
          case Player.Actions.Draw:
          Draw(pe.entity, ref pe.player, game);          
          break;
        }

        pe.player.action = Player.Actions.Nothing;
        // setNextPlayer()
    }

    protected PlayerEntity findActivePlayer() {
      var pe = new PlayerEntity();
      Entities.WithAll<Player, ActivePlayer>().
        ForEach((Entity e, ref Player p)=>{
          pe.entity = e;
          pe.player = p;
          return;
      });        

      return pe;      
    }

    protected void FindStartingPlayer(Game game) {
      var maxMod = -1;
      Entity worstPlayer = new Entity();
      Entities.ForEach((Entity e, ref Player p) => {
        var mod = p.mod;
        if (mod > maxMod) {
          maxMod = mod;
          worstPlayer = e;
        }
      });        
      PostUpdateCommands.AddComponent<ActivePlayer>(worstPlayer,new ActivePlayer());
    }

    protected override void OnUpdate() {
      var game = GetSingleton<Game>();
      switch(game.action) {
        
        case Game.Actions.Nothing:
          return;
        
        case Game.Actions.Start:
          CreatePlayers(game);
          game.action = Game.Actions.Deal;
          break;
        
        case Game.Actions.Deal:
          DealCards(game);
          game.action = Game.Actions.DrawUntilUnequalMods;
          break;
          
        case Game.Actions.DrawUntilUnequalMods:
          DrawUntilUnequalMods(game);
          game.action = Game.Actions.FindStartingPlayer;
          break;
        
        case Game.Actions.FindStartingPlayer:
          FindStartingPlayer(game);
          game.action = Game.Actions.Round;
          break;

        case Game.Actions.Round:
          DoActivePlayerAction(game);
          break;                
      }

      SetSingleton(game);
    }
  }
}