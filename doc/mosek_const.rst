Constants in package ``mosek``
==============================

In the MOSEK C API most conststants are defined as enums. Go has no concept of enums,
all constants in the ``mosek`` package are defined as ``int32`` values.

solveform (``int32``)
~~~~~~~~~~~~~~~~
* `SOLVE_DUAL = 2`. The optimizer should solve the dual problem.
* `SOLVE_FREE = 0`.  The optimizer is free to solve either the primal or the dual problem.
* `SOLVE_PRIMAL = 1`. The optimizer should solve the primal problem.

problemitem (``int32``)
~~~~~~~~~~~~~~~~
* `PI_CON = 1`. Item is a constraint.
* `PI_CONE = 2`. Item is a cone.
* `PI_VAR = 0`. Item is a variable.

accmode (``int32``)
~~~~~~~~~~~~~~~~
* `ACC_CON = 1`. Access data by rows (constraint oriented)
* `ACC_VAR = 0`. Access data by columns (variable oriented)

sensitivitytype (``int32``)
~~~~~~~~~~~~~~~~
* `SENSITIVITY_TYPE_BASIS = 0`.  Basis sensitivity analysis is performed.
* `SENSITIVITY_TYPE_OPTIMAL_PARTITION = 1`.  Optimal partition sensitivity analysis is performed.

uplo (``int32``)
~~~~~~~~~~~~~~~~
* `UPLO_LO = 0`.  Lower part.
* `UPLO_UP = 1`.  Upper part

intpnthotstart (``int32``)
~~~~~~~~~~~~~~~~
* `INTPNT_HOTSTART_DUAL = 2`.  The interior-point optimizer exploits the dual solution only.
* `INTPNT_HOTSTART_NONE = 0`.  The interior-point optimizer performs a coldstart.
* `INTPNT_HOTSTART_PRIMAL = 1`.  The interior-point optimizer exploits the primal solution only.
* `INTPNT_HOTSTART_PRIMAL_DUAL = 3`.  The interior-point optimizer exploits both the primal and dual solution.

sparam (``int32``)
~~~~~~~~~~~~~~~~
* `SPAR_BAS_SOL_FILE_NAME = 0`. Name of the bas solution file.
* `SPAR_DATA_FILE_NAME = 1`. Data are read and written to this file.
* `SPAR_DEBUG_FILE_NAME = 2`. |mosek| debug file.
* `SPAR_INT_SOL_FILE_NAME = 3`. Name of the int solution file.
* `SPAR_ITR_SOL_FILE_NAME = 4`. Name of the itr solution file.
* `SPAR_MIO_DEBUG_STRING = 5`.  For internal use only.
* `SPAR_PARAM_COMMENT_SIGN = 6`.  Solution file comment character.
* `SPAR_PARAM_READ_FILE_NAME = 7`.  Modifications to the parameter database is read from this file.
* `SPAR_PARAM_WRITE_FILE_NAME = 8`. The parameter database is written to this file.
* `SPAR_READ_MPS_BOU_NAME = 9`.  Name of the BOUNDS vector used. An empty name means that the first BOUNDS vector is used.
* `SPAR_READ_MPS_OBJ_NAME = 10`.  Objective name in the MPS file.
* `SPAR_READ_MPS_RAN_NAME = 11`.  Name of the RANGE vector  used. An empty name means that the first RANGE vector is used.
* `SPAR_READ_MPS_RHS_NAME = 12`.  Name of the RHS used. An empty name means that the first RHS vector is used.
* `SPAR_REMOTE_ACCESS_TOKEN = 13`. An access token used to submit tasks to a remote MOSEK server
* `SPAR_SENSITIVITY_FILE_NAME = 14`.  Sensitivity report file name.
* `SPAR_SENSITIVITY_RES_FILE_NAME = 15`.  Name of the sensitivity report output file.
* `SPAR_SOL_FILTER_XC_LOW = 16`.  Solution file filter.
* `SPAR_SOL_FILTER_XC_UPR = 17`.  Solution file filter.
* `SPAR_SOL_FILTER_XX_LOW = 18`.  Solution file filter.
* `SPAR_SOL_FILTER_XX_UPR = 19`.  Solution file filter.
* `SPAR_STAT_FILE_NAME = 20`. Statistics file name.
* `SPAR_STAT_KEY = 21`. Key used when writing the summary file.
* `SPAR_STAT_NAME = 22`. Name used when writing the statistics file.
* `SPAR_WRITE_LP_GEN_VAR_NAME = 23`.  Added variable names in the LP files.

iparam (``int32``)
~~~~~~~~~~~~~~~~
* `IPAR_ANA_SOL_BASIS = 0`.  Controls whether the basis matrix is analyzed in solution analyzer.
* `IPAR_ANA_SOL_PRINT_VIOLATED = 1`.  Controls whether a list of violated constraints is printed.
* `IPAR_AUTO_SORT_A_BEFORE_OPT = 2`. Controls whether the elements in each column of A are sorted before an optimization is performed.
* `IPAR_AUTO_UPDATE_SOL_INFO = 3`.  Controls whether the solution information items are automatically updated after an optimization is performed.
* `IPAR_BASIS_SOLVE_USE_PLUS_ONE = 4`.  Controls the sign of the columns in the basis matrix corresponding to slack variables.
* `IPAR_BI_CLEAN_OPTIMIZER = 5`.  Controls which simplex optimizer is used in the clean-up phase.
* `IPAR_BI_IGNORE_MAX_ITER = 6`.  Turns on basis identification in case the interior-point optimizer is terminated due to maximum number of iterations.
* `IPAR_BI_IGNORE_NUM_ERROR = 7`.  Turns on basis identification in case the interior-point optimizer is terminated due to a numerical problem.
* `IPAR_BI_MAX_ITERATIONS = 8`.  Maximum number of iterations after basis identification.
* `IPAR_CACHE_LICENSE = 9`.  Control license caching.
* `IPAR_CHECK_CONVEXITY = 10`.  Specify the level of convexity check on quadratic problems
* `IPAR_COMPRESS_STATFILE = 11`.  Control compression of stat files.
* `IPAR_INFEAS_GENERIC_NAMES = 12`.  Controls the contents of the infeasibility report.
* `IPAR_INFEAS_PREFER_PRIMAL = 13`.  Controls which certificate is used if both primal- and dual- certificate of infeasibility is available.
* `IPAR_INFEAS_REPORT_AUTO = 14`.  Turns the feasibility report on or off.
* `IPAR_INFEAS_REPORT_LEVEL = 15`.  Controls the contents of the infeasibility report.
* `IPAR_INTPNT_BASIS = 16`.  Controls whether basis identification is performed.
* `IPAR_INTPNT_DIFF_STEP = 17`.  Controls whether different step sizes are allowed in the primal and dual space.
* `IPAR_INTPNT_HOTSTART = 18`.  Currently not in use.
* `IPAR_INTPNT_MAX_ITERATIONS = 19`.  Controls the maximum number of iterations allowed in the interior-point optimizer.
* `IPAR_INTPNT_MAX_NUM_COR = 20`.  Maximum number of correction steps.
* `IPAR_INTPNT_MAX_NUM_REFINEMENT_STEPS = 21`.  Maximum number of steps to be used by the iterative search direction refinement.
* `IPAR_INTPNT_MULTI_THREAD = 22`.  Controls whether the interior-point optimizers are allowed to employ multiple threads if more threads is available.
* `IPAR_INTPNT_OFF_COL_TRH = 23`.  Controls the aggressiveness of the offending column detection.
* `IPAR_INTPNT_ORDER_METHOD = 24`.  Controls the ordering strategy.
* `IPAR_INTPNT_REGULARIZATION_USE = 25`. Controls whether regularization is allowed.
* `IPAR_INTPNT_SCALING = 26`.  Controls how the problem is scaled before the interior-point optimizer is used.
* `IPAR_INTPNT_SOLVE_FORM = 27`.  Controls whether the primal or the dual problem is solved.
* `IPAR_INTPNT_STARTING_POINT = 28`. Starting point used by the interior-point optimizer.
* `IPAR_LICENSE_DEBUG = 29`.  Controls the license manager client debugging behavior.
* `IPAR_LICENSE_PAUSE_TIME = 30`.  Controls license manager client behavior.
* `IPAR_LICENSE_SUPPRESS_EXPIRE_WRNS = 31`.  Controls license manager client behavior.
* `IPAR_LICENSE_TRH_EXPIRY_WRN = 32`.  Controls when expiry warnings are issued.
* `IPAR_LICENSE_WAIT = 33`.  Controls if MOSEK should queue for a license if none is available.
* `IPAR_LOG = 34`.  Controls the amount of log information.
* `IPAR_LOG_ANA_PRO = 35`.  Controls amount of output from the problem analyzer.
* `IPAR_LOG_BI = 36`.  Controls the amount of output printed by the basis identification procedure. A higher level implies that more information is logged.
* `IPAR_LOG_BI_FREQ = 37`.  Controls the logging frequency.
* `IPAR_LOG_CHECK_CONVEXITY = 38`.  Controls logging in convexity check on quadratic problems. Set to a positive value to turn logging on. If a quadratic coefficient matrix is found to violate the requirement of PSD (NSD) then a list of negative (positive) pivot elements is printed. The absolute value of the pivot elements is also shown.
* `IPAR_LOG_CUT_SECOND_OPT = 39`.  Controls the reduction in the log levels for the second and any subsequent optimizations.
* `IPAR_LOG_EXPAND = 40`.  Controls the amount of logging when a data item such as the maximum number constrains is expanded.
* `IPAR_LOG_FACTOR = 41`.  If turned on, then the factor log lines are added to the log.
* `IPAR_LOG_FEAS_REPAIR = 42`.  Controls the amount of output printed when performing feasibility repair. A value higher than one means extensive logging.
* `IPAR_LOG_FILE = 43`.  If turned on, then some log info is printed when a file is written or read.
* `IPAR_LOG_HEAD = 44`.  If turned on, then a header line is added to the log.
* `IPAR_LOG_INFEAS_ANA = 45`.  Controls log level for the infeasibility analyzer.
* `IPAR_LOG_INTPNT = 46`.  Controls the amount of log information from the interior-point optimizers.
* `IPAR_LOG_MIO = 47`.  Controls the amount of log information from the mixed-integer optimizers.
* `IPAR_LOG_MIO_FREQ = 48`.  The mixed-integer optimizer logging frequency.
* `IPAR_LOG_OPTIMIZER = 49`.  Controls the amount of general optimizer information that is logged.
* `IPAR_LOG_ORDER = 50`.  If turned on, then factor lines are added to the log.
* `IPAR_LOG_PRESOLVE = 51`.  Controls amount of output printed by the presolve procedure. A higher level implies that more information is logged.
* `IPAR_LOG_RESPONSE = 52`.  Controls amount of output printed when response codes are reported. A higher level implies that more information is logged.
* `IPAR_LOG_SENSITIVITY = 53`.  Control logging in sensitivity analyzer.
* `IPAR_LOG_SENSITIVITY_OPT = 54`.  Control logging in sensitivity analyzer.
* `IPAR_LOG_SIM = 55`.  Controls the amount of log information from the simplex optimizers.
* `IPAR_LOG_SIM_FREQ = 56`.  Controls simplex logging frequency.
* `IPAR_LOG_SIM_MINOR = 57`.  Currently not in use.
* `IPAR_LOG_STORAGE = 58`.  Controls the memory related log information.
* `IPAR_MAX_NUM_WARNINGS = 59`.  Each warning is shown a limit number times controlled by this parameter. A negative value is identical to infinite number of times.
* `IPAR_MIO_BRANCH_DIR = 60`.  Controls whether the mixed-integer optimizer is branching up or down by default.
* `IPAR_MIO_CONSTRUCT_SOL = 61`.  Controls if an initial mixed integer solution should be constructed from the values of the integer variables.
* `IPAR_MIO_CUT_CLIQUE = 62`.  Controls whether clique cuts should be generated.
* `IPAR_MIO_CUT_CMIR = 63`.  Controls whether mixed integer rounding cuts should be generated.
* `IPAR_MIO_CUT_GMI = 64`.  Controls whether GMI cuts should be generated.
* `IPAR_MIO_CUT_IMPLIED_BOUND = 65`.  Controls whether implied bound cuts should be generated.
* `IPAR_MIO_CUT_KNAPSACK_COVER = 66`.  Controls whether knapsack cover cuts should be generated.
* `IPAR_MIO_CUT_SELECTION_LEVEL = 67`.  Controlls how aggresively generated cuts are selected to be inluded in the relaxation.
* `IPAR_MIO_HEURISTIC_LEVEL = 68`.  Controls the heuristic employed by the mixed-integer optimizer to locate an initial integer feasible solution.
* `IPAR_MIO_MAX_NUM_BRANCHES = 69`.  Maximum number of branches allowed during the branch and bound search.
* `IPAR_MIO_MAX_NUM_RELAXS = 70`.  Maximum number of relaxations in branch and bound search.
* `IPAR_MIO_MAX_NUM_SOLUTIONS = 71`.  Controls how many feasible solutions the mixed-integer optimizer investigates.
* `IPAR_MIO_MODE = 72`.  Turns on/off the mixed-integer mode.
* `IPAR_MIO_MT_USER_CB = 73`.  It true user callbacks are called from each thread used by this optimizer. If false the user callback is only called from a single thread.
* `IPAR_MIO_NODE_OPTIMIZER = 74`.  Controls which optimizer is employed at the non-root nodes in the mixed-integer optimizer.
* `IPAR_MIO_NODE_SELECTION = 75`.  Controls the node selection strategy employed by the mixed-integer optimizer.
* `IPAR_MIO_PERSPECTIVE_REFORMULATE = 76`.  Enables or disables perspective reformulation in presolve.
* `IPAR_MIO_PROBING_LEVEL = 77`.  Controls the amount of probing employed by the mixed-integer optimizer in presolve.
* `IPAR_MIO_RINS_MAX_NODES = 78`.  Maximum number of nodes in each call to RINS.
* `IPAR_MIO_ROOT_OPTIMIZER = 79`.  Controls which optimizer is employed at the root node in the mixed-integer optimizer.
* `IPAR_MIO_ROOT_REPEAT_PRESOLVE_LEVEL = 80`. Controls whether presolve can be repeated at root node.
* `IPAR_MIO_VB_DETECTION_LEVEL = 81`.  Controls how much effort is put into detecting variable bounds.
* `IPAR_MT_SPINCOUNT = 82`.  Set the number of iterations to spin before sleeping.
* `IPAR_NUM_THREADS = 83`.  Controls the number of threads employed by the optimizer. If set to 0 the number of threads used will be equal to the number of cores detected on the machine.
* `IPAR_OPF_MAX_TERMS_PER_LINE = 84`.  The maximum number of terms (linear and quadratic) per line when an OPF file is written.
* `IPAR_OPF_WRITE_HEADER = 85`.  Write a text header with date and |mosek| version in an OPF file.
* `IPAR_OPF_WRITE_HINTS = 86`.  Write a hint section with problem dimensions in the beginning of an OPF file.
* `IPAR_OPF_WRITE_PARAMETERS = 87`. Write a parameter section in an OPF file.
* `IPAR_OPF_WRITE_PROBLEM = 88`.  Write objective, constraints, bounds etc. to an OPF file.
* `IPAR_OPF_WRITE_SOL_BAS = 89`.  Controls what is written to the OPF files.
* `IPAR_OPF_WRITE_SOL_ITG = 90`.  Controls what is written to the OPF files.
* `IPAR_OPF_WRITE_SOL_ITR = 91`.  Controls what is written to the OPF files.
* `IPAR_OPF_WRITE_SOLUTIONS = 92`. Enable inclusion of solutions in the OPF files.
* `IPAR_OPTIMIZER = 93`.  Controls which optimizer is used to optimize the task.
* `IPAR_PARAM_READ_CASE_NAME = 94`.  If turned on, then names in the parameter file are case sensitive.
* `IPAR_PARAM_READ_IGN_ERROR = 95`.  If turned on, then errors in parameter settings is ignored.
* `IPAR_PRESOLVE_ELIMINATOR_MAX_FILL = 96`.  Maximum amount of fill-in created in one pivot during the elimination phase.
* `IPAR_PRESOLVE_ELIMINATOR_MAX_NUM_TRIES = 97`. Control the maximum number of times the eliminator is tried.
* `IPAR_PRESOLVE_LEVEL = 98`. Currently not used.
* `IPAR_PRESOLVE_LINDEP_ABS_WORK_TRH = 99`.  Controls linear dependency check in presolve.
* `IPAR_PRESOLVE_LINDEP_REL_WORK_TRH = 100`.  Controls linear dependency check in presolve.
* `IPAR_PRESOLVE_LINDEP_USE = 101`.  Controls whether the linear constraints are checked for linear dependencies.
* `IPAR_PRESOLVE_MAX_NUM_REDUCTIONS = 102`. Controls the maximum number of reductions performed by the presolve.
* `IPAR_PRESOLVE_USE = 103`. Controls whether the presolve is applied to a problem before it is optimized.
* `IPAR_PRIMAL_REPAIR_OPTIMIZER = 104`.  Controls which optimizer that is used to find the optimal repair.
* `IPAR_READ_DATA_COMPRESSED = 105`.  Controls the input file decompression.
* `IPAR_READ_DATA_FORMAT = 106`. Format of the data file to be read.
* `IPAR_READ_DEBUG = 107`.  Turns on additional debugging information when reading files.
* `IPAR_READ_KEEP_FREE_CON = 108`.  Controls whether the free constraints are included in the problem.
* `IPAR_READ_LP_DROP_NEW_VARS_IN_BOU = 109`.  Controls how the LP files are interpreted.
* `IPAR_READ_LP_QUOTED_NAMES = 110`.  If a name is in quotes when reading an LP file, the quotes will be removed.
* `IPAR_READ_MPS_FORMAT = 111`.  Controls how strictly the MPS file reader interprets the MPS format.
* `IPAR_READ_MPS_WIDTH = 112`.  Controls the maximal number of characters allowed in one line of the MPS file.
* `IPAR_READ_TASK_IGNORE_PARAM = 113`.  Controls what information is used from the task files.
* `IPAR_SENSITIVITY_ALL = 114`.  Controls sensitivity report behavior.
* `IPAR_SENSITIVITY_OPTIMIZER = 115`.  Controls which optimizer is used for optimal partition sensitivity analysis.
* `IPAR_SENSITIVITY_TYPE = 116`.  Controls which type of sensitivity analysis is to be performed.
* `IPAR_SIM_BASIS_FACTOR_USE = 117`.  Controls whether a (LU) factorization of the basis is used in a hot-start. Forcing a refactorization sometimes improves the stability of the simplex optimizers, but in most cases there is a performance penalty.
* `IPAR_SIM_DEGEN = 118`.  Controls how aggressively degeneration is handled.
* `IPAR_SIM_DUAL_CRASH = 119`.  Controls whether crashing is performed in the dual simplex optimizer.
* `IPAR_SIM_DUAL_PHASEONE_METHOD = 120`.  An experimental feature.
* `IPAR_SIM_DUAL_RESTRICT_SELECTION = 121`.  Controls how aggressively restricted selection is used.
* `IPAR_SIM_DUAL_SELECTION = 122`.  Controls the dual simplex strategy.
* `IPAR_SIM_EXPLOIT_DUPVEC = 123`.  Controls if the simplex optimizers are allowed to exploit duplicated columns.
* `IPAR_SIM_HOTSTART = 124`.  Controls the type of hot-start that the simplex optimizer perform.
* `IPAR_SIM_HOTSTART_LU = 125`.  Determines if the simplex optimizer should exploit the initial factorization.
* `IPAR_SIM_INTEGER = 126`.  An experimental feature.
* `IPAR_SIM_MAX_ITERATIONS = 127`.  Maximum number of iterations that can be used by a simplex optimizer.
* `IPAR_SIM_MAX_NUM_SETBACKS = 128`.  Controls how many set-backs that are allowed within a simplex optimizer.
* `IPAR_SIM_NON_SINGULAR = 129`.  Controls if the simplex optimizer ensures a non-singular basis, if possible.
* `IPAR_SIM_PRIMAL_CRASH = 130`.  Controls the simplex crash.
* `IPAR_SIM_PRIMAL_PHASEONE_METHOD = 131`.  An experimental feature.
* `IPAR_SIM_PRIMAL_RESTRICT_SELECTION = 132`.  Controls how aggressively restricted selection is used.
* `IPAR_SIM_PRIMAL_SELECTION = 133`.  Controls the primal simplex strategy.
* `IPAR_SIM_REFACTOR_FREQ = 134`.  Controls the basis refactoring frequency.
* `IPAR_SIM_REFORMULATION = 135`.  Controls if the simplex optimizers are allowed to reformulate the problem.
* `IPAR_SIM_SAVE_LU = 136`.  Controls if the LU factorization stored should be replaced with the LU factorization corresponding to the initial basis.
* `IPAR_SIM_SCALING = 137`.  Controls how much effort is used in scaling the problem before a simplex optimizer is used.
* `IPAR_SIM_SCALING_METHOD = 138`.  Controls how the problem is scaled before a simplex optimizer is used.
* `IPAR_SIM_SOLVE_FORM = 139`.  Controls whether the primal or the dual problem is solved by the primal-/dual-simplex optimizer.
* `IPAR_SIM_STABILITY_PRIORITY = 140`.  Controls how high priority the numerical stability should be given.
* `IPAR_SIM_SWITCH_OPTIMIZER = 141`.  Controls the simplex behavior.
* `IPAR_SOL_FILTER_KEEP_BASIC = 142`.  Controls the license manager client behavior.
* `IPAR_SOL_FILTER_KEEP_RANGED = 143`.  Control the contents of the solution files.
* `IPAR_SOL_READ_NAME_WIDTH = 144`.  Controls the input solution file format.
* `IPAR_SOL_READ_WIDTH = 145`.  Controls the input solution file format.
* `IPAR_SOLUTION_CALLBACK = 146`.  Indicates whether solution call-backs will be performed during the optimization.
* `IPAR_TIMING_LEVEL = 147`.  Controls the a amount of timing performed inside MOSEK.
* `IPAR_WRITE_BAS_CONSTRAINTS = 148`.  Controls the basic solution file format.
* `IPAR_WRITE_BAS_HEAD = 149`.  Controls the basic solution file format.
* `IPAR_WRITE_BAS_VARIABLES = 150`.  Controls the basic solution file format.
* `IPAR_WRITE_DATA_COMPRESSED = 151`.  Controls output file compression.
* `IPAR_WRITE_DATA_FORMAT = 152`.  Controls the output file format.
* `IPAR_WRITE_DATA_PARAM = 153`.  Controls output file data.
* `IPAR_WRITE_FREE_CON = 154`.  Controls the output file data.
* `IPAR_WRITE_GENERIC_NAMES = 155`.  Controls the output file data.
* `IPAR_WRITE_GENERIC_NAMES_IO = 156`. Index origin used in  generic names.
* `IPAR_WRITE_IGNORE_INCOMPATIBLE_ITEMS = 157`.  Controls if the writer ignores incompatible problem items when writing files.
* `IPAR_WRITE_INT_CONSTRAINTS = 158`.  Controls the integer solution file format.
* `IPAR_WRITE_INT_HEAD = 159`.  Controls the integer solution file format.
* `IPAR_WRITE_INT_VARIABLES = 160`.  Controls the integer solution file format.
* `IPAR_WRITE_LP_FULL_OBJ = 161`. Write full linear objective
* `IPAR_WRITE_LP_LINE_WIDTH = 162`.  Controls the LP output file format.
* `IPAR_WRITE_LP_QUOTED_NAMES = 163`.  Controls LP output file format.
* `IPAR_WRITE_LP_STRICT_FORMAT = 164`.  Controls whether LP  output files satisfy the LP format strictly.
* `IPAR_WRITE_LP_TERMS_PER_LINE = 165`.  Controls the LP output file format.
* `IPAR_WRITE_MPS_FORMAT = 166`.  Controls in which format the MPS is written.
* `IPAR_WRITE_MPS_INT = 167`.  Controls the output file data.
* `IPAR_WRITE_PRECISION = 168`.  Controls data precision employed in when writing an MPS file.
* `IPAR_WRITE_SOL_BARVARIABLES = 169`.  Controls the solution file format.
* `IPAR_WRITE_SOL_CONSTRAINTS = 170`.  Controls the solution file format.
* `IPAR_WRITE_SOL_HEAD = 171`.  Controls solution file format.
* `IPAR_WRITE_SOL_IGNORE_INVALID_NAMES = 172`.  Controls whether the user specified names are employed even if they are invalid names.
* `IPAR_WRITE_SOL_VARIABLES = 173`.  Controls the solution file format.
* `IPAR_WRITE_TASK_INC_SOL = 174`.  Controls whether the solutions are  stored in the task file too.
* `IPAR_WRITE_XML_MODE = 175`.  Controls if linear coefficients should be written by row or column when writing in the XML file format.

