(function() {

var canvas = document.getElementById('mainscreen');
var data = canvas.overpowerData;

data.setTargetOne = function(hex, help) {
    var grid = canvas.muleGrid;
    var targetPt;
    var targetInfo = {ships:[], trails:[]};
    var tolerance;
    if (grid.scale > 2.5) {
        tolerance = 2;
    } else if (grid.scale > 1.75) {
        tolerance = 3;
    } else {
        tolerance = 4;
    }
    if (!help) {
        targetPt = hex;
    }
    var closestPl, dist;
    this.planetviews.forEach(function(planet) {
        if (dist === 0) {
            return;
        }
        var steps = grid.stepsBetween(planet.loc, hex);
        if (steps === 0) {
            targetPt = hex;
            targetInfo.planet = planet;
        } else if (!dist || steps < dist) {
            closestPl = planet;
            dist = steps;
        }
    });

    var closestShLoc;
    this.shipviews.forEach(function(ship) {
        if (ship.loc) {
            if (targetPt && grid.ptsEq(targetPt, ship.loc.coord)) {
                targetInfo.ships.push(ship);
            } else {
                var steps = grid.stepsBetween(ship.loc.coord, hex);
                if (!dist || steps < dist) {
                    closestPl = null;
                    closestShLoc = hex;
                    dist = steps;
                }
            }
        }
        if (targetPt) {
            ship.trail.forEach(function(pt) {
                if (grid.ptsEq(pt, targetPt)) {
                    targetInfo.trails.push(ship);
                }
            });
        }
    });
    if (!targetPt) {
        if (closestPl && dist < tolerance) {
            targetPt = closestPl.loc;
            targetInfo.planet = closestPl;
        } else if (closestShLoc && dist < tolerance) {
            targetPt = closestShLoc;
        } else {
            targetPt = hex;
        }
        this.shipviews.forEach(function(ship) {
            if (ship.loc) {
                if (targetPt && grid.ptsEq(targetPt, ship.loc.coord)) {
                    targetInfo.ships.push(ship);
                } 
            }
            ship.trail.forEach(function(pt) {
                if (grid.ptsEq(pt, targetPt)) {
                    targetInfo.trails.push(ship);
                }
            });
        });
    }
    if (targetInfo.planet && this.orders) {
        targetInfo.orders = [];
        this.orders.forEach(function(order) {
            if (order.source === targetInfo.planet.pid) {
                targetInfo.orders.push(order);
            }
        });
    }
    this.targetOne = targetPt;
    this.targetOneInfo = targetInfo;
    if (targetInfo.planet && this.targetTwoInfo && this.targetTwoInfo.planet) {
        this.setTargetOrder(targetInfo.planet, this.targetTwoInfo.planet);
    } else {
        this.setTargetOrder();
    }
};

data.setTargetTwo = function(hex, help, drop) {
    var grid = canvas.muleGrid;
    var tolerance;
    if (grid.scale > 2.5) {
        tolerance = 2;
    } else if (grid.scale > 1.75) {
        tolerance = 3;
    } else {
        tolerance = 4;
    }
    var closestPl, dist;
    this.planetviews.forEach(function(planet) {
        if (dist === 0) {
            return;
        }
        var steps = grid.stepsBetween(planet.loc, hex);
        if (!dist || steps < dist) {
            closestPl = planet;
            dist = steps;
        }
    });
    if (dist === 0 || help && dist < tolerance) {
        this.targetTwo = closestPl.loc;
        this.targetTwoInfo = {planet: closestPl};
    } else if (drop) {
        this.targetTwo = null;
        this.targetTwoInfo = null;
    }
    if (this.targetOneInfo && this.targetOneInfo.planet && this.targetTwoInfo && this.targetTwoInfo.planet) {
        this.setTargetOrder(this.targetOneInfo.planet, this.targetTwoInfo.planet);
    } else {
        this.setTargetOrder();
    }
};

data.setTargetOrder = function(pl1, pl2) {
    if (!pl1) {
        this.targetOrder = null;
        return;
    }
    var order = null;
    for (var i = 0; i< this.orders.length; i++) {
        var o = this.orders[i];
        if (o.source === pl1.pid && o.target === pl2.pid) {
            order = o;
            break;
        }
    }
    if (!order && pl1.avail) {
        order = {"new": true, "gid": this.game.gid, "fid": this.faction.fid, "source":pl1.pid, "target": pl2.pid, "size":0, "targetPl": pl2, "sourcePl": pl1};
    }
    this.targetOrder = order;
};


canvas.parseOPData = function() {
    var data = this.overpowerData;
    var fidMap = new Map();
    data.factions.forEach(function(faction) {
        fidMap.set(faction.fid, faction);
    });
    data.fidMap = fidMap;
    var plidMap = new Map();
    data.planetviews.forEach(function(planet) {
        plidMap.set(planet.pid, planet);
        if (planet.controller === data.faction.fid) {
            planet.avail = planet.parts;
        }
    });
    data.plidMap = plidMap;
    data.orders.forEach(function(order) {
        order.sourcePl = plidMap.get(order.source);
        order.sourcePl.avail -= order.size;
        order.targetPl = plidMap.get(order.target);
    });
    data.map = {"center":[0,0]};
};





canvas.parseOPData();


})();
