using Unity.Collections;
using Unity.Entities;
using Unity.Jobs;
using Unity.Mathematics;
using Unity.Transforms;

public class DealerSystem : JobComponentSystem {
    struct SpawnJob : IJobForEachWithEntity<DealerData, LocalToWorld> {
        public EntityCommandBuffer.Concurrent CommandBuffer;

        public void Execute(Entity entity, int index, [ReadOnly] ref DealerData dealerData, [ReadOnly] ref LocalToWorld location) {
            for (var i = 0; i < dealerData.NumberOfCards; i++) {
                var card = CommandBuffer.Instantiate(index, dealerData.Card);
                var position = math.transform(location.Value, new float3(i, 0, 0));

                CommandBuffer.SetComponent(index, card, new Translation { Value = position });
            };

            // The dealer destroys itself
            CommandBuffer.DestroyEntity(index, entity);
        }
    }


    BeginInitializationEntityCommandBufferSystem entityCommandBufferSystem;

    protected override void OnCreate() {
        entityCommandBufferSystem = World.GetOrCreateSystem<BeginInitializationEntityCommandBufferSystem>();
    }

    protected override JobHandle OnUpdate(JobHandle inputDeps) {
        var job = new SpawnJob {
            CommandBuffer = entityCommandBufferSystem.CreateCommandBuffer().ToConcurrent()
        }.Schedule(this, inputDeps);
        entityCommandBufferSystem.AddJobHandleForProducer(job);
        return job;
    }
}
