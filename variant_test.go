package shrub

import (
	"testing"
)

func TestVariantBuilders(t *testing.T) {
	cases := map[string]func(*testing.T, *Variant){
		"NameSetter": func(t *testing.T, v *Variant) {
			assert(t, v.BuildName == "", "default value")
			v2 := v.Name("foo")
			assert(t, v.BuildName == "foo", "expected value")
			assert(t, v2 == v, "chainable")
		},
		"DisplayNameSetter": func(t *testing.T, v *Variant) {
			assert(t, v.BuildDisplayName == "", "default value")
			v2 := v.DisplayName("foo")
			assert(t, v.BuildDisplayName == "foo", "expected value")
			assert(t, v2 == v, "chainable")
		},
		"BatchTimeSetter": func(t *testing.T, v *Variant) {
			assert(t, v.BatchTimeSecs == 0, "default value")
			v2 := v.BatchTime(12)
			assert(t, v.BatchTimeSecs == 12, "expected value")
			assert(t, v2 == v, "chainable")
		},
		"CronBatchTimeSetter": func(t *testing.T, v *Variant) {
			assert(t, v.CronBatchTime == "", "default value")
			v2 := v.SetCronBatchTime("12")
			assert(t, v.CronBatchTime == "12", "expected value")
			assert(t, v2 == v, "chainable")
		},
		"StepbackSetter": func(t *testing.T, v *Variant) {
			assert(t, v.Stepback == nil, "default value")
			stepbackVal := true
			v2 := v.SetStepback(&stepbackVal)
			assert(t, *v.Stepback == stepbackVal, "expected value")
			assert(t, v2 == v, "chainable")
		},
		"ActivateSetter": func(t *testing.T, v *Variant) {
			assert(t, v.Activate == nil, "default value")
			activateVal := false
			v2 := v.SetActivate(&activateVal)
			require(t, v.Activate != nil, "expected value")
			assert(t, *v.Activate == activateVal, "expected value")
			assert(t, v2 == v, "chainable")
		},
		"DisableSetter": func(t *testing.T, v *Variant) {
			assert(t, v.Disable == nil, "default value")
			disableVal := true
			v2 := v.SetDisable(&disableVal)
			require(t, v.Disable != nil, "expected value")
			assert(t, *v.Disable == disableVal, "expected value")
			assert(t, v2 == v, "chainable")
		},
		"PatchableSetter": func(t *testing.T, v *Variant) {
			assert(t, v.Patchable == nil, "default value")
			patchableVal := true
			v2 := v.SetPatchable(&patchableVal)
			require(t, v.Patchable != nil, "expected value")
			assert(t, *v.Patchable == patchableVal, "expected value")
			assert(t, v2 == v, "chainable")
		},
		"PatchOnlySetter": func(t *testing.T, v *Variant) {
			assert(t, v.PatchOnly == nil, "default value")
			patchOnlyVal := true
			v2 := v.SetPatchOnly(&patchOnlyVal)
			require(t, v.PatchOnly != nil, "expected value")
			assert(t, *v.PatchOnly == patchOnlyVal, "expected value")
			assert(t, v2 == v, "chainable")
		},
		"AllowForGitTagSetter": func(t *testing.T, v *Variant) {
			assert(t, v.AllowForGitTag == nil, "default value")
			allowForGitTagVal := true
			v2 := v.SetAllowForGitTag(&allowForGitTagVal)
			require(t, v.AllowForGitTag != nil, "expected value")
			assert(t, *v.AllowForGitTag == allowForGitTagVal, "expected value")
			assert(t, v2 == v, "chainable")
		},
		"GitTagOnlySetter": func(t *testing.T, v *Variant) {
			assert(t, v.GitTagOnly == nil, "default value")
			gitTagOnlyVal := true
			v2 := v.SetGitTagOnly(&gitTagOnlyVal)
			require(t, v.GitTagOnly != nil, "expected value")
			assert(t, *v.GitTagOnly == gitTagOnlyVal, "expected value")
			assert(t, v2 == v, "chainable")
		},
		"RunOnSetter": func(t *testing.T, v *Variant) {
			assert(t, len(v.DistroRunOn) == 0, "default value")
			v2 := v.RunOn("foo")

			require(t, len(v.DistroRunOn) == 1, "set")
			assert(t, v.DistroRunOn[0] == "foo", "expected value")
			assert(t, v2 == v, "chainable")
		},
		"RunOnMultipleTimes": func(t *testing.T, v *Variant) {
			v2 := v.RunOn("foo").RunOn("bar").RunOn("baz")

			require(t, len(v.DistroRunOn) == 1, "set")
			assert(t, v.DistroRunOn[0] == "baz", "expected value")
			assert(t, v2 == v, "chainable")
		},
		"TaskSpecSetterFirst": func(t *testing.T, v *Variant) {
			v2 := v.TaskSpec(TaskSpec{Name: "foo"})
			require(t, len(v.TaskSpecs) == 1, "set")
			assert(t, v.TaskSpecs[0].Name == "foo", "expected value")
			assert(t, v2 == v, "chainable")
		},
		"TaskSpecSetterSecond": func(t *testing.T, v *Variant) {
			v2 := v.TaskSpec(TaskSpec{Name: "first"}).TaskSpec(TaskSpec{Name: "foo"})
			require(t, len(v.TaskSpecs) == 2, "set")
			assert(t, v.TaskSpecs[0].Name == "first", "expected value")
			assert(t, v.TaskSpecs[1].Name == "foo", "expected value")
			assert(t, v2 == v, "chainable")
		},
		"SetExpansionSetter": func(t *testing.T, v *Variant) {
			v.Expansions = map[string]interface{}{}
			assert(t, v.Expansions != nil)
			v2 := v.SetExpansions(nil)
			assert(t, v2 == v, "chainable")
			assert(t, v.Expansions == nil)
		},
		"SetExpansionOverride": func(t *testing.T, v *Variant) {
			v.Expansions = map[string]interface{}{"b": "one"}
			assert(t, len(v.Expansions) == 1)
			v2 := v.SetExpansions(map[string]interface{}{"a": "two"})
			assert(t, v2 == v, "chainable")
			assert(t, len(v.Expansions) == 1)
			assert(t, v.Expansions["a"] == "two")
		},
		"AddExpansionFirst": func(t *testing.T, v *Variant) {
			assert(t, v.Expansions == nil)
			v2 := v.Expansion("one", 2)
			assert(t, v2 == v, "chainable")
			assert(t, len(v.Expansions) == 1)
			assert(t, v.Expansions["one"] == 2)
		},
		"AddExpansionSecond": func(t *testing.T, v *Variant) {
			v2 := v.Expansion("one", 2).Expansion("two", 42)
			assert(t, v2 == v, "chainable")
			assert(t, len(v.Expansions) == 2)
			assert(t, v.Expansions["two"] == 42)
		},
		"DisplayTaskNil": func(t *testing.T, v *Variant) {
			assert(t, len(v.DisplayTaskSpecs) == 0, "default value")
			v2 := v.DisplayTasks()
			assert(t, v2 == v, "chainable")
			assert(t, len(v.DisplayTaskSpecs) == 0, "length unchanged")
		},
		"DisplayTaskWithValues": func(t *testing.T, v *Variant) {
			v2 := v.DisplayTasks(DisplayTaskDefinition{Name: "one"},
				DisplayTaskDefinition{Name: "two"}).DisplayTasks(
				DisplayTaskDefinition{Name: "3"})

			assert(t, v2 == v, "chainable")
			assert(t, len(v.DisplayTaskSpecs) == 3, "length unchanged")
			assert(t, v.DisplayTaskSpecs[0].Name == "one")
			assert(t, v.DisplayTaskSpecs[1].Name == "two")
			assert(t, v.DisplayTaskSpecs[2].Name == "3")
		},
		"AddNoopTasks": func(t *testing.T, v *Variant) {
			assert(t, len(v.TaskSpecs) == 0, "default value")
			v2 := v.AddTasks("", "", "")
			assert(t, v2 == v, "chainable")
			assert(t, len(v.TaskSpecs) == 0, "no changes")
		},
		"AddSingleTask": func(t *testing.T, v *Variant) {
			assert(t, len(v.TaskSpecs) == 0, "default value")
			v2 := v.AddTasks("taskName")
			assert(t, v2 == v, "chainable")
			assert(t, len(v.TaskSpecs) == 1, "no changes")
			assert(t, v.TaskSpecs[0].Name == "taskName")
		},

		"AddSameTasks": func(t *testing.T, v *Variant) {
			assert(t, len(v.TaskSpecs) == 0, "default value")
			v2 := v.AddTasks("one", "one", "one")
			assert(t, v2 == v, "chainable")
			assert(t, len(v.TaskSpecs) == 3, "state impacted")
		},
		"AddDifferentTasks": func(t *testing.T, v *Variant) {
			assert(t, len(v.TaskSpecs) == 0, "default value")
			v2 := v.AddTasks("one", "two")
			assert(t, v2 == v, "chainable")
			assert(t, len(v.TaskSpecs) == 2, "state impacted")
		},
	}

	for name, test := range cases {
		v := &Variant{}
		t.Run(name, func(t *testing.T) {
			test(t, v)
		})
	}
}

