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
           typeof(Player),
           typeof(Round)
         );
      }

      protected void Start(Game g) {        
        Entities.ForEach((Entity e, ref Player p) => {
          if(!EntityManager.HasComponent<Round>(e)){
            PostUpdateCommands.AddComponent(e, new Round());          
            return;
          }
          Draw(e, p, g.round);
        });
      }

      protected void Draw(Entity pe, Player player, int currentRound) {
        var inefficient = new List<Tuple<Entity, Card>>();
        var playerRound = EntityManager.GetComponentData<Round>(pe);
        if(playerRound.number > currentRound) return;
        playerRound.number++;
        //jesus.
        Entities.ForEach((Entity e, ref Player owner, ref Round r, ref Card c) => {
          if (!player.Equals(owner)) return;
          if (r.number != currentRound) return;
          if (c.faceUp) return;
          //this is the worst;    
          inefficient.Add(Tuple.Create(e, c));
        });

        if (inefficient.Count == 0) return;
        var index = random.Next(0, inefficient.Count);
        var pair = inefficient[index];
        var cardEntity = pair.Item1;
        var isThisYourCard = pair.Item2;
        player.cardCount++;
        player.cardSum += isThisYourCard.value;
        playerRound.number++;

        isThisYourCard.faceUp = true;
        PostUpdateCommands.SetComponent(pe, player);
        PostUpdateCommands.SetComponent(pe, playerRound);
        PostUpdateCommands.CreateEntity(cardArchetype);
        PostUpdateCommands.SetComponent(player);
        PostUpdateCommands.SetComponent(playerRound);
        PostUpdateCommands.SetComponent(isThisYourCard);        
      }

      protected void Steal(Entity pe, Player player, int value, int currentRound) {
        var inefficient = new List<Tuple<Entity, Card>>();

        //jesus.
        Entities.ForEach((Entity e, ref Player owner, ref Round r, ref Card c) => {
          if (c.value != value) return; //wrong one.
          if (r.number != currentRound) return; //this is history
          if (player.Equals(owner)) return; //can't steal my own shit;            
          if (!c.faceUp) return; //can't steal face down cards
                                 //this is the worst;    
          inefficient.Add(Tuple.Create(e, c));
        });

        if (inefficient.Count == 0) return;

        var index = random.Next(0, inefficient.Count);

        var pair = inefficient[index];
        var cardEntity = pair.Item1;
        var isThisYourCard = pair.Item2;
        player.cardCount++;
        player.cardSum -= isThisYourCard.value;
        PostUpdateCommands.CreateEntity(cardArchetype);
        PostUpdateCommands.SetComponent(player);
        PostUpdateCommands.SetComponent(new Round { number = currentRound++ });
        PostUpdateCommands.SetComponent(isThisYourCard);
        PostUpdateCommands.SetComponent(pe, player);
      }

    protected override void OnUpdate() {
      var game = GetSingleton<Game>();
      if (game.action == Game.Actions.Start) {
        Start(game);
      }
    }
  }
}