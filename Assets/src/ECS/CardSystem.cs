using Unity.Entities;
using Unity.Mathematics;
using UnityEngine;
using Unity.Jobs;
using Unity.Collections;

using OurECS;
namespace OurECS {
  [UpdateAfter(typeof(GameSystem))]
  public class CardSystem : ComponentSystem {
    protected EntityArchetype cardInPlayArchetype;

    protected override void OnCreateManager() {      
      RequireSingletonForUpdate<Game>();      
      cardInPlayArchetype =
         EntityManager.CreateArchetype(
           typeof(Card),
           typeof(CardFacedDown)
        );         
    }

    protected void openANewDeckJustLikeVegas() {
      var query = GetEntityQuery(typeof(Card));
      EntityManager.DestroyEntity(query);
    }

    protected void dealCards(Game game) {
      
      Entities.ForEach((Entity pe, ref Player p) => {
        for (int i = 1; i < game.cardCount+1; i++) {
          var e =PostUpdateCommands.CreateEntity(cardInPlayArchetype);
          PostUpdateCommands.SetComponent(e, new Card { Value = i, OriginalPlayer=pe});
        }
      });
    }

    protected override void OnUpdate() {
    }
  }
}