using System;
using Unity.Entities;
using Unity.Collections;

namespace OurECS {
    [Serializable]
    public struct Card : IComponentData
    {
        public int value;
        public Boolean faceUp;
        public int round;
        public Entity owner;   
    }
}