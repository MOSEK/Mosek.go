
The `mosek` package
===================

There are two types defined in the `mosek` package:

`Env`
    Represents the MOSEK environment. Usually only one environment is necessary.
`Task`
    Represents a MOSEK task (problem data, parameters, solutions etc.).

Most functions in the package returns, at least, a response code. If
the code equals `mosek.RES_OK`, no error occurred, otherwise it is an
error code `mosek.RES_ERR_...`. You should always check the response
code before proceeding.


Creating `Env` and `Task`
~~~~~~~~~~~~~~~~~~~~~~~~~

..::

    func MakeEnv() (env Env, r int32)
    func (e *Env) DeleteEnv()

Create and delete an environment. An environment should always be
explicitly deleted, but must not be deleted before all tasks createn
through it.

..::

    func (env *Env) MakeTask() (task Task, r int32) 
    func (t *Task) DeleteTask()
 
Create and delete a task. A task should always be explicitly deleted
when it is not used anymore. The GO garbage collector will collect the
structure but not the underlying MOSEK allocations.


AnalyzeNames
~~~~~~~~~~~~

..::

    func (*Task) AnalyzeNames ( whichstream int32, nametype int32 )
`nametype int32`
    The type of names e.g. valid in MPS or LP files.

Analyze the names and issue an error for the first invalid name.


AnalyzeProblem
~~~~~~~~~~~~~~

..::

    func (*Task) AnalyzeProblem ( whichstream int32 )

Analyze the data of a task.


AnalyzeSolution
~~~~~~~~~~~~~~~

..::

    func (*Task) AnalyzeSolution ( whichstream int32, whichsol int32 )

Print information related to the quality of the solution.


AppendBarvars
~~~~~~~~~~~~~

..::

    func (*Task) AppendBarvars ( dim []int32 )
`dim []int32`
    Dimension of symmetric matrix variables to be added.

Appends a semidefinite  variable of dimension dim to the problem.


AppendCone
~~~~~~~~~~

..::

    func (*Task) AppendCone
        ( ct int32,
          conepar float64,
          submem []int32 )


Appends a new cone constraint to the problem.


AppendConeSeq
~~~~~~~~~~~~~

..::

    func (*Task) AppendConeSeq
        ( ct int32,
          conepar float64,
          nummem int32,
          j int32 )

`nummem int32`
    Dimension of the conic constraint.
`j int32`
    Index of the first variable in the conic constraint.

Appends a new conic constraint to the problem.


AppendConesSeq
~~~~~~~~~~~~~~

..::

    func (*Task) AppendConesSeq
        ( ct []int32,
          conepar []float64,
          nummem []int32,
          j int32 )

`j int32`
    Index of the first variable in the first cone to be appended.

Appends a multiple conic constraints to the problem.


AppendCons
~~~~~~~~~~

..::

    func (*Task) AppendCons ( num int32 )
`num int32`
    Number of constraints which should be appended.

Appends a number of constraints to the optimization task.


AppendSparseSymMat
~~~~~~~~~~~~~~~~~~

..::

    func (*Task) AppendSparseSymMat
        ( dim int32,
          subi []int32,
          subj []int32,
          valij []float64 )
        ( idx int64 )

`dim int32`
    Dimension of the symmetric matrix that is appended.
`subi []int32`
    Row subscript in the triplets.
`subj []int32`
    Column subscripts in the triplets.
`valij []float64`
    Values of each triplet.

Appends a general sparse symmetric matrix to the vector E of symmetric matrixes.


AppendStat
~~~~~~~~~~

..::

    func (*Task) AppendStat (  )

Appends a record the statistics file.


AppendVars
~~~~~~~~~~

..::

    func (*Task) AppendVars ( num int32 )
`num int32`
    Number of variables which should be appended.

Appends a number of variables to the optimization task.


Axpy
~~~~

..::

    func (*Env) Axpy
        ( n int32,
          alpha float64,
          x []float64,
          y []float64 )
        ( y []float64 )

`n int32`
    Length of the vectors.
`alpha float64`
    The scalar that multiplies x.
`x []float64`
    The :math:`x` vector.
`y []float64`
    The :math:`y` vector.

Adds alpha times x to y.


BasisCond
~~~~~~~~~

..::

    func (*Task) BasisCond (  ) ( nrmbasis float64, nrminvbasis float64 )


Computes conditioning information for the basis matrix.


CheckConvexity
~~~~~~~~~~~~~~

..::

    func (*Task) CheckConvexity (  )

Checks if a quadratic optimization problem is convex.


CheckInLicense
~~~~~~~~~~~~~~

..::

    func (*Env) CheckInLicense ( feature int32 )
`feature int32`
    Feature to check in to the license system.

Check in a license feature from the license server ahead of time.


CheckMem
~~~~~~~~

..::

    func (*Task) CheckMem ( file string, line int32 )
`file string`
    File from which the function is called.
`line int32`
    Line in the file from which the function is called.

Checks the memory allocated by the task.


CheckoutLicense
~~~~~~~~~~~~~~~

..::

    func (*Env) CheckoutLicense ( feature int32 )
`feature int32`
    Feature to check out from the license system.

Check out a license feature from the license server ahead of time.


ChgBound
~~~~~~~~

..::

    func (*Task) ChgBound
        ( accmode int32,
          i int32,
          lower int32,
          finite int32,
          value float64 )

`i int32`
    Index of the constraint or variable for which the bounds should be changed.
`lower int32`
    If non-zero, then the lower bound is changed, otherwise the upper bound is changed.
`finite int32`
    If non-zero, then the given value is assumed to be finite.
`value float64`
    New value for the bound.

Changes the bounds for one constraint or variable.


ChgConBound
~~~~~~~~~~~

..::

    func (*Task) ChgConBound
        ( i int32,
          lower int32,
          finite int32,
          value float64 )

`i int32`
    Index of the constraint for which the bounds should be changed.
`lower int32`
    If non-zero, then the lower bound is changed, otherwise the upper bound is changed.
`finite int32`
    If non-zero, then the given value is assumed to be finite.
`value float64`
    New value for the bound.

Changes the bounds for one constraint.


ChgVarBound
~~~~~~~~~~~

..::

    func (*Task) ChgVarBound
        ( j int32,
          lower int32,
          finite int32,
          value float64 )

`j int32`
    Index of the variable for which the bounds should be changed.
`lower int32`
    If non-zero, then the lower bound is changed, otherwise the upper bound is changed.
`finite int32`
    If non-zero, then the given value is assumed to be finite.
`value float64`
    New value for the bound.

Changes the bounds for one variable.


CommitChanges
~~~~~~~~~~~~~

..::

    func (*Task) CommitChanges (  )

Commits all cached problem changes.


DeleteSolution
~~~~~~~~~~~~~~

..::

    func (*Task) DeleteSolution ( whichsol int32 )

Undefine a solution and frees the memory it uses.


Dot
~~~

..::

    func (*Env) Dot
        ( n int32,
          x []float64,
          y []float64 )
        ( xty float64 )

`n int32`
    Length of the vectors.
`x []float64`
    The x vector.
`y []float64`
    The y vector.

Computes the inner product of two vectors.


DualSensitivity
~~~~~~~~~~~~~~~

..::

    func (*Task) DualSensitivity
        ( subj []int32,
          leftpricej []float64,
          rightpricej []float64,
          leftrangej []float64,
          rightrangej []float64 )
        ( leftpricej []float64,
          rightpricej []float64,
          leftrangej []float64,
          rightrangej []float64 )

`subj []int32`
    Index of objective coefficients to analyze.
`leftpricej []float64`
    Left shadow prices for requested coefficients.
`rightpricej []float64`
    Right shadow prices for requested coefficients.
`leftrangej []float64`
    Left range for requested coefficients.
`rightrangej []float64`
    Right range for requested coefficients.

Performs sensitivity analysis on objective coefficients.


EchoIntro
~~~~~~~~~

..::

    func (*Env) EchoIntro ( longver int32 )
`longver int32`
    If non-zero, then the intro is slightly longer.

Prints an intro to message stream.


Gemm
~~~~

..::

    func (*Env) Gemm
        ( transa int32,
          transb int32,
          m int32,
          n int32,
          k int32,
          alpha float64,
          a []float64,
          b []float64,
          beta float64,
          c []float64 )
        ( c []float64 )

`transa int32`
    Indicates whether the matrix A must be transposed.
`transb int32`
    Indicates whether the matrix B must be transposed.
`m int32`
    Indicates the number of rows of matrices A and C.
`n int32`
    Indicates the number of columns of matrices B and C.
`k int32`
    Specifies the number of columns of the matrix A and the number of rows of the matrix B.
`alpha float64`
    A scalar value multipling the result of the matrix multiplication.
`a []float64`
    The pointer to the array storing matrix A in a column-major format.
`b []float64`
    Indicates the number of rows of matrix B and columns of matrix A.
`beta float64`
    A scalar value that multiplies C.
`c []float64`
    The pointer to the array storing matrix C in a column-major format.

Performs a dense matrix multiplication.


Gemv
~~~~

..::

    func (*Env) Gemv
        ( transa int32,
          m int32,
          n int32,
          alpha float64,
          a []float64,
          x []float64,
          beta float64,
          y []float64 )
        ( y []float64 )

`transa int32`
    Indicates whether the matrix A must be transposed.
`m int32`
    Specifies the number of rows of the matrix A.
`n int32`
    Specifies the number of columns of the matrix A.
`alpha float64`
    A scalar value multipling the matrix A.
`a []float64`
    A pointer to the array storing matrix A in a column-major format.
`x []float64`
    A pointer to the array storing the vector x.
`beta float64`
    A scalar value multipling thevector y.
`y []float64`
    A pointer to the array storing the vector y.

Computes dense matrix times a dense vector product.


GetACol
~~~~~~~

..::

    func (*Task) GetACol
        ( j int32,
          subj []int32,
          valj []float64 )
        ( nzj int32,
          subj []int32,
          valj []float64 )

`j int32`
    Index of the column.
`subj []int32`
    Index of the non-zeros in the row obtained.
`valj []float64`
    Numerical values of the column obtained.

Obtains one column of the linear constraint matrix.


GetAColNumNz
~~~~~~~~~~~~

..::

    func (*Task) GetAColNumNz ( i int32 ) ( nzj int32 )

`i int32`
    Index of the column.

Obtains the number of non-zero elements in one column of the linear constraint matrix


