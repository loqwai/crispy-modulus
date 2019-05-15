using Unity.Entities;
using Unity.Mathematics;
using UnityEngine;
using Unity.Jobs;
using Unity.Collections;

using OurECS;
namespace OurECS {
  [UpdateBefore(typeof(GameSystem))]
  public class CardSystem : ComponentSystem {            
    protected EntityArchetype cardArchetype;
    
    protected void openANewDeckJustLikeVegas() {
      var query = GetEntityQuery(typeof(Card));
      EntityManager.DestroyEntity(query);
    }

    protected void dealCards(Game game) {
      var cardArchetype = EntityManager.CreateArchetype(typeof(Card));
      Entities.ForEach((Entity e, ref Player p) => {
        for(int i = 0; i < game.cardCount; i++) {
          PostUpdateCommands.CreateEntity(cardArchetype);   
          PostUpdateCommands.SetComponent(new Card{value=i, faceUp=false, owner=e});       
        }
      });      
    }

    protected override void OnUpdate() {
      if(!HasSingleton<Game>()) return;

      var game = GetSingleton<Game>();      
      if (game.action == Game.Actions.Start) {
          openANewDeckJustLikeVegas();
          dealCards(game);          
      }            
    }
  }
}