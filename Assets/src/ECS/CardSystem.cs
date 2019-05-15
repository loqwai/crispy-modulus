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

    void onCreate() {
      cardArchetype = EntityManager.CreateArchetype(typeof(Card));
    }
    protected void openANewDeckJustLikeVegas() {
      var query = GetEntityQuery(typeof(Card));
      var cardEntities = query.ToEntityArray(Allocator.TempJob);
      foreach(var e in cardEntities ) {
        PostUpdateCommands.DestroyEntity(e);
      }      
    }

    protected void dealCards(Game game) {
      Entities.ForEach( (Entity e, ref Player p) => {
        for(int i = 0; i < game.cardCount; i++) {
          PostUpdateCommands.CreateEntity(cardArchetype);   
          PostUpdateCommands.AddComponent(new Card{value=i, faceUp=false, round=0, owner=e});           
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