solsta (``int32``)
~~~~~~~~~~~~~~~~
* `SOL_STA_DUAL_FEAS = 3`. The solution is dual feasible.
* `SOL_STA_DUAL_ILLPOSED_CER = 15`.  The solution is a certificate that the dual problem is illposed.
* `SOL_STA_DUAL_INFEAS_CER = 6`.  The solution is a certificate of dual infeasibility.
* `SOL_STA_INTEGER_OPTIMAL = 16`. The primal solution is integer optimal.
* `SOL_STA_NEAR_DUAL_FEAS = 10`. The solution is nearly dual feasible.
* `SOL_STA_NEAR_DUAL_INFEAS_CER = 13`.  The solution is almost a certificate of dual infeasibility.
* `SOL_STA_NEAR_INTEGER_OPTIMAL = 17`. The primal solution is near integer optimal.
* `SOL_STA_NEAR_OPTIMAL = 8`. The solution is nearly optimal.
* `SOL_STA_NEAR_PRIM_AND_DUAL_FEAS = 11`.  The solution is nearly both primal and dual feasible.
* `SOL_STA_NEAR_PRIM_FEAS = 9`. The solution is nearly primal feasible.
* `SOL_STA_NEAR_PRIM_INFEAS_CER = 12`.  The solution is almost a certificate of primal infeasibility.
* `SOL_STA_OPTIMAL = 1`. The solution is optimal.
* `SOL_STA_PRIM_AND_DUAL_FEAS = 4`. The solution is both primal and dual feasible.
* `SOL_STA_PRIM_FEAS = 2`. The solution is primal feasible.
* `SOL_STA_PRIM_ILLPOSED_CER = 14`.  The solution is a certificate that the primal problem is illposed.
* `SOL_STA_PRIM_INFEAS_CER = 5`.  The solution is a certificate of primal infeasibility.
* `SOL_STA_UNKNOWN = 0`. Status of the solution is unknown.

objsense (``int32``)
~~~~~~~~~~~~~~~~
* `OBJECTIVE_SENSE_MAXIMIZE = 1`. The problem should be maximized.
* `OBJECTIVE_SENSE_MINIMIZE = 0`. The problem should be minimized.

solitem (``int32``)
~~~~~~~~~~~~~~~~
* `SOL_ITEM_SLC = 3`.  Lagrange multipliers for lower bounds on the constraints.
* `SOL_ITEM_SLX = 5`.  Lagrange multipliers for lower bounds on the variables.
* `SOL_ITEM_SNX = 7`.  Lagrange multipliers corresponding to the conic constraints on the variables.
* `SOL_ITEM_SUC = 4`.  Lagrange multipliers for upper bounds on the constraints.
* `SOL_ITEM_SUX = 6`.  Lagrange multipliers for upper bounds on the variables.
* `SOL_ITEM_XC = 0`. Solution for the constraints.
* `SOL_ITEM_XX = 1`. Variable solution.
* `SOL_ITEM_Y = 2`. Lagrange multipliers for equations.

boundkey (``int32``)
~~~~~~~~~~~~~~~~
* `BK_FR = 3`. The constraint or variable is free.
* `BK_FX = 2`. The constraint or variable is fixed.
* `BK_LO = 0`.  The constraint or variable has a finite lower bound and an infinite upper bound.
* `BK_RA = 4`. The constraint or variable is ranged.
* `BK_UP = 1`.  The constraint or variable has an infinite lower bound and an finite upper bound.

basindtype (``int32``)
~~~~~~~~~~~~~~~~
* `BI_ALWAYS = 1`.  Basis identification is always performed even if the interior-point optimizer terminates abnormally.
* `BI_IF_FEASIBLE = 3`.  Basis identification is not performed if the interior-point optimizer terminates with a problem status saying that the problem is primal or dual infeasible.
* `BI_NEVER = 0`. Never do basis identification.
* `BI_NO_ERROR = 2`.  Basis identification is performed if the interior-point optimizer terminates without an error.
* `BI_RESERVERED = 4`. Not currently in use.

branchdir (``int32``)
~~~~~~~~~~~~~~~~
* `BRANCH_DIR_DOWN = 2`.  The mixed-integer optimizer always chooses the down branch first.
* `BRANCH_DIR_FAR = 4`.  Branch in direction farthest from selected fractional variable.
* `BRANCH_DIR_FREE = 0`.  The mixed-integer optimizer decides which branch to choose.
* `BRANCH_DIR_GUIDED = 6`.  Branch in direction of current incumbent.
* `BRANCH_DIR_NEAR = 3`.  Branch in direction nearest to selected fractional variable.
* `BRANCH_DIR_PSEUDOCOST = 7`.  Branch based on the pseudocost of the variable.
* `BRANCH_DIR_ROOT_LP = 5`.  Chose direction based on root lp value of selected variable.
* `BRANCH_DIR_UP = 1`.  The mixed-integer optimizer always chooses the up branch first.

liinfitem (``int32``)
~~~~~~~~~~~~~~~~
* `LIINF_BI_CLEAN_DUAL_DEG_ITER = 0`.  Number of dual degenerate clean iterations performed in the basis identification.
* `LIINF_BI_CLEAN_DUAL_ITER = 1`.  Number of dual clean iterations performed in the basis identification.
* `LIINF_BI_CLEAN_PRIMAL_DEG_ITER = 2`.  Number of primal degenerate clean iterations performed in the basis identification.
* `LIINF_BI_CLEAN_PRIMAL_DUAL_DEG_ITER = 3`.  Number of primal-dual degenerate clean iterations performed in the basis identification.
* `LIINF_BI_CLEAN_PRIMAL_DUAL_ITER = 4`.  Number of primal-dual clean iterations performed in the basis identification.
* `LIINF_BI_CLEAN_PRIMAL_DUAL_SUB_ITER = 5`.  Number of primal-dual subproblem clean iterations performed in the basis identification.
* `LIINF_BI_CLEAN_PRIMAL_ITER = 6`.  Number of primal clean iterations performed in the basis identification.
* `LIINF_BI_DUAL_ITER = 7`.  Number of dual pivots performed in the basis identification.
* `LIINF_BI_PRIMAL_ITER = 8`.  Number of primal pivots performed in the basis identification.
* `LIINF_INTPNT_FACTOR_NUM_NZ = 9`. Number of non-zeros in factorization.
* `LIINF_MIO_INTPNT_ITER = 10`.  Number of interior-point iterations performed by the mixed-integer optimizer.
* `LIINF_MIO_PRESOLVED_ANZ = 11`.  Number of  non-zero entries in the constraint matrix of presolved problem.
* `LIINF_MIO_SIM_MAXITER_SETBACKS = 12`.  Number of times the the simplex optimizer has hit the maximum iteration limit when re-optimizing.
* `LIINF_MIO_SIMPLEX_ITER = 13`.  Number of simplex iterations performed by the mixed-integer optimizer.
* `LIINF_RD_NUMANZ = 14`. Number of non-zeros in A that is read.
* `LIINF_RD_NUMQNZ = 15`. Number of Q non-zeros.

simhotstart (``int32``)
~~~~~~~~~~~~~~~~
* `SIM_HOTSTART_FREE = 1`.  The simplex optimize chooses the hot-start type.
* `SIM_HOTSTART_NONE = 0`.  The simplex optimizer performs a coldstart.
* `SIM_HOTSTART_STATUS_KEYS = 2`.  Only the status keys of the constraints and variables are used to choose the type of hot-start.

