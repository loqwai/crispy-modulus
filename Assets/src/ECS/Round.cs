using System;
using Unity.Entities;
using Unity.Collections;

namespace OurECS {
    [Serializable]
    public struct Round : IComponentData
    {
        public int number;
    }
}