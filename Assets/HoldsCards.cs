using System.Collections;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;

public class HoldsCards : MonoBehaviour {
  void Update() {
    var sortedCards = cards.OrderBy(c => c.Value).ToArray();
    for (int i = 0; i < sortedCards.Length; i++) {
      var card = sortedCards[i];
      card.transform.localPosition = new Vector3((i - 2) * 6, 0, 0);
    }
  }

  public void AddCard(Card card) {
    card.Reveal();
    card.transform.SetParent(transform);
  }

  Card[] cards {
    get {
      return GetComponentsInChildren<Card>();
    }
  }
}
