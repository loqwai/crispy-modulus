using Unity.Entities;
using Unity.Mathematics;
using UnityEngine;
using Unity.Jobs;
using Unity.Collections;

using OurECS;
namespace OurECS {
  public class GameSystem : ComponentSystem {        

    void onCreate() {
      //this has to happen first, or SetSingleton will silently fail
      EntityManager.CreateEntity(typeof(Game));
      SetSingleton(new Game{
        action=Game.Actions.Nothing,
        numberOfPlayers=2,
        mod=3,
        cardCount=10,
        round=0
      });
    }

    protected void Start(ref Game g) {            
      Entities.ForEach((ref Player p) => {
        p.action = Player.Actions.Draw;
      });      
    }

    protected void initializeGameEntity() {}

    protected override void OnUpdate() {
      if(!HasSingleton<Game>()) return;

      var game = GetSingleton<Game>();      
      if (game.action == Game.Actions.Start) {
        game.action = Game.Actions.DealNewDeck;            
        return;
      }
      
      game.action = Game.Actions.Nothing;
      SetSingleton(game);
    }
  }
}