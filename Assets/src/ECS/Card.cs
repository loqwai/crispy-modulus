using System;
using Unity.Entities;
using Unity.Collections;

namespace OurECS {
  [Serializable]
  public struct Card : IComponentData {
    public int Value;
    public Entity OriginalPlayer;
  }
}