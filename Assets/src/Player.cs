using System;
using Unity.Entities;
using Unity.Collections;

[Serializable]
public struct Player : IComponentData
{
    enum action {
        Draw,
        Steal
    }
    int cardCount;
    NativeArray<int> hand;
    NativeArray<int> deck;
}