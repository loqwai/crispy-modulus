using System;
using Unity.Entities;
using Unity.Collections;

namespace OurECS {
[Serializable]
    public struct Game : IComponentData
    {
        public int cardCount;
        public int playerCount;
        public int mod;
        public int round;    
        public Boolean isDone;        
        public enum Actions {
            Nothing,
            Start,
            Deal,
            DrawUntilUnequalMods,
            FindStartingPlayer,
            Round,            
        }        
        public Actions action;                      
    }
}