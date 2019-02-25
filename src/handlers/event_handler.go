package handler

import (
	"log"
	"github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
)

func handleMotionEvent(event dtypes.Event) dtypes.Event {
	
	var replyEvent dtypes.Event
	if event.EventType=="SendUpdate" {
		replyEvent = dtypes.Event { 
				EventType: "Update",
				Object:event.Object,
				B1Pos: event.B1Pos,
				B2Pos: event.B2Pos,
				B3Pos: event.B3Pos,
				P1Pos: event.P1Pos
				P2Pos: event.P2Pos
				G1Pos: event.G1Pos,
				G2Pos: event.G2Pos,
				G3Pos: event.G3Pos,
				G4Pos: event.G4Pos,
				P1health:GetHealth("p1"),
				P2health:GetHealth("p2")
			}
		return replyEvent
	} 
	if event.EventType=="space"{
		if event.Object=="p1"{
				j:=rand.Intn(10)
				replyEvent = dtypes.Event { 
				EventType: "Teleport",
				Object:event.Object,
				B1Pos: event.B1Pos,
				B2Pos: event.B2Pos,
				B3Pos: event.B3Pos,
				P1Pos: coords.Randompos[j]
				P2Pos: event.P2Pos
				G1Pos: event.G1Pos,
				G2Pos: event.G2Pos,
				G3Pos: event.G3Pos,
				G4Pos: event.G4Pos,
				P1health:GetHealth("p1"),
				P2health:GetHealth("p2")
			}
		}else if event.Object=="p2"{
				j:=rand.Intn(10)
				replyEvent = dtypes.Event { 
				EventType: "Teleport",
				Object:event.Object,
				B1Pos: event.B1Pos,
				B2Pos: event.B2Pos,
				B3Pos: event.B3Pos,
				P1Pos: event.P1Pos
				P2Pos: coords.Randompos[j],
				G1Pos: event.G1Pos,
				G2Pos: event.G2Pos,
				G3Pos: event.G3Pos,
				G4Pos: event.G4Pos,
				P1health: GetHealth("p1"),
				P2health: GetHealth("p2")
			}
		}
		return replyEvent
	}
	direction := [4]string{"Up", "Down", "Left", "Right"}
	var freeFallP1 bool = false
	var freeFallP2 bool = false
	var step int = 2
	dx := [5]int{0, 0, -step, step,}
	dy := [5]int{-step, step, 0, 0,0}
	for i := 0; i < 4 ; i++ {		
		if direction[i] == event.EventType {
			log.Println("Direction detected:", direction[i])
			replyEvent = dtypes.Event { 
				EventType: "Update",
				Object:    event.Object,
				B1Pos:     event.B1Pos,
				B2Pos:     event.B2Pos,
				B3Pos:     event.B3Pos

			}
			log.Println("Set default attr for replyEvent")
			if event.Object == "p1" {
				log.Println("Object of this event is p1")
				replyEvent.P1Pos = dtypes.Position {
					X: event.P1Pos.X + dx[i], 
					Y: event.P1Pos.Y + dy[i]
				}
				var p11 Rect = GetBoundary(event.P1Pos)
				var p22 Rect = GetBoundary(replyEvent.P1Pos)
				var updatedRect Rect
				if i == 0 {
					if(!AllignedWithLadder(p11)) {
						// no change
						log.Println("up but not alligned with ladder")
						updatedRect Rect = p11
					} else if(AllignedWithLadder(p11) && AllignedWithLadder(p22)) {
						// success
						log.Println("up and not alligned with ladder")
						updatedRect Rect = p22
					} else {
						// get ladder top
						log.Println("up and not alligned with ladder restricted")
						updatedRect Rect = SetAccordingToLadderTop(p11)
					}
			} else if i==1 {
				if(!AllignedWithLadder(p11)){
					// no change
					log.Println("down but not alligned with ladder")
					updatedRect Rect=p11
				}else if(AllignedWithLadder(p11) && AllignedWithLadder(p22)){
					// success
					log.Println("down and alligned with ladder")
					updatedRect Rect=p22
				}else{
					// get ladder bottom
						log.Println("down and alligned with ladder restricted")
					updatedRect Rect=SetAccordingToLadderBottom(p11)
				}
			}else if i==2{
				if AllignedWithLadder(p11) && AllignedWithLadder(p22){
					updatedRect Rect=p22
					log.Println("was alligned with ladder on pressing left")
				}else if AllignedWithLadder(p11) && !AllignedWithLadder(p22){
					log.Println("freefall")
					freeFallP1=true;
				}else if !OnPlatform(p11){
					log.Println("not on Platform")
					updatedRect Rect=p11
				}else if(CollidesWithBlockOnLeftMove(p22)){
					log.Println("collided with block on left")
					updatedRect Rect=GetPositionCollidesWithBlockOnLeft(p22)
				}else if(FallsFromBlock(p22)){
						log.Println("fell from block and freefall")
						freeFallP1=true;
				}else{
						log.Println("successfull left move")
					updatedRect Rect=p22
				}
				CollidesGem(updatedRect,"p1")
			}else if i==3{
				if AllignedWithLadder(p11) && AllignedWithLadder(p22){
					updatedRect Rect=p22
					log.Println("was alligned with ladder on pressing right")
				}else if AllignedWithLadder(p11) && !AllignedWithLadder(p22){
					freeFallP1=true
					log.Println("freefall")
				}else if !OnPlatform(p11){
					log.Println("not on Platform")
					updatedRect Rect=p11
				}else if(CollidesWithBlockOnRightMove(p22)){
					log.Println("collided with block on right")
					updatedRect Rect=GetPositionCollidesWithBlockOnRight(p22)
				}else if(FallsFromBlock(p22)){
					log.Println("fell from block and freefall")
					freeFallP1=true
				}else{
					log.Println("successfull right move")
					updatedRect Rect=p22
				}
				CollidesGem(updatedRect,"p1")
			}else if(i==4){

			}
			var  dtypes.Position temporary =getposition(updated)

			if (freeFallP1){
				log.Println("freefall")
				var  dtypes.Position temporary2 ={temporary.X,temporary.Y+2*step}
				var p11 Rect =GetBoundary(temporary)
				var p22 Rect =GetBoundary(temporary2)
				if CollidesWithBlockVertically(p22){
					log.Println("collidedwith block while freefalling")
					freeFallP1=false;
					updatedRect Rect=GetPositionCollidesWithBlockVer(p22)
				}else{
					updatedRect Rect=p22
				}
			}
			replyEvent.P1Pos=getposition(updated)
			replyEvent.P2Pos=event.P2Pos
			replyEvent.P1health=GetHealth("p1")
			replyEvent.P2health=GetHealth("p2")
			replyEvent.G1Pos=getposition(coords.gems[0].pos)
			replyEvent.G2Pos=getposition(coords.gems[1].pos)
			replyEvent.G3Pos=getposition(coords.gems[2].pos)
			replyEvent.G4Pos=getposition(coords.gems[3].pos)
		}

		if event.Object == "p2" {
				log.Println("Object of this event is p2")
				replyEvent.P2Pos = dtypes.Position {
					X: event.P2Pos.X + dx[i], 
					Y: event.P2Pos.Y + dy[i]
				}
				var p11 Rect = GetBoundary(event.P2Pos)
				var p22 Rect = GetBoundary(replyEvent.P2Pos)
				var updatedRect Rect
				if i == 0 {
					if(!AllignedWithLadder(p11)) {
						// no change
						log.Println("up but not alligned with ladder")
						updatedRect Rect = p11
					} else if(AllignedWithLadder(p11) && AllignedWithLadder(p22)) {
						// success
						log.Println("up and not alligned with ladder")
						updatedRect Rect = p22
					} else {
						// get ladder top
						log.Println("up and not alligned with ladder restricted")
						updatedRect Rect = SetAccordingToLadderTop(p11)
					}
			} else if i==1 {
				if(!AllignedWithLadder(p11)){
					// no change
					log.Println("down but not alligned with ladder")
					updatedRect Rect=p11
				}else if(AllignedWithLadder(p11) && AllignedWithLadder(p22)){
					// success
					log.Println("down and alligned with ladder")
					updatedRect Rect=p22
				}else{
					// get ladder bottom
						log.Println("down and alligned with ladder restricted")
					updatedRect Rect=SetAccordingToLadderBottom(p11)
				}
			}else if i==2{
				if AllignedWithLadder(p11) && AllignedWithLadder(p22){
					updatedRect Rect=p22
					log.Println("was alligned with ladder on pressing left")
				}else if AllignedWithLadder(p11) && !AllignedWithLadder(p22){
					log.Println("freefall")
					freeFallP2=true;
				}else if !OnPlatform(p11){
					log.Println("not on Platform")
					updatedRect Rect=p11
				}else if(CollidesWithBlockOnLeftMove(p22)){
					log.Println("collided with block on left")
					updatedRect Rect=GetPositionCollidesWithBlockOnLeft(p22)
				}else if(FallsFromBlock(p22)){
						log.Println("fell from block and freefall")
						freeFallP2=true;
				}else{
						log.Println("successfull left move")
					updatedRect Rect=p22
				}
				CollidesGem(updatedRect,"p2")
			}else if i==3{
				if AllignedWithLadder(p11) && AllignedWithLadder(p22){
					updatedRect Rect=p22
					log.Println("was alligned with ladder on pressing right")
				}else if AllignedWithLadder(p11) && !AllignedWithLadder(p22){
					freeFallP2=true
					log.Println("freefall")
				}else if !OnPlatform(p11){
					log.Println("not on Platform")
					updatedRect Rect=p11
				}else if(CollidesWithBlockOnRightMove(p22)){
					log.Println("collided with block on right")
					updatedRect Rect=GetPositionCollidesWithBlockOnRight(p22)
				}else if(FallsFromBlock(p22)){
					log.Println("fell from block and freefall")
					freeFallP2=true
				}else{
					log.Println("successfull right move")
					updatedRect Rect=p22
				}
				CollidesGem(updatedRect,"p2")
			}
			var  dtypes.Position temporary =getposition(updated)

			if (freeFallP2){
				log.Println("freefall")
				var  dtypes.Position temporary2 ={temporary.X,temporary.Y+2*step}
				var p11 Rect =GetBoundary(temporary)
				var p22 Rect =GetBoundary(temporary2)
				if CollidesWithBlockVertically(p22){
					log.Println("collidedwith block while freefalling")
					freeFallP2=false;
					updatedRect Rect=GetPositionCollidesWithBlockVer(p22)
				}else{
					updatedRect Rect=p22
				}
			}
			replyEvent.P2Pos=getposition(updated)
			replyEvent.P1Pos=event.P1Pos
			replyEvent.P1health=GetHealth("p1")
			replyEvent.P2health=GetHealth("p2")
			replyEvent.G1Pos=getposition(coords.gems[0].pos)
			replyEvent.G2Pos=getposition(coords.gems[1].pos)
			replyEvent.G3Pos=getposition(coords.gems[2].pos)
			replyEvent.G4Pos=getposition(coords.gems[3].pos)
		}
		
	}
 }
	return replyEvent
}

