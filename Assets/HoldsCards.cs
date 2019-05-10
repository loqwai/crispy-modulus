using System.Collections;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;

public class HoldsCards : MonoBehaviour {
  public void AddCard(Card card) {
    card.transform.SetParent(transform);

    var score = cards.Select(c => c.Value).Sum();
    Debug.Log($"Now have {score} points");
  }

  Card[] cards {
    get {
      return GetComponentsInChildren<Card>();
    }
  }
}
