using System.Collections;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;

public class DrawCardOnClick : MonoBehaviour {
  public Transform Hand;

  void OnMouseDown() {
    if (children.Length == 0) return;

    var card = children.First();
    card.SetParent(Hand);
  }

  Transform[] children {
    get {
      var children = new List<Transform>();
      foreach (Transform child in transform) {
        children.Add(child);
      }
      return children.ToArray();
    }
  }

}
