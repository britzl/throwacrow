components {
  id: "script"
  component: "/main/block.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "durability"
    value: "100.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "undamaged"
    value: "elementStone001"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "damaged"
    value: "elementStone007"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "almost_destroyed"
    value: "elementStone045"
    type: PROPERTY_TYPE_HASH
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"\"\ntype: COLLISION_OBJECT_TYPE_DYNAMIC\nmass: 1000.0\nfriction: 1.0\nrestitution: 0.2\ngroup: \"block\"\nmask: \"block\"\nmask: \"bird\"\nmask: \"ground\"\nmask: \"wind\"\nembedded_collision_shape {\n  shapes {\n    shape_type: TYPE_SPHERE\n    position {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n      w: 1.0\n    }\n    index: 0\n    count: 1\n  }\n  data: 35.0\n}\nlinear_damping: 0.0\nangular_damping: 0.0\nlocked_rotation: false\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "factory"
  type: "factory"
  data: "prototype: \"/main/debris.go\"\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/main/throwacrow.atlas\"\ndefault_animation: \"elementStone001\"\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
