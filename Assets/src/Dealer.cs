using System;
using System.Collections.Generic;
using Unity.Entities;
using UnityEngine;

[RequiresEntityConversion]
public class Dealer : MonoBehaviour, IDeclareReferencedPrefabs, IConvertGameObjectToEntity {
    public int NumberOfCards;
    public GameObject Card;

    public void Convert(Entity entity, EntityManager dstManager, GameObjectConversionSystem conversionSystem) {
        DealerData dealerData = new DealerData {
            Card = conversionSystem.GetPrimaryEntity(Card),
            NumberOfCards = NumberOfCards,
        };
        dstManager.AddComponentData(entity, dealerData);
    }

    public void DeclareReferencedPrefabs(List<GameObject> referencedPrefabs) {
        referencedPrefabs.Add(Card);
    }
}
