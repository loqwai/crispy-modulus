using Unity.Entities;
using Unity.Mathematics;
using UnityEngine;
using Unity.Jobs;
using Unity.Collections;

using OurECS;
namespace OurECS {
  public class GameSystem : ComponentSystem {        

    void onCreate() {
      EntityManager.CreateEntity(typeof(Game));
      SetSingleton(new Game{});
    }

    protected void Start(ref Game game) {
      var playerQuery = GetEntityQuery(typeof(Player));
      // var players = playerQuery.ToComponentDataArray<Player>(Allocator.TempJob);   
      // players.Dispose();      
    }
    protected void initializeGameEntity() {}

    protected override void OnUpdate() {
      var game = GetSingleton<Game>();      
      if (game.action == Game.Actions.Start) {
          Start(ref game);          
      }
      
      game.action = Game.Actions.Nothing;
      SetSingleton(game);      

    }
  }
}