callbackcode (``int32``)
~~~~~~~~~~~~~~~~
* `CALLBACK_BEGIN_BI = 0`.  The basis identification procedure has been started.
* `CALLBACK_BEGIN_CONIC = 1`.  The call-back function is called when the conic optimizer is started.
* `CALLBACK_BEGIN_DUAL_BI = 2`.  The call-back function is called from within the basis identification procedure when the dual phase is started.
* `CALLBACK_BEGIN_DUAL_SENSITIVITY = 3`.  Dual sensitivity analysis is started.
* `CALLBACK_BEGIN_DUAL_SETUP_BI = 4`.  The call-back function is called when the dual BI phase is started.
* `CALLBACK_BEGIN_DUAL_SIMPLEX = 5`.  The call-back function is called when the dual simplex optimizer started.
* `CALLBACK_BEGIN_DUAL_SIMPLEX_BI = 6`.  The call-back function is called from within the basis identification procedure when the dual simplex clean-up phase is started.
* `CALLBACK_BEGIN_FULL_CONVEXITY_CHECK = 7`.  Begin full convexity check.
* `CALLBACK_BEGIN_INFEAS_ANA = 8`.  The call-back function is called when the infeasibility analyzer is started.
* `CALLBACK_BEGIN_INTPNT = 9`.  The call-back function is called when the interior-point optimizer is started.
* `CALLBACK_BEGIN_LICENSE_WAIT = 10`.  Begin waiting for license.
* `CALLBACK_BEGIN_MIO = 11`.  The call-back function is called when the mixed-integer optimizer is started.
* `CALLBACK_BEGIN_OPTIMIZER = 12`.  The call-back function is called when the optimizer is started.
* `CALLBACK_BEGIN_PRESOLVE = 13`.  The call-back function is called when the presolve is started.
* `CALLBACK_BEGIN_PRIMAL_BI = 14`.  The call-back function is called from within the basis identification procedure when the primal phase is started.
* `CALLBACK_BEGIN_PRIMAL_DUAL_SIMPLEX = 15`.  The call-back function is called when the primal-dual simplex optimizer is started.
* `CALLBACK_BEGIN_PRIMAL_DUAL_SIMPLEX_BI = 16`.  The call-back function is called from within the basis identification procedure when the primal-dual simplex clean-up phase is started.
* `CALLBACK_BEGIN_PRIMAL_REPAIR = 17`.  Begin primal feasibility repair.
* `CALLBACK_BEGIN_PRIMAL_SENSITIVITY = 18`.  Primal sensitivity analysis is started.
* `CALLBACK_BEGIN_PRIMAL_SETUP_BI = 19`.  The call-back function is called when the primal BI setup is started.
* `CALLBACK_BEGIN_PRIMAL_SIMPLEX = 20`.  The call-back function is called when the primal simplex optimizer is started.
* `CALLBACK_BEGIN_PRIMAL_SIMPLEX_BI = 21`.  The call-back function is called from within the basis identification procedure when the primal simplex clean-up phase is started.
* `CALLBACK_BEGIN_QCQO_REFORMULATE = 22`.  Begin QCQO reformulation.
* `CALLBACK_BEGIN_READ = 23`. |mosek| has started reading a problem file.
* `CALLBACK_BEGIN_ROOT_CUTGEN = 24`.  The call-back function is called when root cut generation is started.
* `CALLBACK_BEGIN_SIMPLEX = 25`.  The call-back function is called when the simplex optimizer is started.
* `CALLBACK_BEGIN_SIMPLEX_BI = 26`.  The call-back function is called from within the basis identification procedure when the simplex clean-up phase is started.
* `CALLBACK_BEGIN_TO_CONIC = 27`.  Begin conic reformulation.
* `CALLBACK_BEGIN_WRITE = 28`. |mosek| has started writing a problem file.
* `CALLBACK_CONIC = 29`.  The call-back function is called from within the conic optimizer after the information database has been updated.
* `CALLBACK_DUAL_SIMPLEX = 30`.  The call-back function is called from within the dual simplex optimizer.
* `CALLBACK_END_BI = 31`.  The call-back function is called when the basis identification procedure is terminated.
* `CALLBACK_END_CONIC = 32`.  The call-back function is called when the conic optimizer is terminated.
* `CALLBACK_END_DUAL_BI = 33`.  The call-back function is called from within the basis identification procedure when the dual phase is terminated.
* `CALLBACK_END_DUAL_SENSITIVITY = 34`.  Dual sensitivity analysis is terminated.
* `CALLBACK_END_DUAL_SETUP_BI = 35`.  The call-back function is called when the dual BI phase is terminated.
* `CALLBACK_END_DUAL_SIMPLEX = 36`.  The call-back function is called when the dual simplex optimizer is terminated.
* `CALLBACK_END_DUAL_SIMPLEX_BI = 37`.  The call-back function is called from within the basis identification procedure when the dual clean-up phase is terminated.
* `CALLBACK_END_FULL_CONVEXITY_CHECK = 38`.  End full convexity check.
* `CALLBACK_END_INFEAS_ANA = 39`.  The call-back function is called when the infeasibility analyzer is terminated.
* `CALLBACK_END_INTPNT = 40`.  The call-back function is called when the interior-point optimizer is terminated.
* `CALLBACK_END_LICENSE_WAIT = 41`.  End waiting for license.
* `CALLBACK_END_MIO = 42`.  The call-back function is called when the mixed-integer optimizer is terminated.
* `CALLBACK_END_OPTIMIZER = 43`.  The call-back function is called when the optimizer is terminated.
* `CALLBACK_END_PRESOLVE = 44`.  The call-back function is called when the presolve is completed.
* `CALLBACK_END_PRIMAL_BI = 45`.  The call-back function is called from within the basis identification procedure when the primal phase is terminated.
* `CALLBACK_END_PRIMAL_DUAL_SIMPLEX = 46`.  The call-back function is called when the primal-dual simplex optimizer is terminated.
* `CALLBACK_END_PRIMAL_DUAL_SIMPLEX_BI = 47`.  The call-back function is called from within the basis identification procedure when the primal-dual clean-up phase is terminated.
* `CALLBACK_END_PRIMAL_REPAIR = 48`.  End primal feasibility repair.
* `CALLBACK_END_PRIMAL_SENSITIVITY = 49`.  Primal sensitivity analysis is terminated.
* `CALLBACK_END_PRIMAL_SETUP_BI = 50`.  The call-back function is called when the primal BI setup is terminated.
* `CALLBACK_END_PRIMAL_SIMPLEX = 51`.  The call-back function is called when the primal simplex optimizer is terminated.
* `CALLBACK_END_PRIMAL_SIMPLEX_BI = 52`.  The call-back function is called from within the basis identification procedure when the primal clean-up phase is terminated.
* `CALLBACK_END_QCQO_REFORMULATE = 53`.  End QCQO reformulation.
* `CALLBACK_END_READ = 54`. |mosek| has finished reading a problem file.
* `CALLBACK_END_ROOT_CUTGEN = 55`.  The call-back function is called when root cut generation is is terminated.
* `CALLBACK_END_SIMPLEX = 56`.  The call-back function is called when the simplex optimizer is terminated.
* `CALLBACK_END_SIMPLEX_BI = 57`.  The call-back function is called from within the basis identification procedure when the simplex clean-up phase is terminated.
* `CALLBACK_END_TO_CONIC = 58`.  End conic reformulation.
* `CALLBACK_END_WRITE = 59`. |mosek| has finished writing a problem file.
* `CALLBACK_IM_BI = 60`.  The call-back function is called from within the basis identification procedure at an intermediate point.
* `CALLBACK_IM_CONIC = 61`.  The call-back function is called at an intermediate stage within the conic optimizer where the information database has not been updated.
* `CALLBACK_IM_DUAL_BI = 62`.  The call-back function is called from within the basis identification procedure at an intermediate point in the dual phase.
* `CALLBACK_IM_DUAL_SENSIVITY = 63`.  The call-back function is called at an intermediate stage of the dual sensitivity analysis.
* `CALLBACK_IM_DUAL_SIMPLEX = 64`.  The call-back function is called at an intermediate point in the dual simplex optimizer.
* `CALLBACK_IM_FULL_CONVEXITY_CHECK = 65`.  The call-back function is called at an intermediate stage of the full convexity check.
* `CALLBACK_IM_INTPNT = 66`.  The call-back function is called at an intermediate stage within the interior-point optimizer where the information database has not been updated.
* `CALLBACK_IM_LICENSE_WAIT = 67`.  |mosek| is waiting for a license.
* `CALLBACK_IM_LU = 68`.  The call-back function is called from within the LU factorization procedure at an intermediate point.
* `CALLBACK_IM_MIO = 69`.  The call-back function is called at an intermediate point in the mixed-integer optimizer.
* `CALLBACK_IM_MIO_DUAL_SIMPLEX = 70`.  The call-back function is called at an intermediate point in the mixed-integer optimizer while running the dual simplex optimizer.
* `CALLBACK_IM_MIO_INTPNT = 71`.  The call-back function is called at an intermediate point in the mixed-integer optimizer while running the interior-point optimizer.
* `CALLBACK_IM_MIO_PRIMAL_SIMPLEX = 72`.  The call-back function is called at an intermediate point in the mixed-integer optimizer while running the primal simplex optimizer.
* `CALLBACK_IM_ORDER = 73`.  The call-back function is called from within the matrix ordering procedure at an intermediate point.
* `CALLBACK_IM_PRESOLVE = 74`.  The call-back function is called from within the presolve procedure at an intermediate stage.
* `CALLBACK_IM_PRIMAL_BI = 75`.  The call-back function is called from within the basis identification procedure at an intermediate point in the primal phase.
* `CALLBACK_IM_PRIMAL_DUAL_SIMPLEX = 76`.  The call-back function is called at an intermediate point in the primal-dual simplex optimizer.
* `CALLBACK_IM_PRIMAL_SENSIVITY = 77`.  The call-back function is called at an intermediate stage of the primal sensitivity analysis.
* `CALLBACK_IM_PRIMAL_SIMPLEX = 78`.  The call-back function is called at an intermediate point in the primal simplex optimizer.
* `CALLBACK_IM_QO_REFORMULATE = 79`.  The call-back function is called at an intermediate stage of the conic quadratic reformulation.
* `CALLBACK_IM_READ = 80`. Intermediate stage in reading.
* `CALLBACK_IM_ROOT_CUTGEN = 81`.  The call-back is called from within root cut generation at an intermediate stage.
* `CALLBACK_IM_SIMPLEX = 82`.  The call-back function is called from within the simplex optimizer at an intermediate point.
* `CALLBACK_IM_SIMPLEX_BI = 83`.  The call-back function is called from within the basis identification procedure at an intermediate point in the simplex clean-up phase.
* `CALLBACK_INTPNT = 84`.  The call-back function is called from within the interior-point optimizer after the information database has been updated.
* `CALLBACK_NEW_INT_MIO = 85`.  The call-back function is called after a new integer solution has been located by the mixed-integer optimizer.
* `CALLBACK_PRIMAL_SIMPLEX = 86`.  The call-back function is called from within the primal simplex optimizer.
* `CALLBACK_READ_OPF = 87`.  The call-back function is called from the OPF reader.
* `CALLBACK_READ_OPF_SECTION = 88`. A chunk of Q non-zeros has been read from a problem file.
* `CALLBACK_SOLVING_REMOTE = 89`.  The call-back function is called while the task is being solved on a remote server.
* `CALLBACK_UPDATE_DUAL_BI = 90`.  The call-back function is called from within the basis identification procedure at an intermediate point in the dual phase.
* `CALLBACK_UPDATE_DUAL_SIMPLEX = 91`.  The call-back function is called in the dual simplex optimizer.
* `CALLBACK_UPDATE_DUAL_SIMPLEX_BI = 92`.  The call-back function is called from within the basis identification procedure at an intermediate point in the dual simplex clean-up phase.
* `CALLBACK_UPDATE_PRESOLVE = 93`.  The call-back function is called from within the presolve procedure.
* `CALLBACK_UPDATE_PRIMAL_BI = 94`.  The call-back function is called from within the basis identification procedure at an intermediate point in the primal phase.
* `CALLBACK_UPDATE_PRIMAL_DUAL_SIMPLEX = 95`.  The call-back function is called  in the primal-dual simplex optimizer.
* `CALLBACK_UPDATE_PRIMAL_DUAL_SIMPLEX_BI = 96`.  The call-back function is called from within the basis identification procedure at an intermediate point in the primal simplex clean-up phase.
* `CALLBACK_UPDATE_PRIMAL_SIMPLEX = 97`.  The call-back function is called  in the primal simplex optimizer.
* `CALLBACK_UPDATE_PRIMAL_SIMPLEX_BI = 98`.  The call-back function is called from within the basis identification procedure at an intermediate point in the primal simplex clean-up phase.
* `CALLBACK_WRITE_OPF = 99`.  The call-back function is called from the OPF writer.

symmattype (``int32``)
~~~~~~~~~~~~~~~~
* `SYMMAT_TYPE_SPARSE = 0`. Sparse symmetric matrix.

feature (``int32``)
~~~~~~~~~~~~~~~~
* `FEATURE_PTON = 1`. Nonlinear extension.
* `FEATURE_PTS = 0`. Base system.

mark (``int32``)
~~~~~~~~~~~~~~~~
* `MARK_LO = 0`.  The lower bound is selected for sensitivity analysis.
* `MARK_UP = 1`.  The upper bound is selected for sensitivity analysis.

conetype (``int32``)
~~~~~~~~~~~~~~~~
* `CT_QUAD = 0`. The cone is a quadratic cone.
* `CT_RQUAD = 1`. The cone is a rotated quadratic cone.

streamtype (``int32``)
~~~~~~~~~~~~~~~~
* `STREAM_ERR = 2`. Error stream. Error messages are written to this stream.
* `STREAM_LOG = 0`. Log stream. Contains the aggregated contents of all other streams. This means that a message written to any other stream will also be written to this stream.
* `STREAM_MSG = 1`. Message stream. Log information relating to performance and progress of the optimization is written to this stream.
* `STREAM_WRN = 3`. Warning stream. Warning messages are written to this stream.

iomode (``int32``)
~~~~~~~~~~~~~~~~
* `IOMODE_READ = 0`. The file is read-only.
* `IOMODE_READWRITE = 2`. The file is to read and written.
* `IOMODE_WRITE = 1`.  The file is write-only. If the file exists then it is truncated when it is opened. Otherwise it is created when it is opened.

simseltype (``int32``)
~~~~~~~~~~~~~~~~
* `SIM_SELECTION_ASE = 2`.  The optimizer uses approximate steepest-edge pricing.
* `SIM_SELECTION_DEVEX = 3`.  The optimizer uses devex steepest-edge pricing (or if it is not available an approximate steep-edge selection).
* `SIM_SELECTION_FREE = 0`. The optimizer chooses the pricing strategy.
* `SIM_SELECTION_FULL = 1`. The optimizer uses full pricing.
* `SIM_SELECTION_PARTIAL = 5`.  The optimizer uses a partial selection approach. The approach is usually beneficial if the number of variables is much larger than  the number of constraints.
* `SIM_SELECTION_SE = 4`.  The optimizer uses steepest-edge selection (or if it is not available an approximate steep-edge selection).

msgkey (``int32``)
~~~~~~~~~~~~~~~~
* `MSG_MPS_SELECTED = 1100`. 
* `MSG_READING_FILE = 1000`. 
* `MSG_WRITING_FILE = 1001`. 

miomode (``int32``)
~~~~~~~~~~~~~~~~
* `MIO_MODE_IGNORED = 0`.  The integer constraints are ignored and the problem is solved as a continuous problem.
* `MIO_MODE_SATISFIED = 1`. Integer restrictions should be satisfied.

