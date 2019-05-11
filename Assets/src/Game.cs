using System;
using Unity.Entities;
using Unity.Collections;

namespace OurECS {
[Serializable]
    public struct Game : IComponentData
    {
        public int CardCount;
        public int CurrentPlayer;
        public int NumberOfPlayers;    
        public int WhoIsWinning;
        public Boolean isDone;
        public Boolean shouldStart;
        public int mod;
    }
}