GetAColSliceTrip
~~~~~~~~~~~~~~~~

..::

    func (*Task) GetAColSliceTrip
        ( first int32,
          last int32,
          subi []int32,
          subj []int32,
          val []float64 )
        ( subi []int32,
          subj []int32,
          val []float64 )

`first int32`
    Index of the first column in the sequence.
`last int32`
    Index of the last column in the sequence plus one.
`subi []int32`
    Constraint subscripts.
`subj []int32`
    Column subscripts.
`val []float64`
    Values.

Obtains a sequence of columns from the coefficient matrix in triplet format.


GetAPieceNumNz
~~~~~~~~~~~~~~

..::

    func (*Task) GetAPieceNumNz
        ( firsti int32,
          lasti int32,
          firstj int32,
          lastj int32 )
        ( numnz int32 )

`firsti int32`
    Index of the first row in the rectangular piece.
`lasti int32`
    Index of the last row plus one in the rectangular piece.
`firstj int32`
    Index of the first column in the rectangular piece.
`lastj int32`
    Index of the last column plus one in the rectangular piece.

Obtains the number non-zeros in a rectangular piece of the linear constraint matrix.


GetARow
~~~~~~~

..::

    func (*Task) GetARow
        ( i int32,
          subi []int32,
          vali []float64 )
        ( nzi int32,
          subi []int32,
          vali []float64 )

`i int32`
    Index of the row or column.
`subi []int32`
    Index of the non-zeros in the row obtained.
`vali []float64`
    Numerical values of the row obtained.

Obtains one row of the linear constraint matrix.


GetARowNumNz
~~~~~~~~~~~~

..::

    func (*Task) GetARowNumNz ( i int32 ) ( nzi int32 )

`i int32`
    Index of the row or column.

Obtains the number of non-zero elements in one row of the linear constraint matrix


GetARowSliceTrip
~~~~~~~~~~~~~~~~

..::

    func (*Task) GetARowSliceTrip
        ( first int32,
          last int32,
          subi []int32,
          subj []int32,
          val []float64 )
        ( subi []int32,
          subj []int32,
          val []float64 )

`first int32`
    Index of the first row or column in the sequence.
`last int32`
    Index of the last row or column in the sequence plus one.
`subi []int32`
    Constraint subscripts.
`subj []int32`
    Column subscripts.
`val []float64`
    Values.

Obtains a sequence of rows from the coefficient matrix in triplet format.


GetASlice
~~~~~~~~~

..::

    func (*Task) GetASlice
        ( accmode int32,
          first int32,
          last int32,
          ptrb []int64,
          ptre []int64,
          sub []int32,
          val []float64 )
        ( ptrb []int64,
          ptre []int64,
          sub []int32,
          val []float64 )

`accmode int32`
    Defines whether a column slice or a row slice is requested.
`first int32`
    Index of the first row or column in the sequence.
`last int32`
    Index of the last row or column in the sequence plus one.
`ptrb []int64`
    Row or column start pointers.
`ptre []int64`
    Row or column end pointers.
`sub []int32`
    Contains the row or column subscripts.
`val []float64`
    Contains the coefficient values.

Obtains a sequence of rows or columns from the coefficient matrix.


GetASliceNumNz
~~~~~~~~~~~~~~

..::

    func (*Task) GetASliceNumNz
        ( accmode int32,
          first int32,
          last int32 )
        ( numnz int64 )

`accmode int32`
    Defines whether non-zeros are counted in a column slice or a row slice.
`first int32`
    Index of the first row or column in the sequence.
`last int32`
    Index of the last row or column plus one in the sequence.

Obtains the number of non-zeros in a slice of rows or columns of the coefficient matrix.


GetAij
~~~~~~

..::

    func (*Task) GetAij ( i int32, j int32 ) ( aij float64 )

`i int32`
    Row index of the coefficient to be returned.
`j int32`
    Column index of the coefficient to be returned.

Obtains a single coefficient in linear constraint matrix.


GetBaraBlockTriplet
~~~~~~~~~~~~~~~~~~~

..::

    func (*Task) GetBaraBlockTriplet
        ( subi []int32,
          subj []int32,
          subk []int32,
          subl []int32,
          valijkl []float64 )
        ( num int64,
          subi []int32,
          subj []int32,
          subk []int32,
          subl []int32,
          valijkl []float64 )

`subi []int32`
    Constraint index.
`subj []int32`
    Symmetric matrix variable index.
`subk []int32`
    Block row index.
`subl []int32`
    Block column index.
`valijkl []float64`
    A list indexes of the elements from symmetric matrix storage that appears in the weighted sum.

Obtains barA in block triplet form.


GetBaraIdx
~~~~~~~~~~

..::

    func (*Task) GetBaraIdx
        ( idx int64,
          sub []int64,
          weights []float64 )
        ( i int32,
          j int32,
          num int64,
          sub []int64,
          weights []float64 )

`idx int64`
    Position of the element in the vectorized form.
`sub []int64`
    A list indexes   of the elements from symmetric matrix storage that appears in the weighted sum.
`weights []float64`
    The weights associated with each term in the weighted sum.

Obtains information about an element barA.


GetBaraIdxIJ
~~~~~~~~~~~~

..::

    func (*Task) GetBaraIdxIJ ( idx int64 ) ( i int32, j int32 )

`idx int64`
    Position of the element in the vectorized form.

Obtains information about an element barA.


GetBaraIdxInfo
~~~~~~~~~~~~~~

..::

    func (*Task) GetBaraIdxInfo ( idx int64 ) ( num int64 )

`idx int64`
    The internal position of the element that should be obtained information for.

Obtains the number terms in the weighted sum that forms a particular element in barA.


GetBaraSparsity
~~~~~~~~~~~~~~~

..::

    func (*Task) GetBaraSparsity ( idxij []int64 ) ( numnz int64, idxij []int64 )

`idxij []int64`
    Position of each nonzero element in the vector representation of barA.

Obtains the sparsity pattern of the barA matrix.


GetBarcBlockTriplet
~~~~~~~~~~~~~~~~~~~

..::

    func (*Task) GetBarcBlockTriplet
        ( subj []int32,
          subk []int32,
          subl []int32,
          valijkl []float64 )
        ( num int64,
          subj []int32,
          subk []int32,
          subl []int32,
          valijkl []float64 )

`subj []int32`
    Symmetric matrix variable index.
`subk []int32`
    Block row index.
`subl []int32`
    Block column index.
`valijkl []float64`
    A list indexes of the elements from symmetric matrix storage that appears in the weighted sum.

Obtains barc in block triplet form.


GetBarcIdx
~~~~~~~~~~

..::

    func (*Task) GetBarcIdx
        ( idx int64,
          sub []int64,
          weights []float64 )
        ( j int32,
          num int64,
          sub []int64,
          weights []float64 )

`idx int64`
    Index of the element that should be obtained information about.
`sub []int64`
    Elements appearing the weighted sum.
`weights []float64`
    Weights of terms in the weighted sum.

Obtains information about an element in barc.


GetBarcIdxInfo
~~~~~~~~~~~~~~

..::

    func (*Task) GetBarcIdxInfo ( idx int64 ) ( num int64 )

`idx int64`
    Index of element that should be obtained information about. The value is an index of a symmetric sparse variable.

Obtains information about an element in barc.


GetBarcIdxJ
~~~~~~~~~~~

..::

    func (*Task) GetBarcIdxJ ( idx int64 ) ( j int32 )

`idx int64`
    Index of the element that should be obtained information about.

Obtains the row index of an element in barc.


GetBarcSparsity
~~~~~~~~~~~~~~~

..::

    func (*Task) GetBarcSparsity ( idxj []int64 ) ( numnz int64, idxj []int64 )

`idxj []int64`
    Internal positions of the nonzeros elements in barc.

Get the positions of the nonzero elements in barc.


GetBarsJ
~~~~~~~~

..::

    func (*Task) GetBarsJ
        ( whichsol int32,
          j int32,
          barsj []float64 )
        ( barsj []float64 )

`j int32`
    Index of the semidefinite variable.
`barsj []float64`
    Value of the j'th variable of barx.

Obtains the dual solution for a semidefinite variable.


GetBarvarName
~~~~~~~~~~~~~

..::

    func (*Task) GetBarvarName ( i int32 ) ( name string )

`i int32`
    Index.

Obtains a name of a semidefinite variable.


GetBarvarNameIndex
~~~~~~~~~~~~~~~~~~

..::

    func (*Task) GetBarvarNameIndex ( somename string ) ( asgn int32, index int32 )

`somename string`
    The requested name is copied to this buffer.

Obtains the index of name of semidefinite variable.


GetBarvarNameLen
~~~~~~~~~~~~~~~~

..::

    func (*Task) GetBarvarNameLen ( i int32 ) ( len int32 )

`i int32`
    Index.

Obtains the length of a name of a semidefinite variable.


GetBarxJ
~~~~~~~~

..::

    func (*Task) GetBarxJ
        ( whichsol int32,
          j int32,
          barxj []float64 )
        ( barxj []float64 )

`j int32`
    Index of the semidefinite variable.
`barxj []float64`
    Value of the j'th variable of barx.

Obtains the primal solution for a semidefinite variable.


GetBound
~~~~~~~~

..::

    func (*Task) GetBound
        ( accmode int32,
          i int32 )
        ( bk int32,
          bl float64,
          bu float64 )

`i int32`
    Index of the constraint or variable for which the bound information should be obtained.

Obtains bound information for one constraint or variable.


GetBoundSlice
~~~~~~~~~~~~~

..::

    func (*Task) GetBoundSlice
        ( accmode int32,
          first int32,
          last int32,
          bk []int32,
          bl []float64,
          bu []float64 )
        ( bk []int32,
          bl []float64,
          bu []float64 )


Obtains bounds information for a sequence of variables or constraints.


GetC
~~~~

..::

    func (*Task) GetC ( c []float64 ) ( c []float64 )


Obtains all objective coefficients.


GetCJ
~~~~~

..::

    func (*Task) GetCJ ( j int32 ) ( cj float64 )

`j int32`
    Index of the variable for which c coefficient should be obtained.

Obtains one coefficient of c.


GetCSlice
~~~~~~~~~

..::

    func (*Task) GetCSlice
        ( first int32,
          last int32,
          c []float64 )
        ( c []float64 )


