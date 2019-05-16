using Unity.Entities;
using Unity.Mathematics;
using UnityEngine;
using Unity.Jobs;
using Unity.Collections;

using OurECS;
namespace OurECS {
  [UpdateBefore(typeof(GameSystem))]
  public class CardSystem : ComponentSystem {
    protected EntityArchetype cardInPlayArchetype;

    protected override void OnCreateManager() {      
      RequireSingletonForUpdate<Game>();      
      cardInPlayArchetype =
         EntityManager.CreateArchetype(
           typeof(Card),
           typeof(Player),
           typeof(Round));           
    }

    protected void openANewDeckJustLikeVegas() {
      var query = GetEntityQuery(typeof(Card));
      EntityManager.DestroyEntity(query);
    }

    protected void dealCards(Game game) {

      Entities.ForEach((Entity e, ref Player p) => {
        for (int i = 0; i < game.cardCount; i++) {
          PostUpdateCommands.CreateEntity(cardInPlayArchetype);
          PostUpdateCommands.SetComponent(new Card { value = i, faceUp= false});
          PostUpdateCommands.SetComponent(p);
          PostUpdateCommands.SetComponent(new Round{number=game.round});
        }
      });
    }

    protected override void OnUpdate() {
      Debug.Log("ECS system online");
      var game = GetSingleton<Game>();
      if (game.action == Game.Actions.Start) {
        openANewDeckJustLikeVegas();
        dealCards(game);        
      }
    }
  }
}