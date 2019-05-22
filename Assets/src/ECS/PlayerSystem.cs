
using UnityEngine;
using Unity.Entities;
using Unity.Mathematics;
using Unity.Collections;
using OurECS;
using System;
using System.Collections;
using System.Collections.Generic;
namespace OurECS {

  [UpdateAfter(typeof(CardSystem))]
  public class PlayerSystem : ComponentSystem {
    System.Random random;
    EntityArchetype cardArchetype;

    protected override void OnCreateManager() {
      RequireSingletonForUpdate<Game>();

        random = new System.Random((int)DateTime.Now.Ticks);
        cardArchetype = EntityManager.CreateArchetype(
           typeof(Card),
           typeof(Player)
         );
      }

      protected void Start(Game g) {        
        Entities.ForEach((Entity e, ref Player p) => {         
          Draw(e, ref p);
        });
      }

      protected void Draw(Entity pe, ref Player player) {
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
          }          
          cards.Dispose();
          entities.Dispose();
          
      }      
      protected void Steal(Entity pe, Player player, int value, int currentRound) {      
      }

    protected override void OnUpdate() {
      return;
      var game = GetSingleton<Game>();
      if (game.action == Game.Actions.Start) {
        Start(game);
      }
    }
  }
}