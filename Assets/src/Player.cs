using System;
using Unity.Entities;
using Unity.Collections;

[Serializable]
public struct Player : IComponentData
{
    // enum action {
    //     Draw,
    //     Steal
    // }
    public int cardCount;
}