using System;
using Unity.Entities;
using Unity.Collections;

[Serializable]
public struct Player : IComponentData
{
    public enum Actions {
        Nothing,
        Draw,
        Steal,
        NewGame,
    }
    
    public Actions action;
    public int cardCount;
}