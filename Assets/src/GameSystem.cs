using Unity.Entities;
using Unity.Mathematics;
using UnityEngine;
public class GameSystem : ComponentSystem
{
    protected override void OnCreate() {
        Debug.Log("I was born");
    }
    protected override void OnUpdate()
    {
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