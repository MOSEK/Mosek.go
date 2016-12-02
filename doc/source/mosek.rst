
The `mosek` package
===================

There are two types defined in the `mosek` package:

``Env``
    Represents the MOSEK environment. Usually only one environment is necessary.
``Task``
    Represents a MOSEK task (problem data, parameters, solutions etc.).

Most functions in the package returns, at least, a response code. If
the code equals ``mosek.RES_OK``, no error occurred, otherwise it is an
error code `mosek.RES_ERR_...`. You should always check the response
code before proceeding.


Creating ``Env`` and ``Task``
~~~~~~~~~~~~~~~~~~~~~~~~~

::

    func MakeEnv(env Env, r int32)
    func (e *Env) DeleteEnv()

Create and delete an environment. An environment should always be
explicitly deleted, but must not be deleted before all tasks createn
through it.::

    func (env *Env) MakeTask(task Task, r int32)
    func (t *Task) DeleteTask()

Create and delete a task. A task should always be explicitly deleted
when it is not used anymore. The GO garbage collector will collect the
structure but not the underlying MOSEK allocations.


AnalyzeNames
~~~~~~~~~~~~

 Analyze the names and issue an error for the first invalid name. ::

    func (*Task) AnalyzeNames ( whichstream int32, nametype int32 )

``nametype int32``
    The type of names e.g. valid in MPS or LP files.


The function analyzes the names and issue an error if a name is invalid.



AnalyzeProblem
~~~~~~~~~~~~~~

 Analyze the data of a task. ::

    func (*Task) AnalyzeProblem ( whichstream int32 )



The function analyzes the data of task and writes out a report.



AnalyzeSolution
~~~~~~~~~~~~~~~

 Print information related to the quality of the solution. ::

    func (*Task) AnalyzeSolution ( whichstream int32, whichsol int32 )



Print information related to the quality of the solution and
other solution statistics.

By default this function prints information about the
largest infeasibilites in the solution, the primal (and
possibly dual) objective value and the solution status.

Following parameters can be used to configure the printed statistics:

* ``IPAR_ANA_SOL_BASIS`` enables or disables printing of statistics specific to the basis solution (condition number, number of basic variables etc.). Default is on.
* ``IPAR_ANA_SOL_PRINT_VIOLATED`` enables or disables listing names of all constraints (both primal and dual) which are violated by the solution. Default is off.
* ``DPAR_ANA_SOL_INFEAS_TOL`` is the tolerance defining when a constraint is considered violated. If a constraint is violated more than this, it will be listed in the summary.




AppendBarvars
~~~~~~~~~~~~~

Appends a semidefinite  variable of dimension dim to the problem. ::

    func (*Task) AppendBarvars ( dim []int32 )

``dim []int32``
     Dimension of symmetric matrix variables to be added.


Appends a positive semidefinite matrix variable of dimension ``dim`` to the problem.



AppendCone
~~~~~~~~~~

Appends a new cone constraint to the problem. ::

    func (*Task) AppendCone
        ( ct int32,
          conepar float64,
          submem []int32 )




Appends a new conic constraint to the problem. Hence, add a constraint

.. math:: \hat{x} \in \K

to the problem where :math:`K` is a convex cone. :math:`\hat{x}` is a
subset of the variables which will be specified by the argument
``submem``.

Depending on the value of ``ct`` this function appends a normal (``CT_QUAD``) or rotated quadratic cone (``CT_RQUAD``).

Define 

.. math:: \hat{x} = x_{\mathtt{submem}[0]},\ldots,x_{\mathtt{submem}[\mathtt{nummem}-1]}.

Depending on the value of ``ct`` this function appends one of the constraints:

* Quadratic cone (``CT_QUAD``) : 

  .. math:: \hat{x}_0 \geq \sqrt{\sum_{i=1}^{i<\mathtt{nummem}} \hat{x}_i^2}

* Rotated quadratic cone (``CT_RQUAD``) : 

  .. math:: 2 \hat{x}_0 \hat{x}_1 \geq \sum_{i=2}^{i<\mathtt{nummem}} \hat{x}^2_i, \quad \hat{x}_{0}, \hat{x}_1 \geq 0

Please note that the sets of variables appearing in different conic constraints must be disjoint.

For an explained code example see Section :ref:`doc.tutorial_cqo`.




AppendConeSeq
~~~~~~~~~~~~~

Appends a new conic constraint to the problem. ::

    func (*Task) AppendConeSeq
        ( ct int32,
          conepar float64,
          nummem int32,
          j int32 )


``nummem int32``
    Dimension of the conic constraint.
``j int32``
    Index of the first variable in the conic constraint.


Appends a new conic constraint to the problem. The function assumes the members of cone are sequential where the first member has index ``j`` and the last ``j+nummem-1``.



AppendConesSeq
~~~~~~~~~~~~~~

Appends a multiple conic constraints to the problem. ::

    func (*Task) AppendConesSeq
        ( ct []int32,
          conepar []float64,
          nummem []int32,
          j int32 )


``j int32``
    Index of the first variable in the first cone to be appended.


Appends a number conic constraints to the problem. The :math:`k`\ th
cone is assumed to be of dimension ``nummem[k]``. Moreover, is assumed
that the first variable of the first cone has index :math:`j` and the
index of the variable in each cone are sequential. Finally, it assumed
in the second cone is the last index of first cone plus one and so
forth.



AppendCons
~~~~~~~~~~

 Appends a number of constraints to the optimization task. ::

    func (*Task) AppendCons ( num int32 )

``num int32``
     Number of constraints which should be appended.


Appends a number of constraints to the model. Appended constraints will be declared free. Please note that MOSEK will automatically expand the problem dimension to accommodate the additional constraints.



AppendSparseSymMat
~~~~~~~~~~~~~~~~~~

Appends a general sparse symmetric matrix to the vector E of symmetric matrixes. ::

    func (*Task) AppendSparseSymMat
        ( dim int32,
          subi []int32,
          subj []int32,
          valij []float64 )
        ( idx int64 )


``dim int32``
     Dimension of the symmetric matrix that is appended.
``subi []int32``
     Row subscript in the triplets.
``subj []int32``
     Column subscripts in the triplets.
``valij []float64``
     Values of each triplet.
``idx int64``
     Unique index assigned to inputted matrix.


MOSEK maintains a storage of symmetric data matrixes that is used to build
the :math:`\bar{c}` and :math:`\bar{A}`. The storage can be thought of as a vector of
symmetric matrixes denoted :math:`E`. Hence, :math:`E_i` is a symmetric matrix of certain
dimension.

This function appends a general sparse symmetric matrix on triplet form to the
vector :math:`E` of symmetric matrixes.  The vectors ``subi``, ``subj``, and
``valij`` contains the row subscripts, column subscripts and values of each
element in the symmetric matrix to be appended.  Since the matrix that is
appended is symmetric then only the lower triangular part should be specified.
Moreover, duplicates are not allowed.

Observe the function reports the index (position) of the appended matrix in
:math:`E`. This index should be used for later references to the appended matrix.



AppendVars
~~~~~~~~~~

 Appends a number of variables to the optimization task. ::

    func (*Task) AppendVars ( num int32 )

``num int32``
     Number of variables which should be appended.


Appends a number of variables to the model. Appended variables will be fixed at zero. Please note that MOSEK will automatically expand the problem dimension to accommodate the additional variables.



Axpy
~~~~

Adds alpha times x to y. ::

    func (*Env) Axpy
        ( n int32,
          alpha float64,
          x []float64,
          y []float64 )
        ( y []float64 )


``n int32``
     Length of the vectors.
``alpha float64``
     The scalar that multiplies x.
``x []float64``
     The :math:`x` vector.
``y []float64``
     The :math:`y` vector.

Adds :math:`\alpha x` to :math:`y`. 


BasisCond
~~~~~~~~~

 Computes conditioning information for the basis matrix. ::

    func (*Task) BasisCond (  ) ( nrmbasis float64, nrminvbasis float64 )


``nrmbasis float64``
     An estimate for the 1 norm of the basis.
``nrminvbasis float64``
     An estimate for the 1 norm of the inverse of the basis.


If a basic solution is available and it defines a nonsingular basis, then
this function computes the 1-norm estimate of the basis matrix and an 1-norm estimate
for the inverse of the basis matrix. The 1-norm estimates are computed using the method
outlined in [STEWART:98:A], pp. 388-391.

By definition the 1-norm condition number of a matrix :math:`B` is defined as

.. math:: \K_1(B) := \|B\|_1 \|B^{-1}|.

Moreover, the larger the condition number is the harder it is to solve
linear equation systems involving :math:`B`.  Given estimates for
:math:`\|B\|_1` and :math:`\|B^{-1}\|_1` it is also possible to
estimate :math:`\kappa_1(B)`.



CheckConvexity
~~~~~~~~~~~~~~

 Checks if a quadratic optimization problem is convex. ::

    func (*Task) CheckConvexity (  )



This function checks if a quadratic optimization problem is convex. The amount of checking is controlled by ``IPAR_CHECK_CONVEXITY``.

The function reports an error if the problem is not convex.



CheckInAll
~~~~~~~~~~

Check in all unsued license features to the license token server.  ::

    func (*Env) CheckInAll (  )



Check in all unsued license features to the license token server. 



CheckInLicense
~~~~~~~~~~~~~~

Check in a license feature from the license server ahead of time. ::

    func (*Env) CheckInLicense ( feature int32 )

``feature int32``
     Feature to check in to the license system.


Check in a license feature to the license server. By default all licenses
consumed by functions using a single environment is kept checked out for the
lifetime of the MOSEK environment. This function checks in a given license
feature to the license server immediately.

If the given license feature is not checked out or is in use by a call to
``optimize`` calling this function has no effect.

Please note that returning a license to the license server incurs a small
overhead, so frequent calls to this function should be avoided.



CheckMem
~~~~~~~~

Checks the memory allocated by the task. ::

    func (*Task) CheckMem ( file string, line int32 )

``file string``
    File from which the function is called.
``line int32``
     Line in the file from which the function is called.

Checks the memory allocated by the task. 


CheckoutLicense
~~~~~~~~~~~~~~~

Check out a license feature from the license server ahead of time. ::

    func (*Env) CheckoutLicense ( feature int32 )

``feature int32``
     Feature to check out from the license system.


Check out a license feature from the license server. Normally the required
license features will be automatically checked out the first time it is needed
by the function ``optimize``. This function can be used to check out one
or more features ahead of time.

The license will remain checked out until the environment is deleted or the function
``checkinlicense`` is called.

If a given feature is already checked out when this function is called, only
one feature will be checked out from the license server.



ChgBound
~~~~~~~~

 Changes the bounds for one constraint or variable. ::

    func (*Task) ChgBound
        ( accmode int32,
          i int32,
          lower int32,
          finite int32,
          value float64 )


``i int32``
     Index of the constraint or variable for which the bounds should be changed.
``lower int32``
     If non-zero, then the lower bound is changed, otherwise the upper bound is changed.
``finite int32``
    If non-zero, then the given value is assumed to be finite.
``value float64``
    New value for the bound.


Changes a bound for one constraint or variable. If ``accmode`` equals ``ACC_CON``, a constraint bound is changed, otherwise a variable
bound is changed.

If ``lower`` is non-zero, then the lower bound is changed as follows:

