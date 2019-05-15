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
        public int round;
        public Entity currentPlayer;
        public Entity whoIsWinning;         
        public Boolean isDone;        
        public enum Actions {
            Nothing,
            Start,
            DealNewDeck,
            SetupPlayers,
            Run,            
        }        
        public Actions action;                      
    }
}