Obtains a sequence of coefficients from the objective.


GetCfix
~~~~~~~

..::

    func (*Task) GetCfix (  ) ( cfix float64 )


Obtains the fixed term in the objective.


GetCodeDesc
~~~~~~~~~~~

..::

    func GetCodeDesc
        ( code int32 )
        ( symname string,
          str string,
          res int32 )

`code int32`
    A valid response code.

Obtains a short description of a response code.


GetConBound
~~~~~~~~~~~

..::

    func (*Task) GetConBound
        ( i int32 )
        ( bk int32,
          bl float64,
          bu float64 )

`i int32`
    Index of the constraint for which the bound information should be obtained.

Obtains bound information for one constraint.


GetConBoundSlice
~~~~~~~~~~~~~~~~

..::

    func (*Task) GetConBoundSlice
        ( first int32,
          last int32,
          bk []int32,
          bl []float64,
          bu []float64 )
        ( bk []int32,
          bl []float64,
          bu []float64 )


Obtains bounds information for a slice of the constraints.


GetConName
~~~~~~~~~~

..::

    func (*Task) GetConName ( i int32 ) ( name string )

`i int32`
    Index.

Obtains a name of a constraint.


GetConNameIndex
~~~~~~~~~~~~~~~

..::

    func (*Task) GetConNameIndex ( somename string ) ( asgn int32, index int32 )

`somename string`
    The name which should be checked.

Checks whether the name somename has been assigned  to any constraint.


GetConNameLen
~~~~~~~~~~~~~

..::

    func (*Task) GetConNameLen ( i int32 ) ( len int32 )

`i int32`
    Index.

Obtains the length of a name of a constraint variable.


GetCone
~~~~~~~

..::

    func (*Task) GetCone
        ( k int32,
          submem []int32 )
        ( ct int32,
          conepar float64,
          nummem int32,
          submem []int32 )

`k int32`
    Index of the cone constraint.

Obtains a conic constraint.


GetConeInfo
~~~~~~~~~~~

..::

    func (*Task) GetConeInfo
        ( k int32 )
        ( ct int32,
          conepar float64,
          nummem int32 )

`k int32`
    Index of the conic constraint.

Obtains information about a conic constraint.


GetConeName
~~~~~~~~~~~

..::

    func (*Task) GetConeName ( i int32 ) ( name string )

`i int32`
    Index.

Obtains a name of a cone.


GetConeNameIndex
~~~~~~~~~~~~~~~~

..::

    func (*Task) GetConeNameIndex ( somename string ) ( asgn int32, index int32 )

`somename string`
    The name which should be checked.

Checks whether the name somename has been assigned  to any cone.


GetConeNameLen
~~~~~~~~~~~~~~

..::

    func (*Task) GetConeNameLen ( i int32 ) ( len int32 )

`i int32`
    Index.

Obtains the length of a name of a cone.


GetDimBarvarJ
~~~~~~~~~~~~~

..::

    func (*Task) GetDimBarvarJ ( j int32 ) ( dimbarvarj int32 )

`j int32`
    Index of the semidefinite variable whose dimension is requested.

Obtains the dimension of a symmetric matrix variable.


GetDouInf
~~~~~~~~~

..::

    func (*Task) GetDouInf ( whichdinf int32 ) ( dvalue float64 )


Obtains a double information item.


GetDouParam
~~~~~~~~~~~

..::

    func (*Task) GetDouParam ( param int32 ) ( parvalue float64 )


Obtains a double parameter.


GetDualObj
~~~~~~~~~~

..::

    func (*Task) GetDualObj ( whichsol int32 ) ( dualobj float64 )


Computes the dual objective value associated with the solution.


GetDviolBarvar
~~~~~~~~~~~~~~

..::

    func (*Task) GetDviolBarvar
        ( whichsol int32,
          sub []int32,
          viol []float64 )
        ( viol []float64 )

`sub []int32`
    An array of indexes of barx variables.
`viol []float64`
    List of violations corresponding to sub.

Computes the violation of dual solution for a set of barx variables.


GetDviolCon
~~~~~~~~~~~

..::

    func (*Task) GetDviolCon
        ( whichsol int32,
          sub []int32,
          viol []float64 )
        ( viol []float64 )

`sub []int32`
    An array of indexes of constraints.
`viol []float64`
    List of violations corresponding to sub.

Computes the violation of a dual solution associated with a set of constraints.


GetDviolCones
~~~~~~~~~~~~~

..::

    func (*Task) GetDviolCones
        ( whichsol int32,
          sub []int32,
          viol []float64 )
        ( viol []float64 )

`sub []int32`
    An array of indexes of barx variables.
`viol []float64`
    List of violations corresponding to sub.

Computes the violation of a solution for set of dual conic constraints.


GetDviolVar
~~~~~~~~~~~

..::

    func (*Task) GetDviolVar
        ( whichsol int32,
          sub []int32,
          viol []float64 )
        ( viol []float64 )

`sub []int32`
    An array of indexes of x variables.
`viol []float64`
    List of violations corresponding to sub.

Computes the violation of a dual solution associated with a set of x variables.


GetInfIndex
~~~~~~~~~~~

..::

    func (*Task) GetInfIndex ( inftype int32, infname string ) ( infindex int32 )


Obtains the index of a named information item.


GetInfMax
~~~~~~~~~

..::

    func (*Task) GetInfMax ( inftype int32, infmax []int32 ) ( infmax []int32 )


Obtains the maximum index of an information of a given type inftype plus 1.


GetInfName
~~~~~~~~~~

..::

    func (*Task) GetInfName ( inftype int32, whichinf int32 ) ( infname string )


Obtains the name of an information item.


GetInfeasibleSubProblem
~~~~~~~~~~~~~~~~~~~~~~~

..::

    func (*Task) GetInfeasibleSubProblem ( whichsol int32 ) ( inftask Task )

`whichsol int32`
    Which solution to use when determining the infeasible subproblem.

Obtains an infeasible sub problem.


GetIntInf
~~~~~~~~~

..::

    func (*Task) GetIntInf ( whichiinf int32 ) ( ivalue int32 )


Obtains an integer information item.


GetIntParam
~~~~~~~~~~~

..::

    func (*Task) GetIntParam ( param int32 ) ( parvalue int32 )


Obtains an integer parameter.


GetLenBarvarJ
~~~~~~~~~~~~~

..::

    func (*Task) GetLenBarvarJ ( j int32 ) ( lenbarvarj int64 )

`j int32`
    Index of the semidefinite variable whose length if requested.

Obtains the length if the j'th semidefinite variables.


GetLintInf
~~~~~~~~~~

..::

    func (*Task) GetLintInf ( whichliinf int32 ) ( ivalue int64 )


Obtains an integer information item.


GetMaxNumANz
~~~~~~~~~~~~

..::

    func (*Task) GetMaxNumANz (  ) ( maxnumanz int64 )


Obtains number of preallocated non-zeros in the linear constraint matrix.


GetMaxNumBarvar
~~~~~~~~~~~~~~~

..::

    func (*Task) GetMaxNumBarvar (  ) ( maxnumbarvar int32 )


Obtains the number of semidefinite variables.


GetMaxNumCon
~~~~~~~~~~~~

..::

    func (*Task) GetMaxNumCon (  ) ( maxnumcon int32 )


Obtains the number of preallocated constraints in the optimization task.


GetMaxNumCone
~~~~~~~~~~~~~

..::

    func (*Task) GetMaxNumCone (  ) ( maxnumcone int32 )


Obtains the number of preallocated cones in the optimization task.


GetMaxNumQNz
~~~~~~~~~~~~

..::

    func (*Task) GetMaxNumQNz (  ) ( maxnumqnz int64 )


Obtains the number of preallocated non-zeros for all quadratic terms in objective and constraints.


GetMaxNumVar
~~~~~~~~~~~~

..::

    func (*Task) GetMaxNumVar (  ) ( maxnumvar int32 )


Obtains the maximum number variables allowed.


GetMemUsage
~~~~~~~~~~~

..::

    func (*Task) GetMemUsage (  ) ( meminuse int64, maxmemuse int64 )


Obtains information about the amount of memory used by a task.


GetNumANz
~~~~~~~~~

..::

    func (*Task) GetNumANz (  ) ( numanz int32 )


Obtains the number of non-zeros in the coefficient matrix.


GetNumANz64
~~~~~~~~~~~

..::

    func (*Task) GetNumANz64 (  ) ( numanz int64 )


Obtains the number of non-zeros in the coefficient matrix.


GetNumBaraBlockTriplets
~~~~~~~~~~~~~~~~~~~~~~~

..::

    func (*Task) GetNumBaraBlockTriplets (  ) ( num int64 )


Obtains an upper bound on the number of scalar elements in the block triplet form of bara.


GetNumBaraNz
~~~~~~~~~~~~

..::

    func (*Task) GetNumBaraNz (  ) ( nz int64 )


Get the number of nonzero elements in barA.


GetNumBarcBlockTriplets
~~~~~~~~~~~~~~~~~~~~~~~

..::

    func (*Task) GetNumBarcBlockTriplets (  ) ( num int64 )


Obtains an upper bound on the number of elements in the block triplet form of barc.


GetNumBarcNz
~~~~~~~~~~~~

..::

    func (*Task) GetNumBarcNz (  ) ( nz int64 )


Obtains the number of nonzero elements in barc.


GetNumBarvar
~~~~~~~~~~~~

..::

    func (*Task) GetNumBarvar (  ) ( numbarvar int32 )


Obtains the number of semidefinite variables.


GetNumCon
~~~~~~~~~

..::

    func (*Task) GetNumCon (  ) ( numcon int32 )


Obtains the number of constraints.


GetNumCone
~~~~~~~~~~

..::

    func (*Task) GetNumCone (  ) ( numcone int32 )


Obtains the number of cones.


GetNumConeMem
~~~~~~~~~~~~~

..::

    func (*Task) GetNumConeMem ( k int32 ) ( nummem int32 )

`k int32`
    Index of the cone.

Obtains the number of members in a cone.


GetNumIntVar
~~~~~~~~~~~~

..::

    func (*Task) GetNumIntVar (  ) ( numintvar int32 )


Obtains the number of integer-constrained variables.


GetNumParam
~~~~~~~~~~~

..::

    func (*Task) GetNumParam ( partype int32 ) ( numparam int32 )


Obtains the number of parameters of a given type.


