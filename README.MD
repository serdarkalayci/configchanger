To create golang image run command below in ConfigChangerGo folder:

```docker build -t configchangergo .```

To run docker golang container run command below in project root folder (where config folder resides)

```docker run --mount type=bind,source="$(pwd)"/config,target=/go/config -p 5200:5200 configchangergo```

To create .NET core image run command below in ConfigChangerNet folder:

```docker build -t configchangernet .```

To run docker .NET core container run command below in project root folder (where config folder resides)

```docker run --mount type=bind,source="$(pwd)"/config,target=/app/config -p 5100:5100 configchangernet```