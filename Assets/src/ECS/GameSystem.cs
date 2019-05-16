using Unity.Entities;
using Unity.Mathematics;
using UnityEngine;
using Unity.Jobs;
using Unity.Collections;
using OurECS;
namespace OurECS {
  public class GameSystem : ComponentSystem {

    protected override void OnCreateManager(){
      RequireSingletonForUpdate<Game>();
    }

    protected void Start(Game game) {      
    }
    
    protected void findStartingPlayer(Game game) {
      var maxMod = -1;
      Entities.ForEach((Entity e, ref Player p) => {
        var mod = p.cardSum % game.mod;
        if (mod > maxMod) {
          maxMod = mod;
          game.currentPlayer = e;
        }
      });
    }

    protected override void OnUpdate() {
      var game = GetSingleton<Game>();

      if (game.action == Game.Actions.Start) {
        Start(game);
        game.action = Game.Actions.Round;
      }

      SetSingleton(game);
      return;
    }
  }
}