GetNumQConKNz
~~~~~~~~~~~~~

..::

    func (*Task) GetNumQConKNz ( k int32 ) ( numqcnz int64 )

`k int32`
    Index of the constraint for which the number quadratic terms should be obtained.

Obtains the number of non-zero quadratic terms in a constraint.


GetNumQObjNz
~~~~~~~~~~~~

..::

    func (*Task) GetNumQObjNz (  ) ( numqonz int64 )


Obtains the number of non-zero quadratic terms in the objective.


GetNumSymMat
~~~~~~~~~~~~

..::

    func (*Task) GetNumSymMat (  ) ( num int64 )


Get the number of symmetric matrixes stored.


GetNumVar
~~~~~~~~~

..::

    func (*Task) GetNumVar (  ) ( numvar int32 )


Obtains the number of variables.


GetObjName
~~~~~~~~~~

..::

    func (*Task) GetObjName (  ) ( objname string )


Obtains the name assigned to the objective function.


GetObjNameLen
~~~~~~~~~~~~~

..::

    func (*Task) GetObjNameLen (  ) ( len int32 )


Obtains the length of the name assigned to the objective function.


GetObjSense
~~~~~~~~~~~

..::

    func (*Task) GetObjSense (  ) ( sense int32 )


Gets the objective sense.


GetParamMax
~~~~~~~~~~~

..::

    func (*Task) GetParamMax ( partype int32 ) ( parammax int32 )


Obtains the maximum index of a parameter of a given type plus 1.


GetParamName
~~~~~~~~~~~~

..::

    func (*Task) GetParamName ( partype int32, param int32 ) ( parname string )


Obtains the name of a parameter.


GetPrimalObj
~~~~~~~~~~~~

..::

    func (*Task) GetPrimalObj ( whichsol int32 ) ( primalobj float64 )


Computes the primal objective value for the desired solution.


GetProSta
~~~~~~~~~

..::

    func (*Task) GetProSta ( whichsol int32 ) ( prosta int32 )


Obtains the problem status.


GetProbType
~~~~~~~~~~~

..::

    func (*Task) GetProbType (  ) ( probtype int32 )


Obtains the problem type.


GetPviolBarvar
~~~~~~~~~~~~~~

..::

    func (*Task) GetPviolBarvar
        ( whichsol int32,
          sub []int32,
          viol []float64 )
        ( viol []float64 )

`sub []int32`
    An array of indexes of barx variables.
`viol []float64`
    List of violations corresponding to sub.

Computes the violation of a primal solution for a list of barx variables.


GetPviolCon
~~~~~~~~~~~

..::

    func (*Task) GetPviolCon
        ( whichsol int32,
          sub []int32,
          viol []float64 )
        ( viol []float64 )

`sub []int32`
    An array of indexes of constraints.
`viol []float64`
    List of violations corresponding to sub.

Computes the violation of a primal solution for a list of xc variables.


GetPviolCones
~~~~~~~~~~~~~

..::

    func (*Task) GetPviolCones
        ( whichsol int32,
          sub []int32,
          viol []float64 )
        ( viol []float64 )

`sub []int32`
    An array of indexes of barx variables.
`viol []float64`
    List of violations corresponding to sub.

Computes the violation of a solution for set of conic constraints.


GetPviolVar
~~~~~~~~~~~

..::

    func (*Task) GetPviolVar
        ( whichsol int32,
          sub []int32,
          viol []float64 )
        ( viol []float64 )

`sub []int32`
    An array of indexes of x variables.
`viol []float64`
    List of violations corresponding to sub.

Computes the violation of a primal solution for a list of x variables.


GetQConK
~~~~~~~~

..::

    func (*Task) GetQConK
        ( k int32,
          qcsubi []int32,
          qcsubj []int32,
          qcval []float64 )
        ( numqcnz int64,
          qcsubi []int32,
          qcsubj []int32,
          qcval []float64 )

`k int32`
    Which constraint.

Obtains all the quadratic terms in a constraint.


GetQObj
~~~~~~~

..::

    func (*Task) GetQObj
        ( qosubi []int32,
          qosubj []int32,
          qoval []float64 )
        ( numqonz int64,
          qosubi []int32,
          qosubj []int32,
          qoval []float64 )


Obtains all the quadratic terms in the objective.


GetQObjIJ
~~~~~~~~~

..::

    func (*Task) GetQObjIJ ( i int32, j int32 ) ( qoij float64 )

`i int32`
    Row index of the coefficient.
`j int32`
    Column index of coefficient.

Obtains one coefficient from the quadratic term of the objective


GetReducedCosts
~~~~~~~~~~~~~~~

..::

    func (*Task) GetReducedCosts
        ( whichsol int32,
          first int32,
          last int32,
          redcosts []float64 )
        ( redcosts []float64 )

`first int32`
    See the documentation for a full description.
`last int32`
    See the documentation for a full description.
`redcosts []float64`
    Returns the requested reduced costs. See documentation for a full description.

Obtains the difference of (slx-sux) for a sequence of variables.


GetSkc
~~~~~~

..::

    func (*Task) GetSkc ( whichsol int32, skc []int32 ) ( skc []int32 )


Obtains the status keys for the constraints.


GetSkcSlice
~~~~~~~~~~~

..::

    func (*Task) GetSkcSlice
        ( whichsol int32,
          first int32,
          last int32,
          skc []int32 )
        ( skc []int32 )


Obtains the status keys for the constraints.


GetSkx
~~~~~~

..::

    func (*Task) GetSkx ( whichsol int32, skx []int32 ) ( skx []int32 )


Obtains the status keys for the scalar variables.


GetSkxSlice
~~~~~~~~~~~

..::

    func (*Task) GetSkxSlice
        ( whichsol int32,
          first int32,
          last int32,
          skx []int32 )
        ( skx []int32 )


Obtains the status keys for the variables.


GetSlc
~~~~~~

..::

    func (*Task) GetSlc ( whichsol int32, slc []float64 ) ( slc []float64 )

`slc []float64`
    The slc vector.

Obtains the slc vector for a solution.


GetSlcSlice
~~~~~~~~~~~

..::

    func (*Task) GetSlcSlice
        ( whichsol int32,
          first int32,
          last int32,
          slc []float64 )
        ( slc []float64 )


Obtains a slice of the slc vector for a solution.


GetSlx
~~~~~~

..::

    func (*Task) GetSlx ( whichsol int32, slx []float64 ) ( slx []float64 )

`slx []float64`
    The slx vector.

Obtains the slx vector for a solution.


GetSlxSlice
~~~~~~~~~~~

..::

    func (*Task) GetSlxSlice
        ( whichsol int32,
          first int32,
          last int32,
          slx []float64 )
        ( slx []float64 )


Obtains a slice of the slx vector for a solution.


GetSnx
~~~~~~

..::

    func (*Task) GetSnx ( whichsol int32, snx []float64 ) ( snx []float64 )

`snx []float64`
    The snx vector.

Obtains the snx vector for a solution.


GetSnxSlice
~~~~~~~~~~~

..::

    func (*Task) GetSnxSlice
        ( whichsol int32,
          first int32,
          last int32,
          snx []float64 )
        ( snx []float64 )


Obtains a slice of the snx vector for a solution.


GetSolSta
~~~~~~~~~

..::

    func (*Task) GetSolSta ( whichsol int32 ) ( solsta int32 )


Obtains the solution status.


GetSolution
~~~~~~~~~~~

..::

    func (*Task) GetSolution
        ( whichsol int32,
          skc []int32,
          skx []int32,
          skn []int32,
          xc []float64,
          xx []float64,
          y []float64,
          slc []float64,
          suc []float64,
          slx []float64,
          sux []float64,
          snx []float64 )
        ( prosta int32,
          solsta int32,
          skc []int32,
          skx []int32,
          skn []int32,
          xc []float64,
          xx []float64,
          y []float64,
          slc []float64,
          suc []float64,
          slx []float64,
          sux []float64,
          snx []float64 )


Obtains the complete solution.


GetSolutionI
~~~~~~~~~~~~

..::

    func (*Task) GetSolutionI
        ( accmode int32,
          i int32,
          whichsol int32 )
        ( sk int32,
          x float64,
          sl float64,
          su float64,
          sn float64 )

`accmode int32`
    Defines whether solution information for a constraint or for a variable is retrieved.
`i int32`
    Index of the constraint or variable.

Obtains the solution for a single constraint or variable.


GetSolutionInfo
~~~~~~~~~~~~~~~

..::

    func (*Task) GetSolutionInfo
        ( whichsol int32 )
        ( pobj float64,
          pviolcon float64,
          pviolvar float64,
          pviolbarvar float64,
          pviolcone float64,
          pviolitg float64,
          dobj float64,
          dviolcon float64,
          dviolvar float64,
          dviolbarvar float64,
          dviolcone float64 )


Obtains information about of a solution.


GetSolutionSlice
~~~~~~~~~~~~~~~~

..::

    func (*Task) GetSolutionSlice
        ( whichsol int32,
          solitem int32,
          first int32,
          last int32,
          values []float64 )
        ( values []float64 )

`first int32`
    Index of the first value in the slice.
`last int32`
    Value of the last index+1 in the slice.
`values []float64`
    The values of the requested solution elements.

Obtains a slice of the solution.


GetSparseSymMat
~~~~~~~~~~~~~~~

..::

    func (*Task) GetSparseSymMat
        ( idx int64,
          subi []int32,
          subj []int32,
          valij []float64 )
        ( subi []int32,
          subj []int32,
          valij []float64 )

`idx int64`
    Index of the matrix to get.
`subi []int32`
    Row subscripts of the matrix non-zero elements.
`subj []int32`
    Column subscripts of the matrix non-zero elements.
`valij []float64`
    Coefficients of the matrix non-zero elements.

Gets a single symmetric matrix from the matrix store.


GetStrParam
~~~~~~~~~~~

..::

    func (*Task) GetStrParam ( param int32 ) ( len int32, parvalue string )


Obtains the value of a string parameter.


GetStrParamLen
~~~~~~~~~~~~~~

..::

    func (*Task) GetStrParamLen ( param int32 ) ( len int32 )


Obtains the length of a string parameter.


GetSuc
~~~~~~

..::

    func (*Task) GetSuc ( whichsol int32, suc []float64 ) ( suc []float64 )

`suc []float64`
    The suc vector.

Obtains the suc vector for a solution.


