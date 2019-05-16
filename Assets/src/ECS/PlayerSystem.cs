
using Unity.Entities;
using Unity.Mathematics;
using Unity.Jobs;
using Unity.Collections;
using OurECS;
using System.Linq;
using System;
using System.Collections;
using System.Collections.Generic;
namespace OurECS {

  [UpdateAfter(typeof(CardSystem))]
  [UpdateBefore(typeof(GameSystem))]
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
          Draw(e, p);
        });
      }

      protected void Draw(Entity pe, Player player) {
        var faceDownCards = new List<Entity>();
        
          var results = GetEntityQuery(typeof(Player), typeof(Card), typeof(CardFaceDown)).ToEntityArray(Allocator.TempJob);
          foreach(Entity e in results){
            var otherPlayer = EntityManager.GetComponentData<Player>(e);
            if(!player.Equals( otherPlayer)) return;
            faceDownCards.Add(e);
        }        
        results.Dispose();
        faceDownCards.Count();       
      }

      protected void Steal(Entity pe, Player player, int value, int currentRound) {      
      }

    protected override void OnUpdate() {
      var game = GetSingleton<Game>();
      if (game.action == Game.Actions.Start) {
        Start(game);
      }
    }
  }
}