dinfitem (``int32``)
~~~~~~~~~~~~~~~~
* `DINF_BI_CLEAN_DUAL_TIME = 0`.  Time  spent within the dual clean-up optimizer of the basis identification procedure since its invocation.
* `DINF_BI_CLEAN_PRIMAL_DUAL_TIME = 1`.  Time spent within the primal-dual clean-up optimizer of the basis identification procedure since its invocation.
* `DINF_BI_CLEAN_PRIMAL_TIME = 2`.  Time spent within the primal clean-up optimizer of the basis identification procedure since its invocation.
* `DINF_BI_CLEAN_TIME = 3`.  Time spent within the clean-up phase of the basis identification procedure since its invocation.
* `DINF_BI_DUAL_TIME = 4`.  Time spent within the dual phase basis identification procedure since its invocation.
* `DINF_BI_PRIMAL_TIME = 5`.  Time  spent within the primal phase of the basis identification procedure since its invocation.
* `DINF_BI_TIME = 6`.  Time spent within the basis identification procedure since its invocation.
* `DINF_INTPNT_DUAL_FEAS = 7`.  Dual feasibility measure reported by the interior-point optimizer. (For the interior-point optimizer this measure does not directly related to the original problem because a homogeneous model is employed.)
* `DINF_INTPNT_DUAL_OBJ = 8`.  Dual objective value reported by the interior-point optimizer.
* `DINF_INTPNT_FACTOR_NUM_FLOPS = 9`.  An estimate of the number of flops used in the factorization.
* `DINF_INTPNT_OPT_STATUS = 10`.  This measure should converge to :math:`+1` if the problem has a primal-dual optimal solution, and converge to :math:`-1` if problem is (strictly) primal or dual infeasible. Furthermore, if the measure converges to 0 the problem is usually ill-posed.
* `DINF_INTPNT_ORDER_TIME = 11`. Order time (in seconds).
* `DINF_INTPNT_PRIMAL_FEAS = 12`.  Primal feasibility measure reported by the interior-point optimizers. (For the interior-point optimizer this measure does not directly related to the original problem because a homogeneous model is employed).
* `DINF_INTPNT_PRIMAL_OBJ = 13`.  Primal objective value reported by the interior-point optimizer.
* `DINF_INTPNT_TIME = 14`.  Time spent within the interior-point optimizer since its invocation.
* `DINF_MIO_CLIQUE_SEPARATION_TIME = 15`.  Seperation time for clique cuts.
* `DINF_MIO_CMIR_SEPARATION_TIME = 16`.  Seperation time for CMIR cuts.
* `DINF_MIO_CONSTRUCT_SOLUTION_OBJ = 17`.  If MOSEK has successfully constructed an integer feasible solution, then this item contains the optimal objective value corresponding to the feasible solution.
* `DINF_MIO_DUAL_BOUND_AFTER_PRESOLVE = 18`.  Value of the dual bound after presolve but before cut generation.
* `DINF_MIO_GMI_SEPARATION_TIME = 19`.  Seperation time for GMI cuts.
* `DINF_MIO_HEURISTIC_TIME = 20`. Total time spent in the optimizer.
* `DINF_MIO_IMPLIED_BOUND_TIME = 21`.  Seperation time for implied bound cuts.
* `DINF_MIO_KNAPSACK_COVER_SEPARATION_TIME = 22`.  Seperation time for knapsack cover.
* `DINF_MIO_OBJ_ABS_GAP = 23`.  If the mixed-integer optimizer has computed a feasible solution and a bound, this contains the absolute gap.
* `DINF_MIO_OBJ_BOUND = 24`.  The best bound on the objective value known.
* `DINF_MIO_OBJ_INT = 25`.  The primal objective value corresponding to the best integer feasible solution.
* `DINF_MIO_OBJ_REL_GAP = 26`.  If the mixed-integer optimizer has computed a feasible solution and a bound, this contains the relative gap.
* `DINF_MIO_OPTIMIZER_TIME = 27`. Total time spent in the optimizer.
* `DINF_MIO_PROBING_TIME = 28`.  Total time for probing.
* `DINF_MIO_ROOT_CUTGEN_TIME = 29`.  Total time for cut generation.
* `DINF_MIO_ROOT_OPTIMIZER_TIME = 30`. Time spent in the optimizer while solving the root relaxation.
* `DINF_MIO_ROOT_PRESOLVE_TIME = 31`. Time spent in while presolving the root relaxation.
* `DINF_MIO_TIME = 32`. Time spent in the mixed-integer optimizer.
* `DINF_MIO_USER_OBJ_CUT = 33`.  If the objective cut is used, then this information item has the value of the cut.
* `DINF_OPTIMIZER_TIME = 34`.  Total time spent in the optimizer since it was invoked.
* `DINF_PRESOLVE_ELI_TIME = 35`.  Total time spent in the eliminator since the presolve was invoked.
* `DINF_PRESOLVE_LINDEP_TIME = 36`.  Total time spent  in the linear dependency checker since the presolve was invoked.
* `DINF_PRESOLVE_TIME = 37`.  Total time (in seconds) spent in the presolve since it was invoked.
* `DINF_PRIMAL_REPAIR_PENALTY_OBJ = 38`.  The optimal objective value of the penalty function.
* `DINF_QCQO_REFORMULATE_MAX_PERTURBATION = 39`.  Maximum absolute diagonal perturbation occuring during the QCQO reformulation.
* `DINF_QCQO_REFORMULATE_TIME = 40`.  Time spent with conic quadratic reformulation.
* `DINF_QCQO_REFORMULATE_WORST_CHOLESKY_COLUMN_SCALING = 41`.  Worst Cholesky column scaling.
* `DINF_QCQO_REFORMULATE_WORST_CHOLESKY_DIAG_SCALING = 42`.  Worst Cholesky diagonal scaling.
* `DINF_RD_TIME = 43`.  Time spent reading the data file.
* `DINF_SIM_DUAL_TIME = 44`.  Time spent in the dual simplex optimizer since invoking it.
* `DINF_SIM_FEAS = 45`.  Feasibility measure reported by the simplex optimizer.
* `DINF_SIM_OBJ = 46`.  Objective value reported by the simplex optimizer.
* `DINF_SIM_PRIMAL_DUAL_TIME = 47`.  Time spent in the primal-dual simplex optimizer since invoking it.
* `DINF_SIM_PRIMAL_TIME = 48`.  Time spent in the primal simplex optimizer since invoking it.
* `DINF_SIM_TIME = 49`.  Time spent in the simplex optimizer since invoking it.
* `DINF_SOL_BAS_DUAL_OBJ = 50`.  Dual objective value of the basic solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_BAS_DVIOLCON = 51`.  Maximal dual bound violation for xx in the basic solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_BAS_DVIOLVAR = 52`.  Maximal dual bound violation for xx in the basic solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_BAS_NRM_BARX = 53`.  Infinity norm of barx in the basic solution.
* `DINF_SOL_BAS_NRM_SLC = 54`.  Infinity norm of slc in the basic solution.
* `DINF_SOL_BAS_NRM_SLX = 55`.  Infinity norm of slx in the basic solution.
* `DINF_SOL_BAS_NRM_SUC = 56`.  Infinity norm of suc in the basic solution.
* `DINF_SOL_BAS_NRM_SUX = 57`.  Infinity norm of sux in the basic solution.
* `DINF_SOL_BAS_NRM_XC = 58`.  Infinity norm of xc in the basic solution.
* `DINF_SOL_BAS_NRM_XX = 59`.  Infinity norm of xx in the basic solution.
* `DINF_SOL_BAS_NRM_Y = 60`.  Infinity norm of Y in the basic solution.
* `DINF_SOL_BAS_PRIMAL_OBJ = 61`.  Primal objective value of the basic solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_BAS_PVIOLCON = 62`.  Maximal primal bound violation for xx in the basic solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_BAS_PVIOLVAR = 63`.  Maximal primal bound violation for xx in the basic solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_ITG_NRM_BARX = 64`.  Infinity norm of barx in the integer solution.
* `DINF_SOL_ITG_NRM_XC = 65`.  Infinity norm of xc in the integer solution.
* `DINF_SOL_ITG_NRM_XX = 66`.  Infinity norm of xx in the integer solution.
* `DINF_SOL_ITG_PRIMAL_OBJ = 67`.  Primal objective value of the integer solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_ITG_PVIOLBARVAR = 68`.  Maximal primal bound violation for barx in the integer solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_ITG_PVIOLCON = 69`.  Maximal primal bound violation for xx in the integer solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_ITG_PVIOLCONES = 70`.  Maximal primal violation for primal conic constraints in the integer solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_ITG_PVIOLITG = 71`.  Maximal violation for the integer constraints in the integer solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_ITG_PVIOLVAR = 72`.  Maximal primal bound violation for xx in the integer solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_ITR_DUAL_OBJ = 73`.  Dual objective value of the interior-point solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_ITR_DVIOLBARVAR = 74`.  Maximal dual bound violation for barx in the interior-point solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_ITR_DVIOLCON = 75`.  Maximal dual bound violation for xx in the interior-point solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_ITR_DVIOLCONES = 76`.  Maximal dual violation for dual conic constraints in the interior-point solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_ITR_DVIOLVAR = 77`.  Maximal dual bound violation for xx in the interior-point solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_ITR_NRM_BARS = 78`.  Infinity norm of bars in the interior-point solution.
* `DINF_SOL_ITR_NRM_BARX = 79`.  Infinity norm of barx in the interior-point solution.
* `DINF_SOL_ITR_NRM_SLC = 80`.  Infinity norm of slc in the interior-point solution.
* `DINF_SOL_ITR_NRM_SLX = 81`.  Infinity norm of slx in the interior-point solution.
* `DINF_SOL_ITR_NRM_SNX = 82`.  Infinity norm of snx in the interior-point solution.
* `DINF_SOL_ITR_NRM_SUC = 83`.  Infinity norm of suc in the interior-point solution.
* `DINF_SOL_ITR_NRM_SUX = 84`.  Infinity norm of sux in the interior-point solution.
* `DINF_SOL_ITR_NRM_XC = 85`.  Infinity norm of xc in the interior-point solution.
* `DINF_SOL_ITR_NRM_XX = 86`.  Infinity norm of xx in the interior-point solution.
* `DINF_SOL_ITR_NRM_Y = 87`.  Infinity norm of Y in the interior-point solution.
* `DINF_SOL_ITR_PRIMAL_OBJ = 88`.  Primal objective value of the interior-point solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_ITR_PVIOLBARVAR = 89`.  Maximal primal bound violation for barx in the interior-point solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_ITR_PVIOLCON = 90`.  Maximal primal bound violation for xx in the interior-point solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_ITR_PVIOLCONES = 91`.  Maximal primal violation for primal conic constraints in the interior-point solution. Updated by the function updatesolutioninfo.
* `DINF_SOL_ITR_PVIOLVAR = 92`.  Maximal primal bound violation for xx in the interior-point solution. Updated by the function updatesolutioninfo.
* `DINF_TO_CONIC_TIME = 93`.  Time spent in the last to conic reformulation.

parametertype (``int32``)
~~~~~~~~~~~~~~~~
* `PAR_DOU_TYPE = 1`. Is a double parameter.
* `PAR_INT_TYPE = 2`. Is an integer parameter.
* `PAR_INVALID_TYPE = 0`. Not a valid parameter.
* `PAR_STR_TYPE = 3`. Is a string parameter.

rescodetype (``int32``)
~~~~~~~~~~~~~~~~
* `RESPONSE_ERR = 3`. The response code is an error.
* `RESPONSE_OK = 0`. The response code is OK.
* `RESPONSE_TRM = 2`.  The response code is an optimizer termination status.
* `RESPONSE_UNK = 4`. The response code does not belong to any class.
* `RESPONSE_WRN = 1`. The response code is a warning.

prosta (``int32``)
~~~~~~~~~~~~~~~~
* `PRO_STA_DUAL_FEAS = 3`. The problem is dual feasible.
* `PRO_STA_DUAL_INFEAS = 5`. The problem is dual infeasible.
* `PRO_STA_ILL_POSED = 7`.  The problem is ill-posed. For example, it may be primal and dual feasible but have a positive duality gap.
* `PRO_STA_NEAR_DUAL_FEAS = 10`. The problem is at least nearly dual feasible.
* `PRO_STA_NEAR_PRIM_AND_DUAL_FEAS = 8`.  The problem is at least nearly primal and dual feasible.
* `PRO_STA_NEAR_PRIM_FEAS = 9`. The problem is at least nearly primal feasible.
* `PRO_STA_PRIM_AND_DUAL_FEAS = 1`. The problem is primal and dual feasible.
* `PRO_STA_PRIM_AND_DUAL_INFEAS = 6`. The problem is primal and dual infeasible.
* `PRO_STA_PRIM_FEAS = 2`. The problem is primal feasible.
* `PRO_STA_PRIM_INFEAS = 4`. The problem is primal infeasible.
* `PRO_STA_PRIM_INFEAS_OR_UNBOUNDED = 11`.  The problem is either primal infeasible or unbounded. This may occur for mixed-integer problems.
* `PRO_STA_UNKNOWN = 0`. Unknown problem status.

scalingtype (``int32``)
~~~~~~~~~~~~~~~~
* `SCALING_AGGRESSIVE = 3`. A very aggressive scaling is performed.
* `SCALING_FREE = 0`. The optimizer chooses the scaling heuristic.
* `SCALING_MODERATE = 2`. A conservative scaling is performed.
* `SCALING_NONE = 1`. No scaling is performed.