GetSucSlice
~~~~~~~~~~~

..::

    func (*Task) GetSucSlice
        ( whichsol int32,
          first int32,
          last int32,
          suc []float64 )
        ( suc []float64 )


Obtains a slice of the suc vector for a solution.


GetSux
~~~~~~

..::

    func (*Task) GetSux ( whichsol int32, sux []float64 ) ( sux []float64 )

`sux []float64`
    The sux vector.

Obtains the sux vector for a solution.


GetSuxSlice
~~~~~~~~~~~

..::

    func (*Task) GetSuxSlice
        ( whichsol int32,
          first int32,
          last int32,
          sux []float64 )
        ( sux []float64 )


Obtains a slice of the sux vector for a solution.


GetSymMatInfo
~~~~~~~~~~~~~

..::

    func (*Task) GetSymMatInfo
        ( idx int64 )
        ( dim int32,
          nz int64,
          type int32 )

`idx int64`
    Index of the matrix that is requested information about.

Obtains information of  a matrix from the symmetric matrix storage E.


GetTaskName
~~~~~~~~~~~

..::

    func (*Task) GetTaskName (  ) ( taskname string )


Obtains the task name.


GetTaskNameLen
~~~~~~~~~~~~~~

..::

    func (*Task) GetTaskNameLen (  ) ( len int32 )


Obtains the length the task name.


GetVarBound
~~~~~~~~~~~

..::

    func (*Task) GetVarBound
        ( i int32 )
        ( bk int32,
          bl float64,
          bu float64 )

`i int32`
    Index of the variable for which the bound information should be obtained.

Obtains bound information for one variable.


GetVarBoundSlice
~~~~~~~~~~~~~~~~

..::

    func (*Task) GetVarBoundSlice
        ( first int32,
          last int32,
          bk []int32,
          bl []float64,
          bu []float64 )
        ( bk []int32,
          bl []float64,
          bu []float64 )


Obtains bounds information for a slice of the variables.


GetVarName
~~~~~~~~~~

..::

    func (*Task) GetVarName ( j int32 ) ( name string )

`j int32`
    Index.

Obtains a name of a variable.


GetVarNameIndex
~~~~~~~~~~~~~~~

..::

    func (*Task) GetVarNameIndex ( somename string ) ( asgn int32, index int32 )

`somename string`
    The name which should be checked.

Checks whether the name somename has been assigned  to any variable.


GetVarNameLen
~~~~~~~~~~~~~

..::

    func (*Task) GetVarNameLen ( i int32 ) ( len int32 )

`i int32`
    Index.

Obtains the length of a name of a variable variable.


GetVarType
~~~~~~~~~~

..::

    func (*Task) GetVarType ( j int32 ) ( vartype int32 )

`j int32`
    Index of the variable.

Gets the variable type of one variable.


GetVarTypeList
~~~~~~~~~~~~~~

..::

    func (*Task) GetVarTypeList ( subj []int32, vartype []int32 ) ( vartype []int32 )

`subj []int32`
    A list of variable indexes.
`vartype []int32`
    Returns the variables types corresponding the variable indexes requested.

Obtains the variable type for one or more variables.


GetVersion
~~~~~~~~~~

..::

    func GetVersion
        (  )
        ( major int32,
          minor int32,
          build int32,
          revision int32,
          res int32 )


Obtains |mosek| version information.


GetXc
~~~~~

..::

    func (*Task) GetXc ( whichsol int32, xc []float64 ) ( xc []float64 )

`xc []float64`
    The xc vector.

Obtains the xc vector for a solution.


GetXcSlice
~~~~~~~~~~

..::

    func (*Task) GetXcSlice
        ( whichsol int32,
          first int32,
          last int32,
          xc []float64 )
        ( xc []float64 )


Obtains a slice of the xc vector for a solution.


GetXx
~~~~~

..::

    func (*Task) GetXx ( whichsol int32, xx []float64 ) ( xx []float64 )

`xx []float64`
    The xx vector.

Obtains the xx vector for a solution.


GetXxSlice
~~~~~~~~~~

..::

    func (*Task) GetXxSlice
        ( whichsol int32,
          first int32,
          last int32,
          xx []float64 )
        ( xx []float64 )


Obtains a slice of the xx vector for a solution.


GetY
~~~~

..::

    func (*Task) GetY ( whichsol int32, y []float64 ) ( y []float64 )

`y []float64`
    The y vector.

Obtains the y vector for a solution.


GetYSlice
~~~~~~~~~

..::

    func (*Task) GetYSlice
        ( whichsol int32,
          first int32,
          last int32,
          y []float64 )
        ( y []float64 )


Obtains a slice of the y vector for a solution.


InitBasisSolve
~~~~~~~~~~~~~~

..::

    func (*Task) InitBasisSolve ( basis []int32 ) ( basis []int32 )

`basis []int32`
    The array of basis indexes to use.

Prepare a task for basis solver.


InputData
~~~~~~~~~

..::

    func (*Task) InputData
        ( maxnumcon int32,
          maxnumvar int32,
          c []float64,
          cfix float64,
          aptrb []int64,
          aptre []int64,
          asub []int32,
          aval []float64,
          bkc []int32,
          blc []float64,
          buc []float64,
          bkx []int32,
          blx []float64,
          bux []float64 )


Input the linear part of an optimization task in one function call.


IsDouParName
~~~~~~~~~~~~

..::

    func (*Task) IsDouParName ( parname string ) ( param int32 )


Checks a double parameter name.


IsIntParName
~~~~~~~~~~~~

..::

    func (*Task) IsIntParName ( parname string ) ( param int32 )


Checks an integer parameter name.


IsStrParName
~~~~~~~~~~~~

..::

    func (*Task) IsStrParName ( parname string ) ( param int32 )


Checks a string parameter name.


Licensecleanup
~~~~~~~~~~~~~~

..::

    func Licensecleanup (  ) ( res int32 )


Stops all threads and delete all handles used by the license system.


LinkFileToStream
~~~~~~~~~~~~~~~~

..::

    func (*Task) LinkFileToStream
        ( whichstream int32,
          filename string,
          append int32 )

`filename string`
    The name of the file where the stream is written.
`append int32`
    If this argument is 0 the output file will be overwritten, otherwise text is append to the output file.

Directs all output from a task stream to a file.


Linkfiletostream
~~~~~~~~~~~~~~~~

..::

    func (*Env) Linkfiletostream
        ( whichstream int32,
          filename string,
          append int32 )

`filename string`
    Name of the file to write stream data to.
`append int32`
    If this argument is non-zero, the output is appended to the file.

Directs all output from a stream to a file.


OneSolutionSummary
~~~~~~~~~~~~~~~~~~

..::

    func (*Task) OneSolutionSummary ( whichstream int32, whichsol int32 )

Prints a short summary for the specified solution.


Optimize
~~~~~~~~

..::

    func (*Task) Optimize (  ) ( trmcode int32 )


Optimizes the problem.


OptimizerSummary
~~~~~~~~~~~~~~~~

..::

    func (*Task) OptimizerSummary ( whichstream int32 )

Prints a short summary with optimizer statistics for last optimization.


Potrf
~~~~~

..::

    func (*Env) Potrf
        ( uplo int32,
          n int32,
          a []float64 )
        ( a []float64 )

`uplo int32`
    Indicates whether the upper or lower triangular part of the matrix is stored.
`n int32`
    Dimension of the symmetric matrix.
`a []float64`
    A symmetric matrix stored in column-major order. Only the lower or the upper triangular part is used, accordingly with the uplo parameter. It will contain the result on exit.

Computes a Cholesky factorization a dense matrix.


PrimalRepair
~~~~~~~~~~~~

..::

    func (*Task) PrimalRepair
        ( wlc []float64,
          wuc []float64,
          wlx []float64,
          wux []float64 )

`wlc []float64`
    Weights associated with relaxing lower bounds on the constraints.
`wuc []float64`
    Weights associated with relaxing the upper bound on the constraints.
`wlx []float64`
    Weights associated with relaxing the lower bounds of the variables.
`wux []float64`
    Weights associated with relaxing the upper bounds of variables.

The function repairs a primal infeasible optimization problem by adjusting the bounds on the constraints and variables.


PrimalSensitivity
~~~~~~~~~~~~~~~~~

..::

    func (*Task) PrimalSensitivity
        ( subi []int32,
          marki []int32,
          subj []int32,
          markj []int32,
          leftpricei []float64,
          rightpricei []float64,
          leftrangei []float64,
          rightrangei []float64,
          leftpricej []float64,
          rightpricej []float64,
          leftrangej []float64,
          rightrangej []float64 )
        ( leftpricei []float64,
          rightpricei []float64,
          leftrangei []float64,
          rightrangei []float64,
          leftpricej []float64,
          rightpricej []float64,
          leftrangej []float64,
          rightrangej []float64 )

`subi []int32`
    Indexes of bounds on constraints to analyze.
`marki []int32`
    Mark which constraint bounds to analyze.
`subj []int32`
    Indexes of bounds on variables to analyze.
`markj []int32`
    Mark which variable bounds to analyze.
`leftpricei []float64`
    Left shadow price for constraints.
`rightpricei []float64`
    Right shadow price for constraints.
`leftrangei []float64`
    Left range for constraints.
`rightrangei []float64`
    Right range for constraints.
`leftpricej []float64`
    Left price for variables.
`rightpricej []float64`
    Right price for variables.
`leftrangej []float64`
    Left range for variables.
`rightrangej []float64`
    Right range for variables.

Perform sensitivity analysis on bounds.


ProStaToStr
~~~~~~~~~~~

..::

    func (*Task) ProStaToStr ( prosta int32 ) ( str string )


Obtains a string containing the name of a problem status given.


ProbTypeToStr
~~~~~~~~~~~~~

..::

    func (*Task) ProbTypeToStr ( probtype int32 ) ( str string )


Obtains a string containing the name of a problem type given.


PutACol
~~~~~~~

..::

    func (*Task) PutACol
        ( j int32,
          subj []int32,
          valj []float64 )

`j int32`
    Column index.
`subj []int32`
    Row indexes of non-zero values in column.
`valj []float64`
    New non-zero values of column.

Replaces all elements in one column of A.


PutAColList
~~~~~~~~~~~

..::

    func (*Task) PutAColList
        ( sub []int32,
          ptrb []int32,
          ptre []int32,
          asub []int32,
          aval []float64 )

