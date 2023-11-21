package helper_test

import (
	"day_02/helper"
	"testing"
)

func TestGetMyPlay(t *testing.T) {
	type args struct {
		elf         string
		matchResult string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Elf_Rock_Lose", args{helper.RockString, helper.LoseString}, helper.Scissors},
		{"Elf_Rock_Draw", args{helper.RockString, helper.DrawString}, helper.Rock},
		{"Elf_Rock_Win", args{helper.RockString, helper.WinString}, helper.Paper},
		{"Elf_Paper_Lose", args{helper.PaperString, helper.LoseString}, helper.Rock},
		{"Elf_Paper_Draw", args{helper.PaperString, helper.DrawString}, helper.Paper},
		{"Elf_Paper_Win", args{helper.PaperString, helper.WinString}, helper.Scissors},
		{"Elf_Scissors_Lose", args{helper.ScissorsString, helper.LoseString}, helper.Paper},
		{"Elf_Scissors_Draw", args{helper.ScissorsString, helper.DrawString}, helper.Scissors},
		{"Elf_Scissors_Win", args{helper.ScissorsString, helper.WinString}, helper.Rock},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helper.GetMyPlay(tt.args.elf, tt.args.matchResult); got != tt.want {
				t.Errorf("GetMyPlay() = %v, want %v", got, tt.want)
			}
		})
	}
}