rescode (``int32``)
~~~~~~~~~~~~~~~~
* `RES_ERR_AD_INVALID_CODELIST = 3102`.  The code list data was invalid.
* `RES_ERR_API_ARRAY_TOO_SMALL = 3001`.  An input array was too short.
* `RES_ERR_API_CB_CONNECT = 3002`.  Failed to connect a callback object.
* `RES_ERR_API_FATAL_ERROR = 3005`.  An internal error occurred in the API. Please report this problem.
* `RES_ERR_API_INTERNAL = 3999`.  An internal fatal error occurred in an interface function.:w
* `RES_ERR_ARG_IS_TOO_LARGE = 1227`.  The value of a argument is too small.
* `RES_ERR_ARG_IS_TOO_SMALL = 1226`.  The value of a argument is too small.
* `RES_ERR_ARGUMENT_DIMENSION = 1201`.  A function argument is of incorrect dimension.
* `RES_ERR_ARGUMENT_IS_TOO_LARGE = 5005`.  The value of a function argument is too large.
* `RES_ERR_ARGUMENT_LENNEQ = 1197`.  Incorrect length of arguments.
* `RES_ERR_ARGUMENT_PERM_ARRAY = 1299`.  An invalid permutation array is specified.
* `RES_ERR_ARGUMENT_TYPE = 1198`.  Incorrect argument type.
* `RES_ERR_BAR_VAR_DIM = 3920`.  The dimension of a symmetric matrix variable has to greater than 0.
* `RES_ERR_BASIS = 1266`.  Invalid basis is specified.
* `RES_ERR_BASIS_FACTOR = 1610`.  The factorization of the basis is invalid.
* `RES_ERR_BASIS_SINGULAR = 1615`.  The basis is singular.
* `RES_ERR_BLANK_NAME = 1070`.  An all blank name has been specified.
* `RES_ERR_CANNOT_CLONE_NL = 2505`.  A task with a nonlinear function call-back cannot be cloned.
* `RES_ERR_CANNOT_HANDLE_NL = 2506`.  A function cannot handle a task with nonlinear function call-backs.
* `RES_ERR_CBF_DUPLICATE_ACOORD = 7116`.  Duplicate index in ACOORD.
* `RES_ERR_CBF_DUPLICATE_BCOORD = 7115`.  Duplicate index in BCOORD.
* `RES_ERR_CBF_DUPLICATE_CON = 7108`.  Duplicate CON keyword.
* `RES_ERR_CBF_DUPLICATE_INT = 7110`.  Duplicate INT keyword.
* `RES_ERR_CBF_DUPLICATE_OBJ = 7107`.  Duplicate OBJ keyword.
* `RES_ERR_CBF_DUPLICATE_OBJACOORD = 7114`.  Duplicate index in OBJCOORD.
* `RES_ERR_CBF_DUPLICATE_VAR = 7109`.  Duplicate VAR keyword.
* `RES_ERR_CBF_INVALID_CON_TYPE = 7112`.  Invalid constraint type.
* `RES_ERR_CBF_INVALID_DOMAIN_DIMENSION = 7113`.  Invalid domain dimension.
* `RES_ERR_CBF_INVALID_INT_INDEX = 7121`.  Invalid INT index.
* `RES_ERR_CBF_INVALID_VAR_TYPE = 7111`.  Invalid variable type.
* `RES_ERR_CBF_NO_VARIABLES = 7102`.  An invalid objective sense is specified.
* `RES_ERR_CBF_NO_VERSION_SPECIFIED = 7105`.  No version specified.
* `RES_ERR_CBF_OBJ_SENSE = 7101`.  An invalid objective sense is specified.
* `RES_ERR_CBF_PARSE = 7100`.  An error occurred while parsing an CBF file.
* `RES_ERR_CBF_SYNTAX = 7106`.  Invalid syntax.
* `RES_ERR_CBF_TOO_FEW_CONSTRAINTS = 7118`.  Too few constraints defined.
* `RES_ERR_CBF_TOO_FEW_INTS = 7119`.  Too ints specified.
* `RES_ERR_CBF_TOO_FEW_VARIABLES = 7117`.  Too few variables defined.
* `RES_ERR_CBF_TOO_MANY_CONSTRAINTS = 7103`.  Too many constraints specified.
* `RES_ERR_CBF_TOO_MANY_INTS = 7120`.  Too ints specified.
* `RES_ERR_CBF_TOO_MANY_VARIABLES = 7104`.  Too many variables specified.
* `RES_ERR_CBF_UNSUPPORTED = 7122`.  Unsupported feature is present.
* `RES_ERR_CON_Q_NOT_NSD = 1294`.  The quadratic constraint matrix is not NSD.
* `RES_ERR_CON_Q_NOT_PSD = 1293`.  The quadratic constraint matrix is not PSD.
* `RES_ERR_CONE_INDEX = 1300`.  An index of a non-existing cone has been specified.
* `RES_ERR_CONE_OVERLAP = 1302`.  One or more of variables in the cone to be added is already member of another cone.
* `RES_ERR_CONE_OVERLAP_APPEND = 1307`.  The cone to be appended has one variable which is already member of another cone.
* `RES_ERR_CONE_REP_VAR = 1303`.  A variable is included multiple times in the cone.
* `RES_ERR_CONE_SIZE = 1301`.  A cone with too few members is specified.
* `RES_ERR_CONE_TYPE = 1305`.  Invalid cone type specified.
* `RES_ERR_CONE_TYPE_STR = 1306`.  Invalid cone type specified.
* `RES_ERR_DATA_FILE_EXT = 1055`.  The data file format cannot be determined from the file name.
* `RES_ERR_DUP_NAME = 1071`.  Duplicate names specified.
* `RES_ERR_DUPLICATE_AIJ = 1385`.  An element in the A matrix is specified twice.
* `RES_ERR_DUPLICATE_BARVARIABLE_NAMES = 4502`.  Two barvariable names are identical.
* `RES_ERR_DUPLICATE_CONE_NAMES = 4503`.  Two cone names are identical.
* `RES_ERR_DUPLICATE_CONSTRAINT_NAMES = 4500`.  Two constraint names are identical.
* `RES_ERR_DUPLICATE_VARIABLE_NAMES = 4501`.  Two variable names are identical.
* `RES_ERR_END_OF_FILE = 1059`.  End of file reached.
* `RES_ERR_FACTOR = 1650`.  An error occurred while factorizing a matrix.
* `RES_ERR_FEASREPAIR_CANNOT_RELAX = 1700`.  An optimization problem cannot be relaxed.
* `RES_ERR_FEASREPAIR_INCONSISTENT_BOUND = 1702`.  The upper bound is less than the lower bound for a variable or a constraint.
* `RES_ERR_FEASREPAIR_SOLVING_RELAXED = 1701`.  The relaxed problem could not be solved to optimality.
* `RES_ERR_FILE_LICENSE = 1007`.  Invalid license file.
* `RES_ERR_FILE_OPEN = 1052`.  An error occurred while opening a file.
* `RES_ERR_FILE_READ = 1053`.  An error occurred while reading file.
* `RES_ERR_FILE_WRITE = 1054`.  An error occurred while writing to a file.
* `RES_ERR_FIRST = 1261`.  Invalid first.
* `RES_ERR_FIRSTI = 1285`.  Invalid firsti.
* `RES_ERR_FIRSTJ = 1287`.  Invalid firstj.
* `RES_ERR_FIXED_BOUND_VALUES = 1425`.  A fixed constraint/variable has been specified using the bound keys but the numerical bounds are different.
* `RES_ERR_FLEXLM = 1014`.  The |flexlm| license manager reported an error.
* `RES_ERR_GLOBAL_INV_CONIC_PROBLEM = 1503`.  The global optimizer can only be applied to problems without semidefinite variables.
* `RES_ERR_HUGE_AIJ = 1380`.  A numerically huge value is specified for an element in A.
* `RES_ERR_HUGE_C = 1375`.  A huge value in absolute size is specified for one an objective coefficient.
* `RES_ERR_IDENTICAL_TASKS = 3101`.  Some tasks related to this function call were identical. Unique tasks were expected.
* `RES_ERR_IN_ARGUMENT = 1200`.  A function argument is incorrect.
* `RES_ERR_INDEX = 1235`.  An index is out of range.
* `RES_ERR_INDEX_ARR_IS_TOO_LARGE = 1222`.  An index in an array argument is too large.
* `RES_ERR_INDEX_ARR_IS_TOO_SMALL = 1221`.  An index in an array argument is too small.
* `RES_ERR_INDEX_IS_TOO_LARGE = 1204`.  An index in an argument is too large.
* `RES_ERR_INDEX_IS_TOO_SMALL = 1203`.  An index in an argument is too small.
* `RES_ERR_INF_DOU_INDEX = 1219`.  A double information index is out of range for the specified type.
* `RES_ERR_INF_DOU_NAME = 1230`.  A double information name is invalid.
* `RES_ERR_INF_INT_INDEX = 1220`.  An integer information index is out of range for the specified type.
* `RES_ERR_INF_INT_NAME = 1231`.  An integer information name is invalid.
* `RES_ERR_INF_LINT_INDEX = 1225`.  A long integer information index is out of range for the specified type.
* `RES_ERR_INF_LINT_NAME = 1234`.  A long integer information name is invalid.
* `RES_ERR_INF_TYPE = 1232`.  The information type is invalid.
* `RES_ERR_INFEAS_UNDEFINED = 3910`.  The requested value is not defined for this solution type.
* `RES_ERR_INFINITE_BOUND = 1400`.  A numerically huge bound value is specified.
* `RES_ERR_INT64_TO_INT32_CAST = 3800`.  An 32 bit integer could not cast to a 64 bit integer.
* `RES_ERR_INTERNAL = 3000`.  An internal error occurred.
* `RES_ERR_INTERNAL_TEST_FAILED = 3500`.  An internal unit test function failed.
* `RES_ERR_INV_APTRE = 1253`.  aptre[j] is strictly smaller than aptrb[j] for some j.
* `RES_ERR_INV_BK = 1255`.  Invalid bound key.
* `RES_ERR_INV_BKC = 1256`.  Invalid bound key is specified for a constraint.
* `RES_ERR_INV_BKX = 1257`.  An invalid bound key is specified for a variable.
* `RES_ERR_INV_CONE_TYPE = 1272`.  Invalid cone type code encountered.
* `RES_ERR_INV_CONE_TYPE_STR = 1271`.  Invalid cone type string encountered.
* `RES_ERR_INV_MARKI = 2501`.  Invalid value in marki.
* `RES_ERR_INV_MARKJ = 2502`.  Invalid value in markj.
* `RES_ERR_INV_NAME_ITEM = 1280`.  An invalid name item code is used.
* `RES_ERR_INV_NUMI = 2503`.  Invalid numi.
* `RES_ERR_INV_NUMJ = 2504`.  Invalid numj.
* `RES_ERR_INV_OPTIMIZER = 1550`.  An invalid optimizer has been chosen for the problem.
* `RES_ERR_INV_PROBLEM = 1500`.  Invalid problem type.
* `RES_ERR_INV_QCON_SUBI = 1405`.  Invalid value in qcsubi.
* `RES_ERR_INV_QCON_SUBJ = 1406`.  Invalid value in qcsubj.
* `RES_ERR_INV_QCON_SUBK = 1404`.  Invalid value in qcsubk.
* `RES_ERR_INV_QCON_VAL = 1407`.  Invalid value in qcval.
* `RES_ERR_INV_QOBJ_SUBI = 1401`.  Invalid value %d at qosubi.
* `RES_ERR_INV_QOBJ_SUBJ = 1402`.  Invalid value in qosubj.
* `RES_ERR_INV_QOBJ_VAL = 1403`.  Invalid value in qoval.
* `RES_ERR_INV_SK = 1270`.  Invalid status key code encountered.
* `RES_ERR_INV_SK_STR = 1269`.  Invalid status key string encountered.
* `RES_ERR_INV_SKC = 1267`.  Invalid value in skc encountered.
* `RES_ERR_INV_SKN = 1274`.  Invalid value in skn encountered.
* `RES_ERR_INV_SKX = 1268`.  Invalid value in skx encountered.
* `RES_ERR_INV_VAR_TYPE = 1258`.  An invalid variable type is specified for a variable.
* `RES_ERR_INVALID_ACCMODE = 2520`.  An invalid access mode is specified.
* `RES_ERR_INVALID_AIJ = 1473`.  a[i,j] contains an invalid floating point value, i.e. a NaN or an infinite value.
* `RES_ERR_INVALID_AMPL_STUB = 3700`.  Invalid AMPL stub.
* `RES_ERR_INVALID_BARVAR_NAME = 1079`.  An invalid symmetric matrix variable name is used.
* `RES_ERR_INVALID_COMPRESSION = 1800`.  Invalid compression type.
* `RES_ERR_INVALID_CON_NAME = 1076`.  An invalid constraint name is used.
* `RES_ERR_INVALID_CONE_NAME = 1078`.  An invalid cone name is used.
* `RES_ERR_INVALID_FILE_FORMAT_FOR_CONES = 4005`.  The file format does not support a problem with conic constraints.
* `RES_ERR_INVALID_FILE_FORMAT_FOR_GENERAL_NL = 4010`.  The file format does not support a problem with general nonlinear terms.
* `RES_ERR_INVALID_FILE_FORMAT_FOR_SYM_MAT = 4000`.  The file format does not support a problem with symmetric matrix variables.
* `RES_ERR_INVALID_FILE_NAME = 1056`.  An invalid file name has been specified.
* `RES_ERR_INVALID_FORMAT_TYPE = 1283`.  Invalid format type.
* `RES_ERR_INVALID_IDX = 1246`.  A specified index is invalid.
* `RES_ERR_INVALID_IOMODE = 1801`.  Invalid io mode.
* `RES_ERR_INVALID_MAX_NUM = 1247`.  A specified index is invalid.
* `RES_ERR_INVALID_NAME_IN_SOL_FILE = 1170`.  An invalid name occurred in a solution file.
* `RES_ERR_INVALID_OBJ_NAME = 1075`.  An invalid objective name is specified.
* `RES_ERR_INVALID_OBJECTIVE_SENSE = 1445`.  An invalid objective sense is specified.
* `RES_ERR_INVALID_PROBLEM_TYPE = 6000`.  An invalid problem type.
* `RES_ERR_INVALID_SOL_FILE_NAME = 1057`.  An invalid file name has been specified.
* `RES_ERR_INVALID_STREAM = 1062`.  An invalid stream is referenced.
* `RES_ERR_INVALID_SURPLUS = 1275`.  Invalid surplus.
* `RES_ERR_INVALID_SYM_MAT_DIM = 3950`.  A sparse symmetric matrix of invalid dimension is specified.
* `RES_ERR_INVALID_TASK = 1064`.  The task is invalid.
* `RES_ERR_INVALID_UTF8 = 2900`.  An invalid UTF8 string is encountered.
* `RES_ERR_INVALID_VAR_NAME = 1077`.  An invalid variable name is used.
* `RES_ERR_INVALID_WCHAR = 2901`.  An invalid wchar string is encountered.
* `RES_ERR_INVALID_WHICHSOL = 1228`.  whichsol is invalid.
* `RES_ERR_JSON_DATA = 1179`.  Inconsistent data in JSON Task file
* `RES_ERR_JSON_FORMAT = 1178`.  Error in an JSON Task file
* `RES_ERR_JSON_MISSING_DATA = 1180`.  Missing data section in JSON task file.
* `RES_ERR_JSON_NUMBER_OVERFLOW = 1177`.  Invalid number entry - wrong type or value overflow.
* `RES_ERR_JSON_STRING = 1176`.  Error in JSON string.
* `RES_ERR_JSON_SYNTAX = 1175`.  Syntax error in an JSON data
* `RES_ERR_LAST = 1262`.  Invalid last.
* `RES_ERR_LASTI = 1286`.  Invalid lasti.
* `RES_ERR_LASTJ = 1288`.  Invalid lastj.
* `RES_ERR_LAU_ARG_K = 7012`.  Invalid argument k.
* `RES_ERR_LAU_ARG_M = 7010`.  Invalid argument m.
* `RES_ERR_LAU_ARG_N = 7011`.  Invalid argument n.
* `RES_ERR_LAU_ARG_TRANS = 7018`.  Invalid argument trans.
* `RES_ERR_LAU_ARG_TRANSA = 7015`.  Invalid argument transa.
* `RES_ERR_LAU_ARG_TRANSB = 7016`.  Invalid argument transb.
* `RES_ERR_LAU_ARG_UPLO = 7017`.  Invalid argument uplo.
* `RES_ERR_LAU_INVALID_LOWER_TRIANGULAR_MATRIX = 7002`.  An invalid lower triangular matrix.
* `RES_ERR_LAU_INVALID_SPARSE_SYMMETRIC_MATRIX = 7019`.  An invalid sparse symmetric matrix is specfified.
* `RES_ERR_LAU_NOT_POSITIVE_DEFINITE = 7001`.  A matrix is not positive definite.
* `RES_ERR_LAU_SINGULAR_MATRIX = 7000`.  A matrix is singular.
* `RES_ERR_LAU_UNKNOWN = 7005`.  An unknown error.
* `RES_ERR_LICENSE = 1000`.  Invalid license.
* `RES_ERR_LICENSE_CANNOT_ALLOCATE = 1020`.  The license system cannot allocate the memory required.
* `RES_ERR_LICENSE_CANNOT_CONNECT = 1021`.  MOSEK cannot connect to the license server.
* `RES_ERR_LICENSE_EXPIRED = 1001`.  The license has expired.
* `RES_ERR_LICENSE_FEATURE = 1018`.  A requested feature is not available in the license file(s).
* `RES_ERR_LICENSE_INVALID_HOSTID = 1025`.  The host ID specified in the license file does not match the host ID of the computer.
* `RES_ERR_LICENSE_MAX = 1016`.  Maximum number of licenses is reached.
* `RES_ERR_LICENSE_MOSEKLM_DAEMON = 1017`.  The MOSEKLM license manager daemon is not up and running.
* `RES_ERR_LICENSE_NO_SERVER_LINE = 1028`.  No SERVER lines in license file.
* `RES_ERR_LICENSE_NO_SERVER_SUPPORT = 1027`.  The license server does not support the requested feature.
* `RES_ERR_LICENSE_SERVER = 1015`.  The license server is not responding.
* `RES_ERR_LICENSE_SERVER_VERSION = 1026`.  The version specified in the checkout request is greater than the highest version number the daemon supports.
* `RES_ERR_LICENSE_VERSION = 1002`.  Invalid license version.
* `RES_ERR_LINK_FILE_DLL = 1040`.  A file cannot be linked to a stream in the DLL version.
* `RES_ERR_LIVING_TASKS = 1066`.  Not all tasks associated with the environment have been deleted.
* `RES_ERR_LOWER_BOUND_IS_A_NAN = 1390`.  The lower bound specified is not a number (nan).
* `RES_ERR_LP_DUP_SLACK_NAME = 1152`.  The name of the slack variable added to a ranged constraint already exists.
* `RES_ERR_LP_EMPTY = 1151`.  The problem cannot be written to an LP formatted file.
* `RES_ERR_LP_FILE_FORMAT = 1157`.  Syntax error in an LP file.
* `RES_ERR_LP_FORMAT = 1160`.  Syntax error in an LP file.
* `RES_ERR_LP_FREE_CONSTRAINT = 1155`.  Free constraints cannot be written in LP file format.
* `RES_ERR_LP_INCOMPATIBLE = 1150`.  The problem cannot be written to an LP formatted file.
* `RES_ERR_LP_INVALID_CON_NAME = 1171`.  A constraint name is invalid when used in an LP formatted file.
* `RES_ERR_LP_INVALID_VAR_NAME = 1154`.  A variable name is invalid when used in an LP formatted file.
* `RES_ERR_LP_WRITE_CONIC_PROBLEM = 1163`.  The problem contains cones that cannot be written to an LP formatted file.
* `RES_ERR_LP_WRITE_GECO_PROBLEM = 1164`.  The problem contains general convex terms that cannot be written to an LP formatted file.
* `RES_ERR_LU_MAX_NUM_TRIES = 2800`.  Could not compute the LU factors of the matrix within the maximum number of allowed tries.
* `RES_ERR_MAX_LEN_IS_TOO_SMALL = 1289`.  An maximum length that is too small has been specified.
* `RES_ERR_MAXNUMBARVAR = 1242`.  The maximum number of semidefinite variables limit is too small.
* `RES_ERR_MAXNUMCON = 1240`.  Invalid maximum number of constraints specified.
* `RES_ERR_MAXNUMCONE = 1304`.  The value specified for maxnumcone is too small.
* `RES_ERR_MAXNUMQNZ = 1243`.  Too small maximum number of non-zeros for the Q matrices is specified.
* `RES_ERR_MAXNUMVAR = 1241`.  The maximum number of variables limit is too small.
* `RES_ERR_MIO_INTERNAL = 5010`.  A fatal error occurred in the mixed integer optimizer.  Please contact MOSEK support.
* `RES_ERR_MIO_INVALID_NODE_OPTIMIZER = 7131`.  An invalid node optimizer was selected for the problem type.
* `RES_ERR_MIO_INVALID_ROOT_OPTIMIZER = 7130`.  An invalid root optimizer was selected for the problem type.
* `RES_ERR_MIO_NO_OPTIMIZER = 1551`.  No optimizer is available for the current class of integer optimization problems.
* `RES_ERR_MIO_NOT_LOADED = 1553`.  The mixed-integer optimizer is not loaded.
* `RES_ERR_MISSING_LICENSE_FILE = 1008`.  A license cannot be located.
* `RES_ERR_MIXED_CONIC_AND_NL = 1501`.  The problem contains both conic and nonlinear constraints.
* `RES_ERR_MPS_CONE_OVERLAP = 1118`.  A variable is specified to be a member of several cones.
* `RES_ERR_MPS_CONE_REPEAT = 1119`.  A variable is repeated within the CSECTION.
* `RES_ERR_MPS_CONE_TYPE = 1117`.  Invalid cone type specified in a  CSECTION.
* `RES_ERR_MPS_DUPLICATE_Q_ELEMENT = 1121`.  Duplicate elements is specified in a Q matrix.
* `RES_ERR_MPS_FILE = 1100`.  An error occurred while reading an MPS file.
* `RES_ERR_MPS_INV_BOUND_KEY = 1108`.  An invalid bound key occurred in an MPS file.
* `RES_ERR_MPS_INV_CON_KEY = 1107`.  An invalid constraint key occurred in an MPS file.
* `RES_ERR_MPS_INV_FIELD = 1101`.  Invalid field occurred while reading an MPS file.
* `RES_ERR_MPS_INV_MARKER = 1102`.  An invalid marker has been specified in the MPS file.
* `RES_ERR_MPS_INV_SEC_NAME = 1109`.  An invalid section name occurred in an MPS file.
* `RES_ERR_MPS_INV_SEC_ORDER = 1115`.  The sections in an MPS file is not in the correct order.
* `RES_ERR_MPS_INVALID_OBJ_NAME = 1128`.  An invalid objective name is specified.
* `RES_ERR_MPS_INVALID_OBJSENSE = 1122`.  An invalid objective sense is specified.
* `RES_ERR_MPS_MUL_CON_NAME = 1112`.  A constraint name is specified multiple times in the ROWS section in an MPS file.
* `RES_ERR_MPS_MUL_CSEC = 1116`.  Multiple CSECTIONs are given the same name.
* `RES_ERR_MPS_MUL_QOBJ = 1114`.  The Q term in the objective is specified multiple times.
* `RES_ERR_MPS_MUL_QSEC = 1113`.  Multiple QSECTIONs are specified for a constraint.
* `RES_ERR_MPS_NO_OBJECTIVE = 1110`.  No objective is defined in an MPS file.
* `RES_ERR_MPS_NON_SYMMETRIC_Q = 1120`.  A non symmetric matrice has been speciefied.
* `RES_ERR_MPS_NULL_CON_NAME = 1103`.  An empty constraint name is used in an MPS file.
* `RES_ERR_MPS_NULL_VAR_NAME = 1104`.  An empty variable name is used in an MPS file.
* `RES_ERR_MPS_SPLITTED_VAR = 1111`.  The non-zero elements in A corresponding to a variable in an MPS file must be specified consecutively.
* `RES_ERR_MPS_TAB_IN_FIELD2 = 1125`.  A tab char occurred in field 2.
* `RES_ERR_MPS_TAB_IN_FIELD3 = 1126`.  A tab char occurred in field 3.
* `RES_ERR_MPS_TAB_IN_FIELD5 = 1127`.  A tab char occurred in field 5.
* `RES_ERR_MPS_UNDEF_CON_NAME = 1105`.  An undefined constraint name occurred in an MPS file.
* `RES_ERR_MPS_UNDEF_VAR_NAME = 1106`.  An undefined variable name occurred in an MPS file.
* `RES_ERR_MUL_A_ELEMENT = 1254`.  An element in A is defined multiple times.
* `RES_ERR_NAME_IS_NULL = 1760`.  The name buffer is a |null| pointer.
* `RES_ERR_NAME_MAX_LEN = 1750`.  A name is longer than the buffer that is supposed to hold it.
* `RES_ERR_NAN_IN_BLC = 1461`.  blc contains an invalid floating point value, i.e. a NaN.
* `RES_ERR_NAN_IN_BLX = 1471`.  blx contains an invalid floating point value, i.e. a NaN.
* `RES_ERR_NAN_IN_BUC = 1462`.  buc contains an invalid floating point value, i.e. a NaN.
* `RES_ERR_NAN_IN_BUX = 1472`.  bux contains an invalid floating point value, i.e. a NaN.
* `RES_ERR_NAN_IN_C = 1470`.  c contains an invalid floating point value, i.e. a NaN.
* `RES_ERR_NAN_IN_DOUBLE_DATA = 1450`.  An invalid floating value was used in some double data.
* `RES_ERR_NEGATIVE_APPEND = 1264`.  Cannot append a negative number.
* `RES_ERR_NEGATIVE_SURPLUS = 1263`.  Negative surplus.
* `RES_ERR_NEWER_DLL = 1036`.  The dynamic link library is newer than the specified version.
* `RES_ERR_NO_BARS_FOR_SOLUTION = 3916`.  There is no bars available for the solution specified.
* `RES_ERR_NO_BARX_FOR_SOLUTION = 3915`.  There is no barx available for the solution specified.
* `RES_ERR_NO_BASIS_SOL = 1600`.  No basic solution is defined.
* `RES_ERR_NO_DUAL_FOR_ITG_SOL = 2950`.  No dual information is available for the integer solution.
* `RES_ERR_NO_DUAL_INFEAS_CER = 2001`.  A certificate of dual infeasibility is not available.
* `RES_ERR_NO_INIT_ENV = 1063`.  Environment is not initialized.
* `RES_ERR_NO_OPTIMIZER_VAR_TYPE = 1552`.  No optimizer is available for this class of optimization problems.
* `RES_ERR_NO_PRIMAL_INFEAS_CER = 2000`.  A certificate of primal infeasibility is not available.
* `RES_ERR_NO_SNX_FOR_BAS_SOL = 2953`.  snx is not available for the basis solution.
* `RES_ERR_NO_SOLUTION_IN_CALLBACK = 2500`.  The required solution is not available.
* `RES_ERR_NON_UNIQUE_ARRAY = 5000`.  An array does not contain unique elements.
* `RES_ERR_NONCONVEX = 1291`.  The optimization problem is nonconvex.
* `RES_ERR_NONLINEAR_EQUALITY = 1290`.  The model contains a nonlinear equality.
* `RES_ERR_NONLINEAR_FUNCTIONS_NOT_ALLOWED = 1428`.  An operation that is invalid for problems with nonlinear functions defined has been attempted.
* `RES_ERR_NONLINEAR_RANGED = 1292`.  The problem contains a nonlinear constraint with inite lower and upper bound.
* `RES_ERR_NR_ARGUMENTS = 1199`.  Incorrect number of function arguments.
* `RES_ERR_NULL_ENV = 1060`.  env is a |null| pointer.
* `RES_ERR_NULL_POINTER = 1065`.  An argument to a function is unexpectedly a |null| pointer.
* `RES_ERR_NULL_TASK = 1061`.  task is a |null| pointer.
* `RES_ERR_NUMCONLIM = 1250`.  Maximum number of constraints limit is exceeded.
* `RES_ERR_NUMVARLIM = 1251`.  Maximum number of variables limit is exceeded.
* `RES_ERR_OBJ_Q_NOT_NSD = 1296`.  The quadratic coefficient matrix in the objective is not NSD.
* `RES_ERR_OBJ_Q_NOT_PSD = 1295`.  The quadratic coefficient matrix in the objective is not PSD.
* `RES_ERR_OBJECTIVE_RANGE = 1260`.  Empty objective range.
* `RES_ERR_OLDER_DLL = 1035`.  The dynamic link library is older than the specified version.
* `RES_ERR_OPEN_DL = 1030`.  A dynamic link library could not be opened.
* `RES_ERR_OPF_FORMAT = 1168`.  Syntax error in an OPF file
* `RES_ERR_OPF_NEW_VARIABLE = 1169`.  Variable not previously defined.
* `RES_ERR_OPF_PREMATURE_EOF = 1172`.  Premature end of file in an OPF file.
* `RES_ERR_OPTIMIZER_LICENSE = 1013`.  The optimizer required is not licensed.
* `RES_ERR_OVERFLOW = 1590`.  A computation produced an overflow.
* `RES_ERR_PARAM_INDEX = 1210`.  Parameter index is out of range.
* `RES_ERR_PARAM_IS_TOO_LARGE = 1215`.  A parameter value is too large.
* `RES_ERR_PARAM_IS_TOO_SMALL = 1216`.  A parameter value is too small.
* `RES_ERR_PARAM_NAME = 1205`.  A parameter name is not correct.
* `RES_ERR_PARAM_NAME_DOU = 1206`.  A parameter name is not correct.
* `RES_ERR_PARAM_NAME_INT = 1207`.  A parameter name is not correct.
* `RES_ERR_PARAM_NAME_STR = 1208`.  A parameter name is not correct.
* `RES_ERR_PARAM_TYPE = 1218`.  A parameter type is invalid.
* `RES_ERR_PARAM_VALUE_STR = 1217`.  A parameter value string is incorrect.
* `RES_ERR_PLATFORM_NOT_LICENSED = 1019`.  A requested license feature is not available for the required platform.
* `RES_ERR_POSTSOLVE = 1580`.  An error occurred during the postsolve.
* `RES_ERR_PRO_ITEM = 1281`.  An invalid problem item is used.
* `RES_ERR_PROB_LICENSE = 1006`.  The software is not licensed to solve the problem.
* `RES_ERR_QCON_SUBI_TOO_LARGE = 1409`.  Invalid value in qcsubi.
* `RES_ERR_QCON_SUBI_TOO_SMALL = 1408`.  Invalid value in qcsubi.
* `RES_ERR_QCON_UPPER_TRIANGLE = 1417`.  An element in the upper triangle of the quadratic term in a constraint.
* `RES_ERR_QOBJ_UPPER_TRIANGLE = 1415`.  An element in the upper triangle of the quadratic term in the objective is specified.
* `RES_ERR_READ_FORMAT = 1090`.  The specified format cannot be read.
* `RES_ERR_READ_LP_MISSING_END_TAG = 1159`.  Syntax error in LP fil. Possibly missing End tag.
* `RES_ERR_READ_LP_NONEXISTING_NAME = 1162`.  A variable never occurred in objective or constraints.
* `RES_ERR_REMOVE_CONE_VARIABLE = 1310`.  A variable cannot be removed because it will make a cone invalid.
* `RES_ERR_REPAIR_INVALID_PROBLEM = 1710`.  The feasibility repair does not support the specified problem type.
* `RES_ERR_REPAIR_OPTIMIZATION_FAILED = 1711`.  Computation the optimal relaxation failed.
* `RES_ERR_SEN_BOUND_INVALID_LO = 3054`.  Analysis of lower bound requested for an index, where no lower bound exists.
* `RES_ERR_SEN_BOUND_INVALID_UP = 3053`.  Analysis of upper bound requested for an index, where no upper bound exists.
* `RES_ERR_SEN_FORMAT = 3050`.  Syntax error in sensitivity analysis file.
* `RES_ERR_SEN_INDEX_INVALID = 3055`.  Invalid range given in the sensitivity file.
* `RES_ERR_SEN_INDEX_RANGE = 3052`.  Index out of range in the sensitivity analysis file.
* `RES_ERR_SEN_INVALID_REGEXP = 3056`.  Syntax error in regexp or regexp longer than 1024.
* `RES_ERR_SEN_NUMERICAL = 3058`.  Numerical difficulties encountered performing the sensitivity analysis.
* `RES_ERR_SEN_SOLUTION_STATUS = 3057`.  No optimal solution found to the original problem given for sensitivity analysis.
* `RES_ERR_SEN_UNDEF_NAME = 3051`.  An undefined name was encountered in the sensitivity analysis file.
* `RES_ERR_SEN_UNHANDLED_PROBLEM_TYPE = 3080`.  Sensitivity analysis cannot be performed for the specified problem.
* `RES_ERR_SERVER_CONNECT = 8000`.  Failed to connect to remote solver server.
* `RES_ERR_SERVER_PROTOCOL = 8001`.  Unexpected message or data from solver server.
* `RES_ERR_SERVER_STATUS = 8002`.  Server returned non-ok status code
* `RES_ERR_SERVER_TOKEN = 8003`.  Invalid job ID
* `RES_ERR_SIZE_LICENSE = 1005`.  The problem is bigger than the license.
* `RES_ERR_SIZE_LICENSE_CON = 1010`.  The problem has too many constraints.
* `RES_ERR_SIZE_LICENSE_INTVAR = 1012`.  The problem contains too many integer variables.
* `RES_ERR_SIZE_LICENSE_NUMCORES = 3900`.  The computer contains more cpu cores than the license allows for.
* `RES_ERR_SIZE_LICENSE_VAR = 1011`.  The problem has too many variables.
* `RES_ERR_SOL_FILE_INVALID_NUMBER = 1350`.  An invalid number is specified in a solution file.
* `RES_ERR_SOLITEM = 1237`.  The solution number  solemn does not exists.
* `RES_ERR_SOLVER_PROBTYPE = 1259`.  Problem type does not match the chosen optimizer.
* `RES_ERR_SPACE = 1051`.  Out of space.
* `RES_ERR_SPACE_LEAKING = 1080`.  MOSEK is leaking memory.
* `RES_ERR_SPACE_NO_INFO = 1081`.  No available information about the space usage.
* `RES_ERR_SYM_MAT_DUPLICATE = 3944`.  A value in a symmetric matric as been specified more than once.
* `RES_ERR_SYM_MAT_HUGE = 1482`.  A numerically huge value is specified for an element in A.
* `RES_ERR_SYM_MAT_INVALID = 1480`. A symmetric matrix contains an invalid floating point value, i.e. a NaN or an infinite value.
* `RES_ERR_SYM_MAT_INVALID_COL_INDEX = 3941`.  A column index specified for sparse symmetric matrix is invalid.
* `RES_ERR_SYM_MAT_INVALID_ROW_INDEX = 3940`.  A row index specified for sparse symmetric matrix is invalid.
* `RES_ERR_SYM_MAT_INVALID_VALUE = 3943`.  The numerical value specified in a sparse symmetric matrix is not a value floating value.
* `RES_ERR_SYM_MAT_NOT_LOWER_TRINGULAR = 3942`.  Only the lower triangular part of sparse symmetric matrix should be specified.
* `RES_ERR_TASK_INCOMPATIBLE = 2560`.  The Task file is incompatible with this platform.
* `RES_ERR_TASK_INVALID = 2561`.  The Task file is invalid.
* `RES_ERR_TASK_WRITE = 2562`.  Failed to write the task file.
* `RES_ERR_THREAD_COND_INIT = 1049`.  Could not initialize a condition.
* `RES_ERR_THREAD_CREATE = 1048`.  Could not create a thread.
* `RES_ERR_THREAD_MUTEX_INIT = 1045`.  Could not initialize a mutex.
* `RES_ERR_THREAD_MUTEX_LOCK = 1046`.  Could not lock a mutex.
* `RES_ERR_THREAD_MUTEX_UNLOCK = 1047`.  Could not unlock a mutex.
* `RES_ERR_TOCONIC_CONSTR_NOT_CONIC = 7153`.  The constraint is not conic representable.
* `RES_ERR_TOCONIC_CONSTR_Q_NOT_PSD = 7150`.  The matrix defining the quadratric part of constraint is not positive semidefinite.
* `RES_ERR_TOCONIC_CONSTRAINT_FX = 7151`.  The quadratic constraint is an equality, thus not convex.
* `RES_ERR_TOCONIC_CONSTRAINT_RA = 7152`.  The quadratic constraint has finite lower and upper bound, and therefore it is not convex.
* `RES_ERR_TOCONIC_OBJECTIVE_NOT_PSD = 7155`.  The matrix defining the quadratric part of the objective function is not positive semidefinite.
* `RES_ERR_TOO_SMALL_MAX_NUM_NZ = 1245`.  The maximum number of non-zeros specified is too small.
* `RES_ERR_TOO_SMALL_MAXNUMANZ = 1252`.  Too small maximum number of non-zeros in A specified.
* `RES_ERR_UNB_STEP_SIZE = 3100`.  A step-size in an optimizer was unexpectedly unbounded.
* `RES_ERR_UNDEF_SOLUTION = 1265`.  The required solution is not defined.
* `RES_ERR_UNDEFINED_OBJECTIVE_SENSE = 1446`.  The objective sense has not been specified before the optimization.
* `RES_ERR_UNHANDLED_SOLUTION_STATUS = 6010`.  Unhandled solution status.
* `RES_ERR_UNKNOWN = 1050`.  Unknown error.
* `RES_ERR_UPPER_BOUND_IS_A_NAN = 1391`.  The upper bound specified is not a number (nan).
* `RES_ERR_UPPER_TRIANGLE = 6020`.  An element in the upper triangle of a lower triangular matrix is specified.
* `RES_ERR_USER_FUNC_RET = 1430`.  An user function reported an error.
* `RES_ERR_USER_FUNC_RET_DATA = 1431`.  An user function returned invalid data.
* `RES_ERR_USER_NLO_EVAL = 1433`.  The user-defined nonlinear function reported an error.
* `RES_ERR_USER_NLO_EVAL_HESSUBI = 1440`.  The user-defined nonlinear function reported an Hessian an invalid subscript.
* `RES_ERR_USER_NLO_EVAL_HESSUBJ = 1441`.  The user-defined nonlinear function reported an invalid subscript in the Hessian.
* `RES_ERR_USER_NLO_FUNC = 1432`.  The user-defined nonlinear function reported an error.
* `RES_ERR_WHICHITEM_NOT_ALLOWED = 1238`.  whichitem is unacceptable.
* `RES_ERR_WHICHSOL = 1236`.  The solution defined by whichsol does not exists.
* `RES_ERR_WRITE_LP_FORMAT = 1158`.  Problem cannot be written as an LP file.
* `RES_ERR_WRITE_LP_NON_UNIQUE_NAME = 1161`.  An auto-generated name is not unique.
* `RES_ERR_WRITE_MPS_INVALID_NAME = 1153`.  An invalid name is created while writing an MPS file.
* `RES_ERR_WRITE_OPF_INVALID_VAR_NAME = 1156`.  Empty variable names cannot be written to OPF files.
* `RES_ERR_WRITING_FILE = 1166`.  An error occurred while writing file
* `RES_ERR_XML_INVALID_PROBLEM_TYPE = 3600`.  The problem type is not supported by the XML format.
* `RES_ERR_Y_IS_UNDEFINED = 1449`.  The solution item y is undefined.
* `RES_OK = 0`.  No error occurred.
* `RES_TRM_INTERNAL = 10030`.  The optimizer terminated due to some internal reason.
* `RES_TRM_INTERNAL_STOP = 10031`.  The optimizer terminated for internal reasons.
* `RES_TRM_MAX_ITERATIONS = 10000`.  The optimizer terminated at the maximum number of iterations.
* `RES_TRM_MAX_NUM_SETBACKS = 10020`.  The optimizer terminated as the maximum number of set-backs was reached.
* `RES_TRM_MAX_TIME = 10001`.  The optimizer terminated at the maximum amount of time.
* `RES_TRM_MIO_NEAR_ABS_GAP = 10004`.  The mixed-integer optimizer terminated because the near optimal absolute gap tolerance was satisfied.
* `RES_TRM_MIO_NEAR_REL_GAP = 10003`.  The mixed-integer optimizer terminated because the near optimal relative gap tolerance was satisfied.
* `RES_TRM_MIO_NUM_BRANCHES = 10009`.  The mixed-integer optimizer terminated as to the maximum number of branches was reached.
* `RES_TRM_MIO_NUM_RELAXS = 10008`.  The mixed-integer optimizer terminated as the maximum number of relaxations was reached.
* `RES_TRM_NUM_MAX_NUM_INT_SOLUTIONS = 10015`.  The mixed-integer optimizer terminated as the maximum number of feasible solutions was reached.
* `RES_TRM_NUMERICAL_PROBLEM = 10025`.  The optimizer terminated due to a numerical problem.
* `RES_TRM_OBJECTIVE_RANGE = 10002`.  The optimizer terminated on the bound of the objective range.
* `RES_TRM_STALL = 10006`.  The optimizer is terminated due to slow progress.
* `RES_TRM_USER_CALLBACK = 10007`.  The user-defined progress call-back function terminated the optimization.
* `RES_WRN_ANA_ALMOST_INT_BOUNDS = 904`.  Warn against almost integral bounds.
* `RES_WRN_ANA_C_ZERO = 901`.  Warn against all objective coefficients being zero.
* `RES_WRN_ANA_CLOSE_BOUNDS = 903`.  Warn against close bounds.
* `RES_WRN_ANA_EMPTY_COLS = 902`.  Warn against empty columns.
* `RES_WRN_ANA_LARGE_BOUNDS = 900`.  Warn against very large bounds.
* `RES_WRN_CONSTRUCT_INVALID_SOL_ITG = 807`.  The initial value for one or more  of the integer variables is not feasible.
* `RES_WRN_CONSTRUCT_NO_SOL_ITG = 810`.  The construct solution requires an integer solution.
* `RES_WRN_CONSTRUCT_SOLUTION_INFEAS = 805`.  After fixing the integer variables at the suggested values then the problem is infeasible.
* `RES_WRN_DROPPED_NZ_QOBJ = 201`.  One or more non-zero elements were dropped in the Q matrix in the objective.
* `RES_WRN_DUPLICATE_BARVARIABLE_NAMES = 852`.  Two barvariable names are identical.
* `RES_WRN_DUPLICATE_CONE_NAMES = 853`.  Two cone names are identical.
* `RES_WRN_DUPLICATE_CONSTRAINT_NAMES = 850`.  Two constraint names are identical.
* `RES_WRN_DUPLICATE_VARIABLE_NAMES = 851`.  Two variable names are identical.
* `RES_WRN_ELIMINATOR_SPACE = 801`.  The eliminator is skipped at least once due to lack of space.
* `RES_WRN_EMPTY_NAME = 502`.  A variable or constraint name is empty. The output file may be invalid.
* `RES_WRN_IGNORE_INTEGER = 250`.  Ignored integer constraints.
* `RES_WRN_INCOMPLETE_LINEAR_DEPENDENCY_CHECK = 800`.  The linear dependency check(s) is incomplete.
* `RES_WRN_LARGE_AIJ = 62`.  A numerically large value is specified for an element in A.
* `RES_WRN_LARGE_BOUND = 51`.  A numerically large bound value is specified.
* `RES_WRN_LARGE_CJ = 57`.  A numerically large value is specified for one element in A.
* `RES_WRN_LARGE_CON_FX = 54`.  A equality constraint is fixed to numerically large value.
* `RES_WRN_LARGE_LO_BOUND = 52`.  A numerically large lower bound value is specified.
* `RES_WRN_LARGE_UP_BOUND = 53`.  A numerically large upper bound value is specified.
* `RES_WRN_LICENSE_EXPIRE = 500`.  The license expires.
* `RES_WRN_LICENSE_FEATURE_EXPIRE = 505`.  The license expires.
* `RES_WRN_LICENSE_SERVER = 501`.  The license server is not responding.
* `RES_WRN_LP_DROP_VARIABLE = 85`.  Ignore a variable because the variable was not previously defined.
* `RES_WRN_LP_OLD_QUAD_FORMAT = 80`.  Missing '/2' after quadratic expressions in bound or objective.
* `RES_WRN_MIO_INFEASIBLE_FINAL = 270`.  The final mixed-integer problem with all the integer variables fixed at their optimal values is infeasible.
* `RES_WRN_MPS_SPLIT_BOU_VECTOR = 72`.  A BOUNDS vector is split into several nonadjacent parts in an MPS file.
* `RES_WRN_MPS_SPLIT_RAN_VECTOR = 71`.  A RANGE vector is split into several nonadjacent parts in an MPS file.
* `RES_WRN_MPS_SPLIT_RHS_VECTOR = 70`.  An RHS vector is split into several nonadjacent parts.
* `RES_WRN_NAME_MAX_LEN = 65`.  A name is longer than the buffer that is supposed to hold it.
* `RES_WRN_NO_DUALIZER = 950`.  No automatic dualizer is available for the specified problem.
* `RES_WRN_NO_GLOBAL_OPTIMIZER = 251`.  No global optimizer is available.
* `RES_WRN_NO_NONLINEAR_FUNCTION_WRITE = 450`.  The problem contains a general nonlinear function that cannot be written to a disk file.
* `RES_WRN_NZ_IN_UPR_TRI = 200`.  Non-zero elements specified in the upper triangle of a matrix were ignored.
* `RES_WRN_OPEN_PARAM_FILE = 50`.  The parameter file could not be opened.
* `RES_WRN_PARAM_IGNORED_CMIO = 516`. A parameter was ignored by the conic mixed integer optimizer.
* `RES_WRN_PARAM_NAME_DOU = 510`. Parameter name not recognized.
* `RES_WRN_PARAM_NAME_INT = 511`. Parameter name not recognized.
* `RES_WRN_PARAM_NAME_STR = 512`. Parameter name not recognized.
* `RES_WRN_PARAM_STR_VALUE = 515`. A parameter value is not correct.
* `RES_WRN_PRESOLVE_OUTOFSPACE = 802`.  The presolve is incomplete due to lack of space.
* `RES_WRN_QUAD_CONES_WITH_ROOT_FIXED_AT_ZERO = 930`.  For at least one quadratic cone the root is fixed at (nearly) zero.
* `RES_WRN_RQUAD_CONES_WITH_ROOT_FIXED_AT_ZERO = 931`.  For at least one rotated quadratic cone the root is fixed at (nearly) zero.
* `RES_WRN_SOL_FILE_IGNORED_CON = 351`.  One or more lines in the constraint section were ignored when reading a solution file.
* `RES_WRN_SOL_FILE_IGNORED_VAR = 352`.  One or more lines in the variable section were ignored when reading a solution file.
* `RES_WRN_SOL_FILTER = 300`.  Invalid solution filter is specified.
* `RES_WRN_SPAR_MAX_LEN = 66`.  A value for a string parameter is longer than the buffer that is supposed to hold it.
* `RES_WRN_SYM_MAT_LARGE = 960`.  A numerically large value is specified for an element in E.
* `RES_WRN_TOO_FEW_BASIS_VARS = 400`.  An incomplete basis is specified.
* `RES_WRN_TOO_MANY_BASIS_VARS = 405`.  A basis with too many variables is specified.
* `RES_WRN_UNDEF_SOL_FILE_NAME = 350`.  Undefined name occurred in a solution.
* `RES_WRN_USING_GENERIC_NAMES = 503`.  Generic names are used because a name is not valid.
* `RES_WRN_WRITE_CHANGED_NAMES = 803`.  Some names were changed because they were invalid for the output file format.
* `RES_WRN_WRITE_DISCARDED_CFIX = 804`.  The fixed objective term was discarded in the output file.
* `RES_WRN_ZERO_AIJ = 63`.  One or more zero elements are specified in A.
* `RES_WRN_ZEROS_IN_SPARSE_COL = 710`.  One or more (near) zero elements are specified in a sparse column of a matrix.
* `RES_WRN_ZEROS_IN_SPARSE_ROW = 705`.  One or more (near) zero elements are specified in a sparse row of a matrix.

