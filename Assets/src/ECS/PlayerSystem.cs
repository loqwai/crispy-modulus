
using UnityEngine;
using Unity.Entities;
using Unity.Mathematics;
using Unity.Collections;
using OurECS;
using System;
using System.Collections;
using System.Collections.Generic;
namespace OurECS {

  [UpdateAfter(typeof(GameSystem))]
  public class PlayerSystem : ComponentSystem {
    System.Random random;
    EntityArchetype activePlayerArchetype;

    protected override void OnCreateManager() {
      RequireSingletonForUpdate<Game>();
      RequireForUpdate(GetEntityQuery(typeof(Player), typeof(ActivePlayer)));
      random = new System.Random((int)DateTime.Now.Ticks);
      
      activePlayerArchetype = EntityManager.CreateArchetype(           
          typeof(Player),
          typeof(ActivePlayer)
        );            
    }      

    protected override void OnUpdate() {      
      Entities.WithAll<Player, ActivePlayer>().
        ForEach((Entity e, ref Player p)=>{          
          Debug.Log("Forcing Active Player to Draw"); 
          p.action = Player.Actions.Draw;          
      });      
    }
  }
}