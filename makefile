main:
	go build -o ./bin .

up:
	scp ./bin/eos_gun ubuntu@3.0.115.46:./mohanson/eos_gun
	scp ./bin/eos_gun ubuntu@13.251.217.145:./mohanson/eos_gun