mionodeseltype (``int32``)
~~~~~~~~~~~~~~~~
* `MIO_NODE_SELECTION_BEST = 2`.  The optimizer employs a best bound node selection strategy.
* `MIO_NODE_SELECTION_FIRST = 1`.  The optimizer employs a depth first node selection strategy.
* `MIO_NODE_SELECTION_FREE = 0`.  The optimizer decides the node selection strategy.
* `MIO_NODE_SELECTION_HYBRID = 4`. The optimizer employs a hybrid strategy.
* `MIO_NODE_SELECTION_PSEUDO = 5`.  The optimizer employs selects the node based on a pseudo cost estimate.
* `MIO_NODE_SELECTION_WORST = 3`.  The optimizer employs a worst bound node selection strategy.

transpose (``int32``)
~~~~~~~~~~~~~~~~
* `TRANSPOSE_NO = 0`.  No transpose is applied.
* `TRANSPOSE_YES = 1`.  A transpose is applied.

onoffkey (``int32``)
~~~~~~~~~~~~~~~~
* `OFF = 0`. Switch the option off.
* `ON = 1`. Switch the option on.

simdegen (``int32``)
~~~~~~~~~~~~~~~~
* `SIM_DEGEN_AGGRESSIVE = 2`.  The simplex optimizer should use an aggressive degeneration strategy.
* `SIM_DEGEN_FREE = 1`.  The simplex optimizer chooses the degeneration strategy.
* `SIM_DEGEN_MINIMUM = 4`.  The simplex optimizer should use a minimum degeneration strategy.
* `SIM_DEGEN_MODERATE = 3`.  The simplex optimizer should use a moderate degeneration strategy.
* `SIM_DEGEN_NONE = 0`.  The simplex optimizer should use no degeneration strategy.

