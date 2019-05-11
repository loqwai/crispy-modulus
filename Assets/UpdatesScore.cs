using System.Collections;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;
using UnityEngine.UI;

public class UpdatesScore : MonoBehaviour {
  public Text ScoreDisplay;

  void Update() {
    ScoreDisplay.text = $"Score: {score}";
  }

  int score {
    get {
      return cards.Select(c => c.Value).Sum();
    }
  }

  Card[] cards {
    get {
      return GetComponentsInChildren<Card>();
    }
  }
}