func TestTaskSpecBuilders(t *testing.T) {
	cases := map[string]func(*testing.T, *TaskSpec){
		"NameSetter": func(t *testing.T, ts *TaskSpec) {
			assert(t, ts.Name == "", "default value")
			ts2 := ts.SetName("foo")
			assert(t, ts.Name == "foo", "expected value")
			assert(t, ts2 == ts, "chainable")
		},
		"StepbackSetter": func(t *testing.T, ts *TaskSpec) {
			assert(t, ts.Stepback == false, "default value")
			ts2 := ts.SetStepback(true)
			assert(t, ts.Stepback, "expected value")
			assert(t, ts2 == ts, "chainable")
		},
		"DistroSetter": func(t *testing.T, ts *TaskSpec) {
			assert(t, len(ts.Distro) == 0, "default value")
			ts2 := ts.SetDistros([]string{"distro"})
			require(t, len(ts.Distro) == 1, "state impacted")
			assert(t, ts.Distro[0] == "distro", "expected value")
			assert(t, ts2 == ts, "chainable")
		},
		"ActivateSetter": func(t *testing.T, ts *TaskSpec) {
			assert(t, ts.Activate == nil, "default value")
			ts2 := ts.SetActivate(&trueVal)
			require(t, ts.Activate != nil)
			assert(t, *ts.Activate, "expected value")
			assert(t, ts2 == ts, "chainable")
		},
	}

	for name, test := range cases {
		ts := &TaskSpec{}
		t.Run(name, func(t *testing.T) {
			test(t, ts)
		})
	}
}
