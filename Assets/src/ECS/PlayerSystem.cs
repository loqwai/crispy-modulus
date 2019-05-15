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
    protected override void OnCreateManager () {
      random = new System.Random((int)DateTime.Now.Ticks);
       cardArchetype = EntityManager.CreateArchetype(
          typeof(Card),
          typeof(Player),
          typeof(Round)
        );
    }    

    protected void Start(Game g) {
      Entities.ForEach((ref Player p) => {
          Draw(p, g.round);
          p.cardCount++;
      });      
    }
  
    protected void Draw(Player player, int currentRound) {
        var inefficient = new List<Tuple<Entity, Card>>();
        
        //jesus.
        Entities.ForEach( (Entity e, ref Player owner, ref Round r, ref Card c) => {
            if(!player.Equals(owner)) return;
            if(r.number != currentRound) return;
            if(c.faceUp) return;        
            //this is the worst;    
            inefficient.Add(Tuple.Create(e, c));
        }); 

        if(inefficient.Count == 0) return;       
        
        var index = random.Next(0, inefficient.Count);        
        var pair = inefficient[index];   
        var cardEntity = pair.Item1;
        var isThisYourCard = pair.Item2;     
        isThisYourCard.faceUp = true;
        PostUpdateCommands.CreateEntity(cardArchetype);                        
        PostUpdateCommands.SetComponent(player);        
        PostUpdateCommands.SetComponent(new Round{number = currentRound++});        
        PostUpdateCommands.SetComponent(isThisYourCard);        
    }    

    protected void Steal(Player player, int value, int currentRound) {
        var inefficient = new List<Tuple<Entity, Card>>();
        
        //jesus.
        Entities.ForEach( (Entity e, ref Player owner, ref Round r, ref Card c) => {
            if(c.value != value) return; //wrong one.
            if(r.number != currentRound) return; //this is history
            if(player.Equals(owner)) return; //can't steal my own shit;            
            if(!c.faceUp) return; //can't steal face down cards
            //this is the worst;    
            inefficient.Add(Tuple.Create(e, c));
        }); 

        if(inefficient.Count == 0) return;       
        
        var index = random.Next(0, inefficient.Count);        

        var pair = inefficient[index];   
        var cardEntity = pair.Item1;
        var isThisYourCard = pair.Item2;     

        PostUpdateCommands.CreateEntity(cardArchetype);                        
        PostUpdateCommands.SetComponent(player);        
        PostUpdateCommands.SetComponent(new Round{number = currentRound++});        
        PostUpdateCommands.SetComponent(isThisYourCard);        
    }  

    protected override void OnUpdate() {
      if(!HasSingleton<Game>()) return;

      var game = GetSingleton<Game>();      
      if (game.action == Game.Actions.Start) {
          Start(game);          
      }
    }
  }
}