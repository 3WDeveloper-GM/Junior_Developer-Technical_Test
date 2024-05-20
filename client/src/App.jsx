const Header = (props) => {
  return <h1>{props.course}</h1>;
};

const Contents = (props) => {
  console.log(props.parts);
  return (
    <>
      <p>
        {" "}
        {props.parts[0].name} {props.parts[0].exercises}{" "}
      </p>
      <p>
        {" "}
        {props.parts[1].name} {props.parts[1].exercises}{" "}
      </p>
      <p>
        {" "}
        {props.parts[2].name} {props.parts[2].exercises}{" "}
      </p>
    </>
  );
};

const Total = (props) => {
  const getTotal = (parts) => {
    let total = 0

    parts.forEach(element => {
      total += element.exercises
    });

    return total
  }  
  return (
    <>
   <p>
      Number of exercises: {getTotal(props.parts)}
   </p>
   </>
  );
};

const App = () => {
  const course = {
    name: "Half Stack Application development",
    part: [
      {
        name: "Fundamentals of React",
        exercises: 10,
      },
      {
        name: "Using props to pass data",
        exercises: 7,
      },
      {
        name: "State of a component",
        exercises: 14,
      },
    ],
  };
  return (
    <div>
      <Header course={course.name} />
      <Contents parts={course.part} />
      <Total parts={course.part} />
    </div>
  );
};

export default App;
