using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class Card : MonoBehaviour {
  public int Value = 0;
  public Text ValueDisplay;
  public GameObject Front;
  public GameObject Back;

  void Start() {
    ValueDisplay.text = $"{Value}";
  }

  public void Reveal() {
    Front.SetActive(true);
    Back.SetActive(false);
  }
}
