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
    public enum Errors {
        Nothing,
        CantDraw,
        CantSteal
    };
    public Actions action;
    public Errors error;
    public int cardCount;
    public int cardSum;
    public int mod;
}