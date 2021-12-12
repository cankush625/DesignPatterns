package strategy_design_pattern

type Bird struct {
	FlyBehaviourStrategy   FlyBehaviourStrategy
	SpeakBehaviourStrategy SpeakBehaviourStrategy
}

func (b Bird) Fly() {
	b.FlyBehaviourStrategy.fly()
}

func (b Bird) Speak() {
	b.SpeakBehaviourStrategy.speak()
}
