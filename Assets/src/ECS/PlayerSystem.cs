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


  public class PlayerSystem : ComponentSystem {
    System.Random random;        

    void onCreate() {
        random = new System.Random((int)DateTime.Now.Ticks);
    }

    protected void Start(Game g) {
      Entities.ForEach((Entity e, ref Player p) => {
          Draw(e, g.round);
          p.cardCount++;
      });      
    }
  
    protected void Draw(Entity playerEntity, int round) {
        List<Entity> inefficient = new List<Entity>();        
        
        //jesus.
        Entities.ForEach( (Entity e, ref Card c) => {
            if(c.owner != playerEntity) return;
            if(c.faceUp) return;
            inefficient.Add(e);
        }); 
        if(inefficient.Count == 0) return;       
        var index = random.Next(0, inefficient.Count);        
        return;
        var cardEntity = inefficient[index];        
        var isThisYourCard = EntityManager.GetComponentData<Card>(cardEntity);
        isThisYourCard.round++;
        isThisYourCard.faceUp = false;
        PostUpdateCommands.SetComponent(cardEntity, isThisYourCard);        
    }

    protected override void OnUpdate() {
      if(!HasSingleton<Game>()) return;

      var game = GetSingleton<Game>();      
      if (game.action == Game.Actions.SetupPlayers) {
          Start(game);          
          game.action = Game.Actions.Nothing;
          SetSingleton(game);  
      }
    }
  }
}