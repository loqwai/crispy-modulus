using Unity.Entities;
using Unity.Mathematics;
using UnityEngine;
using Unity.Jobs;
using Unity.Collections;
using OurECS;
namespace OurECS {
  public class GameSystem : ComponentSystem {

    protected override void OnCreate() {
      EntityManager.CreateEntity(typeof(Game));
      initGame();
    }

    protected Game initGame() {
      var game = new Game{round=0, action=Game.Actions.Start, cardCount=10, mod=5, numberOfPlayers=2};
      SetSingleton(game);
      return game;      
    }

  protected void Start(Game game) {   
      var maxMod = game.cardCount;
      Entities.ForEach((Entity e, ref Player p)=>{
        var mod = p.cardSum % game.mod;        
        if(mod > maxMod) {
          maxMod = mod;                
          game.currentPlayer = e;
        }		
      });       
    }

    protected override void OnUpdate() {
      if (!HasSingleton<Game>()) return;

      var game = GetSingleton<Game>();
      
      if (game.action == Game.Actions.Start) {
        game = initGame();
        Start(game);
        game.action = Game.Actions.Round;
      }

      SetSingleton(game);
      return;
    }
  }
}