dataformat (``int32``)
~~~~~~~~~~~~~~~~
* `DATA_FORMAT_CB = 7`.  Conic benchmark format,
* `DATA_FORMAT_EXTENSION = 0`.  The file extension is used to determine the data file format.
* `DATA_FORMAT_FREE_MPS = 5`.  The data a free MPS formatted file.
* `DATA_FORMAT_JSON_TASK = 8`.  JSON based task format.
* `DATA_FORMAT_LP = 2`. The data file is LP formatted.
* `DATA_FORMAT_MPS = 1`. The data file is MPS formatted.
* `DATA_FORMAT_OP = 3`.  The data file is an optimization problem formatted file.
* `DATA_FORMAT_TASK = 6`.  Generic task dump file.
* `DATA_FORMAT_XML = 4`.  The data file is an XML formatted file.

orderingtype (``int32``)
~~~~~~~~~~~~~~~~
* `ORDER_METHOD_APPMINLOC = 1`.  Approximate minimum local fill-in ordering is employed.
* `ORDER_METHOD_EXPERIMENTAL = 2`.  This option should not be used.
* `ORDER_METHOD_FORCE_GRAPHPAR = 4`. Always use the graph partitioning based ordering even if it is worse than the approximate minimum local fill ordering.
* `ORDER_METHOD_FREE = 0`. The ordering method is chosen automatically.
* `ORDER_METHOD_NONE = 5`. No ordering is used.
* `ORDER_METHOD_TRY_GRAPHPAR = 3`. Always try the graph partitioning based ordering.

problemtype (``int32``)
~~~~~~~~~~~~~~~~
* `PROBTYPE_CONIC = 4`. A conic optimization.
* `PROBTYPE_GECO = 3`. General convex optimization.
* `PROBTYPE_LO = 0`. The problem is a linear optimization problem.
* `PROBTYPE_MIXED = 5`.  General nonlinear constraints and conic constraints. This combination can not be solved by MOSEK.
* `PROBTYPE_QCQO = 2`.  The problem is a quadratically constrained optimization problem.
* `PROBTYPE_QO = 1`. The problem is a quadratic optimization problem.

inftype (``int32``)
~~~~~~~~~~~~~~~~
* `INF_DOU_TYPE = 0`. Is a double information type.
* `INF_INT_TYPE = 1`. Is an integer.
* `INF_LINT_TYPE = 2`. Is a long integer.

dparam (``int32``)
~~~~~~~~~~~~~~~~
* `DPAR_ANA_SOL_INFEAS_TOL = 0`.  If a constraint violates its bound with an amount larger than this value, the constraint name, index and violation will be printed by the solution analyzer.
* `DPAR_BASIS_REL_TOL_S = 1`.  Maximum relative dual bound violation allowed in an optimal basic solution.
* `DPAR_BASIS_TOL_S = 2`.  Maximum absolute dual bound violation in an optimal basic solution.
* `DPAR_BASIS_TOL_X = 3`.  Maximum absolute primal bound violation allowed in an optimal basic solution.
* `DPAR_CHECK_CONVEXITY_REL_TOL = 4`. Convexity check tolerance.
* `DPAR_DATA_SYM_MAT_TOL = 5`.  Zero tolerance threshold for symmetric matrixes.
* `DPAR_DATA_SYM_MAT_TOL_HUGE = 6`.  Data tolerance threshold.
* `DPAR_DATA_SYM_MAT_TOL_LARGE = 7`.  Data tolerance threshold.
* `DPAR_DATA_TOL_AIJ = 8`.  Data tolerance threshold.
* `DPAR_DATA_TOL_AIJ_HUGE = 9`.  Data tolerance threshold.
* `DPAR_DATA_TOL_AIJ_LARGE = 10`.  Data tolerance threshold.
* `DPAR_DATA_TOL_BOUND_INF = 11`.  Data tolerance threshold.
* `DPAR_DATA_TOL_BOUND_WRN = 12`.  Data tolerance threshold.
* `DPAR_DATA_TOL_C_HUGE = 13`.  Data tolerance threshold.
* `DPAR_DATA_TOL_CJ_LARGE = 14`.  Data tolerance threshold.
* `DPAR_DATA_TOL_QIJ = 15`.  Data tolerance threshold.
* `DPAR_DATA_TOL_X = 16`.  Data tolerance threshold.
* `DPAR_INTPNT_CO_TOL_DFEAS = 17`.  Dual feasibility tolerance used by the conic interior-point optimizer.
* `DPAR_INTPNT_CO_TOL_INFEAS = 18`.  Infeasibility tolerance for the conic solver.
* `DPAR_INTPNT_CO_TOL_MU_RED = 19`.  Optimality tolerance for the conic solver.
* `DPAR_INTPNT_CO_TOL_NEAR_REL = 20`.  Optimality tolerance for the conic solver.
* `DPAR_INTPNT_CO_TOL_PFEAS = 21`.  Primal feasibility tolerance used by the conic interior-point optimizer.
* `DPAR_INTPNT_CO_TOL_REL_GAP = 22`.  Relative gap termination tolerance used by the conic interior-point optimizer.
* `DPAR_INTPNT_NL_MERIT_BAL = 23`.  Controls if the complementarity and infeasibility is converging to zero at about equal rates.
* `DPAR_INTPNT_NL_TOL_DFEAS = 24`.  Dual feasibility tolerance used when a nonlinear model is solved.
* `DPAR_INTPNT_NL_TOL_MU_RED = 25`. Relative complementarity gap tolerance.
* `DPAR_INTPNT_NL_TOL_NEAR_REL = 26`.  Nonlinear solver optimality tolerance parameter.
* `DPAR_INTPNT_NL_TOL_PFEAS = 27`.  Primal feasibility tolerance used when a nonlinear model is solved.
* `DPAR_INTPNT_NL_TOL_REL_GAP = 28`.  Relative gap termination tolerance for nonlinear problems.
* `DPAR_INTPNT_NL_TOL_REL_STEP = 29`.  Relative step size to the boundary for general nonlinear optimization problems.
* `DPAR_INTPNT_QO_TOL_DFEAS = 30`.  Dual feasibility tolerance used when the interior-point optimizer is applied to a quadratic optimization problem..
* `DPAR_INTPNT_QO_TOL_INFEAS = 31`.  Infeasibility tolerance employed when a quadratic optimization problem is solved.
* `DPAR_INTPNT_QO_TOL_MU_RED = 32`.  Optimality tolerance employed when a quadratic optimization problem is solved.
* `DPAR_INTPNT_QO_TOL_NEAR_REL = 33`.  Optimality tolerance employed when a quadratic optimization problem is solved.
* `DPAR_INTPNT_QO_TOL_PFEAS = 34`.  Primal feasibility tolerance used when the interior-point optimizer is applied to a quadratic optimization problem.
* `DPAR_INTPNT_QO_TOL_REL_GAP = 35`.  Relative gap termination tolerance used when the interior-point optimizer is applied to a quadratic optimization problem.
* `DPAR_INTPNT_TOL_DFEAS = 36`.  Dual feasibility tolerance used for linear and quadratic optimization problems.
* `DPAR_INTPNT_TOL_DSAFE = 37`.  Controls the interior-point dual starting point.
* `DPAR_INTPNT_TOL_INFEAS = 38`.  Nonlinear solver infeasibility tolerance parameter.
* `DPAR_INTPNT_TOL_MU_RED = 39`. Relative complementarity gap tolerance.
* `DPAR_INTPNT_TOL_PATH = 40`.  interior-point centering aggressiveness.
* `DPAR_INTPNT_TOL_PFEAS = 41`.  Primal feasibility tolerance used for linear and quadratic optimization problems.
* `DPAR_INTPNT_TOL_PSAFE = 42`.  Controls the interior-point primal starting point.
* `DPAR_INTPNT_TOL_REL_GAP = 43`. Relative gap termination tolerance.
* `DPAR_INTPNT_TOL_REL_STEP = 44`.  Relative step size to the boundary for linear and quadratic optimization problems.
* `DPAR_INTPNT_TOL_STEP_SIZE = 45`.  If the step size falls below the value of this parameter, then the interior-point optimizer assumes that it is stalled. In other words the interior-point optimizer does not make any progress and therefore it is better stop.
* `DPAR_LOWER_OBJ_CUT = 46`.  Objective bound.
* `DPAR_LOWER_OBJ_CUT_FINITE_TRH = 47`.  Objective bound.
* `DPAR_MIO_DISABLE_TERM_TIME = 48`.  Certain termination criteria is disabled within the mixed-integer optimizer for period time specified by the parameter.
* `DPAR_MIO_MAX_TIME = 49`.  Time limit for the mixed-integer optimizer.
* `DPAR_MIO_NEAR_TOL_ABS_GAP = 50`.  Relaxed absolute optimality tolerance employed by the mixed-integer optimizer.
* `DPAR_MIO_NEAR_TOL_REL_GAP = 51`.  The mixed-integer optimizer is terminated when this tolerance is satisfied.
* `DPAR_MIO_REL_GAP_CONST = 52`.  This value is used to compute the relative gap for the solution to an integer optimization problem.
* `DPAR_MIO_TOL_ABS_GAP = 53`.  Absolute optimality tolerance employed by the mixed-integer optimizer.
* `DPAR_MIO_TOL_ABS_RELAX_INT = 54`.  Integer constraint tolerance.
* `DPAR_MIO_TOL_FEAS = 55`.  Feasibility tolerance for mixed integer solver.
* `DPAR_MIO_TOL_REL_DUAL_BOUND_IMPROVEMENT = 56`.  Controls cut generation for mixed-integer optimizer.
* `DPAR_MIO_TOL_REL_GAP = 57`.  Relative optimality tolerance employed by the mixed-integer optimizer.
* `DPAR_OPTIMIZER_MAX_TIME = 58`.  Solver time limit.
* `DPAR_PRESOLVE_TOL_ABS_LINDEP = 59`.  Absolute tolerance employed by the linear dependency checker.
* `DPAR_PRESOLVE_TOL_AIJ = 60`.  Absolute zero tolerance employed for constraint coefficients in the presolve.
* `DPAR_PRESOLVE_TOL_REL_LINDEP = 61`.  Relative tolerance employed by the linear dependency checker.
* `DPAR_PRESOLVE_TOL_S = 62`.  Absolute zero tolerance employed for slack variables in the presolve.
* `DPAR_PRESOLVE_TOL_X = 63`.  Absolute zero tolerance employed for variables in the presolve.
* `DPAR_QCQO_REFORMULATE_REL_DROP_TOL = 64`.  This parameter determines when columns are dropped in incomplete Cholesky factorization during reformulation of quadratic problems.
* `DPAR_SEMIDEFINITE_TOL_APPROX = 65`.  Tolerance to define a matrix to be positive semidefinite.
* `DPAR_SIM_LU_TOL_REL_PIV = 66`.  Relative pivot tolerance employed when computing the LU factorization of the basis matrix.
* `DPAR_SIMPLEX_ABS_TOL_PIV = 67`.  Absolute pivot tolerance employed by the simplex optimizers.
* `DPAR_UPPER_OBJ_CUT = 68`.  Objective bound.
* `DPAR_UPPER_OBJ_CUT_FINITE_TRH = 69`.  Objective bound.

