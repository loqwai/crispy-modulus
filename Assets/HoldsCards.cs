using System.Collections;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;

public class HoldsCards : MonoBehaviour {
  public void AddCard(HasValue card) {
    card.transform.SetParent(transform);

    var score = cards.Select(c => c.Value).Sum();
    Debug.Log($"Now have {score} points");
  }

  HasValue[] cards {
    get {
      return GetComponentsInChildren<HasValue>();
    }
  }
}
