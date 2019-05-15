using Unity.Entities;
using Unity.Mathematics;
using UnityEngine;
using Unity.Jobs;
using Unity.Entities;
using Unity.Collections;
using Unity.Entities;

using OurECS;
namespace OurECS {
  public class GameSystem : ComponentSystem {
    protected EntityManager manager;
    protected BeginInitializationEntityCommandBufferSystem commandBufferSystem;

    protected override void OnCreate() {
    }

    protected void Start(ref Game game) {
      var playerQuery = GetEntityQuery(typeof(Player));
      using (var players = playerQuery.ToComponentDataArray<Player>(Allocator.TempJob)) {
        foreach (var p in players) {
          Debug.Log("found a player");
        }
      }


      game.shouldStart = false;
    }
    protected override void OnUpdate() {
      Entities.ForEach((Entity e, ref Game game) => {
        if (game.shouldStart) {
          Start(ref game);
        }
      });
    }
  }
}