simdupvec (``int32``)
~~~~~~~~~~~~~~~~
* `SIM_EXPLOIT_DUPVEC_FREE = 2`.  The simplex optimizer can choose freely.
* `SIM_EXPLOIT_DUPVEC_OFF = 0`.  Disallow the simplex optimizer to exploit duplicated columns.
* `SIM_EXPLOIT_DUPVEC_ON = 1`.  Allow the simplex optimizer to exploit duplicated columns.

compresstype (``int32``)
~~~~~~~~~~~~~~~~
* `COMPRESS_FREE = 1`.  The type of compression used is chosen automatically.
* `COMPRESS_GZIP = 2`. The type of compression used is gzip compatible.
* `COMPRESS_NONE = 0`. No compression is used.

nametype (``int32``)
~~~~~~~~~~~~~~~~
* `NAME_TYPE_GEN = 0`. General names. However, no duplicate and blank names are allowed.
* `NAME_TYPE_LP = 2`. LP type names.
* `NAME_TYPE_MPS = 1`. MPS type names.

mpsformat (``int32``)
~~~~~~~~~~~~~~~~
* `MPS_FORMAT_CPLEX = 3`.  The CPLEX compatible version of the MPS format is employed.
* `MPS_FORMAT_FREE = 2`.  It is assumed that the input file satisfies the free MPS format. This implies that spaces are not allowed in names. Otherwise the format is free.
* `MPS_FORMAT_RELAXED = 1`.  It is assumed that the input file satisfies a slightly relaxed version of the MPS format.
* `MPS_FORMAT_STRICT = 0`.  It is assumed that the input file satisfies the MPS format strictly.

variabletype (``int32``)
~~~~~~~~~~~~~~~~
* `VAR_TYPE_CONT = 0`. Is a continuous variable.
* `VAR_TYPE_INT = 1`. Is an integer variable.

checkconvexitytype (``int32``)
~~~~~~~~~~~~~~~~
* `CHECK_CONVEXITY_FULL = 2`. Perform a full convexity check.
* `CHECK_CONVEXITY_NONE = 0`. No convexity check.
* `CHECK_CONVEXITY_SIMPLE = 1`. Perform simple and fast convexity check.

language (``int32``)
~~~~~~~~~~~~~~~~
* `LANG_DAN = 1`. Danish language selection
* `LANG_ENG = 0`. English language selection

startpointtype (``int32``)
~~~~~~~~~~~~~~~~
* `STARTING_POINT_CONSTANT = 2`.  The optimizer constructs a starting point by assigning a constant value to all primal and dual variables. This starting point is normally robust.
* `STARTING_POINT_FREE = 0`. The starting point is chosen automatically.
* `STARTING_POINT_GUESS = 1`.  The optimizer guesses a starting point.
* `STARTING_POINT_SATISFY_BOUNDS = 3`.  The starting point is chosen to satisfy all the simple bounds on nonlinear variables. If this starting point is employed, then more care than usual should employed when choosing the bounds on the nonlinear variables. In particular very tight bounds should be avoided.

soltype (``int32``)
~~~~~~~~~~~~~~~~
* `SOL_BAS = 1`. The basic solution.
* `SOL_ITG = 2`. The integer solution.
* `SOL_ITR = 0`. The interior solution.

scalingmethod (``int32``)
~~~~~~~~~~~~~~~~
* `SCALING_METHOD_FREE = 1`. The optimizer chooses the scaling heuristic.
* `SCALING_METHOD_POW2 = 0`. Scales only with power of 2 leaving the mantissa untouched.

value (``int32``)
~~~~~~~~~~~~~~~~
* `LICENSE_BUFFER_LENGTH = 20`. The length of a license key buffer.
* `MAX_STR_LEN = 1024`. Maximum string length allowed in MOSEK.

stakey (``int32``)
~~~~~~~~~~~~~~~~
* `SK_BAS = 1`. The constraint or variable is in the basis.
* `SK_FIX = 5`. The constraint or variable is fixed.
* `SK_INF = 6`.  The constraint or variable is infeasible in the bounds.
* `SK_LOW = 3`. The constraint or variable is at its lower bound.
* `SK_SUPBAS = 2`. The constraint or variable is super basic.
* `SK_UNK = 0`.  The status for the constraint or variable is unknown.
* `SK_UPR = 4`. The constraint or variable is at its upper bound.

simreform (``int32``)
~~~~~~~~~~~~~~~~
* `SIM_REFORMULATION_AGGRESSIVE = 3`.  The simplex optimizer should use an aggressive reformulation strategy.
* `SIM_REFORMULATION_FREE = 2`.  The simplex optimizer can choose freely.
* `SIM_REFORMULATION_OFF = 0`.  Disallow the simplex optimizer to reformulate the problem.
* `SIM_REFORMULATION_ON = 1`.  Allow the simplex optimizer to reformulate the problem.

iinfitem (``int32``)
~~~~~~~~~~~~~~~~
* `IINF_ANA_PRO_NUM_CON = 0`.  Number of constraints in the problem.
* `IINF_ANA_PRO_NUM_CON_EQ = 1`.  Number of equality constraints.
* `IINF_ANA_PRO_NUM_CON_FR = 2`.  Number of unbounded constraints.
* `IINF_ANA_PRO_NUM_CON_LO = 3`.  Number of constraints with a lower bound and an infinite upper bound.
* `IINF_ANA_PRO_NUM_CON_RA = 4`.  Number of constraints with finite lower and upper bounds.
* `IINF_ANA_PRO_NUM_CON_UP = 5`.  Number of constraints with an upper bound and an infinite lower bound.
* `IINF_ANA_PRO_NUM_VAR = 6`.  Number of variables in the problem.
* `IINF_ANA_PRO_NUM_VAR_BIN = 7`.  Number of binary variables.
* `IINF_ANA_PRO_NUM_VAR_CONT = 8`.  Number of continuous variables.
* `IINF_ANA_PRO_NUM_VAR_EQ = 9`.  Number of fixed variables.
* `IINF_ANA_PRO_NUM_VAR_FR = 10`.  Number of unbounded constraints.
* `IINF_ANA_PRO_NUM_VAR_INT = 11`.  Number of general integer variables.
* `IINF_ANA_PRO_NUM_VAR_LO = 12`.  Number of variables with a lower bound and an infinite upper bound.
* `IINF_ANA_PRO_NUM_VAR_RA = 13`.  Number of variables with finite lower and upper bounds.
* `IINF_ANA_PRO_NUM_VAR_UP = 14`.  Number of variables with an upper bound and an infinite lower bound.
* `IINF_INTPNT_FACTOR_DIM_DENSE = 15`.  Dimension of the dense sub system in factorization.
* `IINF_INTPNT_ITER = 16`.  Number of interior-point iterations since invoking the interior-point optimizer.
* `IINF_INTPNT_NUM_THREADS = 17`.  Number of threads that the interior-point optimizer is using.
* `IINF_INTPNT_SOLVE_DUAL = 18`.  Non-zero if the interior-point optimizer is solving the dual problem.
* `IINF_MIO_ABSGAP_SATISFIED = 19`.  Non-zero if absolute gap is within tolerances.
* `IINF_MIO_CLIQUE_TABLE_SIZE = 20`.  Size of the clique table.
* `IINF_MIO_CONSTRUCT_NUM_ROUNDINGS = 21`.  Number of values in the integer solution that is rounded to an integer value.
* `IINF_MIO_CONSTRUCT_SOLUTION = 22`.  If this item is positive, then MOSEK successfully constructed an initial integer feasible solution.
* `IINF_MIO_INITIAL_SOLUTION = 23`.  Is non-zero if an initial integer solution is specified.
* `IINF_MIO_NEAR_ABSGAP_SATISFIED = 24`.  Non-zero if absolute gap is within relaxed tolerances.
* `IINF_MIO_NEAR_RELGAP_SATISFIED = 25`.  Non-zero if relative gap is within relaxed tolerances.
* `IINF_MIO_NODE_DEPTH = 26`.  Depth of the last node solved.
* `IINF_MIO_NUM_ACTIVE_NODES = 27`.  Number of active branch bound nodes.
* `IINF_MIO_NUM_BRANCH = 28`.  Number of branches performed during the optimization.
* `IINF_MIO_NUM_CLIQUE_CUTS = 29`.  Number of clique cuts.
* `IINF_MIO_NUM_CMIR_CUTS = 30`.  Number of Complemented Mixed Integer Rounding (CMIR) cuts.
* `IINF_MIO_NUM_GOMORY_CUTS = 31`.  Number of Gomory cuts.
* `IINF_MIO_NUM_IMPLIED_BOUND_CUTS = 32`.  Number of implied bound cuts.
* `IINF_MIO_NUM_INT_SOLUTIONS = 33`.  Number of integer feasible solutions that has been found.
* `IINF_MIO_NUM_KNAPSACK_COVER_CUTS = 34`.  Number of clique cuts.
* `IINF_MIO_NUM_RELAX = 35`.  Number of relaxations solved during the optimization.
* `IINF_MIO_NUM_REPEATED_PRESOLVE = 36`.  Number of times presolve was repeated at root.
* `IINF_MIO_NUMCON = 37`.  Number of constraints in the problem solved by the mixed-integer optimizer.
* `IINF_MIO_NUMINT = 38`.  Number of integer variables in the problem solved be the mixed-integer optimizer.
* `IINF_MIO_NUMVAR = 39`.  Number of variables in the problem solved by the mixed-integer optimizer.
* `IINF_MIO_OBJ_BOUND_DEFINED = 40`.  Non-zero if a valid objective bound has been found, otherwise zero.
* `IINF_MIO_PRESOLVED_NUMBIN = 41`.  Number of binary variables in the problem solved be the mixed-integer optimizer.
* `IINF_MIO_PRESOLVED_NUMCON = 42`.  Number of constraints in the presolved problem.
* `IINF_MIO_PRESOLVED_NUMCONT = 43`.  Number of continuous variables in the problem solved be the mixed-integer optimizer.
* `IINF_MIO_PRESOLVED_NUMINT = 44`.  Number of integer variables in the presolved problem.
* `IINF_MIO_PRESOLVED_NUMVAR = 45`.  Number of variables in the presolved problem.
* `IINF_MIO_RELGAP_SATISFIED = 46`.  Non-zero if relative gap is within tolerances.
* `IINF_MIO_TOTAL_NUM_CUTS = 47`.  Total number of cuts generated by the mixed-integer optimizer.
* `IINF_MIO_USER_OBJ_CUT = 48`.  If it is non-zero, then the objective cut is used.
* `IINF_OPT_NUMCON = 49`.  Number of constraints in the problem solved when the optimizer is called.
* `IINF_OPT_NUMVAR = 50`.  Number of variables in the problem solved when the optimizer is called
* `IINF_OPTIMIZE_RESPONSE = 51`.  The response code returned by optimize.
* `IINF_RD_NUMBARVAR = 52`. Number of variables read.
* `IINF_RD_NUMCON = 53`. Number of constraints read.
* `IINF_RD_NUMCONE = 54`. Number of conic constraints read.
* `IINF_RD_NUMINTVAR = 55`. Number of integer-constrained variables read.
* `IINF_RD_NUMQ = 56`. Number of nonempty Q matrices read.
* `IINF_RD_NUMVAR = 57`. Number of variables read.
* `IINF_RD_PROTYPE = 58`. Problem type.
* `IINF_SIM_DUAL_DEG_ITER = 59`.  The number of dual degenerate iterations.
* `IINF_SIM_DUAL_HOTSTART = 60`.  If 1 then the dual simplex algorithm is solving from an advanced basis.
* `IINF_SIM_DUAL_HOTSTART_LU = 61`. If 1 then a valid basis factorization of full rank was located and used by the dual simplex algorithm.
* `IINF_SIM_DUAL_INF_ITER = 62`.  The number of iterations taken with dual infeasibility.
* `IINF_SIM_DUAL_ITER = 63`.  Number of dual simplex iterations during the last optimization.
* `IINF_SIM_NUMCON = 64`.  Number of constraints in the problem solved by the simplex optimizer.
* `IINF_SIM_NUMVAR = 65`.  Number of variables in the problem solved by the simplex optimizer.
* `IINF_SIM_PRIMAL_DEG_ITER = 66`.  The number of primal degenerate iterations.
* `IINF_SIM_PRIMAL_DUAL_DEG_ITER = 67`.  The number of degenerate major iterations taken by the primal dual simplex algorithm.
* `IINF_SIM_PRIMAL_DUAL_HOTSTART = 68`.  If 1 then the primal dual simplex algorithm is solving from an advanced basis.
* `IINF_SIM_PRIMAL_DUAL_HOTSTART_LU = 69`.  If 1 then a valid basis factorization of full rank was located and used by the primal dual simplex algorithm.
* `IINF_SIM_PRIMAL_DUAL_INF_ITER = 70`.  The number of master iterations with dual infeasibility taken by the primal dual simplex algorithm.
* `IINF_SIM_PRIMAL_DUAL_ITER = 71`.  Number of primal dual simplex iterations during the last optimization.
* `IINF_SIM_PRIMAL_HOTSTART = 72`.  If 1 then the primal simplex algorithm is solving from an advanced basis.
* `IINF_SIM_PRIMAL_HOTSTART_LU = 73`.  If 1 then a valid basis factorization of full rank was located and used by the primal simplex algorithm.
* `IINF_SIM_PRIMAL_INF_ITER = 74`.  The number of iterations taken with primal infeasibility.
* `IINF_SIM_PRIMAL_ITER = 75`.  Number of primal simplex iterations during the last optimization.
* `IINF_SIM_SOLVE_DUAL = 76`. Is non-zero if dual problem is solved.
* `IINF_SOL_BAS_PROSTA = 77`.  Problem status of the basic solution. Updated after each optimization.
* `IINF_SOL_BAS_SOLSTA = 78`.  Solution status of the basic solution. Updated after each optimization.
* `IINF_SOL_ITG_PROSTA = 79`.  Problem status of the integer solution. Updated after each optimization.
* `IINF_SOL_ITG_SOLSTA = 80`.  Solution status of the integer solution. Updated after each optimization.
* `IINF_SOL_ITR_PROSTA = 81`.  Problem status of the interior-point solution. Updated after each optimization.
* `IINF_SOL_ITR_SOLSTA = 82`.  Solution status of the interior-point solution. Updated after each optimization.
* `IINF_STO_NUM_A_REALLOC = 83`.  Number of times the storage for storing the linear coefficient matrix has been changed.

xmlwriteroutputtype (``int32``)
~~~~~~~~~~~~~~~~
* `WRITE_XML_MODE_COL = 1`. Write in column order.
* `WRITE_XML_MODE_ROW = 0`. Write in row order.

optimizertype (``int32``)
~~~~~~~~~~~~~~~~
* `OPTIMIZER_CONIC = 0`. The optimizer for problems having conic constraints.
* `OPTIMIZER_DUAL_SIMPLEX = 1`. The dual simplex optimizer is used.
* `OPTIMIZER_FREE = 2`. The optimizer is chosen automatically.
* `OPTIMIZER_FREE_SIMPLEX = 3`.  One of the simplex optimizers is used.
* `OPTIMIZER_INTPNT = 4`. The interior-point optimizer is used.
* `OPTIMIZER_MIXED_INT = 5`. The mixed-integer optimizer.
* `OPTIMIZER_PRIMAL_SIMPLEX = 6`. The primal simplex optimizer is used.

presolvemode (``int32``)
~~~~~~~~~~~~~~~~
* `PRESOLVE_MODE_FREE = 2`.  It is decided automatically whether to presolve before the problem is optimized.
* `PRESOLVE_MODE_OFF = 0`. The problem is not presolved before it is optimized.
* `PRESOLVE_MODE_ON = 1`.  The problem is presolved before it is optimized.

miocontsoltype (``int32``)
~~~~~~~~~~~~~~~~
* `MIO_CONT_SOL_ITG = 2`.  The reported interior-point and basic solutions are a solution to the problem with all integer variables fixed at the value they have in the integer solution. A solution is only reported in case the problem has a primal feasible solution.
* `MIO_CONT_SOL_ITG_REL = 3`.  In case the problem is primal feasible then the reported interior-point and basic solutions are a solution to the problem with all integer variables fixed at the value they have in the integer solution. If the problem is primal infeasible, then the solution to the root node problem is reported.
* `MIO_CONT_SOL_NONE = 0`.  No interior-point or basic solution are reported when the mixed-integer optimizer is used.
* `MIO_CONT_SOL_ROOT = 1`.  The reported interior-point and basic solutions are a solution to the root node problem when mixed-integer optimizer is used.
