using Unity.Entities;
using Unity.Mathematics;
using UnityEngine;
using Unity.Jobs;
using Unity.Collections;
using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using OurECS;
namespace OurECS {
  public class GameSystem : ComponentSystem {
    protected EntityArchetype undrawnCardArchetype;
    System.Random random;
    
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
      var query = 
        GetEntityQuery(
          typeof(Card),
          typeof(CardFacedDown)
        );
        var ourCards = new List<Entity>();
        var cards = query.ToComponentDataArray<Card>(Allocator.TempJob);
        var entities = query.ToEntityArray(Allocator.TempJob);          
        
        for(int i = 0; i < cards.Length; i++) {
          if(cards[i].OriginalPlayer != pe) continue;
          ourCards.Add(entities[i]);
        }

        if(ourCards.Count != 0){
          var cardToDraw = random.Next(0, ourCards.Count);
          var drawnCard = ourCards[cardToDraw];
          PostUpdateCommands.RemoveComponent<CardFacedDown>(drawnCard);
          player.cardCount++;
          player.cardSum += cards[cardToDraw].Value;
          player.mod = player.cardSum % game.mod;
        }          
        cards.Dispose();
        entities.Dispose();
        
    } 

    protected void Start(Game game) {     
      CreatePlayers(game);
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
          Start(game);
          game.action = Game.Actions.Deal;
          SetSingleton(game);
          break;
        
        case Game.Actions.Deal:
          DealCards(game);
          game.action = Game.Actions.DrawUntilUnequalMods;
          SetSingleton(game);
          break;
          
        case Game.Actions.DrawUntilUnequalMods:
          DrawUntilUnequalMods(game);
          game.action = Game.Actions.FindStartingPlayer;
          SetSingleton(game);
          break;
        
        case Game.Actions.FindStartingPlayer:
          FindStartingPlayer(game);
          game.action = Game.Actions.Nothing;
          SetSingleton(game);
          break;
      }
    }
  }
}