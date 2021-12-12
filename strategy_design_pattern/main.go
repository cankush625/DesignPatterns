package strategy_design_pattern

func CreateBird(fly FlyBehaviourStrategy, speak SpeakBehaviourStrategy) Bird {
	bird := Bird{FlyBehaviourStrategy: fly, SpeakBehaviourStrategy: speak}
	return bird
}
