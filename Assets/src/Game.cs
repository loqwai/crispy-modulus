using System;
using Unity.Entities;
using Unity.Collections;

[Serializable]
public struct Game : IComponentData
{
    public int CardCount;
    public int CurrentPlayer;
    public int NumberOfPlayers;    
    public int WhoIsWinning;
    public bool isDone
    {
        get { return _isDone == 1; }
        set { _isDone = value == true ? 1 : 0; }
    }
    private int _isDone;
}