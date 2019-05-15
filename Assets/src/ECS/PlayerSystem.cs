using Unity.Entities;
using Unity.Mathematics;
using UnityEngine;
using Unity.Jobs;
using Unity.Collections;
using OurECS;
namespace OurECS {
  public class PlayerSystem : ComponentSystem {        

    void onCreate() {
    }

    protected void Start(ref Game g) {
      Entities.ForEach((ref Player p) => {
        p.action = Player.Actions.NewGame;
      });      
    }

    protected override void OnUpdate() {
      if(!HasSingleton<Game>()) return;

      var game = GetSingleton<Game>();      
      if (game.action == Game.Actions.Start) {
          Start(ref game);          
      }
      
      game.action = Game.Actions.Nothing;
      SetSingleton(game);
    }
  }
}