.. math::

    \mbox{new lower bound} =
        \left\{
            \begin{array}{ll}
                - \infty,     & \mathtt{finite}=0, \\
                \mathtt{value} & \mbox{otherwise}. 
            \end{array}
        \right.


Otherwise if ``lower`` is zero, then

.. math:: 

    \mbox{new upper bound} = 
        \left\{ 
            \begin{array}{ll}
                \infty,     & \mathtt{finite}=0, \\
                \mathtt{value} & \mbox{otherwise}. 
            \end{array}
        \right.


Please note that this function automatically updates the bound key for  bound, in particular, if the lower and upper bounds are identical, the  bound key is changed to ``fixed``.




ChgConBound
~~~~~~~~~~~

 Changes the bounds for one constraint. ::

    func (*Task) ChgConBound
        ( i int32,
          lower int32,
          finite int32,
          value float64 )


``i int32``
     Index of the constraint for which the bounds should be changed.
``lower int32``
     If non-zero, then the lower bound is changed, otherwise the upper bound is changed.
``finite int32``
    If non-zero, then the given value is assumed to be finite.
``value float64``
    New value for the bound.


Changes a bound for one constraint.

If ``lower`` is non-zero, then the lower bound is changed as follows:

.. math::

    \mbox{new lower bound} =
      \left\{
        \begin{array}{ll}
          - \infty,       & \mathtt{finite}=0, \\
          \mathtt{value}  & \mbox{otherwise}. 
        \end{array}
      \right.

Otherwise if ``lower`` is zero, then

.. math::

    \mbox{new upper bound} = 
      \left\{
        \begin{array}{ll}
          \infty,        & \mathtt{finite}=0, \\
          \mathtt{value} & \mbox{otherwise}. 
        \end{array}
      \right.


Please note that this function automatically updates the bound key for
bound, in particular, if the lower and upper bounds are identical, the
bound key is changed to ``fixed``.



ChgVarBound
~~~~~~~~~~~

 Changes the bounds for one variable. ::

    func (*Task) ChgVarBound
        ( j int32,
          lower int32,
          finite int32,
          value float64 )


``j int32``
     Index of the variable for which the bounds should be changed.
``lower int32``
     If non-zero, then the lower bound is changed, otherwise the upper bound is changed.
``finite int32``
    If non-zero, then the given value is assumed to be finite.
``value float64``
    New value for the bound.


Changes a bound for on variable.

If ``lower`` is non-zero, then the lower bound is changed as follows:

.. math::

    \mbox{new lower bound} =
      \left\{
        \begin{array}{ll}
          - \infty,     & \mathtt{finite}=0, \\
          \mathtt{value} & \mbox{otherwise}. 
        \end{array}
      \right.

Otherwise if ``lower`` is zero, then

.. math::

    \mbox{new upper bound} = 
      \left\{
        \begin{array}{ll}
          \infty,     & \mathtt{finite}=0, \\
          \mathtt{value} & \mbox{otherwise}. 
        \end{array}
      \right.

Please note that this function automatically updates the bound key for bound,
in particular, if the lower and upper bounds are identical, the bound key is
changed to ``fixed``.



CommitChanges
~~~~~~~~~~~~~

Commits all cached problem changes. ::

    func (*Task) CommitChanges (  )



Commits all cached problem changes to the task. It is usually not necessary explicitly to call this function since changes will be committed automatically when required.



DeleteSolution
~~~~~~~~~~~~~~

Undefine a solution and frees the memory it uses. ::

    func (*Task) DeleteSolution ( whichsol int32 )


Undefine a solution and frees the memory it uses. 


Dot
~~~

Computes the inner product of two vectors. ::

    func (*Env) Dot
        ( n int32,
          x []float64,
          y []float64 )
        ( xty float64 )


``n int32``
     Length of the vectors.
``x []float64``
    The x vector.
``y []float64``
    The y vector.
``xty float64``
    The result of the inner product.


Computes the inner product of two vectors :math:`x,y` of lenght :math:`n\geq 0`, i.e

.. math:: x\cdot y= \sum_{i=1}^n x_i y_i.

Note that if :math:`n=0`, then the results of the operation is 0.



DualSensitivity
~~~~~~~~~~~~~~~

 Performs sensitivity analysis on objective coefficients. ::

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


``subj []int32``
    Index of objective coefficients to analyze.
``leftpricej []float64``
    Left shadow prices for requested coefficients.
``rightpricej []float64``
    Right shadow prices for requested coefficients.
``leftrangej []float64``
    Left range for requested coefficients.
``rightrangej []float64``
    Right range for requested coefficients.


Calculates sensitivity information for objective coefficients. The indexes of the coefficients to analyze are

.. math:: \{\mathtt{subj}[i] | i \in 0,\ldots,\mathtt{numj}-1\}

The results are returned so that e.g :math:`\mathtt{leftprice}[j]` is the left
shadow price of the objective coefficient with index :math:`\mathtt{subj}[j]`.

The type of sensitivity analysis to perform (basis or optimal partition) is controlled by the parameter ``IPAR_SENSITIVITY_TYPE``.

For an example, please see Section :ref:`doc.shared.sensitivity_example`.




EchoIntro
~~~~~~~~~

Prints an intro to message stream. ::

    func (*Env) EchoIntro ( longver int32 )

``longver int32``
    If non-zero, then the intro is slightly longer.

Prints an intro to message stream. 


Gemm
~~~~

Performs a dense matrix multiplication.::

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


``transa int32``
    Indicates whether the matrix A must be transposed.
``transb int32``
    Indicates whether the matrix B must be transposed.
``m int32``
    Indicates the number of rows of matrices A and C.
``n int32``
    Indicates the number of columns of matrices B and C.
``k int32``
    Specifies the number of columns of the matrix A and the number of rows of the matrix B.
``alpha float64``
    A scalar value multipling the result of the matrix multiplication.
``a []float64``
    The pointer to the array storing matrix A in a column-major format.
``b []float64``
    Indicates the number of rows of matrix B and columns of matrix A.
``beta float64``
    A scalar value that multiplies C.
``c []float64``
    The pointer to the array storing matrix C in a column-major format.


Performs a matrix multiplication plus addition of dense matrices. Given
:math:`A`, :math:`B` and :math:`C` of compatible dimensions, this function
computes 

.. math:: C:= \alpha op(A)op(B) + \beta C

where :math:`\alpha,\beta` are two scalar values. The function :math:`op(X)`
return :math:`X` if transX is YES, or :math:`X^T` if set to NO. Dimensions of
:math:`A,b` must therefore match those of :math:`C`.

The result of this operation is stored in :math:`C`.                  



Gemv
~~~~

Computes dense matrix times a dense vector product.::

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


``transa int32``
    Indicates whether the matrix A must be transposed.
``m int32``
     Specifies the number of rows of the matrix A.
``n int32``
     Specifies the number of columns of the matrix A.
``alpha float64``
    A scalar value multipling the matrix A.
``a []float64``
    A pointer to the array storing matrix A in a column-major format.
``x []float64``
    A pointer to the array storing the vector x.
``beta float64``
    A scalar value multipling thevector y.
``y []float64``
    A pointer to the array storing the vector y.


Computes the multiplication of a scaled dense matrix times a dense vector product, plus a scaled dense vector. In formula

.. math:: y = \alpha A x + \beta y,

or if trans is set to transpose.yes

.. math:: y = \alpha A^T x + \beta y,

where :math:`\alpha,\beta` are scalar values. :math:`A` is an :math:`n\times m`
matrix, :math:`x\in \mathbf{R}}^m` and :math:`y\in \mathbf{R}}^n`.

Note that the result is stored overwriting :math:`y`.



GetACol
~~~~~~~

Obtains one column of the linear constraint matrix. ::

    func (*Task) GetACol
        ( j int32,
          subj []int32,
          valj []float64 )
        ( nzj int32,
          subj []int32,
          valj []float64 )


``j int32``
    Index of the column.
``subj []int32``
    Index of the non-zeros in the row obtained.
``valj []float64``
    Numerical values of the column obtained.
``nzj int32``
    Number of non-zeros in the column obtained.

Obtains one column of :math:`A` in a sparse format.  


GetAColNumNz
~~~~~~~~~~~~

 Obtains the number of non-zero elements in one column of the linear constraint matrix ::

    func (*Task) GetAColNumNz ( i int32 ) ( nzj int32 )


``i int32``
    Index of the column.
``nzj int32``
    Number of non-zeros in the j'th row or column of (A).

Obtains the number of non-zero elements in one column of :math:`A`.  


GetAColSliceTrip
~~~~~~~~~~~~~~~~

 Obtains a sequence of columns from the coefficient matrix in triplet format. ::

    func (*Task) GetAColSliceTrip
        ( first int32,
          last int32,
          subi []int32,
          subj []int32,
          val []float64 )
        ( subi []int32,
          subj []int32,
          val []float64 )


``first int32``
     Index of the first column in the sequence.
``last int32``
    Index of the last column in the sequence plus one.
``subi []int32``
    Constraint subscripts.
``subj []int32``
    Column subscripts.
``val []float64``
    Values.


Obtains a sequence of columns  from :math:`A` in a sparse triplet format.

.. only: c

   Define :math:`p^1` as

   .. math:: p^1 = \mathtt{maxnumnz-surp[0]}

   when the function is called and :math:`p^2` by

   .. math:: p^2 = \mathtt{maxnumnz-surp[0]},

   where ``surp[0]`` is the value upon termination. Using this notation then

   .. math:: \mathtt{val}[k] = a_{\mathtt{subi[k]},\mathtt{subj[k]}}, \quad k=p^1,\ldots,p^2-1.





GetAPieceNumNz
~~~~~~~~~~~~~~

 Obtains the number non-zeros in a rectangular piece of the linear constraint matrix. ::

    func (*Task) GetAPieceNumNz
        ( firsti int32,
          lasti int32,
          firstj int32,
          lastj int32 )
        ( numnz int32 )


``firsti int32``
    Index of the first row in the rectangular piece.
``lasti int32``
     Index of the last row plus one in the rectangular piece.
``firstj int32``
     Index of the first column in the rectangular piece.
``lastj int32``
     Index of the last column plus one in the rectangular piece.
``numnz int32``
    Number of non-zero elements in the rectangular piece of the linear constraint matrix.


Obtains the number non-zeros in a rectangular piece of :math:`A`, i.e. the number

.. math:: \left| (i,j): ~ a_{i,j} \neq 0,~ \mathtt{firsti} \leq i \leq \mathtt{lasti}-1, ~\mathtt{firstj} \leq j \leq \mathtt{lastj}-1\} \right|

where :math:`|\mathcal{I}|` means the number of elements in the set :math:`\mathcal{I}`.

This function is not an efficient way to obtain the number of non-zeros in one
row or column. In that case use the function ``getarownumnz`` or ``getacolnumnz``.



GetARow
~~~~~~~

Obtains one row of the linear constraint matrix. ::

    func (*Task) GetARow
        ( i int32,
          subi []int32,
          vali []float64 )
        ( nzi int32,
          subi []int32,
          vali []float64 )


``i int32``
    Index of the row or column.
``subi []int32``
    Index of the non-zeros in the row obtained.
``vali []float64``
    Numerical values of the row obtained.
``nzi int32``
    Number of non-zeros in the row obtained.

Obtains one row of :math:`A` in a sparse format.  


GetARowNumNz
~~~~~~~~~~~~

 Obtains the number of non-zero elements in one row of the linear constraint matrix ::

    func (*Task) GetARowNumNz ( i int32 ) ( nzi int32 )


``i int32``
    Index of the row or column.
``nzi int32``
    Number of non-zeros in the i'th row of `A`.

Obtains the number of non-zero elements in one row of :math:`A`.  


GetARowSliceTrip
~~~~~~~~~~~~~~~~

 Obtains a sequence of rows from the coefficient matrix in triplet format. ::

    func (*Task) GetARowSliceTrip
        ( first int32,
          last int32,
          subi []int32,
          subj []int32,
          val []float64 )
        ( subi []int32,
          subj []int32,
          val []float64 )


``first int32``
     Index of the first row or column in the sequence.
``last int32``
    Index of the last row or column in the sequence plus one.
``subi []int32``
    Constraint subscripts.
``subj []int32``
    Column subscripts.
``val []float64``
    Values.


Obtains a sequence of rows  from :math:`A` in a sparse triplets format.

.. only: c

   Define :math:`p^1` as

   .. math:: p^1 = \texttt{maxnumnz-surp[0]}

   when the function is called and :math:`p^2` by

   .. math:: p^2 = \mathtt{maxnumnz-surp[0]}

   where ``surp[0]`` is the value upon termination. Using this notation then

   .. math:: \mathtt{val}[k] = a_{\mathtt{subi[k]},\mathtt{subj[k]}}, \quad k=p^1,\ldots,p^2-1.





GetASlice
~~~~~~~~~

Obtains a sequence of rows or columns from the coefficient matrix. ::

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


``accmode int32``
     Defines whether a column slice or a row slice is requested.
``first int32``
     Index of the first row or column in the sequence.
``last int32``
     Index of the last row or column in the sequence plus one.
``ptrb []int64``
     Row or column start pointers.
``ptre []int64``
     Row or column end pointers.
``sub []int32``
    Contains the row or column subscripts.
``val []float64``
    Contains the coefficient values.

Obtains a sequence of rows or columns from :math:`A` in sparse format.  


GetASliceNumNz
~~~~~~~~~~~~~~

 Obtains the number of non-zeros in a slice of rows or columns of the coefficient matrix. ::

    func (*Task) GetASliceNumNz
        ( accmode int32,
          first int32,
          last int32 )
        ( numnz int64 )


``accmode int32``
     Defines whether non-zeros are counted in a column slice or a row slice.
``first int32``
     Index of the first row or column in the sequence.
``last int32``
    Index of the last row or column plus one in the sequence.
``numnz int64``
    Number of non-zeros in the slice.

Obtains the number of non-zeros in a slice of rows or columns of :math:`A`.  


GetAij
~~~~~~

Obtains a single coefficient in linear constraint matrix. ::

    func (*Task) GetAij ( i int32, j int32 ) ( aij float64 )


``i int32``
    Row index of the coefficient to be returned.
``j int32``
    Column index of the coefficient to be returned.
``aij float64``
    Returns the requested coefficient.

Obtains a single coefficient in :math:`A`. 


GetBaraBlockTriplet
~~~~~~~~~~~~~~~~~~~

Obtains barA in block triplet form. ::

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


``subi []int32``
     Constraint index.
``subj []int32``
     Symmetric matrix variable index.
``subk []int32``
     Block row index.
``subl []int32``
     Block column index.
``valijkl []float64``
    A list indexes of the elements from symmetric matrix storage that appears in the weighted sum.
``num int64``
     Number of elements in the block triplet form.

Obtains :math:`\bar{A}` in block triplet form.  


GetBaraIdx
~~~~~~~~~~

Obtains information about an element barA. ::

    func (*Task) GetBaraIdx
        ( idx int64,
          sub []int64,
          weights []float64 )
        ( i int32,
          j int32,
          num int64,
          sub []int64,
          weights []float64 )


``idx int64``
     Position of the element in the vectorized form.
``sub []int64``
     A list indexes   of the elements from symmetric matrix storage that appears in the weighted sum.
``weights []float64``
     The weights associated with each term in the weighted sum.
``i int32``
     Row index of the element at position idx.
``j int32``
    Column index of the element at position idx.
``num int64``
    Number of terms in weighted sum that forms the element.


Obtains information about an element in :math:`\bar{A}`. Since :math:`\bar{A}`
is a sparse matrix of symmetric matrixes then only the nonzero elements in
:math:`\bar{A}` are stored in order to save space. Now :math:`\bar{A}` is
stored vectorized form i.e. as one long vector.  This function makes it
possible to obtain information such as the row index and the column index of a
particular element of the vectorized form of :math:`\bar{A}`.

Please observe if one element of :math:`\bar{A}` is inputted multiple times
then it may be stored several times in vectorized form. In that case the
element with the highest index is the one that is used.



GetBaraIdxIJ
~~~~~~~~~~~~

Obtains information about an element barA. ::

    func (*Task) GetBaraIdxIJ ( idx int64 ) ( i int32, j int32 )


``idx int64``
     Position of the element in the vectorized form.
``i int32``
     Row index of the element at position idx.
``j int32``
     Column index of the element at position idx.


Obtains information about an element in :math:`\bar{A}`. Since :math:`\bar{A}`
is a sparse matrix of symmetric matrixes only the nonzero elements in
:math:`\bar{A}` are stored in order to save space. Now :math:`\bar{A}` is
stored vectorized form i.e. as one long vector.  This function makes it
possible to obtain information such as the row index and the column index of a
particular element of the vectorized form of :math:`\bar{A}`.

Please note that if one element of :math:`\bar{A}` is inputted multiple times
then it may be stored several times in vectorized form. In that case the
element with the highest index is the one that is used.



GetBaraIdxInfo
~~~~~~~~~~~~~~

Obtains the number terms in the weighted sum that forms a particular element in barA. ::

    func (*Task) GetBaraIdxInfo ( idx int64 ) ( num int64 )


``idx int64``
     The internal position of the element that should be obtained information for.
``num int64``
     Number of terms in the weighted sum that forms the specified element in barA.


Each nonzero element in :math:`\bar{A}_{ij}` is formed as a weighted sum of
symmetric matrices. Using this function the number terms in the weighted sum
can be obtained. See description of ``appendsparsesymmat`` for details
about the weighted sum.  



GetBaraSparsity
~~~~~~~~~~~~~~~

Obtains the sparsity pattern of the barA matrix. ::

    func (*Task) GetBaraSparsity ( idxij []int64 ) ( numnz int64, idxij []int64 )


``idxij []int64``
    Position of each nonzero element in the vector representation of barA.
``numnz int64``
    Number of nonzero elements in barA.


The matrix :math:`\bar{A}` is assumed to be a sparse matrix of symmetric matrices.
This implies that many of elements in :math:`\bar{A}` is likely to be zero matrixes.
Therefore, in order to save space only nonzero elements in :math:`\bar{A}` are stored
on vectorized form. This function is used to obtain the sparsity pattern of
:math:`\bar{A}` and the position of each nonzero element in the vectorized form of
:math:`\bar{A}`.



GetBarcBlockTriplet
~~~~~~~~~~~~~~~~~~~

Obtains barc in block triplet form. ::

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


``subj []int32``
     Symmetric matrix variable index.
``subk []int32``
     Block row index.
``subl []int32``
     Block column index.
``valijkl []float64``
     A list indexes of the elements from symmetric matrix storage that appears in the weighted sum.
``num int64``
     Number of elements in the block triplet form.

Obtains :math:`\bar{C}` in block triplet form.  


GetBarcIdx
~~~~~~~~~~

Obtains information about an element in barc.::

    func (*Task) GetBarcIdx
        ( idx int64,
          sub []int64,
          weights []float64 )
        ( j int32,
          num int64,
          sub []int64,
          weights []float64 )


``idx int64``
     Index of the element that should be obtained information about.
``sub []int64``
     Elements appearing the weighted sum.
``weights []float64``
     Weights of terms in the weighted sum.
``j int32``
     Row index in barc.
``num int64``
     Number of terms in the weighted sum.

Obtains information about an element in :math:`\bar{c}`.  


GetBarcIdxInfo
~~~~~~~~~~~~~~

Obtains information about an element in barc. ::

    func (*Task) GetBarcIdxInfo ( idx int64 ) ( num int64 )


``idx int64``
     Index of element that should be obtained information about. The value is an index of a symmetric sparse variable.
``num int64``
     Number of terms that appears in weighted that forms the requested element.

Obtains information about the :math:`\bar{c}_{ij}`.  


GetBarcIdxJ
~~~~~~~~~~~

Obtains the row index of an element in barc. ::

    func (*Task) GetBarcIdxJ ( idx int64 ) ( j int32 )


``idx int64``
     Index of the element that should be obtained information about.
``j int32``
     Row index in barc.

Obtains the row index of an element in :math:`\bar{c}`.  


GetBarcSparsity
~~~~~~~~~~~~~~~

Get the positions of the nonzero elements in barc. ::

    func (*Task) GetBarcSparsity ( idxj []int64 ) ( numnz int64, idxj []int64 )


``idxj []int64``
    Internal positions of the nonzeros elements in barc.
``numnz int64``
    Number of nonzero elements in barc.


Internally only the nonzero elements of :math:`\bar{c}` is stored 

in a vector. This function returns which elements :math:`\bar{c}` that are
nonzero (in ``subj``) and their internal position (in ``idx``). Using the
position detailed information about each nonzero :math:`\bar{C}_j` can be
obtained using ``getbarcidxinfo`` and ``getbarcidx``.



GetBarsJ
~~~~~~~~

 Obtains the dual solution for a semidefinite variable. ::

    func (*Task) GetBarsJ
        ( whichsol int32,
          j int32,
          barsj []float64 )
        ( barsj []float64 )


``j int32``
    Index of the semidefinite variable.
``barsj []float64``
    Value of the j'th variable of barx.

Obtains the dual solution for a semidefinite variable. Only the lower triangle part of :math:`\bar{s}_j` is returned because the matrix by construction is symmetric. The format is that the columns are stored sequentially in the natural order.  


GetBarvarName
~~~~~~~~~~~~~

 Obtains a name of a semidefinite variable. ::

    func (*Task) GetBarvarName ( i int32 ) ( name string )


``i int32``
    Index.
``name string``
    The requested name is copied to this buffer.


Obtains a name of a semidefinite variable.



GetBarvarNameIndex
~~~~~~~~~~~~~~~~~~

 Obtains the index of name of semidefinite variable. ::

    func (*Task) GetBarvarNameIndex ( somename string ) ( asgn int32, index int32 )


``somename string``
    The requested name is copied to this buffer.
``asgn int32``
    Is non-zero if name somename is assigned to a semidefinite variable.
``index int32``
     If the name somename is assigned to a semidefinite variable, then index is the name of the constraint.


Obtains the index of name of semidefinite variable.



GetBarvarNameLen
~~~~~~~~~~~~~~~~

 Obtains the length of a name of a semidefinite variable. ::

    func (*Task) GetBarvarNameLen ( i int32 ) ( len int32 )


``i int32``
    Index.
``len int32``
    Returns the length of the indicated name.


Obtains the length of a name of a semidefinite variable.



GetBarxJ
~~~~~~~~

 Obtains the primal solution for a semidefinite variable. ::

    func (*Task) GetBarxJ
        ( whichsol int32,
          j int32,
          barxj []float64 )
        ( barxj []float64 )


``j int32``
    Index of the semidefinite variable.
``barxj []float64``
    Value of the j'th variable of barx.

Obtains the primal solution for a semidefinite variable. Only the lower triangle part of :math:`\bar{x}_j` is returned because the matrix by construction is symmetric. The format is that the columns are stored sequentially in the natural order.  


GetBound
~~~~~~~~

 Obtains bound information for one constraint or variable. ::

    func (*Task) GetBound
        ( accmode int32,
          i int32 )
        ( bk int32,
          bl float64,
          bu float64 )


``i int32``
     Index of the constraint or variable for which the bound information should be obtained.


Obtains bound information for one constraint or variable.



GetBoundSlice
~~~~~~~~~~~~~

 Obtains bounds information for a sequence of variables or constraints. ::

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

Obtains all objective coefficients. ::

    func (*Task) GetC ( c []float64 ) ( c []float64 )



Obtains all objective coefficients :math:`c`. 


GetCJ
~~~~~

Obtains one coefficient of c. ::

    func (*Task) GetCJ ( j int32 ) ( cj float64 )


``j int32``
    Index of the variable for which c coefficient should be obtained.
``cj float64``
    The c coefficient value.

Obtains one coefficient of :math:`c`.  


GetCSlice
~~~~~~~~~

Obtains a sequence of coefficients from the objective. ::

    func (*Task) GetCSlice
        ( first int32,
          last int32,
          c []float64 )
        ( c []float64 )



Obtains a sequence of elements in :math:`c`. 


GetCfix
~~~~~~~

Obtains the fixed term in the objective. ::

    func (*Task) GetCfix (  ) ( cfix float64 )



Obtains the fixed term in the objective. 


GetCodeDesc
~~~~~~~~~~~

Obtains a short description of a response code. ::

    func GetCodeDesc
        ( code int32 )
        ( symname string,
          str string,
          res int32 )


``code int32``
    A valid response code.
``symname string``
    Symbolic name corresponding to the code.
``str string``
    Obtains a short description of a response code.
``res int32``
    Response code (see `mosek.RES_...`

Obtains a short description of the meaning of the response code given by ``code``.  


GetConBound
~~~~~~~~~~~

 Obtains bound information for one constraint. ::

    func (*Task) GetConBound
        ( i int32 )
        ( bk int32,
          bl float64,
          bu float64 )


``i int32``
     Index of the constraint for which the bound information should be obtained.


Obtains bound information for one constraint.



GetConBoundSlice
~~~~~~~~~~~~~~~~

 Obtains bounds information for a slice of the constraints. ::

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

 Obtains a name of a constraint. ::

    func (*Task) GetConName ( i int32 ) ( name string )


``i int32``
    Index.
``name string``
    Is assigned the required name.


Obtains a name of a constraint.



GetConNameIndex
~~~~~~~~~~~~~~~

 Checks whether the name somename has been assigned  to any constraint. ::

    func (*Task) GetConNameIndex ( somename string ) ( asgn int32, index int32 )


``somename string``
    The name which should be checked.
``asgn int32``
    Is non-zero if name somename is assigned to a constraint.
``index int32``
     If the name somename is assigned to a constraint, then index is the name of the constraint.


Checks whether the name ``somename`` has been assigned to any constraint. If it has been assigned to constraint, then index of the constraint is reported.



GetConNameLen
~~~~~~~~~~~~~

 Obtains the length of a name of a constraint variable. ::

    func (*Task) GetConNameLen ( i int32 ) ( len int32 )


``i int32``
    Index.
``len int32``
    Returns the length of the indicated name.


Obtains the length of a name of a constraint variable.



GetCone
~~~~~~~

Obtains a conic constraint. ::

    func (*Task) GetCone
        ( k int32,
          submem []int32 )
        ( ct int32,
          conepar float64,
          nummem int32,
          submem []int32 )


``k int32``
    Index of the cone constraint.

Obtains a conic constraint. 


GetConeInfo
~~~~~~~~~~~

Obtains information about a conic constraint. ::

    func (*Task) GetConeInfo
        ( k int32 )
        ( ct int32,
          conepar float64,
          nummem int32 )


``k int32``
    Index of the conic constraint.

Obtains information about a conic constraint. 


GetConeName
~~~~~~~~~~~

 Obtains a name of a cone. ::

    func (*Task) GetConeName ( i int32 ) ( name string )


``i int32``
    Index.
``name string``
    Is assigned the required name.


Obtains a name of a cone.



GetConeNameIndex
~~~~~~~~~~~~~~~~

 Checks whether the name somename has been assigned  to any cone. ::

    func (*Task) GetConeNameIndex ( somename string ) ( asgn int32, index int32 )


``somename string``
    The name which should be checked.
``asgn int32``
    Is non-zero if name somename is assigned to a cone.
``index int32``
    If the name somename is assigned to a cone, then index is the name of the cone.

Checks whether the name ``somename`` has been assigned  to any cone. If it has been assigned to cone, then index of the cone is reported.  


GetConeNameLen
~~~~~~~~~~~~~~

 Obtains the length of a name of a cone. ::

    func (*Task) GetConeNameLen ( i int32 ) ( len int32 )


``i int32``
    Index.
``len int32``
    Returns the length of the indicated name.


Obtains the length of a name of a cone.



GetDimBarvarJ
~~~~~~~~~~~~~

Obtains the dimension of a symmetric matrix variable.::

    func (*Task) GetDimBarvarJ ( j int32 ) ( dimbarvarj int32 )


``j int32``
    Index of the semidefinite variable whose dimension is requested.
``dimbarvarj int32``
    The dimension of the j'th semidefinite variable.

Obtains the dimension of a symmetric matrix variable.


GetDouInf
~~~~~~~~~

Obtains a double information item. ::

    func (*Task) GetDouInf ( whichdinf int32 ) ( dvalue float64 )


``dvalue float64``
     The value of the required double information item.


Obtains a double information item from the task information database.



GetDouParam
~~~~~~~~~~~

Obtains a double parameter. ::

    func (*Task) GetDouParam ( param int32 ) ( parvalue float64 )



Obtains the value of a double parameter. 


GetDualObj
~~~~~~~~~~

Computes the dual objective value associated with the solution. ::

    func (*Task) GetDualObj ( whichsol int32 ) ( dualobj float64 )



Computes the dual objective value associated with the solution. Note if the solution is a primal infeasibility certificate, then the fixed term in the objective value is not included. 
                      Moreover, since there is no dual solution associated with integer solution, then an error will be reported if the dual objective value is requested for the integer solution. 


GetDualSolutionNorms
~~~~~~~~~~~~~~~~~~~~

Compute norms of the primal solution. ::

    func (*Task) GetDualSolutionNorms
        ( whichsol int32 )
        ( nrmy float64,
          nrmslc float64,
          nrmsuc float64,
          nrmslx float64,
          nrmsux float64,
          nrmsnx float64,
          nrmbars float64 )


``nrmy float64``
     The norm of the y vector.
``nrmslc float64``
     The norm of the slc vector.
``nrmsuc float64``
     The norm of the suc vector.
``nrmslx float64``
     The norm of the slx vector.
``nrmsux float64``
     The norm of the sux vector.
``nrmsnx float64``
     The norm of the snx vector.
``nrmbars float64``
     The norm of the bars vector.

Compute norms of the primal solution.
               


GetDviolBarvar
~~~~~~~~~~~~~~

Computes the violation of dual solution for a set of barx variables. ::

    func (*Task) GetDviolBarvar
        ( whichsol int32,
          sub []int32,
          viol []float64 )
        ( viol []float64 )


``sub []int32``
    An array of indexes of barx variables.
``viol []float64``
    List of violations corresponding to sub.


Let :math:`(\bar{S}_j)^*` be the value of variable :math:`\bar{S}_j` for the
specified solution.  Then the dual violation of the solution associated with
variable :math:`\bar{S}_j` is given by

.. math:: \max(-\lambda_{\min}(\bar{S}_j),0.0).

Both when the solution is a certificate of primal infeasibility or when it is
dual feasibleness solution the violation should be small.



GetDviolCon
~~~~~~~~~~~

Computes the violation of a dual solution associated with a set of constraints. ::

    func (*Task) GetDviolCon
        ( whichsol int32,
          sub []int32,
          viol []float64 )
        ( viol []float64 )


``sub []int32``
    An array of indexes of constraints.
``viol []float64``
    List of violations corresponding to sub.


The violation of the dual solution associated with the :math:`i`\'th constraint
is computed as follows

.. math:: \max( \rho( (s_l^c)_i^*,(b_l^c)_i ), \rho( (s_u^c)_i^*, -(b_u^c)_i) , |-y_i+(s_l^c)_i^*-(s_u^c)_i^*| )

where

.. math::

    \rho(x,l) =
      \left\{
        \begin{array}{ll}
           -x,   & l > -\infty , \\
           |x|, &  \mbox{otherwise}\\
        \end{array}
      \right.
 
Both when the solution is a certificate of primal infeasibility or it is a dual feasibleness solution the violation should be small.                 



GetDviolCones
~~~~~~~~~~~~~

Computes the violation of a solution for set of dual conic constraints. ::

    func (*Task) GetDviolCones
        ( whichsol int32,
          sub []int32,
          viol []float64 )
        ( viol []float64 )


``sub []int32``
    An array of indexes of barx variables.
``viol []float64``
    List of violations corresponding to sub.


Let :math:`(s_n^x)^*` be the value of variable :math:`(s_n^x)` for the
specified solution. For simplicity let us assume that :math:`s_n^x` is a member
of quadratic cone, then the violation is computed as follows

.. math::
    
    \left\{
      \begin{array}{ll}
        \max(0,\|(s_n^x\|_{2;n}^*-(s_n^x)_1^*) / \sqrt{2}, & (s_n^x)^* \geq -\|(s_n^x)_{2:n}^*\|, \\
        \|(s_n^x)^*\|, & \mbox{otherwise.}
      \end{array}
    \right.

Both when the solution is a certificate of primal infeasibility or when it is a
dual feasibleness solution the violation should be small.



GetDviolVar
~~~~~~~~~~~

Computes the violation of a dual solution associated with a set of x variables. ::

    func (*Task) GetDviolVar
        ( whichsol int32,
          sub []int32,
          viol []float64 )
        ( viol []float64 )


``sub []int32``
    An array of indexes of x variables.
``viol []float64``
    List of violations corresponding to sub.


The violation of the dual solution associated with the :math:`j`'th variable is
computed as follows

.. math:: \max \left(\rho((s_l^x)_i^*,(b_l^x)_i),\rho((s_u^x)_i^*,-(b_u^x)_i),|\sum_{j=\idxbeg}^{\idxend{numcon}} a_{ij} y_i+(s_l^x)_i^*-(s_u^x)_i^* - \tau c_j| \right)

where

.. math::

  \rho(x,l) =
    \left\{
      \begin{array}{ll}
         -x,   & l > -\infty , \\
         |x|, &  \mbox{otherwise}
      \end{array}
    \right.


:math:`\tau=0` if the solution is certificate of dual infeasibility and :math:`\tau=1` otherwise. The formula for computing the violation is only shown
for linear case but is generalized appropriately for the more general problems.



GetInfIndex
~~~~~~~~~~~

Obtains the index of a named information item. ::

    func (*Task) GetInfIndex ( inftype int32, infname string ) ( infindex int32 )


``infindex int32``
    The item index.

Obtains the index of a named information item. 


GetInfMax
~~~~~~~~~

 Obtains the maximum index of an information of a given type inftype plus 1. ::

    func (*Task) GetInfMax ( inftype int32, infmax []int32 ) ( infmax []int32 )


``infmax []int32``
    The maximum index requested.

Obtains the maximum index of an information of a given type ``inftype`` plus 1.  


GetInfName
~~~~~~~~~~

Obtains the name of an information item. ::

    func (*Task) GetInfName ( inftype int32, whichinf int32 ) ( infname string )



Obtains the name of an information item. 


GetInfeasibleSubProblem
~~~~~~~~~~~~~~~~~~~~~~~

Obtains an infeasible sub problem. ::

    func (*Task) GetInfeasibleSubProblem ( whichsol int32 ) ( inftask Task )


``whichsol int32``
     Which solution to use when determining the infeasible subproblem.
``inftask Task``
     A new task containing the infeasible subproblem.


Given the solution is a certificate of primal or dual infeasibility then a
primal or dual infeasible subproblem is obtained respectively.  The subproblem
tend to be much smaller than the original problem and hence it easier to locate
the infeasibility inspecting the subproblem than the original problem.

For the procedure to be useful then it is important to assigning meaningful
names to constraints, variables etc. in the original task because those names
will be duplicated in the subproblem.

The function is only applicable to linear and conic quadratic optimization
problems.

For more information see Section :ref:`doc.shared.feas_repair`.



GetIntInf
~~~~~~~~~

Obtains an integer information item. ::

    func (*Task) GetIntInf ( whichiinf int32 ) ( ivalue int32 )


``ivalue int32``
     The value of the required integer information item.


Obtains an integer information item from the task information database.



GetIntParam
~~~~~~~~~~~

Obtains an integer parameter. ::

    func (*Task) GetIntParam ( param int32 ) ( parvalue int32 )



Obtains the value of an integer parameter. 


GetLenBarvarJ
~~~~~~~~~~~~~

Obtains the length if the j'th semidefinite variables. ::

    func (*Task) GetLenBarvarJ ( j int32 ) ( lenbarvarj int64 )


``j int32``
    Index of the semidefinite variable whose length if requested.
``lenbarvarj int64``
    Number of scalar elements in the lower triangular part of the semidefinite variable.

Obtains the length of the :math:`j`\ th semidefinite variable i.e. the number of elements in the triangular part. 


GetLintInf
~~~~~~~~~~

Obtains an integer information item. ::

    func (*Task) GetLintInf ( whichliinf int32 ) ( ivalue int64 )


``ivalue int64``
     The value of the required integer information item.


Obtains an integer information item from the task information database.



GetMaxNumANz
~~~~~~~~~~~~

 Obtains number of preallocated non-zeros in the linear constraint matrix. ::

    func (*Task) GetMaxNumANz (  ) ( maxnumanz int64 )



Obtains number of preallocated non-zeros in :math:`A`. When this number of non-zeros is reached MOSEK will automatically allocate more space for :math:`A`.  


GetMaxNumBarvar
~~~~~~~~~~~~~~~

Obtains the number of semidefinite variables. ::

    func (*Task) GetMaxNumBarvar (  ) ( maxnumbarvar int32 )


``maxnumbarvar int32``
    Obtains maximum number of semidefinite variable currently allowed.

Obtains the number of semidefinite variables. 


GetMaxNumCon
~~~~~~~~~~~~

Obtains the number of preallocated constraints in the optimization task. ::

    func (*Task) GetMaxNumCon (  ) ( maxnumcon int32 )



Obtains the number of preallocated constraints in the optimization task. When this number of constraints is reached MOSEK will automatically allocate more space for constraints.  


GetMaxNumCone
~~~~~~~~~~~~~

Obtains the number of preallocated cones in the optimization task. ::

    func (*Task) GetMaxNumCone (  ) ( maxnumcone int32 )




Obtains the number of preallocated cones in the optimization task. When this
number of cones is reached MOSEK will automatically allocate space for more
cones.



GetMaxNumQNz
~~~~~~~~~~~~

 Obtains the number of preallocated non-zeros for all quadratic terms in objective and constraints. ::

    func (*Task) GetMaxNumQNz (  ) ( maxnumqnz int64 )




Obtains the number of preallocated non-zeros for :math:`Q` (both objective and
constraints). When this number of non-zeros is reached MOSEK will
automatically allocate more space for :math:`Q`.



GetMaxNumVar
~~~~~~~~~~~~

Obtains the maximum number variables allowed. ::

    func (*Task) GetMaxNumVar (  ) ( maxnumvar int32 )



Obtains the number of preallocated variables in the optimization task. When this number of variables is reached MOSEK will automatically allocate more space for constraints.  


GetMemUsage
~~~~~~~~~~~

Obtains information about the amount of memory used by a task. ::

    func (*Task) GetMemUsage (  ) ( meminuse int64, maxmemuse int64 )


``meminuse int64``
    Amount of memory currently used by the task.
``maxmemuse int64``
    Maximum amount of memory used by the task until now.

Obtains information about the amount of memory used by a task. 


GetNumANz
~~~~~~~~~

Obtains the number of non-zeros in the coefficient matrix. ::

    func (*Task) GetNumANz (  ) ( numanz int32 )



Obtains the number of non-zeros in :math:`A`. 


GetNumANz64
~~~~~~~~~~~

Obtains the number of non-zeros in the coefficient matrix. ::

    func (*Task) GetNumANz64 (  ) ( numanz int64 )



Obtains the number of non-zeros in :math:`A`. 


GetNumBaraBlockTriplets
~~~~~~~~~~~~~~~~~~~~~~~

Obtains an upper bound on the number of scalar elements in the block triplet form of bara. ::

    func (*Task) GetNumBaraBlockTriplets (  ) ( num int64 )


``num int64``
     Number elements in the block triplet form of bara.

Obtains an upper bound on the number of elements in the block triplet form of :math:`\bar{A}`.  


GetNumBaraNz
~~~~~~~~~~~~

Get the number of nonzero elements in barA. ::

    func (*Task) GetNumBaraNz (  ) ( nz int64 )


``nz int64``
    The number of nonzero block elements in barA.

Get the number of nonzero elements in :math:`\bar{A}`.  


GetNumBarcBlockTriplets
~~~~~~~~~~~~~~~~~~~~~~~

Obtains an upper bound on the number of elements in the block triplet form of barc. ::

    func (*Task) GetNumBarcBlockTriplets (  ) ( num int64 )


``num int64``
     An upper bound on the number elements in the block trip let form of barc.

Obtains an upper bound on the number of elements in the block triplet form of :math:`\bar{C}`.  


GetNumBarcNz
~~~~~~~~~~~~

Obtains the number of nonzero elements in barc.::

    func (*Task) GetNumBarcNz (  ) ( nz int64 )


``nz int64``
    The number of nonzero elements in barc.

Obtains the number of nonzero elements in :math:`\bar{c}`.  


GetNumBarvar
~~~~~~~~~~~~

Obtains the number of semidefinite variables. ::

    func (*Task) GetNumBarvar (  ) ( numbarvar int32 )


``numbarvar int32``
    Number of semidefinite variable in the problem.

Obtains the number of semidefinite variables. 


GetNumCon
~~~~~~~~~

Obtains the number of constraints. ::

    func (*Task) GetNumCon (  ) ( numcon int32 )



Obtains the number of constraints. 


GetNumCone
~~~~~~~~~~

Obtains the number of cones. ::

    func (*Task) GetNumCone (  ) ( numcone int32 )


``numcone int32``
    Number conic constraints.

Obtains the number of cones. 


GetNumConeMem
~~~~~~~~~~~~~

Obtains the number of members in a cone. ::

    func (*Task) GetNumConeMem ( k int32 ) ( nummem int32 )


``k int32``
    Index of the cone.

Obtains the number of members in a cone. 


GetNumIntVar
~~~~~~~~~~~~

 Obtains the number of integer-constrained variables. ::

    func (*Task) GetNumIntVar (  ) ( numintvar int32 )


``numintvar int32``
    Number of integer variables.


Obtains the number of integer-constrained variables.



GetNumParam
~~~~~~~~~~~

Obtains the number of parameters of a given type. ::

    func (*Task) GetNumParam ( partype int32 ) ( numparam int32 )


``numparam int32``
    Returns the number of parameters of the requested type.

Obtains the number of parameters of a given type. 


GetNumQConKNz
~~~~~~~~~~~~~

 Obtains the number of non-zero quadratic terms in a constraint. ::

    func (*Task) GetNumQConKNz ( k int32 ) ( numqcnz int64 )


``k int32``
     Index of the constraint for which the number quadratic terms should be obtained.


Obtains the number of non-zero quadratic terms in a constraint.



GetNumQObjNz
~~~~~~~~~~~~

 Obtains the number of non-zero quadratic terms in the objective. ::

    func (*Task) GetNumQObjNz (  ) ( numqonz int64 )




Obtains the number of non-zero quadratic terms in the objective.



GetNumSymMat
~~~~~~~~~~~~

Get the number of symmetric matrixes stored. ::

    func (*Task) GetNumSymMat (  ) ( num int64 )


``num int64``
     Returns the number of symmetric sparse matrixes.

Get the number of symmetric matrixes stored in the vector :math:`E`.  


GetNumVar
~~~~~~~~~

Obtains the number of variables. ::

    func (*Task) GetNumVar (  ) ( numvar int32 )



Obtains the number of variables. 


GetObjName
~~~~~~~~~~

 Obtains the name assigned to the objective function. ::

    func (*Task) GetObjName (  ) ( objname string )


``objname string``
    Assigned the objective name.


Obtains the name assigned to the objective function.



GetObjNameLen
~~~~~~~~~~~~~

 Obtains the length of the name assigned to the objective function. ::

    func (*Task) GetObjNameLen (  ) ( len int32 )


``len int32``
    Assigned the length of the objective name.


Obtains the length of the name assigned to the objective function.



GetObjSense
~~~~~~~~~~~

Gets the objective sense. ::

    func (*Task) GetObjSense (  ) ( sense int32 )


``sense int32``
     The returned objective sense.


Gets the objective sense of the task.



GetParamMax
~~~~~~~~~~~

 Obtains the maximum index of a parameter of a given type plus 1. ::

    func (*Task) GetParamMax ( partype int32 ) ( parammax int32 )


``parammax int32``
    The maximum index of the given parameter type.


Obtains the maximum index of a parameter of a given type plus 1.



GetParamName
~~~~~~~~~~~~

Obtains the name of a parameter. ::

    func (*Task) GetParamName ( partype int32, param int32 ) ( parname string )



Obtains the name for a parameter ``param`` of type ``partype``.  


GetPrimalObj
~~~~~~~~~~~~

 Computes the primal objective value for the desired solution. ::

    func (*Task) GetPrimalObj ( whichsol int32 ) ( primalobj float64 )




Computes the primal objective value for the desired solution. Note if the solution is an infeasibility certificate, then the fixed term in the objective is not included.



GetPrimalSolutionNorms
~~~~~~~~~~~~~~~~~~~~~~

Compute norms of the primal solution. ::

    func (*Task) GetPrimalSolutionNorms
        ( whichsol int32 )
        ( nrmxc float64,
          nrmxx float64,
          nrmbarx float64 )


``nrmxc float64``
     The norm of xc vector.
``nrmxx float64``
     The norm of xx vector.
``nrmbarx float64``
     The norm of barx vector.

Compute norms of the primal solution.
               


GetProSta
~~~~~~~~~

 Obtains the problem status. ::

    func (*Task) GetProSta ( whichsol int32 ) ( prosta int32 )




Obtains the problem status.



GetProbType
~~~~~~~~~~~

 Obtains the problem type. ::

    func (*Task) GetProbType (  ) ( probtype int32 )


``probtype int32``
     The problem type.


Obtains the problem type.



GetPviolBarvar
~~~~~~~~~~~~~~

Computes the violation of a primal solution for a list of barx variables. ::

    func (*Task) GetPviolBarvar
        ( whichsol int32,
          sub []int32,
          viol []float64 )
        ( viol []float64 )


``sub []int32``
    An array of indexes of barx variables.
``viol []float64``
    List of violations corresponding to sub.


Let :math:`(\bar{X}_j)^*` be the value of variable :math:`\bar{X}_j` for the
specified solution.  Then the primal violation of the solution associated with
variable :math:`\bar{X}_j` is given by

.. math:: \max(-\lambda_{\min}(\barX_j),0.0).



GetPviolCon
~~~~~~~~~~~

Computes the violation of a primal solution for a list of xc variables. ::

    func (*Task) GetPviolCon
        ( whichsol int32,
          sub []int32,
          viol []float64 )
        ( viol []float64 )


``sub []int32``
    An array of indexes of constraints.
``viol []float64``
    List of violations corresponding to sub.


The primal violation of the solution associated of constraint is computed by

.. math:: \max(l_i^c \tau - (x_i^c)^*),(x_i^c)^* \tau - u_i^c\tau,|\sum_{j=\idxbeg}^{\idxend{numvar}} a_{ij} x_j^* - x_i^c|)

where :math:`\tau` is defined as follows. If the solution is a certificate of
dual infeasibility, then :math:`\tau=0` and otherwise :math:`\tau=1`. Both when
the solution is a valid certificate of dual infeasibility or when it is primal
feasibleness solution the violation should be small. The above is only shown
for linear case but is appropriately generalized for the other cases.



GetPviolCones
~~~~~~~~~~~~~

Computes the violation of a solution for set of conic constraints. ::

    func (*Task) GetPviolCones
        ( whichsol int32,
          sub []int32,
          viol []float64 )
        ( viol []float64 )


``sub []int32``
    An array of indexes of barx variables.
``viol []float64``
    List of violations corresponding to sub.


Let :math:`x^*` be the value of variable :math:`x` for the specified solution.
For simplicity let us assume that :math:`x` is a member of quadratic cone, then
the violation is computed as follows

.. math::

  \left\{
    \begin{array}{ll}
      \max(0,\|x_{2;n}\|-x_1) / \sqrt{2}, & x_1 \geq -\|x_{2:n}\|, \\
      \|x\|, & \mbox{otherwise.}
    \end{array}
  \right.

Both when the solution is a certificate of dual infeasibility or when it is a
primal feasibleness solution the violation should be small.



GetPviolVar
~~~~~~~~~~~

Computes the violation of a primal solution for a list of x variables. ::

    func (*Task) GetPviolVar
        ( whichsol int32,
          sub []int32,
          viol []float64 )
        ( viol []float64 )


``sub []int32``
    An array of indexes of x variables.
``viol []float64``
    List of violations corresponding to sub.


Let :math:`x_j^*` be the value of variable :math:`x_j` for the specified
solution.  Then the primal violation of the solution associated with variable
:math:`x_j` is given by

.. math:: \max(l_j^x \tau - x_j^*,x_j^* - u_j^x\tau).

where :math:`\tau` is defined as follows. If the solution is a certificate of
dual infeasibility, then :math:`\tau=0` and otherwise :math:`\tau=1`. Both when
the solution is a valid certificate of dual infeasibility or when it is primal
feasibleness solution the violation should be small.



GetQConK
~~~~~~~~

Obtains all the quadratic terms in a constraint. ::

    func (*Task) GetQConK
        ( k int32,
          qcsubi []int32,
          qcsubj []int32,
          qcval []float64 )
        ( numqcnz int64,
          qcsubi []int32,
          qcsubj []int32,
          qcval []float64 )


``k int32``
    Which constraint.


Obtains all the quadratic terms in a constraint. The quadratic
terms are stored sequentially ``qcsubi``, ``qcsubj``, and ``qcval``.



GetQObj
~~~~~~~

Obtains all the quadratic terms in the objective. ::

    func (*Task) GetQObj
        ( qosubi []int32,
          qosubj []int32,
          qoval []float64 )
        ( numqonz int64,
          qosubi []int32,
          qosubj []int32,
          qoval []float64 )




Obtains the quadratic terms in the objective. The required quadratic terms
are stored sequentially in ``qosubi``, ``qosubj``, and ``qoval``.



GetQObjIJ
~~~~~~~~~

 Obtains one coefficient from the quadratic term of the objective ::

    func (*Task) GetQObjIJ ( i int32, j int32 ) ( qoij float64 )


``i int32``
    Row index of the coefficient.
``j int32``
    Column index of coefficient.
``qoij float64``
    The required coefficient.

Obtains one coefficient :math:`q_{ij}^o` in the quadratic term of the objective.  


GetReducedCosts
~~~~~~~~~~~~~~~

Obtains the difference of (slx-sux) for a sequence of variables. ::

    func (*Task) GetReducedCosts
        ( whichsol int32,
          first int32,
          last int32,
          redcosts []float64 )
        ( redcosts []float64 )


``first int32``
    See the documentation for a full description.
``last int32``
    See the documentation for a full description.
``redcosts []float64``
    Returns the requested reduced costs. See documentation for a full description.


Computes the reduced costs for a sequence of variables and return them in the variable ``redcosts`` i.e.

.. math::
    :label: ais-eq-redcost

    \mathtt{redcosts}[j-\mathtt{first}] = (s_l^x)_j-(s_u^x)_j, ~j=\mathtt{first},\ldots,\idxend{last}




GetSkc
~~~~~~

 Obtains the status keys for the constraints. ::

    func (*Task) GetSkc ( whichsol int32, skc []int32 ) ( skc []int32 )




Obtains the status keys for the constraints.



GetSkcSlice
~~~~~~~~~~~

 Obtains the status keys for the constraints. ::

    func (*Task) GetSkcSlice
        ( whichsol int32,
          first int32,
          last int32,
          skc []int32 )
        ( skc []int32 )




Obtains the status keys for the constraints.



GetSkx
~~~~~~

 Obtains the status keys for the scalar variables. ::

    func (*Task) GetSkx ( whichsol int32, skx []int32 ) ( skx []int32 )




Obtains the status keys for the scalar variables.



GetSkxSlice
~~~~~~~~~~~

 Obtains the status keys for the variables. ::

    func (*Task) GetSkxSlice
        ( whichsol int32,
          first int32,
          last int32,
          skx []int32 )
        ( skx []int32 )




Obtains the status keys for the variables.



GetSlc
~~~~~~

 Obtains the slc vector for a solution. ::

    func (*Task) GetSlc ( whichsol int32, slc []float64 ) ( slc []float64 )


``slc []float64``
    The slc vector.

Obtains the :math:`s_l^c` vector for a solution.  


GetSlcSlice
~~~~~~~~~~~

 Obtains a slice of the slc vector for a solution. ::

    func (*Task) GetSlcSlice
        ( whichsol int32,
          first int32,
          last int32,
          slc []float64 )
        ( slc []float64 )



Obtains a slice of the :math:`s_l^c` vector for a solution.  


GetSlx
~~~~~~

 Obtains the slx vector for a solution. ::

    func (*Task) GetSlx ( whichsol int32, slx []float64 ) ( slx []float64 )


``slx []float64``
    The slx vector.

Obtains the :math:`s_l^x` vector for a solution. 


GetSlxSlice
~~~~~~~~~~~

 Obtains a slice of the slx vector for a solution. ::

    func (*Task) GetSlxSlice
        ( whichsol int32,
          first int32,
          last int32,
          slx []float64 )
        ( slx []float64 )



Obtains a slice of the :math:`s_l^x` vector for a solution.  


GetSnx
~~~~~~

 Obtains the snx vector for a solution. ::

    func (*Task) GetSnx ( whichsol int32, snx []float64 ) ( snx []float64 )


``snx []float64``
    The snx vector.

Obtains the :math:`s_n^x` vector for a solution.  


GetSnxSlice
~~~~~~~~~~~

 Obtains a slice of the snx vector for a solution. ::

    func (*Task) GetSnxSlice
        ( whichsol int32,
          first int32,
          last int32,
          snx []float64 )
        ( snx []float64 )



Obtains a slice of the :math:`s_n^x` vector for a solution.  


GetSolSta
~~~~~~~~~

 Obtains the solution status. ::

    func (*Task) GetSolSta ( whichsol int32 ) ( solsta int32 )




Obtains the solution status.



GetSolution
~~~~~~~~~~~

Obtains the complete solution. ::

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

Consider the case of linear programming. The primal problem is given by

.. math::

   \begin{array}{lccccl}
     \mbox{minimize}              &      &      & c^T x+c^f &      &        \\
     \mbox{subject to} &  l^c & \leq & A x       & \leq & u^c,     \\
     &  l^x & \leq & x         & \leq & u^x.   \\
   \end{array}


and the corresponding dual problem is

.. math::

   \begin{array}{lccl}
     \mbox{maximize}   & (l^c)^T s_l^c - (u^c)^T s_u^c         &  \\
     & + (l^x)^T s_l^x - (u^x)^T s_u^x + c^f &  \\
     \mbox{subject to} & A^T y + s_l^x - s_u^x                 & = & c, \\
     & -y    + s_l^c - s_u^c                 & = & 0, \\
     & s_l^c,s_u^c,s_l^x,s_u^x \geq 0.       &   &    \\
   \end{array}


In this case the mapping between variables and arguments to the function is as
follows:
  
* ``xx`` : Corresponds to variable :math:`x`.
* ``y``  : Corresponds to variable :math:`y`.
* ``slc``: Corresponds to variable :math:`s_l^c`.
* ``suc``: Corresponds to variable :math:`s_u^c`.
* ``slx``: Corresponds to variable :math:`s_l^x`.
* ``sux``: Corresponds to variable :math:`s_u^x`.
* ``xc`` : Corresponds to :math:`Ax`.

The meaning of the values returned by this function depend on the *solution status* returned in the argument ``solsta``. The most important possible values  of ``solsta`` are:

* ``SOL_STA_OPTIMAL`` : An optimal solution satisfying the optimality criteria for continuous problems is returned.

* ``SOL_STA_INTEGER_OPTIMAL`` : An optimal solution satisfying the optimality criteria for integer problems is returned.

* ``SOL_STA_PRIM_FEAS`` : A solution satisfying the feasibility criteria.

* ``SOL_STA_PRIM_INFEAS_CER`` : A primal certificate of infeasibility is returned.

* ``SOL_STA_DUAL_INFEAS_CER`` : A dual certificate of infeasibility is returned.




GetSolutionI
~~~~~~~~~~~~

 Obtains the solution for a single constraint or variable. ::

    func (*Task) GetSolutionI
        ( accmode int32,
          i int32,
          whichsol int32 )
        ( sk int32,
          x float64,
          sl float64,
          su float64,
          sn float64 )


``accmode int32``
     Defines whether solution information for a constraint or for a variable is retrieved.
``i int32``
    Index of the constraint or variable.
``sk int32``
    Status key of the constraint of variable.
``x float64``
    Solution value of the primal variable.
``sl float64``
     Solution value of the dual variable associated with the lower bound.
``su float64``
     Solution value of the dual variable associated with the upper bound.
``sn float64``
     Solution value of the dual variable associated with the cone constraint.


Obtains the primal and dual solution information for a single constraint or variable.



GetSolutionInfo
~~~~~~~~~~~~~~~

Obtains information about of a solution. ::

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


``pobj float64``
     The primal objective value.
``pviolcon float64``
     Maximal primal bound violation for a xc variable.
``pviolvar float64``
     Maximal primal bound violation for a xx variable.
``pviolbarvar float64``
     Maximal primal bound violation for a barx variable.
``pviolcone float64``
     Maximal primal violation of the solution with respect to the conic constraints.
``pviolitg float64``
     Maximal violation in the integer constraints.
``dobj float64``
     Dual objective value.
``dviolcon float64``
     Maximal dual bound violation a xc variable.
``dviolvar float64``
     Maximal dual bound violation xx variable.
``dviolbarvar float64``
     Maximal dual bound violation for a bars variable.
``dviolcone float64``
     Maximum violation of the dual solution in the dual conic constraints .

Obtains information about a solution.
               


GetSolutionSlice
~~~~~~~~~~~~~~~~

Obtains a slice of the solution. ::

    func (*Task) GetSolutionSlice
        ( whichsol int32,
          solitem int32,
          first int32,
          last int32,
          values []float64 )
        ( values []float64 )


``first int32``
    Index of the first value in the slice.
``last int32``
     Value of the last index+1 in the slice.
``values []float64``
     The values of the requested solution elements.


Obtains a slice of the solution.

Consider the case of linear programming. The primal problem is given by

.. math::

  \begin{array}{lccccl}
    \mbox{minimize}              &      &      & c^T x+c^f &      &        \\
    \mbox{subject to} &  l^c & \leq & A x       & \leq & u^c,     \\
    &  l^x & \leq & x         & \leq & u^x.   \\
  \end{array}

and the corresponding dual problem is

.. math::
  
  \begin{array}{lccl}
    \mbox{maximize}   & (l^c)^T s_l^c - (u^c)^T s_u^c         &  \\
    & + (l^x)^T s_l^x - (u^x)^T s_u^x + c^f &  \\
    \mbox{subject to} & A^T y + s_l^x - s_u^x                 & = & c, \\
    & -y    + s_l^c - s_u^c                 & = & 0, \\
    & s_l^c,s_u^c,s_l^x,s_u^x \geq 0.       &   &    \\
  \end{array}

The ``solitem`` argument determines which part of the solution is returned:
  
* ``SOL_ITEM_XX``  : The variable ``values`` return :math:`x`.
* ``SOL_ITEM_Y``   : The variable ``values`` return :math:`y`.
* ``SOL_ITEM_SLC`` : The variable ``values`` return :math:`s_l^c`.
* ``SOL_ITEM_SUC`` : The variable ``values`` return :math:`s_u^c`.
* ``SOL_ITEM_SLX`` : The variable ``values`` return :math:`s_l^x`.
* ``SOL_ITEM_SUX`` : The variable ``values`` return :math:`s_u^x`.

A conic optimization problem has the same primal variables as in the linear case. Recall that the dual of a conic optimization problem is given by:

.. math::
  
  \begin{array}{lccccc}
    \mbox{maximize}   & (l^c)^T s_l^c - (u^c)^T s_u^c         &      &    \\
    & +(l^x)^T s_l^x - (u^x)^T s_u^x + c^f  &      &    \\
    \mbox{subject to} & A^T y + s_l^x - s_u^x + s_n^x         & =    & c, \\
    & -y + s_l^c - s_u^c                    & =    & 0, \\
    & s_l^c,s_u^c,s_l^x,s_u^x               & \geq & 0, \\
    & s_n^x \in \K^*                        &      &    \\
  \end{array}

This introduces one additional dual variable :math:`s_n^x`. This variable can be acceded by selecting ``solitem`` as ``SOL_ITEM_SNX``.

The meaning of the values returned by this function also depends on the *solution status* which can be obtained with ``getsolsta``.
Depending on the solution status ``value`` will be:
    
* ``SOL_STA_OPTIMAL``  A part of the  optimal solution satisfying the optimality criteria for continuous problems.
* ``SOL_STA_INTEGER_OPTIMAL``  A part of the  optimal solution satisfying the optimality criteria for integer problems.
* ``SOL_STA_PRIM_FEAS``        A part of the solution satisfying the feasibility criteria.
* ``SOL_STA_PRIM_INFEAS_CER``   A part of the primal certificate of infeasibility.
* ``SOL_STA_DUAL_INFEAS_CER``   A part of the dual certificate of infeasibility.




GetSparseSymMat
~~~~~~~~~~~~~~~

Gets a single symmetric matrix from the matrix store. ::

    func (*Task) GetSparseSymMat
        ( idx int64,
          subi []int32,
          subj []int32,
          valij []float64 )
        ( subi []int32,
          subj []int32,
          valij []float64 )


``idx int64``
     Index of the matrix to get.
``subi []int32``
     Row subscripts of the matrix non-zero elements.
``subj []int32``
     Column subscripts of the matrix non-zero elements.
``valij []float64``
     Coefficients of the matrix non-zero elements.


Get a single symmetric matrix from the matrix store.



GetStrParam
~~~~~~~~~~~

Obtains the value of a string parameter. ::

    func (*Task) GetStrParam ( param int32 ) ( len int32, parvalue string )


``len int32``
    The length of the parameter value.
``parvalue string``
     If this is not |null|, the parameter value is stored here.

Obtains the value of a string parameter. 


GetStrParamLen
~~~~~~~~~~~~~~

Obtains the length of a string parameter. ::

    func (*Task) GetStrParamLen ( param int32 ) ( len int32 )


``len int32``
    The length of the parameter value.

Obtains the length of a string parameter. 


GetSuc
~~~~~~

 Obtains the suc vector for a solution. ::

    func (*Task) GetSuc ( whichsol int32, suc []float64 ) ( suc []float64 )


``suc []float64``
    The suc vector.

Obtains the :math:`s_u^c` vector for a solution.  


GetSucSlice
~~~~~~~~~~~

 Obtains a slice of the suc vector for a solution. ::

    func (*Task) GetSucSlice
        ( whichsol int32,
          first int32,
          last int32,
          suc []float64 )
        ( suc []float64 )



Obtains a slice of the :math:`s_u^c` vector for a solution.  


GetSux
~~~~~~

 Obtains the sux vector for a solution. ::

    func (*Task) GetSux ( whichsol int32, sux []float64 ) ( sux []float64 )


``sux []float64``
    The sux vector.

Obtains the :math:`s_u^x` vector for a solution.  


GetSuxSlice
~~~~~~~~~~~

 Obtains a slice of the sux vector for a solution. ::

    func (*Task) GetSuxSlice
        ( whichsol int32,
          first int32,
          last int32,
          sux []float64 )
        ( sux []float64 )



Obtains a slice of the :math:`s_u^x` vector for a solution.  


GetSymMatInfo
~~~~~~~~~~~~~

Obtains information of  a matrix from the symmetric matrix storage E. ::

    func (*Task) GetSymMatInfo
        ( idx int64 )
        ( dim int32,
          nz int64,
          type int32 )


``idx int64``
     Index of the matrix that is requested information about.
``dim int32``
     Returns the dimension of the requested matrix.
``nz int64``
     Returns the number of non-zeros in the requested matrix.
``type int32``
     Returns the type of the requested matrix.


MOSEK maintains a vector denoted by :math:`E` of symmetric data matrixes. This function makes it possible to obtain important information about an data matrix in :math:`E`.



GetTaskName
~~~~~~~~~~~

Obtains the task name. ::

    func (*Task) GetTaskName (  ) ( taskname string )


``taskname string``
    Is assigned the task name.

Obtains the name assigned to the task. 


GetTaskNameLen
~~~~~~~~~~~~~~

 Obtains the length the task name. ::

    func (*Task) GetTaskNameLen (  ) ( len int32 )


``len int32``
    Returns the length of the task name.


Obtains the length the task name.



GetVarBound
~~~~~~~~~~~

 Obtains bound information for one variable. ::

    func (*Task) GetVarBound
        ( i int32 )
        ( bk int32,
          bl float64,
          bu float64 )


``i int32``
     Index of the variable for which the bound information should be obtained.


Obtains bound information for one variable.



GetVarBoundSlice
~~~~~~~~~~~~~~~~

 Obtains bounds information for a slice of the variables. ::

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

 Obtains a name of a variable. ::

    func (*Task) GetVarName ( j int32 ) ( name string )


``j int32``
    Index.
``name string``
    Returns the required name.


Obtains a name of a variable.



GetVarNameIndex
~~~~~~~~~~~~~~~

Checks whether the name somename has been assigned  to any variable.  ::

    func (*Task) GetVarNameIndex ( somename string ) ( asgn int32, index int32 )


``somename string``
    The name which should be checked.
``asgn int32``
    Is non-zero if name somename is assigned to a variable.
``index int32``
     If the name somename is assigned to a variable, then index is the name of the variable.

Checks whether the name ``somename`` has been assigned  to any variable. If it has been assigned to variable, then index of the variable is reported.  


GetVarNameLen
~~~~~~~~~~~~~

 Obtains the length of a name of a variable variable. ::

    func (*Task) GetVarNameLen ( i int32 ) ( len int32 )


``i int32``
    Index.
``len int32``
    Returns the length of the indicated name.


Obtains the length of a name of a variable variable.



GetVarType
~~~~~~~~~~

Gets the variable type of one variable. ::

    func (*Task) GetVarType ( j int32 ) ( vartype int32 )


``j int32``
    Index of the variable.
``vartype int32``
    Variable type of variable index j.

Gets the variable type of one variable. 


GetVarTypeList
~~~~~~~~~~~~~~

 Obtains the variable type for one or more variables. ::

    func (*Task) GetVarTypeList ( subj []int32, vartype []int32 ) ( vartype []int32 )


``subj []int32``
    A list of variable indexes.
``vartype []int32``
    Returns the variables types corresponding the variable indexes requested.


Obtains the variable type of one or more variables.

Upon return ``vartype[k]`` is the variable type of variable ``subj[k]``.



GetVersion
~~~~~~~~~~

Obtains MOSEK version information. ::

    func GetVersion
        (  )
        ( major int32,
          minor int32,
          build int32,
          revision int32,
          res int32 )


``major int32``
    Major version number.
``minor int32``
    Minor version number.
``build int32``
    Build number.
``revision int32``
    Revision number.
``res int32``
    Response code (see `mosek.RES_...`

Obtains MOSEK version information. 


GetXc
~~~~~

 Obtains the xc vector for a solution. ::

    func (*Task) GetXc ( whichsol int32, xc []float64 ) ( xc []float64 )


``xc []float64``
    The xc vector.

Obtains the :math:`x^c` vector for a solution.  


GetXcSlice
~~~~~~~~~~

 Obtains a slice of the xc vector for a solution. ::

    func (*Task) GetXcSlice
        ( whichsol int32,
          first int32,
          last int32,
          xc []float64 )
        ( xc []float64 )




Obtains a slice of the :math:`x^c` vector for a solution. 



GetXx
~~~~~

 Obtains the xx vector for a solution. ::

    func (*Task) GetXx ( whichsol int32, xx []float64 ) ( xx []float64 )


``xx []float64``
    The xx vector.

Obtains the :math:`x^x` vector for a solution.  


GetXxSlice
~~~~~~~~~~

 Obtains a slice of the xx vector for a solution. ::

    func (*Task) GetXxSlice
        ( whichsol int32,
          first int32,
          last int32,
          xx []float64 )
        ( xx []float64 )



Obtains a slice of the :math:`x^x` vector for a solution.  


GetY
~~~~

 Obtains the y vector for a solution. ::

    func (*Task) GetY ( whichsol int32, y []float64 ) ( y []float64 )


``y []float64``
    The y vector.

Obtains the :math:`y` vector for a solution.  


GetYSlice
~~~~~~~~~

 Obtains a slice of the y vector for a solution. ::

    func (*Task) GetYSlice
        ( whichsol int32,
          first int32,
          last int32,
          y []float64 )
        ( y []float64 )



Obtains a slice of the :math:`y` vector for a solution.  


InitBasisSolve
~~~~~~~~~~~~~~

 Prepare a task for basis solver. ::

    func (*Task) InitBasisSolve ( basis []int32 ) ( basis []int32 )


``basis []int32``
     The array of basis indexes to use.


Prepare a task for use with the ``solvewithbasis`` function.

This function should be called

* immediately before the first call to ``solvewithbasis``, and
* immediately before any subsequent call to ``solvewithbasis`` if the task has been modified. 

If the basis is singular i.e. not invertible, then the error ``RES_ERR_BASIS_SINGULAR`` is reported.




InputData
~~~~~~~~~

Input the linear part of an optimization task in one function call. ::

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




Input the linear part of an optimization problem.


The non-zeros of :math:`A` are inputted column-wise in the format described in Section :ref:`doc.optimizer.cmo_rmo_matrix`.

For an explained code example see Section :ref:`doc.tutorial_lo` and Section :ref:`doc.optimizer.matrix_formats`.




IsDouParName
~~~~~~~~~~~~

Checks a double parameter name. ::

    func (*Task) IsDouParName ( parname string ) ( param int32 )



Checks whether ``parname`` is a valid double parameter name.  


IsIntParName
~~~~~~~~~~~~

Checks an integer parameter name. ::

    func (*Task) IsIntParName ( parname string ) ( param int32 )



Checks whether ``parname`` is a valid integer parameter name.  


IsStrParName
~~~~~~~~~~~~

Checks a string parameter name. ::

    func (*Task) IsStrParName ( parname string ) ( param int32 )



Checks whether ``parname`` is a valid string parameter name.  


Licensecleanup
~~~~~~~~~~~~~~

Stops all threads and delete all handles used by the license system. ::

    func Licensecleanup (  ) ( res int32 )


``res int32``
    Response code (see `mosek.RES_...`


Stops all threads and delete all handles used by the license system. If this
function is called, it must be called as the last MOSEK API call. No other
MOSEK API calls are valid after this.



LinkFileToStream
~~~~~~~~~~~~~~~~

Directs all output from a task stream to a file. ::

    func (*Task) LinkFileToStream
        ( whichstream int32,
          filename string,
          append int32 )


``filename string``
    The name of the file where the stream is written.
``append int32``
     If this argument is 0 the output file will be overwritten, otherwise text is append to the output file.

Directs all output from a task stream to a file. 


Linkfiletostream
~~~~~~~~~~~~~~~~

Directs all output from a stream to a file. ::

    func (*Env) Linkfiletostream
        ( whichstream int32,
          filename string,
          append int32 )


``filename string``
    Name of the file to write stream data to.
``append int32``
     If this argument is non-zero, the output is appended to the file.

Directs all output from a stream to a file. 


OneSolutionSummary
~~~~~~~~~~~~~~~~~~

Prints a short summary for the specified solution. ::

    func (*Task) OneSolutionSummary ( whichstream int32, whichsol int32 )



Prints a short summary for a specified solution. 



Optimize
~~~~~~~~

Optimizes the problem. ::

    func (*Task) Optimize (  ) ( trmcode int32 )


``trmcode int32``
    Is either OK or a termination response code.


Calls the optimizer. Depending on the problem type and the selected optimizer
this will call one of the optimizers in MOSEK. By default the interior point
optimizer will be selected for continuous problems.  The optimizer may be
selected manually by setting the parameter ``IPAR_OPTIMIZER``.

.. only: c

   This function is equivalent to ``optimize`` except in the case where
   ``optimize`` would have returned a termination response code such as

   * ``RES_TRM_MAX_ITERATIONS`` or
   * ``RES_TRM_STALL``.

   Response codes comes in three categories:

   *  Errors: Indicate that an error has occurred during the optimization. E.g  that the optimizer has run out of memory (``RES_ERR_SPACE``). 
   *  Warnings: Less fatal than errors. E.g ``RES_WRN_LARGE_CJ`` indicating possibly problematic problem data
   *  Termination codes: Relaying information about the conditions under which the optimizer terminated. E.g ``RES_TRM_MAX_ITERATIONS`` indicates that
      the optimizer finished because it reached the maximum number of iterations specified by the user. 

This function returns errors on the left hand side. Warnings are not returned and termination codes are returned in the separate argument ``trmcode``.




OptimizerSummary
~~~~~~~~~~~~~~~~

Prints a short summary with optimizer statistics for last optimization. ::

    func (*Task) OptimizerSummary ( whichstream int32 )



Prints a short summary with optimizer statistics for last optimization.



Potrf
~~~~~

Computes a Cholesky factorization a dense matrix. ::

    func (*Env) Potrf
        ( uplo int32,
          n int32,
          a []float64 )
        ( a []float64 )


``uplo int32``
    Indicates whether the upper or lower triangular part of the matrix is stored.
``n int32``
     Dimension of the symmetric matrix.
``a []float64``
     A symmetric matrix stored in column-major order. Only the lower or the upper triangular part is used, accordingly with the uplo parameter. It will contain the result on exit.


Computes a Cholesky factorization of a real symmetric positive definite dense matrix.



PrimalRepair
~~~~~~~~~~~~

 The function repairs a primal infeasible optimization problem by adjusting the bounds on the constraints and variables. ::

    func (*Task) PrimalRepair
        ( wlc []float64,
          wuc []float64,
          wlx []float64,
          wux []float64 )


``wlc []float64``
    Weights associated with relaxing lower bounds on the constraints.
``wuc []float64``
     Weights associated with relaxing the upper bound on the constraints.
``wlx []float64``
     Weights associated with relaxing the lower bounds of the variables.
``wux []float64``
     Weights associated with relaxing the upper bounds of variables.


The function repairs a primal infeasible optimization problem by adjusting the bounds on the constraints and variables where the adjustment
is computed as the minimal weighted sum relaxation to the bounds on the constraints and variables. Observe the function only repairs the problem but does not
compute an optimal solution to the repaired problem. If an optimal solution is required the problem should be optimized after the repair.

The function is applicable to linear and conic problems possibly having integer constrained variables.

Observe that when computing the minimal weighted relaxation then the termination tolerance specified by the parameters of the task is employed. For instance
the parameter ``IPAR_MIO_MODE`` can be used make MOSEK ignore the integer constraints during the repair which usually leads to a much faster repair.
However, the drawback is of course that the repaired problem may not have an integer feasible solution.

Note the function modifies the bounds on the constraints and variables. If this is not a desired feature, then
apply the function to a cloned task. 



PrimalSensitivity
~~~~~~~~~~~~~~~~~

Perform sensitivity analysis on bounds. ::

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


``subi []int32``
    Indexes of bounds on constraints to analyze.
``marki []int32``
    Mark which constraint bounds to analyze.
``subj []int32``
    Indexes of bounds on variables to analyze.
``markj []int32``
    Mark which variable bounds to analyze.
``leftpricei []float64``
    Left shadow price for constraints.
``rightpricei []float64``
    Right shadow price for constraints.
``leftrangei []float64``
    Left range for constraints.
``rightrangei []float64``
    Right range for constraints.
``leftpricej []float64``
    Left price for variables.
``rightpricej []float64``
    Right price for variables.
``leftrangej []float64``
    Left range for variables.
``rightrangej []float64``
    Right range for variables.


Calculates sensitivity information for bounds on variables and constraints.

For details on sensitivity analysis and the definitions of *shadow price* and
*linearity interval* see Section :ref:`doc.sensitivity_analysis`.

The constraints for which sensitivity analysis is performed are given by the
data structures:


#. ``subi`` Index of constraint to analyze.
#. ``marki`` Indicate for which bound of constraint ``subi[i]`` sensitivity
   analysis is performed.  If ``marki[i]`` = ``MARK_UP`` the upper bound of
   constraint ``subi[i]`` is analyzed, and if ``marki[i]`` = ``MARK_LO`` the
   lower bound is analyzed.  If ``subi[i]`` is an equality constraint, either
   ``MARK_LO`` or ``MARK_UP`` can be used to select the
   constraint for sensitivity analysis.

Consider the problem:

.. math::

    \begin{array}{lccl}
    \mbox{minimize}   & x_1 + x_2 &  &\\
    \mbox{subject to} -1 \leq & x_1 - x_2                  & \leq & 1, \\
                      & x_1                       & = & 0, \\
                      & x_1 \geq 0,x_2 \geq 0  & &
    \end{array}

Suppose that

* ``numi = 1;``
* ``subi = [0];``
* ``marki`` = [``MARK_UP``]

then

``leftpricei[0]``, ``rightpricei[0]``, ``leftrangei[0]`` and ``rightrangei[0]``
will contain the sensitivity information for the upper bound on constraint :math:`0`
given by the expression:

.. math:: x_1 - x_2 \leq  1

Similarly, the variables for which to perform sensitivity analysis are given by
the structures:
  
#. ``subj`` Index of variables to analyze.

#. ``markj`` Indicate for which bound of variable ``subi[j]`` sensitivity
   analysis is performed.  If ``markj[j]`` = ``MARK_UP`` the upper bound of
   constraint ``subi[j]`` is analyzed, and if ``markj[j]`` = ``MARK_LO`` the
   lower bound is analyzed.

#. If ``subi[j]`` is an equality constraint, either ``MARK_LO`` or
   ``MARK_UP`` can be used to select the constraint for sensitivity
   analysis.


For an example, please see Section :ref:`doc.shared.sensitivity_example`.

The type of sensitivity analysis to be performed (basis or optimal partition)
is controlled by the parameter ``IPAR_SENSITIVITY_TYPE``.




ProStaToStr
~~~~~~~~~~~

 Obtains a string containing the name of a problem status given. ::

    func (*Task) ProStaToStr ( prosta int32 ) ( str string )


``str string``
    String corresponding to the status key.


Obtains a string containing the name of a problem status given.



ProbTypeToStr
~~~~~~~~~~~~~

Obtains a string containing the name of a problem type given.::

    func (*Task) ProbTypeToStr ( probtype int32 ) ( str string )


``str string``
    String corresponding to the problem type key.


Obtains a string containing the name of a problem type given.



PutACol
~~~~~~~

 Replaces all elements in one column of A. ::

    func (*Task) PutACol
        ( j int32,
          subj []int32,
          valj []float64 )


``j int32``
     Column index.
``subj []int32``
    Row indexes of non-zero values in column.
``valj []float64``
    New non-zero values of column.


Resets all the elements in column :math:`j` to zero and then do
   
.. math:: A_{\mathtt{subj}[k],\mathtt{j}} = \mathtt{valj}[k], \quad k=0,\ldots,\mathtt{nzj}-1. 



PutAColList
~~~~~~~~~~~

 Replaces all elements in several columns the linear constraint matrix by new values. ::

    func (*Task) PutAColList
        ( sub []int32,
          ptrb []int32,
          ptre []int32,
          asub []int32,
          aval []float64 )


``sub []int32``
     Indexes of columns that should be replaced.
``ptrb []int32``
     Array of pointers to the first element in the columns.
``ptre []int32``
     Array of pointers to the last element plus one in the columns.
``asub []int32``
    Variable indexes.


Replaces all elements in a set of columns of :math:`A`. The elements are replaced as follows  

.. math::

    \begin{array}{rl}
      \mathtt{for} & i=\idxbeg,\ldots,\idxend{num}\\
                  & a_{\mathtt{asub}[k],\mathtt{sub}[i]} = \mathtt{aval}[k],\quad k=\mathtt{aptrb}[i],\ldots,\mathtt{aptre}[i]-1. 
    \end{array}



PutAColSlice
~~~~~~~~~~~~

 Replaces all elements in several columns the linear constraint matrix by new values. ::

    func (*Task) PutAColSlice
        ( first int32,
          last int32,
          ptrb []int64,
          ptre []int64,
          asub []int32,
          aval []float64 )


``first int32``
    First column in the slice.
``last int32``
     Last column plus one in the slice.
``ptrb []int64``
     Array of pointers to the first element in the columns.
``ptre []int64``
     Array of pointers to the last element plus one in the columns.
``asub []int32``
    Variable indexes.

Replaces all elements in a set of columns of :math:`A`.  


PutARow
~~~~~~~

 Replaces all elements in one row of A. ::

    func (*Task) PutARow
        ( i int32,
          subi []int32,
          vali []float64 )


``i int32``
     row index.
``subi []int32``
    Row indexes of non-zero values in row.
``vali []float64``
    New non-zero values of row.


Resets all the elements in row :math:`i` to zero and then do

.. math:: A_{\mathtt{i},\mathtt{subi}[k]} = \mathtt{vali}[k], \quad k=0,\ldots,\mathtt{nzi}-1. 



PutARowList
~~~~~~~~~~~

 Replaces all elements in several rows the linear constraint matrix by new values. ::

    func (*Task) PutARowList
        ( sub []int32,
          aptrb []int32,
          aptre []int32,
          asub []int32,
          aval []float64 )


``sub []int32``
     Indexes of rows or columns that should be replaced.
``aptrb []int32``
     Array of pointers to the first element in the rows or columns.
``aptre []int32``
     Array of pointers to the last element plus one in the rows or columns.
``asub []int32``
    Variable indexes.


Replaces all elements in a set of rows of :math:`A`. The elements are replaced as follows  

.. math::

    \begin{array}{rl}
      \mathtt{for} & i=\idxbeg,\ldots,\idxend{num} \\
                   & a_{\mathtt{sub}[i],\mathtt{asub}[k]} = \mathtt{aval}[k],\quad k=\mathtt{aptrb}[i],\ldots,\mathtt{aptre}[i]-1. 
    \end{array}



PutARowSlice
~~~~~~~~~~~~

 Replaces all elements in several rows the linear constraint matrix by new values. ::

    func (*Task) PutARowSlice
        ( first int32,
          last int32,
          ptrb []int64,
          ptre []int64,
          asub []int32,
          aval []float64 )


``first int32``
    First row in the slice.
``last int32``
     Last row plus one in the slice.
``ptrb []int64``
     Array of pointers to the first element in the rows.
``ptre []int64``
     Array of pointers to the last element plus one in the rows.
``asub []int32``
    Variable indexes.


Replaces all elements in a set of rows of :math:`A`. The elements is replaced as follows

.. math::

    \begin{array}{rl}
      \mathtt{for} & i=\mathtt{first},\ldots,\mathtt{last}-1 \\
                  & a_{\mathtt{sub}[i],\mathtt{asub}[k]} = \mathtt{aval}[k],\quad k=\mathtt{aptrb}[i],\ldots,\mathtt{aptre}[i]-1. 
    \end{array}




PutAij
~~~~~~

Changes a single value in the linear coefficient matrix. ::

    func (*Task) PutAij
        ( i int32,
          j int32,
          aij float64 )


``i int32``
     Index of the constraint in which the change should occur.
``j int32``
     Index of the variable in which the change should occur.
``aij float64``
    New coefficient.


Changes a coefficient in :math:`A` using the method

.. math:: a_{\mathtt{i}\mathtt{j}} = \mathtt{aij}.



PutAijList
~~~~~~~~~~

Changes one or more coefficients in the linear constraint matrix. ::

    func (*Task) PutAijList
        ( subi []int32,
          subj []int32,
          valij []float64 )


``subi []int32``
     Constraint indexes in which the change should occur.
``subj []int32``
     Variable indexes in which the change should occur.
``valij []float64``
    New coefficient values.


Changes one or more coefficients in :math:`A` using the method

.. math:: a_{\mathtt{subi[k]},\mathtt{subj[k]}} = \mathtt{valij[k]}, \quad k=0,\ldots,\mathtt{num}-1.



PutBaraBlockTriplet
~~~~~~~~~~~~~~~~~~~

Inputs barA in block triplet form. ::

    func (*Task) PutBaraBlockTriplet
        ( num int64,
          subi []int32,
          subj []int32,
          subk []int32,
          subl []int32,
          valijkl []float64 )


``num int64``
     Number of elements in the block triplet form.
``subi []int32``
     Constraint index.
``subj []int32``
     Symmetric matrix variable index.
``subk []int32``
     Block row index.
``subl []int32``
     Block column index.
``valijkl []float64``
     The numerical value associated with the block triplet.

Inputs the :math:`\bar{A}` in block triplet form.  


PutBaraIj
~~~~~~~~~

Inputs an element of barA. ::

    func (*Task) PutBaraIj
        ( i int32,
          j int32,
          sub []int64,
          weights []float64 )


``i int32``
     Row index of barA.
``j int32``
     Column index of barA.
``sub []int64``
     See argument weights for an explanation.
``weights []float64``
     Weights in the weighted sum.


This function puts one element associated with :math:`\bar{X}_j` in the :math:`\bar{A}` matrix.

Each element in the :math:`\bar{A}` matrix is a weighted sum of
symmetric matrixes, i.e. :math:`\bar{A}_{ij}` is a symmetric matrix
with dimensions as :math:`\bar{X}_j`. By default all elements in
:math:`\bar{A}` are 0, so only non-zero elements need be added.

Setting the same elements again will overwrite the earlier entry. 

The symmetric matrixes themselves are defined separately
using the function ``appendsparsesymmat``.



PutBarcBlockTriplet
~~~~~~~~~~~~~~~~~~~

Inputs barC in block triplet form. ::

    func (*Task) PutBarcBlockTriplet
        ( num int64,
          subj []int32,
          subk []int32,
          subl []int32,
          valjkl []float64 )


``num int64``
     Number of elements in the block triplet form.
``subj []int32``
     Symmetric matrix variable index.
``subk []int32``
     Block row index.
``subl []int32``
     Block column index.
``valjkl []float64``
     The numerical value associated with the block triplet.

Inputs the :math:`\bar{C}` in block triplet form.  


PutBarcJ
~~~~~~~~

Changes one element in barc.::

    func (*Task) PutBarcJ
        ( j int32,
          sub []int64,
          weights []float64 )


``j int32``
     Index of the element in barc` that should be changed.
``sub []int64``
     sub is list of indexes of those symmetric matrices appearing in sum.
``weights []float64``
     The weights of the terms in the weighted sum.


This function puts one element associated with :math:`\bar{X}_j` in the :math:`\bar{c}` vector. 

Each element in the :math:`\bar{c}` vector is a weighted sum of symmetric
matrixes, i.e. :math:`\bar{c}_j` is a symmetric matrix with dimensions as
:math:`\bar{X}_j`. By default all elements in :math:`\bar{c}` are :math:`0`, so only non-zero elements need be added.

Setting the same elements again will overwrite the earlier entry. 

The symmetric matrixes themselves are defined separately using the function
``appendsparsesymmat``.



PutBarsJ
~~~~~~~~

 Sets the dual solution for a semidefinite variable. ::

    func (*Task) PutBarsJ
        ( whichsol int32,
          j int32,
          barsj []float64 )


``j int32``
    Index of the semidefinite variable.
``barsj []float64``
    Value of the j'th variable of barx.


Sets the dual solution for a semidefinite variable.



PutBarvarName
~~~~~~~~~~~~~

 Puts the name of a semidefinite variable. ::

    func (*Task) PutBarvarName ( j int32, name string )

``j int32``
    Index of the variable.
``name string``
    The variable name.


Puts the name of a semidefinite variable.



PutBarxJ
~~~~~~~~

 Sets the primal solution for a semidefinite variable. ::

    func (*Task) PutBarxJ
        ( whichsol int32,
          j int32,
          barxj []float64 )


``j int32``
    Index of the semidefinite variable.
``barxj []float64``
    Value of the j'th variable of barx.


Sets the primal solution for a semidefinite variable.



PutBound
~~~~~~~~

 Changes the bound for either one constraint or one variable. ::

    func (*Task) PutBound
        ( accmode int32,
          i int32,
          bk int32,
          bl float64,
          bu float64 )


``accmode int32``
     Defines whether the bound for a constraint or a variable is changed.
``i int32``
    Index of the constraint or variable.
``bk int32``
    New bound key.
``bl float64``
    New lower bound.
``bu float64``
    New upper bound.


Changes the bounds for either one constraint or one variable.

If the a bound value specified is numerically larger than
``DPAR_DATA_TOL_BOUND_INF`` it is considered infinite and the bound key is
changed accordingly. If a bound value is numerically larger than
``DPAR_DATA_TOL_BOUND_WRN``, a warning will be displayed, but the bound is
inputted as specified.



PutBoundList
~~~~~~~~~~~~

Changes the bounds of constraints or variables. ::

    func (*Task) PutBoundList
        ( accmode int32,
          sub []int32,
          bk []int32,
          bl []float64,
          bu []float64 )


``accmode int32``
     Defines whether to access bounds on variables or constraints.
``sub []int32``
    Subscripts of the bounds that should be changed.
``bk []int32``
     Bound keys for variables or constraints.
``bl []float64``
     Bound keys for variables or constraints.
``bu []float64``
     Constraint or variable upper bounds.


Changes the bounds for either some constraints or variables.
If multiple bound changes are specified for
a constraint or a variable, only the last change takes effect.



PutBoundSlice
~~~~~~~~~~~~~

Modifies bounds. ::

    func (*Task) PutBoundSlice
        ( con int32,
          first int32,
          last int32,
          bk []int32,
          bl []float64,
          bu []float64 )


``con int32``
     Determines whether variables or constraints are modified.


Changes the bounds for a sequence of variables or constraints.



PutCJ
~~~~~

Modifies one linear coefficient in the objective. ::

    func (*Task) PutCJ ( j int32, cj float64 )

``j int32``
    Index of the variable whose objective coefficient should be changed.
``cj float64``
    New coefficient value.


Modifies one coefficient in the linear objective vector :math:`c`, i.e.

.. math:: c_{\mathtt{j}} = \mathtt{cj}.




PutCList
~~~~~~~~

Modifies a part of the linear objective coefficients. ::

    func (*Task) PutCList ( subj []int32, val []float64 )

``subj []int32``
    Index of variables for which objective coefficients should be changed.
``val []float64``
    New numerical values for the objective coefficients that should be modified.


Modifies elements in the linear term :math:`c` in the objective using the principle

.. math:: c_{\mathtt{subj[t]}} = \mathtt{val[t]}, \quad t=0,\ldots,\mathtt{num}-1.

If a variable index is specified multiple times in ``subj`` only the last entry is used.



PutCSlice
~~~~~~~~~

Modifies a slice of the linear objective coefficients. ::

    func (*Task) PutCSlice
        ( first int32,
          last int32,
          slice []float64 )


``first int32``
    First element in the slice of c.
``last int32``
    Last element plus 1 of the slice in c to be changed.
``slice []float64``
    New numerical values for the objective coefficients that should be modified.


Modifies a slice in the linear term :math:`c` in the objective using the principle

.. math:: c_{\mathtt{j}} = \mathtt{slice[j-first]}, \quad j=first,..,\idxend{last}




PutCallbackFunc
~~~~~~~~~~~~~~~

::

    func (t *Task) PutCallbackFunc ( fun func(int32) int )


Add a callback function to the task.

The callback function takes one integer argument that indicates the
  progress of the solver (``mosek.CALLBACK_...``). It returns an integer
value: ``0`` means that the solver should just continue, anything else
means that the solver will stop.


PutCfix
~~~~~~~

Replaces the fixed term in the objective. ::

    func (*Task) PutCfix ( cfix float64 )



Replaces the fixed term in the objective by a new one.



PutConBound
~~~~~~~~~~~

 Changes the bound for one constraint. ::

    func (*Task) PutConBound
        ( i int32,
          bk int32,
          bl float64,
          bu float64 )


``i int32``
    Index of the constraint.
``bk int32``
    New bound key.
``bl float64``
    New lower bound.
``bu float64``
    New upper bound.


Changes the bounds for one constraint.

If the a bound value specified is numerically larger than ``DPAR_DATA_TOL_BOUND_INF`` it is considered infinite and the bound key is
changed accordingly. If a bound value is numerically larger than ``DPAR_DATA_TOL_BOUND_WRN``, a warning will be displayed, but the bound is inputted as specified.



PutConBoundList
~~~~~~~~~~~~~~~

Changes the bounds of a list of constraints. ::

    func (*Task) PutConBoundList
        ( sub []int32,
          bkc []int32,
          blc []float64,
          buc []float64 )


``sub []int32``
    List constraints indexes.
``bkc []int32``
     New bound keys.
``blc []float64``
     New lower bound values.
``buc []float64``
     New upper bounds values.


Changes the bounds for a list of constraints. If multiple bound changes are specified for a constraint, then only the last change takes effect.



PutConBoundSlice
~~~~~~~~~~~~~~~~

 Changes the bounds for a slice of the constraints. ::

    func (*Task) PutConBoundSlice
        ( first int32,
          last int32,
          bk []int32,
          bl []float64,
          bu []float64 )


``first int32``
    Index of the first constraint in the slice.
``last int32``
    Index of the last constraint in the slice plus 1.
``bk []int32``
    New bound keys.
``bl []float64``
    New lower bounds.
``bu []float64``
    New upper bounds.


Changes the bounds for a slice of the constraints.



PutConName
~~~~~~~~~~

 Puts the name of a constraint. ::

    func (*Task) PutConName ( i int32, name string )

``i int32``
    Index of the constraint.
``name string``
    The variable name.


Puts the name of a constraint.



PutCone
~~~~~~~

 Replaces a conic constraint. ::

    func (*Task) PutCone
        ( k int32,
          ct int32,
          conepar float64,
          submem []int32 )


``k int32``
    Index of the cone.


Replaces a conic constraint.



PutConeName
~~~~~~~~~~~

 Puts the name of a cone. ::

    func (*Task) PutConeName ( j int32, name string )

``j int32``
    Index of the cone.
``name string``
    The variable name.


Puts the name of a cone.



PutDouParam
~~~~~~~~~~~

Sets a double parameter. ::

    func (*Task) PutDouParam ( param int32, parvalue float64 )


Sets the value of a double parameter. 


PutInfoCallbackFunc
~~~~~~~~~~~~~~~~~~~

::

    func (t *Task) PutInfoCallbackFunc ( fun func(int32) int )


Add an information callback function to the task.

The callback function takes four arguments: ``(code,dinf,iinf,liinf)``

Callback function arguments:

``code``
    Indicates the progress of the solver (``mosek.CALLBACK_...``).
``dinf``
    An array of ``float64`` information items. The indexes correspond to ``mosek.DINF_...``
``iinf``
    An array of ``int32`` information items. The indexes correspond to ``mosek.IINF_...``
``liinf``
    An array of ``int64`` information items. The indexes correspond to ``mosek.LIINF_...``


Callback function returns: Non-zero to indicate that the solver should stop.


PutIntParam
~~~~~~~~~~~

Sets an integer parameter. ::

    func (*Task) PutIntParam ( param int32, parvalue int32 )



Sets the value of an integer parameter.


   Please notice that some parameters take values that are defined in Enum
   classes. This function accepts only integer values, so to use e.g. the value
   ``ON``, is necessary to use the member ``.value``. For example: ::

       task.putintparam(Env.iparam.opf_write_problem,Env.onoffkey.on.value)





PutLicenseCode
~~~~~~~~~~~~~~

The purpose of this function is to input a runtime license code. ::

    func (*Env) PutLicenseCode ( code []int32 )

``code []int32``
     A license key string.


The purpose of this function is to input a runtime license code.



PutLicenseDebug
~~~~~~~~~~~~~~~

Enables debug information for the license system. ::

    func (*Env) PutLicenseDebug ( licdebug int32 )

``licdebug int32``
    Enable output of license check-out debug information.

If ``licdebug`` is  non-zero, then MOSEK will print debug info regarding the license checkout.  


PutLicensePath
~~~~~~~~~~~~~~

Set the path to the license file. ::

    func (*Env) PutLicensePath ( licensepath string )

``licensepath string``
    A path specifycing where to search for the license.


Set the path to the license file.



PutLicenseWait
~~~~~~~~~~~~~~

Control whether mosek should wait for an available license if no license is available. ::

    func (*Env) PutLicenseWait ( licwait int32 )

``licwait int32``
    Enable waiting for a license.


If ``licwait`` is non-zero, then MOSEK will wait for a license if no license
is available. Moreover, ``licwait-1`` is the number of milliseconds to wait between each check for an available license.



PutMaxNumANz
~~~~~~~~~~~~

 The function changes the size of the preallocated storage for linear coefficients. ::

    func (*Task) PutMaxNumANz ( maxnumanz int64 )

``maxnumanz int64``
    New size of the storage reserved for storing the linear coefficient matrix.


MOSEK stores only the non-zero elements in :math:`A`.  Therefore1 MOSEK
cannot predict how much storage is required to store :math:`A`. Using this
function it is possible to specify the number of non-zeros to preallocate for
storing :math:`A`.

If the number of non-zeros in the problem is known, it is a good idea to set
``maxnumanz`` slightly larger than this number, otherwise a rough estimate can
be used. In general, if :math:`A` is inputted in many small chunks, setting
this value may speed up the data input phase.

It is not mandatory to call this function, since MOSEK will reallocate
internal structures whenever it is necessary.

Observe the function call has no effect if both ``maxnumcon`` and ``maxnumvar``
is zero.



PutMaxNumBarvar
~~~~~~~~~~~~~~~

Sets the number of preallocated symmetric matrix variables in the optimization task. ::

    func (*Task) PutMaxNumBarvar ( maxnumbarvar int32 )

``maxnumbarvar int32``
    The maximum number of semidefinite variables.


Sets the number of preallocated symmetric matrix variables in the optimization
task. When this number of variables is reached MOSEK will automatically
allocate more space for variables.

It is not mandatory to call this function, since its only function is to give a
hint of the amount of data to preallocate for efficiency reasons.

Please note that ``maxnumbarvar`` must be larger than the current number of
variables in the task.



PutMaxNumCon
~~~~~~~~~~~~

Sets the number of preallocated constraints in the optimization task. ::

    func (*Task) PutMaxNumCon ( maxnumcon int32 )



Sets the number of preallocated constraints in the optimization task. When this
number of constraints is reached MOSEK will automatically allocate more space
for constraints.

It is never mandatory to call this function, since MOSEK will reallocate any
internal structures whenever it is required.

Please note that ``maxnumcon`` must be larger than the current number of
constraints in the task.



PutMaxNumCone
~~~~~~~~~~~~~

 Sets the number of preallocated conic constraints in the optimization task. ::

    func (*Task) PutMaxNumCone ( maxnumcone int32 )



Sets the number of preallocated conic constraints in the optimization task.
When this number of conic constraints is reached MOSEK will automatically
allocate more space for conic constraints.

It is never mandatory to call this function, since MOSEK will reallocate any
internal structures whenever it is required.

Please note that ``maxnumcon`` must be larger than the current number of
constraints in the task.



PutMaxNumQNz
~~~~~~~~~~~~

Changes the size of the preallocated storage for quadratic terms. ::

    func (*Task) PutMaxNumQNz ( maxnumqnz int64 )



MOSEK stores only the non-zero elements in :math:`Q`. Therefore, MOSEK
cannot predict how much storage is required to store :math:`Q`. Using this
function it is possible to specify the number non-zeros to preallocate for
storing :math:`Q` (both objective and constraints).

It may be advantageous to reserve more non-zeros for :math:`Q` than actually
needed since it may improve the internal efficiency of MOSEK, however, it is
never worthwhile to specify more than the double of the anticipated number of
non-zeros in :math:`Q`.

It is never mandatory to call this function, since its only function is to give
a hint of the amount of data to preallocate for efficiency reasons.



PutMaxNumVar
~~~~~~~~~~~~

Sets the number of preallocated variables in the optimization task. ::

    func (*Task) PutMaxNumVar ( maxnumvar int32 )



Sets the number of preallocated variables in the optimization task. When this
number of variables is reached MOSEK will automatically allocate more space
for variables.

It is never mandatory to call this function, since its only function is to give
a hint of the amount of data to preallocate for efficiency reasons.

Please note that ``maxnumvar`` must be larger than the current number of
variables in the task.



PutNaDouParam
~~~~~~~~~~~~~

Sets a double parameter. ::

    func (*Task) PutNaDouParam ( paramname string, parvalue float64 )


Sets the value of a named double parameter. 


PutNaIntParam
~~~~~~~~~~~~~

Sets an integer parameter. ::

    func (*Task) PutNaIntParam ( paramname string, parvalue int32 )


Sets the value of a named integer parameter. 


PutNaStrParam
~~~~~~~~~~~~~

Sets a string parameter. ::

    func (*Task) PutNaStrParam ( paramname string, parvalue string )


Sets the value of a named string parameter. 


PutObjName
~~~~~~~~~~

Assigns a new name to the objective. ::

    func (*Task) PutObjName ( objname string )


Assigns the name given by ``objname`` to the objective function.  


PutObjSense
~~~~~~~~~~~

Sets the objective sense. ::

    func (*Task) PutObjSense ( sense int32 )

``sense int32``
    The objective sense of the task

Sets the objective sense of the task.  


PutParam
~~~~~~~~

Modifies the value of parameter. ::

    func (*Task) PutParam ( parname string, parvalue string )



Checks if a ``parname`` is valid parameter name. If it is, the parameter is
assigned the value specified by ``parvalue``.



PutQCon
~~~~~~~

Replaces all quadratic terms in constraints. ::

    func (*Task) PutQCon
        ( qcsubk []int32,
          qcsubi []int32,
          qcsubj []int32,
          qcval []float64 )




Replace all quadratic entries in the constraints. consider constraints on the form:

.. math:: l_k^c \leq  \half \sum_{i=0}^{\idxend{numvar}} \sum_{j=0}^{\idxend{numvar}} q_{ij}^k x_i x_j + \sum_{j=0}^{\idxend{numvar}} a_{kj} x_j \leq u_k^c, ~  k=0,\ldots,m-1.

the function assigns values to :math:`q` such that:

.. math:: q_{\mathtt{qcsubi[t]},\mathtt{qcsubj[t]}}^{\mathtt{qcsubk[t]}} = \mathtt{qcval[t]},~t=0,\ldots,\mathtt{numqcnz}-1.

and

.. math:: q_{\mathtt{\mathtt{qcsubj[t]},qcsubi[t]}}^{\mathtt{qcsubk[t]}} = \mathtt{qcval[t]},~t=0,\ldots,\mathtt{numqcnz}-1.

values not assigned are set to zero.

Please note that duplicate entries are added together.



PutQConK
~~~~~~~~

Replaces all quadratic terms in a single constraint. ::

    func (*Task) PutQConK
        ( k int32,
          qcsubi []int32,
          qcsubj []int32,
          qcval []float64 )


``k int32``
    The constraint in which the new quadratic elements are inserted.


Replaces all the quadratic entries in one constraint :math:`k` of the form:

 .. math:: l_k^c \leq  \half \sum_{i=\idxbeg}^{\idxend{numvar}} \sum_{j=\idxbeg}^{\idxend{numvar}} q_{ij}^k x_i x_j + \sum_{j=\idxbeg}^{\idxend{numvar}} a_{kj} x_j \leq u_k^c.

It is assumed that :math:`Q^k` is symmetric, i.e. :math:`q^k_{ij} = q^k_{ji}`,and therefore, only the values of :math:`q^k_{ij}` for which :math:`i \geq j`
should be inputted to MOSEK.  To be precise, MOSEK uses the following procedure

.. math::

    \begin{array}{ll}
    1. & Q^{k}  = 0\\
    2. & \mbox{for } t=\idxbeg \mbox{ to }\idxend{numqcnz} \\
    3. & \qquad q_{\mathtt{qcsubi[t]},\mathtt{qcsubj[t]}}^k = q_{\mathtt{qcsubi[t]},\mathtt{qcsubj[t]}}^{k} + \mathtt{qcval[t]} \\
    3. & \qquad q_{\mathtt{qcsubj[t]},\mathtt{qcsubi[t]}}^k = q_{\mathtt{qcsubj[t]},\mathtt{qcsubi[t]}}^{k} + \mathtt{qcval[t]} \\
    \end{array}

Please note that:

*   For large problems it is essential for the efficiency that the function
    ``putmaxnumqnz`` is employed to specify an appropriate
    ``maxnumqnz``.
*   Only the lower triangular part should be specified because :math:`Q^k` is
    symmetric. Specifying values for :math:`q^k_{ij}` where :math:`i < j`
    will result in an error. 
*   Only non-zero elements should be specified.
*   The order in which the non-zero elements are specified is insignificant.
*   Duplicate elements are added together. Hence, it is recommended not to
    specify the same element multiple times in ``qcsubi``,
    ``qcsubj``, and ``qcval``.

For a code example see Section :ref:`doc.tutorial_qo`



PutQObj
~~~~~~~

Replaces all quadratic terms in the objective. ::

    func (*Task) PutQObj
        ( qosubi []int32,
          qosubj []int32,
          qoval []float64 )




Replaces all the quadratic terms in the objective

.. math:: \half \sum_{i=\idxbeg}^{\idxend{numvar}} \sum_{j=\idxbeg}^{\idxend{numvar}} q_{ij}^o x_i x_j + \sum_{j=\idxbeg}^{\idxend{numvar}} c_j x_j + c^f.

It is assumed that :math:`Q^o` is symmetric, i.e. :math:`q^o_{ij} = q^o_{ji}`, and therefore, only the values of :math:`q^o_{ij}` for which :math:`i \geq j` should be specified.  To be precise, MOSEK uses the following procedure

.. math::

   \begin{array}{ll}
    1. & Q^o = 0\\
    2. & \mbox{for } t=\idxbeg \mbox{ to } \idxend{numqonz} \\
    3. & \qquad q_{\mathtt{qosubi[t]},\mathtt{qosubj[t]}}^o = q_{\mathtt{qosubi[t]},\mathtt{qosubj[t]}}^o + \mathtt{qoval[t]} \\
    3. & \qquad q_{\mathtt{qosubj[t]},\mathtt{qosubi[t]}}^o = q_{\mathtt{qosubj[t]},\mathtt{qosubi[t]}}^o + \mathtt{qoval[t]} \\
    \end{array}

Please note that:

* Only the lower triangular part should be specified because :math:`Q^o` is symmetric. Specifying values for :math:`q^o_{ij}` where :math:`i < j` will result in an error. 

* Only non-zero elements should be specified.

* The order in which the non-zero elements are specified is insignificant.

* Duplicate entries are added to together.

For a code example see Section :ref:`doc.tutorial_qo`.




PutQObjIJ
~~~~~~~~~

 Replaces one coefficient in the quadratic term in the objective. ::

    func (*Task) PutQObjIJ
        ( i int32,
          j int32,
          qoij float64 )


``i int32``
    Row index for the coefficient to be replaced.
``j int32``
    Column index for the coefficient to be replaced.
``qoij float64``
    The new coefficient value.


Replaces one coefficient in the quadratic term in the objective. The function
performs the assignment

.. math:: q_{\mathtt{i}\mathtt{j}}^o = \mathtt{qoij}.

Only the elements in the lower triangular part are accepted. Setting
:math:`q_{ij}` with :math:`j>i` will cause an error.

Please note that replacing all quadratic element, one at a time, is more
computationally expensive than replacing all elements at once. Use
``putqobj`` instead whenever possible.



PutSkc
~~~~~~

 Sets the status keys for the constraints. ::

    func (*Task) PutSkc ( whichsol int32, skc []int32 )



Sets the status keys for the constraints.



PutSkcSlice
~~~~~~~~~~~

 Sets the status keys for the constraints. ::

    func (*Task) PutSkcSlice
        ( whichsol int32,
          first int32,
          last int32,
          skc []int32 )




Sets the status keys for the constraints.



PutSkx
~~~~~~

 Sets the status keys for the scalar variables. ::

    func (*Task) PutSkx ( whichsol int32, skx []int32 )



Sets the status keys for the scalar variables.



PutSkxSlice
~~~~~~~~~~~

 Sets the status keys for the variables. ::

    func (*Task) PutSkxSlice
        ( whichsol int32,
          first int32,
          last int32,
          skx []int32 )




Sets the status keys for the variables.



PutSlc
~~~~~~

 Sets the slc vector for a solution. ::

    func (*Task) PutSlc ( whichsol int32, slc []float64 )

``slc []float64``
    The slc vector.

Sets the :math:`s_l^c` vector for a solution.  


PutSlcSlice
~~~~~~~~~~~

 Sets a slice of the slc vector for a solution. ::

    func (*Task) PutSlcSlice
        ( whichsol int32,
          first int32,
          last int32,
          slc []float64 )



Sets a slice of the :math:`s_l^c` vector for a solution.  


PutSlx
~~~~~~

 Sets the slx vector for a solution. ::

    func (*Task) PutSlx ( whichsol int32, slx []float64 )

``slx []float64``
    The slx vector.

Sets the :math:`s_l^x` vector for a solution.  


PutSlxSlice
~~~~~~~~~~~

 Sets a slice of the slx vector for a solution. ::

    func (*Task) PutSlxSlice
        ( whichsol int32,
          first int32,
          last int32,
          slx []float64 )



Sets a slice of the :math:`s_l^x` vector for a solution.  


PutSnx
~~~~~~

 Sets the snx vector for a solution. ::

    func (*Task) PutSnx ( whichsol int32, sux []float64 )

``sux []float64``
    The snx vector.

Sets the :math:`s_n^x` vector for a solution.  


PutSnxSlice
~~~~~~~~~~~

 Sets a slice of the snx vector for a solution. ::

    func (*Task) PutSnxSlice
        ( whichsol int32,
          first int32,
          last int32,
          snx []float64 )



Sets a slice of the :math:`s_n^x` vector for a solution.  


PutSolution
~~~~~~~~~~~

Inserts a solution. ::

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



Inserts a solution into the task. 


PutSolutionI
~~~~~~~~~~~~

 Sets the primal and dual solution information for a single constraint or variable. ::

    func (*Task) PutSolutionI
        ( accmode int32,
          i int32,
          whichsol int32,
          sk int32,
          x float64,
          sl float64,
          su float64,
          sn float64 )


``accmode int32``
     Defines whether solution information for a constraint or for a variable is modified.
``i int32``
    Index of the constraint or variable.
``sk int32``
    Status key of the constraint or variable.
``x float64``
    Solution value of the primal constraint or variable.
``sl float64``
     Solution value of the dual variable associated with the lower bound.
``su float64``
     Solution value of the dual variable associated with the upper bound.
``sn float64``
     Solution value of the dual variable associated with the cone constraint.


Sets the primal and dual solution information for a single
constraint or variable.



PutSolutionYI
~~~~~~~~~~~~~

 Inputs the dual variable of a solution. ::

    func (*Task) PutSolutionYI
        ( i int32,
          whichsol int32,
          y float64 )


``i int32``
    Index of the dual variable.
``y float64``
    Solution value of the dual variable.


Inputs the dual variable of a solution.



PutStrParam
~~~~~~~~~~~

Sets a string parameter. ::

    func (*Task) PutStrParam ( param int32, parvalue string )


Sets the value of a string parameter. 


PutStreamFunc
~~~~~~~~~~~~~

::

    func (t *Task) PutStreamFunc ( whichstream int32, fun func(string) )


Add a stream printer function to the task. ``whichstream`` should be a ``mosek.STREAM_...`` constant.


PutSuc
~~~~~~

 Sets the suc vector for a solution. ::

    func (*Task) PutSuc ( whichsol int32, suc []float64 )

``suc []float64``
    The suc vector.

Sets the :math:`s_u^c` vector for a solution.  


PutSucSlice
~~~~~~~~~~~

 Sets a slice of the suc vector for a solution. ::

    func (*Task) PutSucSlice
        ( whichsol int32,
          first int32,
          last int32,
          suc []float64 )



Sets a slice of the :math:`s_u^c` vector for a solution.  


PutSux
~~~~~~

 Sets the sux vector for a solution. ::

    func (*Task) PutSux ( whichsol int32, sux []float64 )

``sux []float64``
    The sux vector.

Sets the :math:`s_u^x` vector for a solution.  


PutSuxSlice
~~~~~~~~~~~

 Sets a slice of the sux vector for a solution. ::

    func (*Task) PutSuxSlice
        ( whichsol int32,
          first int32,
          last int32,
          sux []float64 )



Sets a slice of the :math:`s_u^x` vector for a solution.  


PutTaskName
~~~~~~~~~~~

Assigns a new name to the task. ::

    func (*Task) PutTaskName ( taskname string )


Assigns the name ``taskname`` to the task. 


PutVarBound
~~~~~~~~~~~

 Changes the bound for one variable. ::

    func (*Task) PutVarBound
        ( j int32,
          bk int32,
          bl float64,
          bu float64 )


``j int32``
    Index of the variable.
``bk int32``
    New bound key.
``bl float64``
    New lower bound.
``bu float64``
    New upper bound.


Changes the bounds for one variable.

If the a bound value specified is numerically larger than
``DPAR_DATA_TOL_BOUND_INF`` it is considered infinite and the bound key is
changed accordingly. If a bound value is numerically larger than
``DPAR_DATA_TOL_BOUND_WRN``, a warning will be displayed, but the bound is
inputted as specified.



PutVarBoundList
~~~~~~~~~~~~~~~

Changes the bounds of a list of variables. ::

    func (*Task) PutVarBoundList
        ( sub []int32,
          bkx []int32,
          blx []float64,
          bux []float64 )


``sub []int32``
    List of variable indexes.
``bkx []int32``
     New bound keys.
``blx []float64``
     New lower bound values.
``bux []float64``
     New upper bounds values.


Changes the bounds for one or more variables. If multiple bound changes are specified for a variable, then only the last change takes effect.



PutVarBoundSlice
~~~~~~~~~~~~~~~~

 Changes the bounds for a slice of the variables. ::

    func (*Task) PutVarBoundSlice
        ( first int32,
          last int32,
          bk []int32,
          bl []float64,
          bu []float64 )


``first int32``
    Index of the first variable in the slice.
``last int32``
    Index of the last variable in the slice plus 1.
``bk []int32``
    New bound keys.
``bl []float64``
    New lower bounds.
``bu []float64``
    New upper bounds.


Changes the bounds for a slice of the variables.



PutVarName
~~~~~~~~~~

 Puts the name of a variable. ::

    func (*Task) PutVarName ( j int32, name string )

``j int32``
    Index of the variable.
``name string``
    The variable name.


Puts the name of a variable.



PutVarType
~~~~~~~~~~

Sets the variable type of one variable. ::

    func (*Task) PutVarType ( j int32, vartype int32 )

``j int32``
    Index of the variable.
``vartype int32``
    The new variable type.

Sets the variable type of one variable. 


PutVarTypeList
~~~~~~~~~~~~~~

Sets the variable type for one or more variables. ::

    func (*Task) PutVarTypeList ( subj []int32, vartype []int32 )

``subj []int32``
     A list of variable indexes for which the variable type should be changed.
``vartype []int32``
     A list of variable types.


Sets the variable type for one or more variables, i.e.  variable number
:math:`\mathtt{subj}[k]` is assigned the variable type
:math:`\mathtt{vartype}[k]`.

If the same index is specified multiple times in ``subj`` only the last entry
takes effect.



PutXc
~~~~~

 Sets the xc vector for a solution. ::

    func (*Task) PutXc ( whichsol int32, xc []float64 ) ( xc []float64 )


``xc []float64``
    The xc vector.

Sets the :math:`x^c` vector for a solution.  


PutXcSlice
~~~~~~~~~~

 Sets a slice of the xc vector for a solution. ::

    func (*Task) PutXcSlice
        ( whichsol int32,
          first int32,
          last int32,
          xc []float64 )



Sets a slice of the :math:`x^c` vector for a solution.  


PutXx
~~~~~

 Sets the xx vector for a solution. ::

    func (*Task) PutXx ( whichsol int32, xx []float64 )

``xx []float64``
    The xx vector.

Sets the :math:`x^x` vector for a solution.  


PutXxSlice
~~~~~~~~~~

 Obtains a slice of the xx vector for a solution. ::

    func (*Task) PutXxSlice
        ( whichsol int32,
          first int32,
          last int32,
          xx []float64 )




Obtains a slice of the :math:`x^x` vector for a solution. 



PutY
~~~~

 Sets the y vector for a solution. ::

    func (*Task) PutY ( whichsol int32, y []float64 )

``y []float64``
    The y vector.

Sets the :math:`y` vector for a solution.  


PutYSlice
~~~~~~~~~

 Sets a slice of the y vector for a solution. ::

    func (*Task) PutYSlice
        ( whichsol int32,
          first int32,
          last int32,
          y []float64 )



Sets a slice of the :math:`y` vector for a solution.  


ReadData
~~~~~~~~

Reads problem data from a file. ::

    func (*Task) ReadData ( filename string )

``filename string``
     Input data file name.

Reads an optimization problem and associated data from a file.


ReadDataFormat
~~~~~~~~~~~~~~

Reads problem data from a file. ::

    func (*Task) ReadDataFormat
        ( filename string,
          format int32,
          compress int32 )


``filename string``
     Input data file name.
``format int32``
    File data format.
``compress int32``
    File compression type.

Reads an optimization problem and associated data from a file.


ReadParamFile
~~~~~~~~~~~~~

Reads a parameter file. ::

    func (*Task) ReadParamFile ( filename string )

``filename string``
     Input data file name.

Reads a parameter file. 


ReadSolution
~~~~~~~~~~~~

Reads a solution from a file. ::

    func (*Task) ReadSolution ( whichsol int32, filename string )


Reads a solution file and inserts the solution into the solution ``whichsol``.  


ReadSummary
~~~~~~~~~~~

Prints information about last file read. ::

    func (*Task) ReadSummary ( whichstream int32 )



Prints a short summary of last file that was read.



ReadTask
~~~~~~~~

 Load task data from a file. ::

    func (*Task) ReadTask ( filename string )

``filename string``
    Input file name.


Load task data from a file, replacing any data that already is in the task
object. All problem data are resorted, but if the file contains solutions, the
solution status after loading a file is still unknown, even if it was optimal
or otherwise well-defined when the file was dumped.

See section :ref:`doc.shared.taskformat` for a description of the Task format.



RemoveBarvars
~~~~~~~~~~~~~

 The function removes a number of symmetric matrix. ::

    func (*Task) RemoveBarvars ( subset []int32 )

``subset []int32``
     Indexes of symmetric matrix which should be removed.


The function removes a subset of the symmetric matrix 
from the optimization task. This implies that the existing
symmetric matrix are renumbered, for instance if
constraint 5 is removed then constraint 6 becomes
constraint 5 and so forth.



RemoveCones
~~~~~~~~~~~

Removes a conic constraint from the problem. ::

    func (*Task) RemoveCones ( subset []int32 )

``subset []int32``
     Indexes of cones which should be removed.


Removes a number conic constraint from the problem. 
In general, it is much more efficient to remove a cone with a high index than a low index.



RemoveCons
~~~~~~~~~~

 The function removes a number of constraints. ::

    func (*Task) RemoveCons ( subset []int32 )

``subset []int32``
     Indexes of constraints which should be removed.


The function removes a subset of the constraints 
from the optimization task. This implies that the existing
constraints are renumbered, for instance if
constraint 5 is removed then constraint 6 becomes
constraint 5 and so forth.



RemoveVars
~~~~~~~~~~

 The function removes a number of variables. ::

    func (*Task) RemoveVars ( subset []int32 )

``subset []int32``
     Indexes of variables which should be removed.


The function removes a subset of the variables 
from the optimization task. This implies that the existing
variables are renumbered, for instance if
constraint 5 is removed then constraint 6 becomes
constraint 5 and so forth.



ResizeTask
~~~~~~~~~~

Resizes an optimization task. ::

    func (*Task) ResizeTask
        ( maxnumcon int32,
          maxnumvar int32,
          maxnumcone int32,
          maxnumanz int64,
          maxnumqnz int64 )


``maxnumcon int32``
    New maximum number of constraints.
``maxnumvar int32``
    New maximum number of variables.
``maxnumcone int32``
    New maximum number of cones.
``maxnumanz int64``
    New maximum number of linear non-zero elements.
``maxnumqnz int64``
    New maximum number of quadratic non-zeros elements.


Sets the amount of preallocated space assigned for each type of data in an
optimization task.

It is never mandatory to call this function, since its only function is to give
a hint of the amount of data to preallocate for efficiency reasons.

Please note that the procedure is **destructive** in the sense that all
existing data stored in the task is destroyed.



SensitivityReport
~~~~~~~~~~~~~~~~~

Creates a sensitivity report. ::

    func (*Task) SensitivityReport ( whichstream int32 )



Reads a sensitivity format file from a location given by
``SPAR_SENSITIVITY_FILE_NAME`` and writes the result to the stream
``whichstream``. If ``SPAR_SENSITIVITY_RES_FILE_NAME`` is set to a non-empty
string, then the sensitivity report is also written to a file of this name.



SetDefaults
~~~~~~~~~~~

Resets all parameters values. ::

    func (*Task) SetDefaults (  )



Resets all the parameters to their default values.



SkToStr
~~~~~~~

Obtains a status key string. ::

    func (*Task) SkToStr ( sk int32 ) ( str string )


``sk int32``
    A valid status key.
``str string``
    String corresponding to the status key.


Obtains an explanatory string corresponding to a status key.



SolStaToStr
~~~~~~~~~~~

Obtains a solution status string. ::

    func (*Task) SolStaToStr ( solsta int32 ) ( str string )


``str string``
    String corresponding to the solution status.


Obtains an explanatory string corresponding to a solution status.



SolutionDef
~~~~~~~~~~~

Checks whether a solution is defined. ::

    func (*Task) SolutionDef ( whichsol int32 ) ( isdef bool )


``isdef bool``
    Is non-zero if the requested solution is defined.

Checks whether a solution is defined. 


SolutionSummary
~~~~~~~~~~~~~~~

Prints a short summary of the current solutions. ::

    func (*Task) SolutionSummary ( whichstream int32 )


Prints a short summary of the current solutions.  


SolveWithBasis
~~~~~~~~~~~~~~

 Solve a linear equation system involving a basis matrix. ::

    func (*Task) SolveWithBasis
        ( transp int32,
          numnz int32,
          sub []int32,
          val []float64 )
        ( numnz int32,
          sub []int32,
          val []float64 )


``transp int32``
     Controls which problem formulation is solved.
``numnz int32``
     Input (number of non-zeros in right-hand side) and output (number of non-zeros in solution vector).
``sub []int32``
     Input (indexes of non-zeros in right-hand side) and output (indexes of non-zeros in solution vector).
``val []float64``
     Input (right-hand side values) and output (solution vector values).


If a basic solution is available, then exactly :math:`\mathtt{numcon}`
basis variables are defined.  These :math:`\mathtt{numcon}` basis
variables are denoted the basis.  Associated with the basis is a basis
matrix denoted :math:`B`.  This function solves either the linear
equation system

.. math:: 
   :label: ais-eq-Bxb

   B \barX = b                       

or the system

.. math::
   :label: ais-eq-Btxb

   B^T \barX = b

for the unknowns :math:`\bar{X}`, with :math:`b` being a user-defined  vector.

In order to make sense of the solution :math:`\bar{X}` it is important
to know the ordering of the variables in the basis because the
ordering specifies how :math:`B` is constructed. When calling
``initbasissolve`` an ordering of the basis variables is
obtained, which can be used to deduce how MOSEK has constructed
:math:`B`. Indeed if the :math:`k`\ th basis variable is variable
:math:`x_j` it implies that


.. math:: B_{i,k} = A_{i,j}, ~i=\idxbeg,\ldots,\idxend{numcon}.


Otherwise if the :math:`k`\ th basis variable is variable :math:`x_j^c` it implies that

.. math::
    
  B_{i,k} = \left\{ \begin{array}{ll}
                          -1, & i = j, \\
                          0 , & i \neq j. \\
                      \end{array} 
              \right.


Given the knowledge of how :math:`B` is constructed it is possible to interpret the solution :math:`\bar{X}` correctly.

Please note that this function exploits the sparsity in the vector :math:`b` to speed up the computations.




StrToConeType
~~~~~~~~~~~~~

Obtains a cone type code. ::

    func (*Task) StrToConeType ( str string ) ( conetype int32 )


``str string``
    String corresponding to the cone type code.
``conetype int32``
    The cone type corresponding to str.


Obtains cone type code corresponding to a cone type string.



StrToSk
~~~~~~~

Obtains a status key. ::

    func (*Task) StrToSk ( str string ) ( sk int32 )


``str string``
    Status key string.
``sk int32``
    Status key corresponding to the string.


Obtains the status key corresponding to an explanatory string.



Syeig
~~~~~

Computes all eigenvalues of a symmetric dense matrix.::

    func (*Env) Syeig
        ( uplo int32,
          n int32,
          a []float64,
          w []float64 )
        ( w []float64 )


``uplo int32``
    Indicates whether the upper or lower triangular part is used.
``n int32``
     Dimension of the symmetric input matrix.
``a []float64``
     A symmetric matrix stored in column-major order. Only the lower-triangular part is used.
``w []float64``
     Array of minimum dimension n where eigenvalues will be stored.

Computes all eigenvalues of a real symmetric matrix :math:`A`. Eigenvalues are stored in the :math:`w` array.  


Syevd
~~~~~

Computes all the eigenvalue and eigenvectors of a symmetric dense matrix, and thus its eigenvalue decomposition.::

    func (*Env) Syevd
        ( uplo int32,
          n int32,
          a []float64,
          w []float64 )
        ( a []float64,
          w []float64 )


``uplo int32``
    Indicates whether the upper or lower triangular part is used.
``n int32``
     Dimension of symmetric input matrix.
``a []float64``
     A symmetric matrix stored in column-major order. Only the lower-triangular part is used. It will be overwritten on exit.
``w []float64``
     An array where eigenvalues will be stored. Its lenght must be at least the dimension of the input matrix.


Computes all the eigenvalues and eigenvectors a real symmetric matrix. 

Given the input matrix :math:`A\in \mathbf{R}}^{n\times n}`, this function returns a
vector :math:`w\in \mathbf{R}}^n`  containing the eigenvalues of :math:`A` and the
corresponding eigenvectors, stored in :math:`A` as well.

Therefore, this function compute the eigenvalue decomposition of :math:`A` as 

.. math:: A= U V U^T,

where :math:`V=diag(w)` and :math:`U` contains the eigen-vectors of :math:`A`.



Syrk
~~~~

Performs a rank-k update of a symmetric matrix.::

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


``uplo int32``
    Indicates whether the upper or lower triangular part of C is stored.
``trans int32``
    Indicates whether the matrix A must be transposed.
``n int32``
    Specifies the order of C.
``k int32``
    Indicates the number of rows or columns of A, and its rank.
``alpha float64``
    A scalar value multipling the result of the matrix multiplication.
``a []float64``
    The pointer to the array storing matrix A in a column-major format.
``beta float64``
    A scalar value that multiplies C.
``c []float64``
    The pointer to the array storing matrix C in a column-major format.


Performs a symmetric rank-:math:`k` update for a symmetric matrix. 

Given a symmetric matrix :math:`C\in \mathbf{R}}^{n\times n}`, two scalars
:math:`\alpha,\beta` and a matrix :math:`A` of rank :math:`k\leq n`, it
computes either 

.. math:: C= \alpha A A^T + \beta C,

or 

.. math:: C= \alpha A^T A + \beta C.

In the first case :math:`A\in \mathbf{R}}^{k\times n}`, in the second :math:`A\in
\mathbf{R}}^{n\times k}`.

Note that the results overwrite the matrix :math:`C`.



Toconic
~~~~~~~

Inplace reformulation of a QCQP to a COP::

    func (*Task) Toconic (  )



This function tries to reformulate a given Quadratically Constrained Quadratic
Optimization problem (QCQP) as a Conic Quadratic Optimization problem (CQO).
The first step of the reformulation is to convert the quadratic term of the
objective function as a constraint, if any. Then the following steps are
repeated for each quadratic constraint:

* a conic constraint is added along with a suitable number of auxiliary variables and constraints;
* the original quadratic constraint is not removed, but all its coefficients are zeroed out.


Note that the reformulation preserves all the original variables.

The conversion is performed in-place, i.e. the task passed as argument is
modified on exit. That also means that if the reformulation fails, i.e. the
given QCQP is not representable as a CQO, then the task has an undefined state.
In some cases, users may want to clone the task to ensure a clean copy is
preserved.



UpdateSolutionInfo
~~~~~~~~~~~~~~~~~~

Update the information items related to the solution.::

    func (*Task) UpdateSolutionInfo ( whichsol int32 )


Update the information items related to the solution.   


WriteData
~~~~~~~~~

Writes problem data to a file. ::

    func (*Task) WriteData ( filename string )

``filename string``
    Output file name.


Writes problem data associated with the optimization task to a file in one of
the supported formats. See Section :ref:`doc.shared.file_formats` for the complete list.

By default the data file format is determined by the file name extension. This
behaviour can be overridden by setting the ``IPAR_WRITE_DATA_FORMAT``
parameter.

MOSEK is able to read and write files in a compressed format (gzip). To write
in the compressed format append the extension ``.gz``.  E.g to write a gzip
compressed MPS file use the extension ``mps.gz``.

Please note that MPS, LP and OPF files require all variables to have unique
names. If a task contains no names, it is possible to write the file with
automatically generated anonymous names by setting the
``IPAR_WRITE_GENERIC_NAMES`` parameter to ``ON``.

Please note that if a general nonlinear function appears in the problem then
such function *cannot* be written to file and MOSEK will issue a warning.




WriteJsonSol
~~~~~~~~~~~~

Write a solution to a file. ::

    func (*Task) WriteJsonSol ( filename string )



Saves the current  solutions and solver information items in a JSON file.



WriteParamFile
~~~~~~~~~~~~~~

Writes all the parameters to a parameter file. ::

    func (*Task) WriteParamFile ( filename string )

``filename string``
    The name of parameter file.

Writes all the parameters to a parameter file. 


WriteSolution
~~~~~~~~~~~~~

Write a solution to a file. ::

    func (*Task) WriteSolution ( whichsol int32, filename string )



Saves the current basic, interior-point, or integer solution to a file.



WriteTask
~~~~~~~~~

Write a complete binary dump of the task data.::

    func (*Task) WriteTask ( filename string )

``filename string``
    Output file name.


Write a binary dump of the task data. This format saves all problem data, but
not callback-functions and general non-linear terms.

See section :ref:`doc.shared.taskformat` for a description of the Task format.