`sub []int32`
    Indexes of columns that should be replaced.
`ptrb []int32`
    Array of pointers to the first element in the columns.
`ptre []int32`
    Array of pointers to the last element plus one in the columns.
`asub []int32`
    Variable indexes.

Replaces all elements in several columns the linear constraint matrix by new values.


PutAColSlice
~~~~~~~~~~~~

..::

    func (*Task) PutAColSlice
        ( first int32,
          last int32,
          ptrb []int64,
          ptre []int64,
          asub []int32,
          aval []float64 )

`first int32`
    First column in the slice.
`last int32`
    Last column plus one in the slice.
`ptrb []int64`
    Array of pointers to the first element in the columns.
`ptre []int64`
    Array of pointers to the last element plus one in the columns.
`asub []int32`
    Variable indexes.

Replaces all elements in several columns the linear constraint matrix by new values.


PutARow
~~~~~~~

..::

    func (*Task) PutARow
        ( i int32,
          subi []int32,
          vali []float64 )

`i int32`
    row index.
`subi []int32`
    Row indexes of non-zero values in row.
`vali []float64`
    New non-zero values of row.

Replaces all elements in one row of A.


PutARowList
~~~~~~~~~~~

..::

    func (*Task) PutARowList
        ( sub []int32,
          aptrb []int32,
          aptre []int32,
          asub []int32,
          aval []float64 )

`sub []int32`
    Indexes of rows or columns that should be replaced.
`aptrb []int32`
    Array of pointers to the first element in the rows or columns.
`aptre []int32`
    Array of pointers to the last element plus one in the rows or columns.
`asub []int32`
    Variable indexes.

Replaces all elements in several rows the linear constraint matrix by new values.


PutARowSlice
~~~~~~~~~~~~

..::

    func (*Task) PutARowSlice
        ( first int32,
          last int32,
          ptrb []int64,
          ptre []int64,
          asub []int32,
          aval []float64 )

`first int32`
    First row in the slice.
`last int32`
    Last row plus one in the slice.
`ptrb []int64`
    Array of pointers to the first element in the rows.
`ptre []int64`
    Array of pointers to the last element plus one in the rows.
`asub []int32`
    Variable indexes.

Replaces all elements in several rows the linear constraint matrix by new values.


PutAij
~~~~~~

..::

    func (*Task) PutAij
        ( i int32,
          j int32,
          aij float64 )

`i int32`
    Index of the constraint in which the change should occur.
`j int32`
    Index of the variable in which the change should occur.
`aij float64`
    New coefficient.

Changes a single value in the linear coefficient matrix.


PutAijList
~~~~~~~~~~

..::

    func (*Task) PutAijList
        ( subi []int32,
          subj []int32,
          valij []float64 )

`subi []int32`
    Constraint indexes in which the change should occur.
`subj []int32`
    Variable indexes in which the change should occur.
`valij []float64`
    New coefficient values.

Changes one or more coefficients in the linear constraint matrix.


PutBaraBlockTriplet
~~~~~~~~~~~~~~~~~~~

..::

    func (*Task) PutBaraBlockTriplet
        ( num int64,
          subi []int32,
          subj []int32,
          subk []int32,
          subl []int32,
          valijkl []float64 )

`num int64`
    Number of elements in the block triplet form.
`subi []int32`
    Constraint index.
`subj []int32`
    Symmetric matrix variable index.
`subk []int32`
    Block row index.
`subl []int32`
    Block column index.
`valijkl []float64`
    The numerical value associated with the block triplet.

Inputs barA in block triplet form.


PutBaraIj
~~~~~~~~~

..::

    func (*Task) PutBaraIj
        ( i int32,
          j int32,
          sub []int64,
          weights []float64 )

`i int32`
    Row index of barA.
`j int32`
    Column index of barA.
`sub []int64`
    See argument weights for an explanation.
`weights []float64`
    Weights in the weighted sum.

Inputs an element of barA.


PutBarcBlockTriplet
~~~~~~~~~~~~~~~~~~~

..::

    func (*Task) PutBarcBlockTriplet
        ( num int64,
          subj []int32,
          subk []int32,
          subl []int32,
          valjkl []float64 )

`num int64`
    Number of elements in the block triplet form.
`subj []int32`
    Symmetric matrix variable index.
`subk []int32`
    Block row index.
`subl []int32`
    Block column index.
`valjkl []float64`
    The numerical value associated with the block triplet.

Inputs barC in block triplet form.


PutBarcJ
~~~~~~~~

..::

    func (*Task) PutBarcJ
        ( j int32,
          sub []int64,
          weights []float64 )

