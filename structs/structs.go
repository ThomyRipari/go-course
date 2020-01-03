package structs

import "fmt"

type Course struct {
	Name string
	Slug string
	Skills [] string
}

type Carrer struct {
	Name string
	Courses [] Course
}

func (c Course) newCourseMessage(name string) {
	fmt.Printf("Bienvenido al curso de %s, %s \n", c.Name, name)
}

func createNewCourse() {
	rubyCourse := new(Course)

	rubyCourse.Name, rubyCourse.Slug, rubyCourse.Skills = "Ruby Programming", "/ruby", [] string {
		"Programming",
		"Ruby",
		"Backend",
		"Apps Development",
	}

	rubyCourse.newCourseMessage("Thomas")

	defer fmt.Printf("Listo para aprender %s \n", rubyCourse.Name)
}

func ManipulateStructs() {
	goCourse := Course {Name: "Go Course", Slug: "/go" , Skills: [] string {"Go", "Backend", "Web"}}

	fmt.Println(goCourse)

	cCourse := new(Course)

	cCourse.Name, cCourse.Slug, cCourse.Skills = "C Programming Course", "/C", [] string {
		"C",
		"Programming",
		"App Development",
	}

	fmt.Println(cCourse)

	programmingCarrer := new(Carrer)

	Courses := [] Course {goCourse, *cCourse}

	programmingCarrer.Name, programmingCarrer.Courses = "Programming Carrer", Courses

	fmt.Println(programmingCarrer)

	createNewCourse()
}