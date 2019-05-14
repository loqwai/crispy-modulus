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
        manager = World.Active.GetOrCreateManager<EntityManager>();
    }
        
    protected void Start(ref Game game) {
      var entities = new NativeArray<Entity>(game.numberOfPlayers, Allocator.Temp);
      for (int i = 0; i < game.numberOfPlayers; i++) {
        entities[i] = PostUpdateCommands.Instantiate<Player>();
        PostUpdateCommands.SetComponent(entities[i], new Player());
      }  
        game.shouldStart = false;        
    }

    protected override void OnUpdate() {
      Entities.ForEach((Entity e, ref Game game) => {
        if(game.shouldStart) {
          Start(ref game);
        }
      });
      }        
    }
    }