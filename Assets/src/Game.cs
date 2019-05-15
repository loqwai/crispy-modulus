using System;
using Unity.Entities;
using Unity.Collections;

namespace OurECS {
[Serializable]
    public struct Game : IComponentData
    {
        public int cardCount;
        public int numberOfPlayers;
        public int mod;
        public Entity currentPlayer;
        public Entity whoIsWinning;         
        public Boolean isDone;        
        public enum Actions {
            Nothing,
            Start,
            Run,            
        }
        //This is clearly wrong, because I'm just relying on initialization to set action=nothing.
        public Actions action;                      
    }
}