`j int32`
    Index of the element in barc` that should be changed.
`sub []int64`
    sub is list of indexes of those symmetric matrices appearing in sum.
`weights []float64`
    The weights of the terms in the weighted sum.

Changes one element in barc.


PutBarsJ
~~~~~~~~

..::

    func (*Task) PutBarsJ
        ( whichsol int32,
          j int32,
          barsj []float64 )

`j int32`
    Index of the semidefinite variable.
`barsj []float64`
    Value of the j'th variable of barx.

Sets the dual solution for a semidefinite variable.


PutBarvarName
~~~~~~~~~~~~~

..::

    func (*Task) PutBarvarName ( j int32, name string )
`j int32`
    Index of the variable.
`name string`
    The variable name.

Puts the name of a semidefinite variable.


PutBarxJ
~~~~~~~~

..::

    func (*Task) PutBarxJ
        ( whichsol int32,
          j int32,
          barxj []float64 )

`j int32`
    Index of the semidefinite variable.
`barxj []float64`
    Value of the j'th variable of barx.

Sets the primal solution for a semidefinite variable.


PutBound
~~~~~~~~

..::

    func (*Task) PutBound
        ( accmode int32,
          i int32,
          bk int32,
          bl float64,
          bu float64 )

`accmode int32`
    Defines whether the bound for a constraint or a variable is changed.
`i int32`
    Index of the constraint or variable.
`bk int32`
    New bound key.
`bl float64`
    New lower bound.
`bu float64`
    New upper bound.

Changes the bound for either one constraint or one variable.


PutBoundList
~~~~~~~~~~~~

..::

    func (*Task) PutBoundList
        ( accmode int32,
          sub []int32,
          bk []int32,
          bl []float64,
          bu []float64 )

`accmode int32`
    Defines whether to access bounds on variables or constraints.
`sub []int32`
    Subscripts of the bounds that should be changed.
`bk []int32`
    Bound keys for variables or constraints.
`bl []float64`
    Bound keys for variables or constraints.
`bu []float64`
    Constraint or variable upper bounds.

Changes the bounds of constraints or variables.


PutBoundSlice
~~~~~~~~~~~~~

..::

    func (*Task) PutBoundSlice
        ( con int32,
          first int32,
          last int32,
          bk []int32,
          bl []float64,
          bu []float64 )

`con int32`
    Determines whether variables or constraints are modified.

Modifies bounds.


PutCJ
~~~~~

..::

    func (*Task) PutCJ ( j int32, cj float64 )
`j int32`
    Index of the variable whose objective coefficient should be changed.
`cj float64`
    New coefficient value.

Modifies one linear coefficient in the objective.


PutCList
~~~~~~~~

..::

    func (*Task) PutCList ( subj []int32, val []float64 )
`subj []int32`
    Index of variables for which objective coefficients should be changed.
`val []float64`
    New numerical values for the objective coefficients that should be modified.

Modifies a part of the linear objective coefficients.


PutCSlice
~~~~~~~~~

..::

    func (*Task) PutCSlice
        ( first int32,
          last int32,
          slice []float64 )

`first int32`
    First element in the slice of c.
`last int32`
    Last element plus 1 of the slice in c to be changed.
`slice []float64`
    New numerical values for the objective coefficients that should be modified.

Modifies a slice of the linear objective coefficients.


PutCallbackFunc
~~~~~~~~~~~~~~~

..::

    func (t *Task) PutCallbackFunc ( fun func(int32) int )

Add a callback function to the task.

The callback function takes one integer argument that indicates the
progress of the solver (`mosek.CALLBACK_...`). It returns an integer
value: `0` means that the solver should just continue, anything else
means that the solver will stop.


PutCfix
~~~~~~~

..::

    func (*Task) PutCfix ( cfix float64 )

Replaces the fixed term in the objective.


PutConBound
~~~~~~~~~~~

..::

    func (*Task) PutConBound
        ( i int32,
          bk int32,
          bl float64,
          bu float64 )

`i int32`
    Index of the constraint.
`bk int32`
    New bound key.
`bl float64`
    New lower bound.
`bu float64`
    New upper bound.

Changes the bound for one constraint.


PutConBoundList
~~~~~~~~~~~~~~~

..::

    func (*Task) PutConBoundList
        ( sub []int32,
          bkc []int32,
          blc []float64,
          buc []float64 )

`sub []int32`
    List constraints indexes.
`bkc []int32`
    New bound keys.
`blc []float64`
    New lower bound values.
`buc []float64`
    New upper bounds values.

Changes the bounds of a list of constraints.


PutConBoundSlice
~~~~~~~~~~~~~~~~

..::

    func (*Task) PutConBoundSlice
        ( first int32,
          last int32,
          bk []int32,
          bl []float64,
          bu []float64 )

`first int32`
    Index of the first constraint in the slice.
`last int32`
    Index of the last constraint in the slice plus 1.
`bk []int32`
    New bound keys.
`bl []float64`
    New lower bounds.
`bu []float64`
    New upper bounds.

Changes the bounds for a slice of the constraints.


PutConName
~~~~~~~~~~

..::

    func (*Task) PutConName ( i int32, name string )
`i int32`
    Index of the constraint.
`name string`
    The variable name.

Puts the name of a constraint.


PutCone
~~~~~~~

..::

    func (*Task) PutCone
        ( k int32,
          ct int32,
          conepar float64,
          submem []int32 )

`k int32`
    Index of the cone.

Replaces a conic constraint.


PutConeName
~~~~~~~~~~~

..::

    func (*Task) PutConeName ( j int32, name string )
`j int32`
    Index of the cone.
`name string`
    The variable name.

Puts the name of a cone.


PutDouParam
~~~~~~~~~~~

..::

    func (*Task) PutDouParam ( param int32, parvalue float64 )

Sets a double parameter.


PutInfoCallbackFunc
~~~~~~~~~~~~~~~~~~~

..::

    func (t *Task) PutInfoCallbackFunc ( fun func(int32) int )

Add an information callback function to the task.

The callback function takes four arguments: `(code,dinf,iinf,liinf)`

Callback function arguments:

`code`
    Indicates the progress of the solver (`mosek.CALLBACK_...`).
`dinf`
    An array of `float64` information items. The indexes correspond to `mosek.DINF_...`
`iinf`
    An array of `int32` information items. The indexes correspond to `mosek.IINF_...`
`liinf`
    An array of `int64` information items. The indexes correspond to `mosek.LIINF_...`
    

Callback function returns: Non-zero to indicate that the solver should stop.


PutIntParam
~~~~~~~~~~~

..::

    func (*Task) PutIntParam ( param int32, parvalue int32 )

Sets an integer parameter.


PutKeepDlls
~~~~~~~~~~~

..::

    func (*Env) PutKeepDlls ( keepdlls int32 )
`keepdlls int32`
    Controls whether explicitly loaded DLLs should be kept.

Controls whether explicitly loaded DLLs should be kept.


PutLicenseCode
~~~~~~~~~~~~~~

..::

    func (*Env) PutLicenseCode ( code []int32 )
`code []int32`
    A license key string.

The purpose of this function is to input a runtime license code.


PutLicenseDebug
~~~~~~~~~~~~~~~

..::

    func (*Env) PutLicenseDebug ( licdebug int32 )
`licdebug int32`
    Enable output of license check-out debug information.

Enables debug information for the license system.


PutLicensePath
~~~~~~~~~~~~~~

..::

    func (*Env) PutLicensePath ( licensepath string )
`licensepath string`
    A path specifycing where to search for the license.

Set the path to the license file.


PutLicenseWait
~~~~~~~~~~~~~~

..::

    func (*Env) PutLicenseWait ( licwait int32 )
`licwait int32`
    Enable waiting for a license.

Control whether mosek should wait for an available license if no license is available.


PutMaxNumANz
~~~~~~~~~~~~

..::

    func (*Task) PutMaxNumANz ( maxnumanz int64 )
`maxnumanz int64`
    New size of the storage reserved for storing the linear coefficient matrix.

The function changes the size of the preallocated storage for linear coefficients.


PutMaxNumBarvar
~~~~~~~~~~~~~~~

..::

    func (*Task) PutMaxNumBarvar ( maxnumbarvar int32 )
`maxnumbarvar int32`
    The maximum number of semidefinite variables.

Sets the number of preallocated symmetric matrix variables in the optimization task.


PutMaxNumCon
~~~~~~~~~~~~

..::

    func (*Task) PutMaxNumCon ( maxnumcon int32 )

Sets the number of preallocated constraints in the optimization task.


PutMaxNumCone
~~~~~~~~~~~~~

..::

    func (*Task) PutMaxNumCone ( maxnumcone int32 )

Sets the number of preallocated conic constraints in the optimization task.


PutMaxNumQNz
~~~~~~~~~~~~

..::

    func (*Task) PutMaxNumQNz ( maxnumqnz int64 )

Changes the size of the preallocated storage for quadratic terms.


PutMaxNumVar
~~~~~~~~~~~~

..::

    func (*Task) PutMaxNumVar ( maxnumvar int32 )

Sets the number of preallocated variables in the optimization task.


PutNaDouParam
~~~~~~~~~~~~~

..::

    func (*Task) PutNaDouParam ( paramname string, parvalue float64 )

Sets a double parameter.


PutNaIntParam
~~~~~~~~~~~~~

..::

    func (*Task) PutNaIntParam ( paramname string, parvalue int32 )

Sets an integer parameter.


PutNaStrParam
~~~~~~~~~~~~~

..::

    func (*Task) PutNaStrParam ( paramname string, parvalue string )

Sets a string parameter.


PutObjName
~~~~~~~~~~

..::

    func (*Task) PutObjName ( objname string )

Assigns a new name to the objective.


PutObjSense
~~~~~~~~~~~

..::

    func (*Task) PutObjSense ( sense int32 )
`sense int32`
    The objective sense of the task

Sets the objective sense.


PutParam
~~~~~~~~

..::

    func (*Task) PutParam ( parname string, parvalue string )

Modifies the value of parameter.


PutQCon
~~~~~~~

..::

    func (*Task) PutQCon
        ( qcsubk []int32,
          qcsubi []int32,
          qcsubj []int32,
          qcval []float64 )


Replaces all quadratic terms in constraints.


PutQConK
~~~~~~~~

..::

    func (*Task) PutQConK
        ( k int32,
          qcsubi []int32,
          qcsubj []int32,
          qcval []float64 )

`k int32`
    The constraint in which the new quadratic elements are inserted.

Replaces all quadratic terms in a single constraint.


PutQObj
~~~~~~~

..::

    func (*Task) PutQObj
        ( qosubi []int32,
          qosubj []int32,
          qoval []float64 )


Replaces all quadratic terms in the objective.


PutQObjIJ
~~~~~~~~~

..::

    func (*Task) PutQObjIJ
        ( i int32,
          j int32,
          qoij float64 )

`i int32`
    Row index for the coefficient to be replaced.
`j int32`
    Column index for the coefficient to be replaced.
`qoij float64`
    The new coefficient value.

Replaces one coefficient in the quadratic term in the objective.


PutSkc
~~~~~~

..::

    func (*Task) PutSkc ( whichsol int32, skc []int32 )

Sets the status keys for the constraints.


PutSkcSlice
~~~~~~~~~~~

..::

    func (*Task) PutSkcSlice
        ( whichsol int32,
          first int32,
          last int32,
          skc []int32 )


Sets the status keys for the constraints.


PutSkx
~~~~~~

..::

    func (*Task) PutSkx ( whichsol int32, skx []int32 )

Sets the status keys for the scalar variables.


PutSkxSlice
~~~~~~~~~~~

..::

    func (*Task) PutSkxSlice
        ( whichsol int32,
          first int32,
          last int32,
          skx []int32 )


Sets the status keys for the variables.


PutSlc
~~~~~~

..::

    func (*Task) PutSlc ( whichsol int32, slc []float64 )
`slc []float64`
    The slc vector.

Sets the slc vector for a solution.


PutSlcSlice
~~~~~~~~~~~

..::

    func (*Task) PutSlcSlice
        ( whichsol int32,
          first int32,
          last int32,
          slc []float64 )


Sets a slice of the slc vector for a solution.


PutSlx
~~~~~~

..::

    func (*Task) PutSlx ( whichsol int32, slx []float64 )
`slx []float64`
    The slx vector.

Sets the slx vector for a solution.


PutSlxSlice
~~~~~~~~~~~

..::

    func (*Task) PutSlxSlice
        ( whichsol int32,
          first int32,
          last int32,
          slx []float64 )


Sets a slice of the slx vector for a solution.


PutSnx
~~~~~~

..::

    func (*Task) PutSnx ( whichsol int32, sux []float64 )
`sux []float64`
    The snx vector.

Sets the snx vector for a solution.


PutSnxSlice
~~~~~~~~~~~

..::

    func (*Task) PutSnxSlice
        ( whichsol int32,
          first int32,
          last int32,
          snx []float64 )


Sets a slice of the snx vector for a solution.


PutSolution
~~~~~~~~~~~

..::

    func (*Task) PutSolution
        ( whichsol int32,
          skc []int32,
          skx []int32,
          skn []int32,
          xc []float64,
          xx []float64,
          y []float64,
          slc []float64,
          suc []float64,
          slx []float64,
          sux []float64,
          snx []float64 )


Inserts a solution.


PutSolutionI
~~~~~~~~~~~~

..::

    func (*Task) PutSolutionI
        ( accmode int32,
          i int32,
          whichsol int32,
          sk int32,
          x float64,
          sl float64,
          su float64,
          sn float64 )

`accmode int32`
    Defines whether solution information for a constraint or for a variable is modified.
`i int32`
    Index of the constraint or variable.
`sk int32`
    Status key of the constraint or variable.
`x float64`
    Solution value of the primal constraint or variable.
`sl float64`
    Solution value of the dual variable associated with the lower bound.
`su float64`
    Solution value of the dual variable associated with the upper bound.
`sn float64`
    Solution value of the dual variable associated with the cone constraint.

Sets the primal and dual solution information for a single constraint or variable.


PutSolutionYI
~~~~~~~~~~~~~

..::

    func (*Task) PutSolutionYI
        ( i int32,
          whichsol int32,
          y float64 )

`i int32`
    Index of the dual variable.
`y float64`
    Solution value of the dual variable.

Inputs the dual variable of a solution.


PutStrParam
~~~~~~~~~~~

..::

    func (*Task) PutStrParam ( param int32, parvalue string )

Sets a string parameter.


PutStreamFunc
~~~~~~~~~~~~~

..::

    func (t *Task) PutStreamFunc ( whichstream int32, fun func(string) )

Add a stream printer function to the task. `whichstream` should be a `mosek.STREAM_...` constant.


PutSuc
~~~~~~

..::

    func (*Task) PutSuc ( whichsol int32, suc []float64 )
`suc []float64`
    The suc vector.

Sets the suc vector for a solution.


PutSucSlice
~~~~~~~~~~~

..::

    func (*Task) PutSucSlice
        ( whichsol int32,
          first int32,
          last int32,
          suc []float64 )


Sets a slice of the suc vector for a solution.


PutSux
~~~~~~

..::

    func (*Task) PutSux ( whichsol int32, sux []float64 )
`sux []float64`
    The sux vector.

Sets the sux vector for a solution.


PutSuxSlice
~~~~~~~~~~~

..::

    func (*Task) PutSuxSlice
        ( whichsol int32,
          first int32,
          last int32,
          sux []float64 )


Sets a slice of the sux vector for a solution.


PutTaskName
~~~~~~~~~~~

..::

    func (*Task) PutTaskName ( taskname string )

Assigns a new name to the task.


PutVarBound
~~~~~~~~~~~

..::

    func (*Task) PutVarBound
        ( j int32,
          bk int32,
          bl float64,
          bu float64 )

`j int32`
    Index of the variable.
`bk int32`
    New bound key.
`bl float64`
    New lower bound.
`bu float64`
    New upper bound.

Changes the bound for one variable.


PutVarBoundList
~~~~~~~~~~~~~~~

..::

    func (*Task) PutVarBoundList
        ( sub []int32,
          bkx []int32,
          blx []float64,
          bux []float64 )

`sub []int32`
    List of variable indexes.
`bkx []int32`
    New bound keys.
`blx []float64`
    New lower bound values.
`bux []float64`
    New upper bounds values.

Changes the bounds of a list of variables.


PutVarBoundSlice
~~~~~~~~~~~~~~~~

..::

    func (*Task) PutVarBoundSlice
        ( first int32,
          last int32,
          bk []int32,
          bl []float64,
          bu []float64 )

`first int32`
    Index of the first variable in the slice.
`last int32`
    Index of the last variable in the slice plus 1.
`bk []int32`
    New bound keys.
`bl []float64`
    New lower bounds.
`bu []float64`
    New upper bounds.

Changes the bounds for a slice of the variables.


PutVarName
~~~~~~~~~~

..::

    func (*Task) PutVarName ( j int32, name string )
`j int32`
    Index of the variable.
`name string`
    The variable name.

Puts the name of a variable.


PutVarType
~~~~~~~~~~

..::

    func (*Task) PutVarType ( j int32, vartype int32 )
`j int32`
    Index of the variable.
`vartype int32`
    The new variable type.

Sets the variable type of one variable.


PutVarTypeList
~~~~~~~~~~~~~~

..::

    func (*Task) PutVarTypeList ( subj []int32, vartype []int32 )
`subj []int32`
    A list of variable indexes for which the variable type should be changed.
`vartype []int32`
    A list of variable types.

Sets the variable type for one or more variables.


PutXc
~~~~~

..::

    func (*Task) PutXc ( whichsol int32, xc []float64 ) ( xc []float64 )

`xc []float64`
    The xc vector.

Sets the xc vector for a solution.


PutXcSlice
~~~~~~~~~~

..::

    func (*Task) PutXcSlice
        ( whichsol int32,
          first int32,
          last int32,
          xc []float64 )


Sets a slice of the xc vector for a solution.


PutXx
~~~~~

..::

    func (*Task) PutXx ( whichsol int32, xx []float64 )
`xx []float64`
    The xx vector.

Sets the xx vector for a solution.


PutXxSlice
~~~~~~~~~~

..::

    func (*Task) PutXxSlice
        ( whichsol int32,
          first int32,
          last int32,
          xx []float64 )


Obtains a slice of the xx vector for a solution.


PutY
~~~~

..::

    func (*Task) PutY ( whichsol int32, y []float64 )
`y []float64`
    The y vector.

Sets the y vector for a solution.


PutYSlice
~~~~~~~~~

..::

    func (*Task) PutYSlice
        ( whichsol int32,
          first int32,
          last int32,
          y []float64 )


Sets a slice of the y vector for a solution.


ReadData
~~~~~~~~

..::

    func (*Task) ReadData ( filename string )
`filename string`
    Input data file name.

Reads problem data from a file.


ReadDataFormat
~~~~~~~~~~~~~~

..::

    func (*Task) ReadDataFormat
        ( filename string,
          format int32,
          compress int32 )

`filename string`
    Input data file name.
`format int32`
    File data format.
`compress int32`
    File compression type.

Reads problem data from a file.


ReadParamFile
~~~~~~~~~~~~~

..::

    func (*Task) ReadParamFile ( filename string )
`filename string`
    Input data file name.

Reads a parameter file.


ReadSolution
~~~~~~~~~~~~

..::

    func (*Task) ReadSolution ( whichsol int32, filename string )

Reads a solution from a file.


ReadSummary
~~~~~~~~~~~

..::

    func (*Task) ReadSummary ( whichstream int32 )

Prints information about last file read.


ReadTask
~~~~~~~~

..::

    func (*Task) ReadTask ( filename string )
`filename string`
    Input file name.

Load task data from a file.


RemoveBarvars
~~~~~~~~~~~~~

..::

    func (*Task) RemoveBarvars ( subset []int32 )
`subset []int32`
    Indexes of symmetric matrix which should be removed.

The function removes a number of symmetric matrix.


RemoveCones
~~~~~~~~~~~

..::

    func (*Task) RemoveCones ( subset []int32 )
`subset []int32`
    Indexes of cones which should be removed.

Removes a conic constraint from the problem.


RemoveCons
~~~~~~~~~~

..::

    func (*Task) RemoveCons ( subset []int32 )
`subset []int32`
    Indexes of constraints which should be removed.

The function removes a number of constraints.


RemoveVars
~~~~~~~~~~

..::

    func (*Task) RemoveVars ( subset []int32 )
`subset []int32`
    Indexes of variables which should be removed.

The function removes a number of variables.


ResizeTask
~~~~~~~~~~

..::

    func (*Task) ResizeTask
        ( maxnumcon int32,
          maxnumvar int32,
          maxnumcone int32,
          maxnumanz int64,
          maxnumqnz int64 )

`maxnumcon int32`
    New maximum number of constraints.
`maxnumvar int32`
    New maximum number of variables.
`maxnumcone int32`
    New maximum number of cones.
`maxnumanz int64`
    New maximum number of linear non-zero elements.
`maxnumqnz int64`
    New maximum number of quadratic non-zeros elements.

Resizes an optimization task.


SensitivityReport
~~~~~~~~~~~~~~~~~

..::

    func (*Task) SensitivityReport ( whichstream int32 )

Creates a sensitivity report.


SetDefaults
~~~~~~~~~~~

..::

    func (*Task) SetDefaults (  )

Resets all parameters values.


SkToStr
~~~~~~~

..::

    func (*Task) SkToStr ( sk int32 ) ( str string )

`sk int32`
    A valid status key.

Obtains a status key string.


SolStaToStr
~~~~~~~~~~~

..::

    func (*Task) SolStaToStr ( solsta int32 ) ( str string )


Obtains a solution status string.


SolutionDef
~~~~~~~~~~~

..::

    func (*Task) SolutionDef ( whichsol int32 ) ( isdef bool )


Checks whether a solution is defined.


SolutionSummary
~~~~~~~~~~~~~~~

..::

    func (*Task) SolutionSummary ( whichstream int32 )

Prints a short summary of the current solutions.


SolveWithBasis
~~~~~~~~~~~~~~

..::

    func (*Task) SolveWithBasis
        ( transp int32,
          numnz int32,
          sub []int32,
          val []float64 )
        ( numnz int32,
          sub []int32,
          val []float64 )

`transp int32`
    Controls which problem formulation is solved.
`numnz int32`
    Input (number of non-zeros in right-hand side) and output (number of non-zeros in solution vector).
`sub []int32`
    Input (indexes of non-zeros in right-hand side) and output (indexes of non-zeros in solution vector).
`val []float64`
    Input (right-hand side values) and output (solution vector values).

Solve a linear equation system involving a basis matrix.


StartStat
~~~~~~~~~

..::

    func (*Task) StartStat (  )

Starts the statistics file.


StopStat
~~~~~~~~

..::

    func (*Task) StopStat (  )

Stops the statistics file.


StrToConeType
~~~~~~~~~~~~~

..::

    func (*Task) StrToConeType ( str string ) ( conetype int32 )

`str string`
    String corresponding to the cone type code.

Obtains a cone type code.


StrToSk
~~~~~~~

..::

    func (*Task) StrToSk ( str string ) ( sk int32 )

`str string`
    Status key string.

Obtains a status key.


Syeig
~~~~~

..::

    func (*Env) Syeig
        ( uplo int32,
          n int32,
          a []float64,
          w []float64 )
        ( w []float64 )

`uplo int32`
    Indicates whether the upper or lower triangular part is used.
`n int32`
    Dimension of the symmetric input matrix.
`a []float64`
    A symmetric matrix stored in column-major order. Only the lower-triangular part is used.
`w []float64`
    Array of minimum dimension n where eigenvalues will be stored.

Computes all eigenvalues of a symmetric dense matrix.


Syevd
~~~~~

..::

    func (*Env) Syevd
        ( uplo int32,
          n int32,
          a []float64,
          w []float64 )
        ( a []float64,
          w []float64 )

`uplo int32`
    Indicates whether the upper or lower triangular part is used.
`n int32`
    Dimension of symmetric input matrix.
`a []float64`
    A symmetric matrix stored in column-major order. Only the lower-triangular part is used. It will be overwritten on exit.
`w []float64`
    An array where eigenvalues will be stored. Its lenght must be at least the dimension of the input matrix.

Computes all the eigenvalue and eigenvectors of a symmetric dense matrix, and thus its eigenvalue decomposition.


Syrk
~~~~

..::

    func (*Env) Syrk
        ( uplo int32,
          trans int32,
          n int32,
          k int32,
          alpha float64,
          a []float64,
          beta float64,
          c []float64 )
        ( c []float64 )

`uplo int32`
    Indicates whether the upper or lower triangular part of C is stored.
`trans int32`
    Indicates whether the matrix A must be transposed.
`n int32`
    Specifies the order of C.
`k int32`
    Indicates the number of rows or columns of A, and its rank.
`alpha float64`
    A scalar value multipling the result of the matrix multiplication.
`a []float64`
    The pointer to the array storing matrix A in a column-major format.
`beta float64`
    A scalar value that multiplies C.
`c []float64`
    The pointer to the array storing matrix C in a column-major format.

Performs a rank-k update of a symmetric matrix.


Toconic
~~~~~~~

..::

    func (*Task) Toconic (  )

Inplace reformulation of a QCQP to a COP


UpdateSolutionInfo
~~~~~~~~~~~~~~~~~~

..::

    func (*Task) UpdateSolutionInfo ( whichsol int32 )

Update the information items related to the solution.


WriteData
~~~~~~~~~

..::

    func (*Task) WriteData ( filename string )
`filename string`
    Output file name.

Writes problem data to a file.


WriteJsonSol
~~~~~~~~~~~~

..::

    func (*Task) WriteJsonSol ( filename string )

Write a solution to a file.


WriteParamFile
~~~~~~~~~~~~~~

..::

    func (*Task) WriteParamFile ( filename string )
`filename string`
    The name of parameter file.

Writes all the parameters to a parameter file.


WriteSolution
~~~~~~~~~~~~~

..::

    func (*Task) WriteSolution ( whichsol int32, filename string )

Write a solution to a file.


WriteTask
~~~~~~~~~

..::

    func (*Task) WriteTask ( filename string )
`filename string`
    Output file name.

Write a complete binary dump of the task data.

