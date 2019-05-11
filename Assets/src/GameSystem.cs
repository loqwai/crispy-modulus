using Unity.Entities;
using Unity.Mathematics;
using Unity.Collections;
using UnityEngine;
using OurECS;
namespace OurECS {
    public class GameSystem : ComponentSystem
    {
        protected EntityManager manager;
        protected override void OnCreate() {
            manager = World.Active.GetOrCreateManager<EntityManager>();
        }
        
        protected void Start(ref Game game) {        
            for (int i = 0; i < game.NumberOfPlayers; i++) {
                var entity = manager.CreateEntity();
                var player = new Player() {
                    cardCount = 0
                };            
                manager.AddComponentData(entity, player);
            }
            game.shouldStart = false;        
        }

        protected override void OnUpdate()
        {
            Entities.ForEach((ref Game game) => {
                if(game.shouldStart){
                    Start(ref game);
                }
            });
            // Entities.ForEach processes each set of ComponentData on the main thread. This is not the recommended
            // method for best performance. However, we start with it here to demonstrate the clearer separation
            // between ComponentSystem Update (logic) and ComponentData (data).
            // There is no update logic on the individual ComponentData.
            // Entities.ForEach((ref RotationSpeed rotationSpeed, ref Rotation rotation) =>
            // {
            //     var deltaTime = Time.deltaTime;
            //     rotation.Value = math.mul(math.normalize(rotation.Value),
            //         quaternion.AxisAngle(math.up(), rotationSpeed.RadiansPerSecond * deltaTime));
            // });
        }
    }
}