using System.Collections;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;

public class DrawCardOnClick : MonoBehaviour {
  public HoldsCards Hand;

  void OnMouseDown() {
    if (cards.Length == 0) return;

    Hand.AddCard(cards.First());
  }

  HasValue[] cards {
    get {
      return GetComponentsInChildren<HasValue>();
    }
  }
}
