## Radius of gyration from VTK file

One file script for radius of gyration calculating using structured VTK file. 


## Compilation

Put it to appropriate place and run:

```
go build rg_calc.go
```

## Usage

The script expects a path to structured VTK file.

Simply run:

`
./rg_calc <path/to/structured/vtk/file>
`

The output should look similar to this:

```
==== GO [VTK->Rg^2] =====
v. 0.0.1
You provided next path:  anobject.vtk
Box size: [66, 66, 66]
Object mass: 600.00
Center of mass: [30, 30, 39]
Rg2: 169.57610
```

This is `anobject.vtk`:

![Anobject](./fig/anobject.png)


## License 

Feel free to use/contribute.



