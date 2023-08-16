package mosek

// /*<comment>*/

// #include <stdlib.h>
// #include <stdint.h>
// #cgo LDFLAGS: -lmosek64
//
// typedef void * MSKuserhandle_t;
// typedef const int8_t * string_t;
// typedef void * MSKtask_t;
// typedef void * MSKenv_t;
// typedef void  (* MSKstreamfunc) ( void *, const string_t);
// typedef int32_t  (* MSKcallbackfunc) ( MSKtask_t, MSKuserhandle_t, int32_t, const double *, const int32_t *, const int64_t *);
// 
// extern void streamfunc_log(void *, char*);
// extern void streamfunc_wrn(void *, char*);
// extern void streamfunc_msg(void *, char*);
// extern void streamfunc_err(void *, char*);
// extern int callbackfunc(void *, void *, int, double*,int32_t*, int64_t*);
// extern int MSK_makeenv(MSKenv_t *, const string_t);
// extern int MSK_makeemptytask(MSKenv_t, MSKtask_t *);
// extern int MSK_deleteenv(MSKenv_t *);
// extern int MSK_deletetask(MSKtask_t *);
// extern int MSK_getlasterror64(MSKtask_t,int*,int64_t,int64_t*,char*);
// extern int MSK_linkfunctotaskstream(MSKtask_t task,int32_t whichstream, MSKuserhandle_t handle, MSKstreamfunc func);
// int32_t MSK_putcallbackfunc(MSKtask_t task,MSKcallbackfunc func,MSKuserhandle_t handle);
// extern int MSK_analyzenames(MSKtask_t,int32_t,int32_t);
// extern int MSK_analyzeproblem(MSKtask_t,int32_t);
// extern int MSK_analyzesolution(MSKtask_t,int32_t,int32_t);
// extern int MSK_appendacc(MSKtask_t,int64_t,int64_t,int64_t *,double *);
// extern int MSK_appendaccs(MSKtask_t,int64_t,int64_t *,int64_t,int64_t *,double *);
// extern int MSK_appendaccseq(MSKtask_t,int64_t,int64_t,int64_t,double *);
// extern int MSK_appendaccsseq(MSKtask_t,int64_t,int64_t *,int64_t,int64_t,double *);
// extern int MSK_appendafes(MSKtask_t,int64_t);
// extern int MSK_appendbarvars(MSKtask_t,int32_t,int32_t *);
// extern int MSK_appendcone(MSKtask_t,int32_t,double,int32_t,int32_t *);
// extern int MSK_appendconeseq(MSKtask_t,int32_t,double,int32_t,int32_t);
// extern int MSK_appendconesseq(MSKtask_t,int32_t,int32_t *,double *,int32_t *,int32_t);
// extern int MSK_appendcons(MSKtask_t,int32_t);
// extern int MSK_appenddjcs(MSKtask_t,int64_t);
// extern int MSK_appenddualexpconedomain(MSKtask_t,int64_t *);
// extern int MSK_appenddualgeomeanconedomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appenddualpowerconedomain(MSKtask_t,int64_t,int64_t,double *,int64_t *);
// extern int MSK_appendprimalexpconedomain(MSKtask_t,int64_t *);
// extern int MSK_appendprimalgeomeanconedomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appendprimalpowerconedomain(MSKtask_t,int64_t,int64_t,double *,int64_t *);
// extern int MSK_appendquadraticconedomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appendrdomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appendrminusdomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appendrplusdomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appendrquadraticconedomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appendrzerodomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appendsparsesymmat(MSKtask_t,int32_t,int64_t,int32_t *,int32_t *,double *,int64_t *);
// extern int MSK_appendsparsesymmatlist(MSKtask_t,int32_t,int32_t *,int64_t *,int32_t *,int32_t *,double *,int64_t *);
// extern int MSK_appendsvecpsdconedomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_appendvars(MSKtask_t,int32_t);
// extern int MSK_basiscond(MSKtask_t,double *,double *);
// extern int MSK_checkmemtask(MSKtask_t,const char *,int32_t);
// extern int MSK_chgconbound(MSKtask_t,int32_t,int32_t,int32_t,double);
// extern int MSK_chgvarbound(MSKtask_t,int32_t,int32_t,int32_t,double);
// extern int MSK_commitchanges(MSKtask_t);
// extern int MSK_deletesolution(MSKtask_t,int32_t);
// extern int MSK_dualsensitivity(MSKtask_t,int32_t,int32_t *,double *,double *,double *,double *);
// extern int MSK_emptyafebarfrow(MSKtask_t,int64_t);
// extern int MSK_emptyafebarfrowlist(MSKtask_t,int64_t,int64_t *);
// extern int MSK_emptyafefcol(MSKtask_t,int32_t);
// extern int MSK_emptyafefcollist(MSKtask_t,int64_t,int32_t *);
// extern int MSK_emptyafefrow(MSKtask_t,int64_t);
// extern int MSK_emptyafefrowlist(MSKtask_t,int64_t,int64_t *);
// extern int MSK_evaluateacc(MSKtask_t,int32_t,int64_t,double *);
// extern int MSK_evaluateaccs(MSKtask_t,int32_t,double *);
// extern int MSK_generateaccnames(MSKtask_t,int64_t,int64_t *,const char *,int32_t,int32_t *,int64_t *,int32_t,int32_t *,int64_t,const char **);
// extern int MSK_generatebarvarnames(MSKtask_t,int32_t,int32_t *,const char *,int32_t,int32_t *,int64_t *,int32_t,int32_t *,int64_t,const char **);
// extern int MSK_generateconenames(MSKtask_t,int32_t,int32_t *,const char *,int32_t,int32_t *,int64_t *,int32_t,int32_t *,int64_t,const char **);
// extern int MSK_generateconnames(MSKtask_t,int32_t,int32_t *,const char *,int32_t,int32_t *,int64_t *,int32_t,int32_t *,int64_t,const char **);
// extern int MSK_generatedjcnames(MSKtask_t,int64_t,int64_t *,const char *,int32_t,int32_t *,int64_t *,int32_t,int32_t *,int64_t,const char **);
// extern int MSK_generatevarnames(MSKtask_t,int32_t,int32_t *,const char *,int32_t,int32_t *,int64_t *,int32_t,int32_t *,int64_t,const char **);
// extern int MSK_getaccafeidxlist(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getaccb(MSKtask_t,int64_t,double *);
// extern int MSK_getaccbarfblocktriplet(MSKtask_t,int64_t,int64_t *,int64_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_getaccbarfnumblocktriplets(MSKtask_t,int64_t *);
// extern int MSK_getaccdomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getaccdoty(MSKtask_t,int32_t,int64_t,double *);
// extern int MSK_getaccdotys(MSKtask_t,int32_t,double *);
// extern int MSK_getaccfnumnz(MSKtask_t,int64_t *);
// extern int MSK_getaccftrip(MSKtask_t,int64_t *,int32_t *,double *);
// extern int MSK_getaccgvector(MSKtask_t,double *);
// extern int MSK_getaccn(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getaccname(MSKtask_t,int64_t,int32_t,unsigned char *);
// extern int MSK_getaccnamelen(MSKtask_t,int64_t,int32_t *);
// extern int MSK_getaccntot(MSKtask_t,int64_t *);
// extern int MSK_getaccs(MSKtask_t,int64_t *,int64_t *,double *);
// extern int MSK_getacol(MSKtask_t,int32_t,int32_t *,int32_t *,double *);
// extern int MSK_getacolnumnz(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getacolslice64(MSKtask_t,int32_t,int32_t,int64_t,int64_t *,int64_t *,int32_t *,double *);
// extern int MSK_getacolslicenumnz64(MSKtask_t,int32_t,int32_t,int64_t *);
// extern int MSK_getacolslicetrip(MSKtask_t,int32_t,int32_t,int64_t,int32_t *,int32_t *,double *);
// extern int MSK_getafebarfblocktriplet(MSKtask_t,int64_t,int64_t *,int64_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_getafebarfnumblocktriplets(MSKtask_t,int64_t *);
// extern int MSK_getafebarfnumrowentries(MSKtask_t,int64_t,int32_t *);
// extern int MSK_getafebarfrow(MSKtask_t,int64_t,int32_t *,int64_t *,int64_t *,int64_t *,double *);
// extern int MSK_getafebarfrowinfo(MSKtask_t,int64_t,int32_t *,int64_t *);
// extern int MSK_getafefnumnz(MSKtask_t,int64_t *);
// extern int MSK_getafefrow(MSKtask_t,int64_t,int32_t *,int32_t *,double *);
// extern int MSK_getafefrownumnz(MSKtask_t,int64_t,int32_t *);
// extern int MSK_getafeftrip(MSKtask_t,int64_t *,int32_t *,double *);
// extern int MSK_getafeg(MSKtask_t,int64_t,double *);
// extern int MSK_getafegslice(MSKtask_t,int64_t,int64_t,double *);
// extern int MSK_getaij(MSKtask_t,int32_t,int32_t,double *);
// extern int MSK_getapiecenumnz(MSKtask_t,int32_t,int32_t,int32_t,int32_t,int32_t *);
// extern int MSK_getarow(MSKtask_t,int32_t,int32_t *,int32_t *,double *);
// extern int MSK_getarownumnz(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getarowslice64(MSKtask_t,int32_t,int32_t,int64_t,int64_t *,int64_t *,int32_t *,double *);
// extern int MSK_getarowslicenumnz64(MSKtask_t,int32_t,int32_t,int64_t *);
// extern int MSK_getarowslicetrip(MSKtask_t,int32_t,int32_t,int64_t,int32_t *,int32_t *,double *);
// extern int MSK_getatrip(MSKtask_t,int64_t,int32_t *,int32_t *,double *);
// extern int MSK_getatruncatetol(MSKtask_t,double *);
// extern int MSK_getbarablocktriplet(MSKtask_t,int64_t,int64_t *,int32_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_getbaraidx(MSKtask_t,int64_t,int64_t,int32_t *,int32_t *,int64_t *,int64_t *,double *);
// extern int MSK_getbaraidxij(MSKtask_t,int64_t,int32_t *,int32_t *);
// extern int MSK_getbaraidxinfo(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getbarasparsity(MSKtask_t,int64_t,int64_t *,int64_t *);
// extern int MSK_getbarcblocktriplet(MSKtask_t,int64_t,int64_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_getbarcidx(MSKtask_t,int64_t,int64_t,int32_t *,int64_t *,int64_t *,double *);
// extern int MSK_getbarcidxinfo(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getbarcidxj(MSKtask_t,int64_t,int32_t *);
// extern int MSK_getbarcsparsity(MSKtask_t,int64_t,int64_t *,int64_t *);
// extern int MSK_getbarsj(MSKtask_t,int32_t,int32_t,double *);
// extern int MSK_getbarsslice(MSKtask_t,int32_t,int32_t,int32_t,int64_t,double *);
// extern int MSK_getbarvarname(MSKtask_t,int32_t,int32_t,unsigned char *);
// extern int MSK_getbarvarnameindex(MSKtask_t,const char *,int32_t *,int32_t *);
// extern int MSK_getbarvarnamelen(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getbarxj(MSKtask_t,int32_t,int32_t,double *);
// extern int MSK_getbarxslice(MSKtask_t,int32_t,int32_t,int32_t,int64_t,double *);
// extern int MSK_getc(MSKtask_t,double *);
// extern int MSK_getcfix(MSKtask_t,double *);
// extern int MSK_getcj(MSKtask_t,int32_t,double *);
// extern int MSK_getclist(MSKtask_t,int32_t,int32_t *,double *);
// extern int MSK_getconbound(MSKtask_t,int32_t,int32_t *,double *,double *);
// extern int MSK_getconboundslice(MSKtask_t,int32_t,int32_t,int32_t *,double *,double *);
// extern int MSK_getcone(MSKtask_t,int32_t,int32_t *,double *,int32_t *,int32_t *);
// extern int MSK_getconeinfo(MSKtask_t,int32_t,int32_t *,double *,int32_t *);
// extern int MSK_getconename(MSKtask_t,int32_t,int32_t,unsigned char *);
// extern int MSK_getconenameindex(MSKtask_t,const char *,int32_t *,int32_t *);
// extern int MSK_getconenamelen(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getconname(MSKtask_t,int32_t,int32_t,unsigned char *);
// extern int MSK_getconnameindex(MSKtask_t,const char *,int32_t *,int32_t *);
// extern int MSK_getconnamelen(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getcslice(MSKtask_t,int32_t,int32_t,double *);
// extern int MSK_getdimbarvarj(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getdjcafeidxlist(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getdjcb(MSKtask_t,int64_t,double *);
// extern int MSK_getdjcdomainidxlist(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getdjcname(MSKtask_t,int64_t,int32_t,unsigned char *);
// extern int MSK_getdjcnamelen(MSKtask_t,int64_t,int32_t *);
// extern int MSK_getdjcnumafe(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getdjcnumafetot(MSKtask_t,int64_t *);
// extern int MSK_getdjcnumdomain(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getdjcnumdomaintot(MSKtask_t,int64_t *);
// extern int MSK_getdjcnumterm(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getdjcnumtermtot(MSKtask_t,int64_t *);
// extern int MSK_getdjcs(MSKtask_t,int64_t *,int64_t *,double *,int64_t *,int64_t *);
// extern int MSK_getdjctermsizelist(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getdomainn(MSKtask_t,int64_t,int64_t *);
// extern int MSK_getdomainname(MSKtask_t,int64_t,int32_t,unsigned char *);
// extern int MSK_getdomainnamelen(MSKtask_t,int64_t,int32_t *);
// extern int MSK_getdomaintype(MSKtask_t,int64_t,int32_t *);
// extern int MSK_getdouinf(MSKtask_t,int32_t,double *);
// extern int MSK_getdouparam(MSKtask_t,int32_t,double *);
// extern int MSK_getdualobj(MSKtask_t,int32_t,double *);
// extern int MSK_getdualsolutionnorms(MSKtask_t,int32_t,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_getdviolacc(MSKtask_t,int32_t,int64_t,int64_t *,double *);
// extern int MSK_getdviolbarvar(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_getdviolcon(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_getdviolcones(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_getdviolvar(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_getinfindex(MSKtask_t,int32_t,const char *,int32_t *);
// extern int MSK_getinfmax(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getinfname(MSKtask_t,int32_t,int32_t,unsigned char *);
// extern int MSK_getintinf(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getintparam(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getlenbarvarj(MSKtask_t,int32_t,int64_t *);
// extern int MSK_getlintinf(MSKtask_t,int32_t,int64_t *);
// extern int MSK_getmaxnumanz64(MSKtask_t,int64_t *);
// extern int MSK_getmaxnumbarvar(MSKtask_t,int32_t *);
// extern int MSK_getmaxnumcon(MSKtask_t,int32_t *);
// extern int MSK_getmaxnumcone(MSKtask_t,int32_t *);
// extern int MSK_getmaxnumqnz64(MSKtask_t,int64_t *);
// extern int MSK_getmaxnumvar(MSKtask_t,int32_t *);
// extern int MSK_getmemusagetask(MSKtask_t,int64_t *,int64_t *);
// extern int MSK_getnumacc(MSKtask_t,int64_t *);
// extern int MSK_getnumafe(MSKtask_t,int64_t *);
// extern int MSK_getnumanz(MSKtask_t,int32_t *);
// extern int MSK_getnumanz64(MSKtask_t,int64_t *);
// extern int MSK_getnumbarablocktriplets(MSKtask_t,int64_t *);
// extern int MSK_getnumbaranz(MSKtask_t,int64_t *);
// extern int MSK_getnumbarcblocktriplets(MSKtask_t,int64_t *);
// extern int MSK_getnumbarcnz(MSKtask_t,int64_t *);
// extern int MSK_getnumbarvar(MSKtask_t,int32_t *);
// extern int MSK_getnumcon(MSKtask_t,int32_t *);
// extern int MSK_getnumcone(MSKtask_t,int32_t *);
// extern int MSK_getnumconemem(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getnumdjc(MSKtask_t,int64_t *);
// extern int MSK_getnumdomain(MSKtask_t,int64_t *);
// extern int MSK_getnumintvar(MSKtask_t,int32_t *);
// extern int MSK_getnumparam(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getnumqconknz64(MSKtask_t,int32_t,int64_t *);
// extern int MSK_getnumqobjnz64(MSKtask_t,int64_t *);
// extern int MSK_getnumsymmat(MSKtask_t,int64_t *);
// extern int MSK_getnumvar(MSKtask_t,int32_t *);
// extern int MSK_getobjname(MSKtask_t,int32_t,unsigned char *);
// extern int MSK_getobjnamelen(MSKtask_t,int32_t *);
// extern int MSK_getobjsense(MSKtask_t,int32_t *);
// extern int MSK_getparammax(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getparamname(MSKtask_t,int32_t,int32_t,unsigned char *);
// extern int MSK_getpowerdomainalpha(MSKtask_t,int64_t,double *);
// extern int MSK_getpowerdomaininfo(MSKtask_t,int64_t,int64_t *,int64_t *);
// extern int MSK_getprimalobj(MSKtask_t,int32_t,double *);
// extern int MSK_getprimalsolutionnorms(MSKtask_t,int32_t,double *,double *,double *);
// extern int MSK_getprobtype(MSKtask_t,int32_t *);
// extern int MSK_getprosta(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getpviolacc(MSKtask_t,int32_t,int64_t,int64_t *,double *);
// extern int MSK_getpviolbarvar(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_getpviolcon(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_getpviolcones(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_getpvioldjc(MSKtask_t,int32_t,int64_t,int64_t *,double *);
// extern int MSK_getpviolvar(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_getqconk64(MSKtask_t,int32_t,int64_t,int64_t *,int32_t *,int32_t *,double *);
// extern int MSK_getqobj64(MSKtask_t,int64_t,int64_t *,int32_t *,int32_t *,double *);
// extern int MSK_getqobjij(MSKtask_t,int32_t,int32_t,double *);
// extern int MSK_getreducedcosts(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_getskc(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getskcslice(MSKtask_t,int32_t,int32_t,int32_t,int32_t *);
// extern int MSK_getskn(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getskx(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getskxslice(MSKtask_t,int32_t,int32_t,int32_t,int32_t *);
// extern int MSK_getslc(MSKtask_t,int32_t,double *);
// extern int MSK_getslcslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_getslx(MSKtask_t,int32_t,double *);
// extern int MSK_getslxslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_getsnx(MSKtask_t,int32_t,double *);
// extern int MSK_getsnxslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_getsolsta(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getsolution(MSKtask_t,int32_t,int32_t *,int32_t *,int32_t *,int32_t *,int32_t *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_getsolutioninfo(MSKtask_t,int32_t,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_getsolutioninfonew(MSKtask_t,int32_t,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_getsolutionnew(MSKtask_t,int32_t,int32_t *,int32_t *,int32_t *,int32_t *,int32_t *,double *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_getsolutionslice(MSKtask_t,int32_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_getsparsesymmat(MSKtask_t,int64_t,int64_t,int32_t *,int32_t *,double *);
// extern int MSK_getstrparam(MSKtask_t,int32_t,int32_t,int32_t *,unsigned char *);
// extern int MSK_getstrparamlen(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getsuc(MSKtask_t,int32_t,double *);
// extern int MSK_getsucslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_getsux(MSKtask_t,int32_t,double *);
// extern int MSK_getsuxslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_getsymmatinfo(MSKtask_t,int64_t,int32_t *,int64_t *,int32_t *);
// extern int MSK_gettaskname(MSKtask_t,int32_t,unsigned char *);
// extern int MSK_gettasknamelen(MSKtask_t,int32_t *);
// extern int MSK_getvarbound(MSKtask_t,int32_t,int32_t *,double *,double *);
// extern int MSK_getvarboundslice(MSKtask_t,int32_t,int32_t,int32_t *,double *,double *);
// extern int MSK_getvarname(MSKtask_t,int32_t,int32_t,unsigned char *);
// extern int MSK_getvarnameindex(MSKtask_t,const char *,int32_t *,int32_t *);
// extern int MSK_getvarnamelen(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getvartype(MSKtask_t,int32_t,int32_t *);
// extern int MSK_getvartypelist(MSKtask_t,int32_t,int32_t *,int32_t *);
// extern int MSK_getxc(MSKtask_t,int32_t,double *);
// extern int MSK_getxcslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_getxx(MSKtask_t,int32_t,double *);
// extern int MSK_getxxslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_gety(MSKtask_t,int32_t,double *);
// extern int MSK_getyslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_infeasibilityreport(MSKtask_t,int32_t,int32_t);
// extern int MSK_initbasissolve(MSKtask_t,int32_t *);
// extern int MSK_inputdata64(MSKtask_t,int32_t,int32_t,int32_t,int32_t,double *,double,int64_t *,int64_t *,int32_t *,double *,int32_t *,double *,double *,int32_t *,double *,double *);
// extern int MSK_isdouparname(MSKtask_t,const char *,int32_t *);
// extern int MSK_isintparname(MSKtask_t,const char *,int32_t *);
// extern int MSK_isstrparname(MSKtask_t,const char *,int32_t *);
// extern int MSK_linkfiletotaskstream(MSKtask_t,int32_t,const char *,int32_t);
// extern int MSK_onesolutionsummary(MSKtask_t,int32_t,int32_t);
// extern int MSK_optimizersummary(MSKtask_t,int32_t);
// extern int MSK_optimizetrm(MSKtask_t,int32_t *);
// extern int MSK_primalrepair(MSKtask_t,double *,double *,double *,double *);
// extern int MSK_primalsensitivity(MSKtask_t,int32_t,int32_t *,int32_t *,int32_t,int32_t *,int32_t *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_probtypetostr(MSKtask_t,int32_t,unsigned char *);
// extern int MSK_prostatostr(MSKtask_t,int32_t,unsigned char *);
// extern int MSK_putacc(MSKtask_t,int64_t,int64_t,int64_t,int64_t *,double *);
// extern int MSK_putaccb(MSKtask_t,int64_t,int64_t,double *);
// extern int MSK_putaccbj(MSKtask_t,int64_t,int64_t,double);
// extern int MSK_putaccdoty(MSKtask_t,int32_t,int64_t,double *);
// extern int MSK_putacclist(MSKtask_t,int64_t,int64_t *,int64_t *,int64_t,int64_t *,double *);
// extern int MSK_putaccname(MSKtask_t,int64_t,const char *);
// extern int MSK_putacol(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_putacollist(MSKtask_t,int32_t,int32_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_putacolslice64(MSKtask_t,int32_t,int32_t,int64_t *,int64_t *,int32_t *,double *);
// extern int MSK_putafebarfblocktriplet(MSKtask_t,int64_t,int64_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_putafebarfentry(MSKtask_t,int64_t,int32_t,int64_t,int64_t *,double *);
// extern int MSK_putafebarfentrylist(MSKtask_t,int64_t,int64_t *,int32_t *,int64_t *,int64_t *,int64_t,int64_t *,double *);
// extern int MSK_putafebarfrow(MSKtask_t,int64_t,int32_t,int32_t *,int64_t *,int64_t *,int64_t,int64_t *,double *);
// extern int MSK_putafefcol(MSKtask_t,int32_t,int64_t,int64_t *,double *);
// extern int MSK_putafefentry(MSKtask_t,int64_t,int32_t,double);
// extern int MSK_putafefentrylist(MSKtask_t,int64_t,int64_t *,int32_t *,double *);
// extern int MSK_putafefrow(MSKtask_t,int64_t,int32_t,int32_t *,double *);
// extern int MSK_putafefrowlist(MSKtask_t,int64_t,int64_t *,int32_t *,int64_t *,int64_t,int32_t *,double *);
// extern int MSK_putafeg(MSKtask_t,int64_t,double);
// extern int MSK_putafeglist(MSKtask_t,int64_t,int64_t *,double *);
// extern int MSK_putafegslice(MSKtask_t,int64_t,int64_t,double *);
// extern int MSK_putaij(MSKtask_t,int32_t,int32_t,double);
// extern int MSK_putaijlist64(MSKtask_t,int64_t,int32_t *,int32_t *,double *);
// extern int MSK_putarow(MSKtask_t,int32_t,int32_t,int32_t *,double *);
// extern int MSK_putarowlist64(MSKtask_t,int32_t,int32_t *,int64_t *,int64_t *,int32_t *,double *);
// extern int MSK_putarowslice64(MSKtask_t,int32_t,int32_t,int64_t *,int64_t *,int32_t *,double *);
// extern int MSK_putatruncatetol(MSKtask_t,double);
// extern int MSK_putbarablocktriplet(MSKtask_t,int64_t,int32_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_putbaraij(MSKtask_t,int32_t,int32_t,int64_t,int64_t *,double *);
// extern int MSK_putbaraijlist(MSKtask_t,int32_t,int32_t *,int32_t *,int64_t *,int64_t *,int64_t *,double *);
// extern int MSK_putbararowlist(MSKtask_t,int32_t,int32_t *,int64_t *,int64_t *,int32_t *,int64_t *,int64_t *,double *);
// extern int MSK_putbarcblocktriplet(MSKtask_t,int64_t,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_putbarcj(MSKtask_t,int32_t,int64_t,int64_t *,double *);
// extern int MSK_putbarsj(MSKtask_t,int32_t,int32_t,double *);
// extern int MSK_putbarvarname(MSKtask_t,int32_t,const char *);
// extern int MSK_putbarxj(MSKtask_t,int32_t,int32_t,double *);
// extern int MSK_putcfix(MSKtask_t,double);
// extern int MSK_putcj(MSKtask_t,int32_t,double);
// extern int MSK_putclist(MSKtask_t,int32_t,int32_t *,double *);
// extern int MSK_putconbound(MSKtask_t,int32_t,int32_t,double,double);
// extern int MSK_putconboundlist(MSKtask_t,int32_t,int32_t *,int32_t *,double *,double *);
// extern int MSK_putconboundlistconst(MSKtask_t,int32_t,int32_t *,int32_t,double,double);
// extern int MSK_putconboundslice(MSKtask_t,int32_t,int32_t,int32_t *,double *,double *);
// extern int MSK_putconboundsliceconst(MSKtask_t,int32_t,int32_t,int32_t,double,double);
// extern int MSK_putcone(MSKtask_t,int32_t,int32_t,double,int32_t,int32_t *);
// extern int MSK_putconename(MSKtask_t,int32_t,const char *);
// extern int MSK_putconname(MSKtask_t,int32_t,const char *);
// extern int MSK_putconsolutioni(MSKtask_t,int32_t,int32_t,int32_t,double,double,double);
// extern int MSK_putcslice(MSKtask_t,int32_t,int32_t,double *);
// extern int MSK_putdjc(MSKtask_t,int64_t,int64_t,int64_t *,int64_t,int64_t *,double *,int64_t,int64_t *);
// extern int MSK_putdjcname(MSKtask_t,int64_t,const char *);
// extern int MSK_putdjcslice(MSKtask_t,int64_t,int64_t,int64_t,int64_t *,int64_t,int64_t *,double *,int64_t,int64_t *,int64_t *);
// extern int MSK_putdomainname(MSKtask_t,int64_t,const char *);
// extern int MSK_putdouparam(MSKtask_t,int32_t,double);
// extern int MSK_putintparam(MSKtask_t,int32_t,int32_t);
// extern int MSK_putmaxnumacc(MSKtask_t,int64_t);
// extern int MSK_putmaxnumafe(MSKtask_t,int64_t);
// extern int MSK_putmaxnumanz(MSKtask_t,int64_t);
// extern int MSK_putmaxnumbarvar(MSKtask_t,int32_t);
// extern int MSK_putmaxnumcon(MSKtask_t,int32_t);
// extern int MSK_putmaxnumcone(MSKtask_t,int32_t);
// extern int MSK_putmaxnumdjc(MSKtask_t,int64_t);
// extern int MSK_putmaxnumdomain(MSKtask_t,int64_t);
// extern int MSK_putmaxnumqnz(MSKtask_t,int64_t);
// extern int MSK_putmaxnumvar(MSKtask_t,int32_t);
// extern int MSK_putnadouparam(MSKtask_t,const char *,double);
// extern int MSK_putnaintparam(MSKtask_t,const char *,int32_t);
// extern int MSK_putnastrparam(MSKtask_t,const char *,const char *);
// extern int MSK_putobjname(MSKtask_t,const char *);
// extern int MSK_putobjsense(MSKtask_t,int32_t);
// extern int MSK_putoptserverhost(MSKtask_t,const char *);
// extern int MSK_putparam(MSKtask_t,const char *,const char *);
// extern int MSK_putqcon(MSKtask_t,int32_t,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_putqconk(MSKtask_t,int32_t,int32_t,int32_t *,int32_t *,double *);
// extern int MSK_putqobj(MSKtask_t,int32_t,int32_t *,int32_t *,double *);
// extern int MSK_putqobjij(MSKtask_t,int32_t,int32_t,double);
// extern int MSK_putskc(MSKtask_t,int32_t,int32_t *);
// extern int MSK_putskcslice(MSKtask_t,int32_t,int32_t,int32_t,int32_t *);
// extern int MSK_putskx(MSKtask_t,int32_t,int32_t *);
// extern int MSK_putskxslice(MSKtask_t,int32_t,int32_t,int32_t,int32_t *);
// extern int MSK_putslc(MSKtask_t,int32_t,double *);
// extern int MSK_putslcslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_putslx(MSKtask_t,int32_t,double *);
// extern int MSK_putslxslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_putsnx(MSKtask_t,int32_t,double *);
// extern int MSK_putsnxslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_putsolution(MSKtask_t,int32_t,int32_t *,int32_t *,int32_t *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_putsolutionnew(MSKtask_t,int32_t,int32_t *,int32_t *,int32_t *,double *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_putsolutionyi(MSKtask_t,int32_t,int32_t,double);
// extern int MSK_putstrparam(MSKtask_t,int32_t,const char *);
// extern int MSK_putsuc(MSKtask_t,int32_t,double *);
// extern int MSK_putsucslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_putsux(MSKtask_t,int32_t,double *);
// extern int MSK_putsuxslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_puttaskname(MSKtask_t,const char *);
// extern int MSK_putvarbound(MSKtask_t,int32_t,int32_t,double,double);
// extern int MSK_putvarboundlist(MSKtask_t,int32_t,int32_t *,int32_t *,double *,double *);
// extern int MSK_putvarboundlistconst(MSKtask_t,int32_t,int32_t *,int32_t,double,double);
// extern int MSK_putvarboundslice(MSKtask_t,int32_t,int32_t,int32_t *,double *,double *);
// extern int MSK_putvarboundsliceconst(MSKtask_t,int32_t,int32_t,int32_t,double,double);
// extern int MSK_putvarname(MSKtask_t,int32_t,const char *);
// extern int MSK_putvarsolutionj(MSKtask_t,int32_t,int32_t,int32_t,double,double,double,double);
// extern int MSK_putvartype(MSKtask_t,int32_t,int32_t);
// extern int MSK_putvartypelist(MSKtask_t,int32_t,int32_t *,int32_t *);
// extern int MSK_putxc(MSKtask_t,int32_t,double *);
// extern int MSK_putxcslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_putxx(MSKtask_t,int32_t,double *);
// extern int MSK_putxxslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_puty(MSKtask_t,int32_t,double *);
// extern int MSK_putyslice(MSKtask_t,int32_t,int32_t,int32_t,double *);
// extern int MSK_readbsolution(MSKtask_t,const char *,int32_t);
// extern int MSK_readdataautoformat(MSKtask_t,const char *);
// extern int MSK_readdataformat(MSKtask_t,const char *,int32_t,int32_t);
// extern int MSK_readjsonsol(MSKtask_t,const char *);
// extern int MSK_readjsonstring(MSKtask_t,const char *);
// extern int MSK_readlpstring(MSKtask_t,const char *);
// extern int MSK_readopfstring(MSKtask_t,const char *);
// extern int MSK_readparamfile(MSKtask_t,const char *);
// extern int MSK_readptfstring(MSKtask_t,const char *);
// extern int MSK_readsolution(MSKtask_t,int32_t,const char *);
// extern int MSK_readsolutionfile(MSKtask_t,const char *);
// extern int MSK_readsummary(MSKtask_t,int32_t);
// extern int MSK_readtask(MSKtask_t,const char *);
// extern int MSK_removebarvars(MSKtask_t,int32_t,int32_t *);
// extern int MSK_removecones(MSKtask_t,int32_t,int32_t *);
// extern int MSK_removecons(MSKtask_t,int32_t,int32_t *);
// extern int MSK_removevars(MSKtask_t,int32_t,int32_t *);
// extern int MSK_resizetask(MSKtask_t,int32_t,int32_t,int32_t,int64_t,int64_t);
// extern int MSK_sensitivityreport(MSKtask_t,int32_t);
// extern int MSK_setdefaults(MSKtask_t);
// extern int MSK_sktostr(MSKtask_t,int32_t,unsigned char *);
// extern int MSK_solstatostr(MSKtask_t,int32_t,unsigned char *);
// extern int MSK_solutiondef(MSKtask_t,int32_t,int *);
// extern int MSK_solutionsummary(MSKtask_t,int32_t);
// extern int MSK_solvewithbasis(MSKtask_t,int,int32_t,int32_t *,double *,int32_t *);
// extern int MSK_strtoconetype(MSKtask_t,const char *,int32_t *);
// extern int MSK_strtosk(MSKtask_t,const char *,int32_t *);
// extern int MSK_toconic(MSKtask_t);
// extern int MSK_updatesolutioninfo(MSKtask_t,int32_t);
// extern int MSK_writebsolution(MSKtask_t,const char *,int32_t);
// extern int MSK_writedata(MSKtask_t,const char *);
// extern int MSK_writejsonsol(MSKtask_t,const char *);
// extern int MSK_writeparamfile(MSKtask_t,const char *);
// extern int MSK_writesolution(MSKtask_t,int32_t,const char *);
// extern int MSK_writesolutionfile(MSKtask_t,const char *);
// extern int MSK_writetask(MSKtask_t,const char *);
// extern int MSK_axpy(MSKenv_t,int32_t,double,double *,double *);
// extern int MSK_checkinall(MSKenv_t);
// extern int MSK_checkinlicense(MSKenv_t,int32_t);
// extern int MSK_checkoutlicense(MSKenv_t,int32_t);
// extern int MSK_dot(MSKenv_t,int32_t,double *,double *,double *);
// extern int MSK_echointro(MSKenv_t,int32_t);
// extern int MSK_expirylicenses(MSKenv_t,int64_t *);
// extern int MSK_gemm(MSKenv_t,int32_t,int32_t,int32_t,int32_t,int32_t,double,double *,double *,double,double *);
// extern int MSK_gemv(MSKenv_t,int32_t,int32_t,int32_t,double,double *,double *,double,double *);
// extern int MSK_getcodedesc(int32_t,unsigned char *,unsigned char *);
// extern int MSK_getversion(int32_t *,int32_t *,int32_t *);
// extern int MSK_licensecleanup();
// extern int MSK_linkfiletoenvstream(MSKenv_t,int32_t,const char *,int32_t);
// extern int MSK_potrf(MSKenv_t,int32_t,int32_t,double *);
// extern int MSK_putlicensecode(MSKenv_t,int32_t *);
// extern int MSK_putlicensedebug(MSKenv_t,int32_t);
// extern int MSK_putlicensepath(MSKenv_t,const char *);
// extern int MSK_putlicensewait(MSKenv_t,int32_t);
// extern int MSK_resetexpirylicenses(MSKenv_t);
// extern int MSK_syeig(MSKenv_t,int32_t,int32_t,double *,double *);
// extern int MSK_syevd(MSKenv_t,int32_t,int32_t,double *,double *);
// extern int MSK_syrk(MSKenv_t,int32_t,int32_t,int32_t,int32_t,double,double *,double,double *);
import "C"

import (
    "unsafe"
    "fmt"
    "bytes"
)
// Constants
type Basindtype int32
const (
    MSK_BI_NEVER Basindtype = 0
    MSK_BI_ALWAYS Basindtype = 1
    MSK_BI_NO_ERROR Basindtype = 2
    MSK_BI_IF_FEASIBLE Basindtype = 3
    MSK_BI_RESERVERED Basindtype = 4
    MSK_BI_END Basindtype = 4
)
type Boundkey int32
const (
    MSK_BK_LO Boundkey = 0
    MSK_BK_UP Boundkey = 1
    MSK_BK_FX Boundkey = 2
    MSK_BK_FR Boundkey = 3
    MSK_BK_RA Boundkey = 4
    MSK_BK_END Boundkey = 4
)
type Mark int32
const (
    MSK_MARK_LO Mark = 0
    MSK_MARK_UP Mark = 1
    MSK_MARK_END Mark = 1
)
type Simdegen int32
const (
    MSK_SIM_DEGEN_NONE Simdegen = 0
    MSK_SIM_DEGEN_FREE Simdegen = 1
    MSK_SIM_DEGEN_AGGRESSIVE Simdegen = 2
    MSK_SIM_DEGEN_MODERATE Simdegen = 3
    MSK_SIM_DEGEN_MINIMUM Simdegen = 4
    MSK_SIM_DEGEN_END Simdegen = 4
)
type Transpose int32
const (
    MSK_TRANSPOSE_NO Transpose = 0
    MSK_TRANSPOSE_YES Transpose = 1
    MSK_TRANSPOSE_END Transpose = 1
)
type Uplo int32
const (
    MSK_UPLO_LO Uplo = 0
    MSK_UPLO_UP Uplo = 1
    MSK_UPLO_END Uplo = 1
)
type Simreform int32
const (
    MSK_SIM_REFORMULATION_OFF Simreform = 0
    MSK_SIM_REFORMULATION_ON Simreform = 1
    MSK_SIM_REFORMULATION_FREE Simreform = 2
    MSK_SIM_REFORMULATION_AGGRESSIVE Simreform = 3
    MSK_SIM_REFORMULATION_END Simreform = 3
)
type Simdupvec int32
const (
    MSK_SIM_EXPLOIT_DUPVEC_OFF Simdupvec = 0
    MSK_SIM_EXPLOIT_DUPVEC_ON Simdupvec = 1
    MSK_SIM_EXPLOIT_DUPVEC_FREE Simdupvec = 2
    MSK_SIM_EXPLOIT_DUPVEC_END Simdupvec = 2
)
type Simhotstart int32
const (
    MSK_SIM_HOTSTART_NONE Simhotstart = 0
    MSK_SIM_HOTSTART_FREE Simhotstart = 1
    MSK_SIM_HOTSTART_STATUS_KEYS Simhotstart = 2
    MSK_SIM_HOTSTART_END Simhotstart = 2
)
type Intpnthotstart int32
const (
    MSK_INTPNT_HOTSTART_NONE Intpnthotstart = 0
    MSK_INTPNT_HOTSTART_PRIMAL Intpnthotstart = 1
    MSK_INTPNT_HOTSTART_DUAL Intpnthotstart = 2
    MSK_INTPNT_HOTSTART_PRIMAL_DUAL Intpnthotstart = 3
    MSK_INTPNT_HOTSTART_END Intpnthotstart = 3
)
type Purify int32
const (
    MSK_PURIFY_NONE Purify = 0
    MSK_PURIFY_PRIMAL Purify = 1
    MSK_PURIFY_DUAL Purify = 2
    MSK_PURIFY_PRIMAL_DUAL Purify = 3
    MSK_PURIFY_AUTO Purify = 4
    MSK_PURIFY_END Purify = 4
)
type Callbackcode int32
const (
    MSK_CALLBACK_BEGIN_BI Callbackcode = 0
    MSK_CALLBACK_BEGIN_CONIC Callbackcode = 1
    MSK_CALLBACK_BEGIN_DUAL_BI Callbackcode = 2
    MSK_CALLBACK_BEGIN_DUAL_SENSITIVITY Callbackcode = 3
    MSK_CALLBACK_BEGIN_DUAL_SETUP_BI Callbackcode = 4
    MSK_CALLBACK_BEGIN_DUAL_SIMPLEX Callbackcode = 5
    MSK_CALLBACK_BEGIN_DUAL_SIMPLEX_BI Callbackcode = 6
    MSK_CALLBACK_BEGIN_INFEAS_ANA Callbackcode = 7
    MSK_CALLBACK_BEGIN_INTPNT Callbackcode = 8
    MSK_CALLBACK_BEGIN_LICENSE_WAIT Callbackcode = 9
    MSK_CALLBACK_BEGIN_MIO Callbackcode = 10
    MSK_CALLBACK_BEGIN_OPTIMIZER Callbackcode = 11
    MSK_CALLBACK_BEGIN_PRESOLVE Callbackcode = 12
    MSK_CALLBACK_BEGIN_PRIMAL_BI Callbackcode = 13
    MSK_CALLBACK_BEGIN_PRIMAL_REPAIR Callbackcode = 14
    MSK_CALLBACK_BEGIN_PRIMAL_SENSITIVITY Callbackcode = 15
    MSK_CALLBACK_BEGIN_PRIMAL_SETUP_BI Callbackcode = 16
    MSK_CALLBACK_BEGIN_PRIMAL_SIMPLEX Callbackcode = 17
    MSK_CALLBACK_BEGIN_PRIMAL_SIMPLEX_BI Callbackcode = 18
    MSK_CALLBACK_BEGIN_QCQO_REFORMULATE Callbackcode = 19
    MSK_CALLBACK_BEGIN_READ Callbackcode = 20
    MSK_CALLBACK_BEGIN_ROOT_CUTGEN Callbackcode = 21
    MSK_CALLBACK_BEGIN_SIMPLEX Callbackcode = 22
    MSK_CALLBACK_BEGIN_SIMPLEX_BI Callbackcode = 23
    MSK_CALLBACK_BEGIN_SOLVE_ROOT_RELAX Callbackcode = 24
    MSK_CALLBACK_BEGIN_TO_CONIC Callbackcode = 25
    MSK_CALLBACK_BEGIN_WRITE Callbackcode = 26
    MSK_CALLBACK_CONIC Callbackcode = 27
    MSK_CALLBACK_DUAL_SIMPLEX Callbackcode = 28
    MSK_CALLBACK_END_BI Callbackcode = 29
    MSK_CALLBACK_END_CONIC Callbackcode = 30
    MSK_CALLBACK_END_DUAL_BI Callbackcode = 31
    MSK_CALLBACK_END_DUAL_SENSITIVITY Callbackcode = 32
    MSK_CALLBACK_END_DUAL_SETUP_BI Callbackcode = 33
    MSK_CALLBACK_END_DUAL_SIMPLEX Callbackcode = 34
    MSK_CALLBACK_END_DUAL_SIMPLEX_BI Callbackcode = 35
    MSK_CALLBACK_END_INFEAS_ANA Callbackcode = 36
    MSK_CALLBACK_END_INTPNT Callbackcode = 37
    MSK_CALLBACK_END_LICENSE_WAIT Callbackcode = 38
    MSK_CALLBACK_END_MIO Callbackcode = 39
    MSK_CALLBACK_END_OPTIMIZER Callbackcode = 40
    MSK_CALLBACK_END_PRESOLVE Callbackcode = 41
    MSK_CALLBACK_END_PRIMAL_BI Callbackcode = 42
    MSK_CALLBACK_END_PRIMAL_REPAIR Callbackcode = 43
    MSK_CALLBACK_END_PRIMAL_SENSITIVITY Callbackcode = 44
    MSK_CALLBACK_END_PRIMAL_SETUP_BI Callbackcode = 45
    MSK_CALLBACK_END_PRIMAL_SIMPLEX Callbackcode = 46
    MSK_CALLBACK_END_PRIMAL_SIMPLEX_BI Callbackcode = 47
    MSK_CALLBACK_END_QCQO_REFORMULATE Callbackcode = 48
    MSK_CALLBACK_END_READ Callbackcode = 49
    MSK_CALLBACK_END_ROOT_CUTGEN Callbackcode = 50
    MSK_CALLBACK_END_SIMPLEX Callbackcode = 51
    MSK_CALLBACK_END_SIMPLEX_BI Callbackcode = 52
    MSK_CALLBACK_END_SOLVE_ROOT_RELAX Callbackcode = 53
    MSK_CALLBACK_END_TO_CONIC Callbackcode = 54
    MSK_CALLBACK_END_WRITE Callbackcode = 55
    MSK_CALLBACK_IM_BI Callbackcode = 56
    MSK_CALLBACK_IM_CONIC Callbackcode = 57
    MSK_CALLBACK_IM_DUAL_BI Callbackcode = 58
    MSK_CALLBACK_IM_DUAL_SENSIVITY Callbackcode = 59
    MSK_CALLBACK_IM_DUAL_SIMPLEX Callbackcode = 60
    MSK_CALLBACK_IM_INTPNT Callbackcode = 61
    MSK_CALLBACK_IM_LICENSE_WAIT Callbackcode = 62
    MSK_CALLBACK_IM_LU Callbackcode = 63
    MSK_CALLBACK_IM_MIO Callbackcode = 64
    MSK_CALLBACK_IM_MIO_DUAL_SIMPLEX Callbackcode = 65
    MSK_CALLBACK_IM_MIO_INTPNT Callbackcode = 66
    MSK_CALLBACK_IM_MIO_PRIMAL_SIMPLEX Callbackcode = 67
    MSK_CALLBACK_IM_ORDER Callbackcode = 68
    MSK_CALLBACK_IM_PRESOLVE Callbackcode = 69
    MSK_CALLBACK_IM_PRIMAL_BI Callbackcode = 70
    MSK_CALLBACK_IM_PRIMAL_SENSIVITY Callbackcode = 71
    MSK_CALLBACK_IM_PRIMAL_SIMPLEX Callbackcode = 72
    MSK_CALLBACK_IM_QO_REFORMULATE Callbackcode = 73
    MSK_CALLBACK_IM_READ Callbackcode = 74
    MSK_CALLBACK_IM_ROOT_CUTGEN Callbackcode = 75
    MSK_CALLBACK_IM_SIMPLEX Callbackcode = 76
    MSK_CALLBACK_IM_SIMPLEX_BI Callbackcode = 77
    MSK_CALLBACK_INTPNT Callbackcode = 78
    MSK_CALLBACK_NEW_INT_MIO Callbackcode = 79
    MSK_CALLBACK_PRIMAL_SIMPLEX Callbackcode = 80
    MSK_CALLBACK_READ_OPF Callbackcode = 81
    MSK_CALLBACK_READ_OPF_SECTION Callbackcode = 82
    MSK_CALLBACK_RESTART_MIO Callbackcode = 83
    MSK_CALLBACK_SOLVING_REMOTE Callbackcode = 84
    MSK_CALLBACK_UPDATE_DUAL_BI Callbackcode = 85
    MSK_CALLBACK_UPDATE_DUAL_SIMPLEX Callbackcode = 86
    MSK_CALLBACK_UPDATE_DUAL_SIMPLEX_BI Callbackcode = 87
    MSK_CALLBACK_UPDATE_PRESOLVE Callbackcode = 88
    MSK_CALLBACK_UPDATE_PRIMAL_BI Callbackcode = 89
    MSK_CALLBACK_UPDATE_PRIMAL_SIMPLEX Callbackcode = 90
    MSK_CALLBACK_UPDATE_PRIMAL_SIMPLEX_BI Callbackcode = 91
    MSK_CALLBACK_UPDATE_SIMPLEX Callbackcode = 92
    MSK_CALLBACK_WRITE_OPF Callbackcode = 93
    MSK_CALLBACK_END Callbackcode = 93
)
type Compresstype int32
const (
    MSK_COMPRESS_NONE Compresstype = 0
    MSK_COMPRESS_FREE Compresstype = 1
    MSK_COMPRESS_GZIP Compresstype = 2
    MSK_COMPRESS_ZSTD Compresstype = 3
    MSK_COMPRESS_END Compresstype = 3
)
type Conetype int32
const (
    MSK_CT_QUAD Conetype = 0
    MSK_CT_RQUAD Conetype = 1
    MSK_CT_PEXP Conetype = 2
    MSK_CT_DEXP Conetype = 3
    MSK_CT_PPOW Conetype = 4
    MSK_CT_DPOW Conetype = 5
    MSK_CT_ZERO Conetype = 6
    MSK_CT_END Conetype = 6
)
type Domaintype int32
const (
    MSK_DOMAIN_R Domaintype = 0
    MSK_DOMAIN_RZERO Domaintype = 1
    MSK_DOMAIN_RPLUS Domaintype = 2
    MSK_DOMAIN_RMINUS Domaintype = 3
    MSK_DOMAIN_QUADRATIC_CONE Domaintype = 4
    MSK_DOMAIN_RQUADRATIC_CONE Domaintype = 5
    MSK_DOMAIN_PRIMAL_EXP_CONE Domaintype = 6
    MSK_DOMAIN_DUAL_EXP_CONE Domaintype = 7
    MSK_DOMAIN_PRIMAL_POWER_CONE Domaintype = 8
    MSK_DOMAIN_DUAL_POWER_CONE Domaintype = 9
    MSK_DOMAIN_PRIMAL_GEO_MEAN_CONE Domaintype = 10
    MSK_DOMAIN_DUAL_GEO_MEAN_CONE Domaintype = 11
    MSK_DOMAIN_SVEC_PSD_CONE Domaintype = 12
    MSK_DOMAIN_END Domaintype = 12
)
type Nametype int32
const (
    MSK_NAME_TYPE_GEN Nametype = 0
    MSK_NAME_TYPE_MPS Nametype = 1
    MSK_NAME_TYPE_LP Nametype = 2
    MSK_NAME_TYPE_END Nametype = 2
)
type Symmattype int32
const (
    MSK_SYMMAT_TYPE_SPARSE Symmattype = 0
    MSK_SYMMAT_TYPE_END Symmattype = 0
)
type Dataformat int32
const (
    MSK_DATA_FORMAT_EXTENSION Dataformat = 0
    MSK_DATA_FORMAT_MPS Dataformat = 1
    MSK_DATA_FORMAT_LP Dataformat = 2
    MSK_DATA_FORMAT_OP Dataformat = 3
    MSK_DATA_FORMAT_FREE_MPS Dataformat = 4
    MSK_DATA_FORMAT_TASK Dataformat = 5
    MSK_DATA_FORMAT_PTF Dataformat = 6
    MSK_DATA_FORMAT_CB Dataformat = 7
    MSK_DATA_FORMAT_JSON_TASK Dataformat = 8
    MSK_DATA_FORMAT_END Dataformat = 8
)
type Solformat int32
const (
    MSK_SOL_FORMAT_EXTENSION Solformat = 0
    MSK_SOL_FORMAT_B Solformat = 1
    MSK_SOL_FORMAT_TASK Solformat = 2
    MSK_SOL_FORMAT_JSON_TASK Solformat = 3
    MSK_SOL_FORMAT_END Solformat = 3
)
type Dinfitem int32
const (
    MSK_DINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_DENSITY Dinfitem = 0
    MSK_DINF_BI_CLEAN_DUAL_TIME Dinfitem = 1
    MSK_DINF_BI_CLEAN_PRIMAL_TIME Dinfitem = 2
    MSK_DINF_BI_CLEAN_TIME Dinfitem = 3
    MSK_DINF_BI_DUAL_TIME Dinfitem = 4
    MSK_DINF_BI_PRIMAL_TIME Dinfitem = 5
    MSK_DINF_BI_TIME Dinfitem = 6
    MSK_DINF_INTPNT_DUAL_FEAS Dinfitem = 7
    MSK_DINF_INTPNT_DUAL_OBJ Dinfitem = 8
    MSK_DINF_INTPNT_FACTOR_NUM_FLOPS Dinfitem = 9
    MSK_DINF_INTPNT_OPT_STATUS Dinfitem = 10
    MSK_DINF_INTPNT_ORDER_TIME Dinfitem = 11
    MSK_DINF_INTPNT_PRIMAL_FEAS Dinfitem = 12
    MSK_DINF_INTPNT_PRIMAL_OBJ Dinfitem = 13
    MSK_DINF_INTPNT_TIME Dinfitem = 14
    MSK_DINF_MIO_CLIQUE_SELECTION_TIME Dinfitem = 15
    MSK_DINF_MIO_CLIQUE_SEPARATION_TIME Dinfitem = 16
    MSK_DINF_MIO_CMIR_SELECTION_TIME Dinfitem = 17
    MSK_DINF_MIO_CMIR_SEPARATION_TIME Dinfitem = 18
    MSK_DINF_MIO_CONSTRUCT_SOLUTION_OBJ Dinfitem = 19
    MSK_DINF_MIO_DUAL_BOUND_AFTER_PRESOLVE Dinfitem = 20
    MSK_DINF_MIO_GMI_SELECTION_TIME Dinfitem = 21
    MSK_DINF_MIO_GMI_SEPARATION_TIME Dinfitem = 22
    MSK_DINF_MIO_IMPLIED_BOUND_SELECTION_TIME Dinfitem = 23
    MSK_DINF_MIO_IMPLIED_BOUND_SEPARATION_TIME Dinfitem = 24
    MSK_DINF_MIO_INITIAL_FEASIBLE_SOLUTION_OBJ Dinfitem = 25
    MSK_DINF_MIO_KNAPSACK_COVER_SELECTION_TIME Dinfitem = 26
    MSK_DINF_MIO_KNAPSACK_COVER_SEPARATION_TIME Dinfitem = 27
    MSK_DINF_MIO_LIPRO_SELECTION_TIME Dinfitem = 28
    MSK_DINF_MIO_LIPRO_SEPARATION_TIME Dinfitem = 29
    MSK_DINF_MIO_OBJ_ABS_GAP Dinfitem = 30
    MSK_DINF_MIO_OBJ_BOUND Dinfitem = 31
    MSK_DINF_MIO_OBJ_INT Dinfitem = 32
    MSK_DINF_MIO_OBJ_REL_GAP Dinfitem = 33
    MSK_DINF_MIO_PROBING_TIME Dinfitem = 34
    MSK_DINF_MIO_ROOT_CUT_SELECTION_TIME Dinfitem = 35
    MSK_DINF_MIO_ROOT_CUT_SEPARATION_TIME Dinfitem = 36
    MSK_DINF_MIO_ROOT_OPTIMIZER_TIME Dinfitem = 37
    MSK_DINF_MIO_ROOT_PRESOLVE_TIME Dinfitem = 38
    MSK_DINF_MIO_ROOT_TIME Dinfitem = 39
    MSK_DINF_MIO_SYMMETRY_DETECTION_TIME Dinfitem = 40
    MSK_DINF_MIO_SYMMETRY_FACTOR Dinfitem = 41
    MSK_DINF_MIO_TIME Dinfitem = 42
    MSK_DINF_MIO_USER_OBJ_CUT Dinfitem = 43
    MSK_DINF_OPTIMIZER_TICKS Dinfitem = 44
    MSK_DINF_OPTIMIZER_TIME Dinfitem = 45
    MSK_DINF_PRESOLVE_ELI_TIME Dinfitem = 46
    MSK_DINF_PRESOLVE_LINDEP_TIME Dinfitem = 47
    MSK_DINF_PRESOLVE_TIME Dinfitem = 48
    MSK_DINF_PRESOLVE_TOTAL_PRIMAL_PERTURBATION Dinfitem = 49
    MSK_DINF_PRIMAL_REPAIR_PENALTY_OBJ Dinfitem = 50
    MSK_DINF_QCQO_REFORMULATE_MAX_PERTURBATION Dinfitem = 51
    MSK_DINF_QCQO_REFORMULATE_TIME Dinfitem = 52
    MSK_DINF_QCQO_REFORMULATE_WORST_CHOLESKY_COLUMN_SCALING Dinfitem = 53
    MSK_DINF_QCQO_REFORMULATE_WORST_CHOLESKY_DIAG_SCALING Dinfitem = 54
    MSK_DINF_READ_DATA_TIME Dinfitem = 55
    MSK_DINF_REMOTE_TIME Dinfitem = 56
    MSK_DINF_SIM_DUAL_TIME Dinfitem = 57
    MSK_DINF_SIM_FEAS Dinfitem = 58
    MSK_DINF_SIM_OBJ Dinfitem = 59
    MSK_DINF_SIM_PRIMAL_TIME Dinfitem = 60
    MSK_DINF_SIM_TIME Dinfitem = 61
    MSK_DINF_SOL_BAS_DUAL_OBJ Dinfitem = 62
    MSK_DINF_SOL_BAS_DVIOLCON Dinfitem = 63
    MSK_DINF_SOL_BAS_DVIOLVAR Dinfitem = 64
    MSK_DINF_SOL_BAS_NRM_BARX Dinfitem = 65
    MSK_DINF_SOL_BAS_NRM_SLC Dinfitem = 66
    MSK_DINF_SOL_BAS_NRM_SLX Dinfitem = 67
    MSK_DINF_SOL_BAS_NRM_SUC Dinfitem = 68
    MSK_DINF_SOL_BAS_NRM_SUX Dinfitem = 69
    MSK_DINF_SOL_BAS_NRM_XC Dinfitem = 70
    MSK_DINF_SOL_BAS_NRM_XX Dinfitem = 71
    MSK_DINF_SOL_BAS_NRM_Y Dinfitem = 72
    MSK_DINF_SOL_BAS_PRIMAL_OBJ Dinfitem = 73
    MSK_DINF_SOL_BAS_PVIOLCON Dinfitem = 74
    MSK_DINF_SOL_BAS_PVIOLVAR Dinfitem = 75
    MSK_DINF_SOL_ITG_NRM_BARX Dinfitem = 76
    MSK_DINF_SOL_ITG_NRM_XC Dinfitem = 77
    MSK_DINF_SOL_ITG_NRM_XX Dinfitem = 78
    MSK_DINF_SOL_ITG_PRIMAL_OBJ Dinfitem = 79
    MSK_DINF_SOL_ITG_PVIOLACC Dinfitem = 80
    MSK_DINF_SOL_ITG_PVIOLBARVAR Dinfitem = 81
    MSK_DINF_SOL_ITG_PVIOLCON Dinfitem = 82
    MSK_DINF_SOL_ITG_PVIOLCONES Dinfitem = 83
    MSK_DINF_SOL_ITG_PVIOLDJC Dinfitem = 84
    MSK_DINF_SOL_ITG_PVIOLITG Dinfitem = 85
    MSK_DINF_SOL_ITG_PVIOLVAR Dinfitem = 86
    MSK_DINF_SOL_ITR_DUAL_OBJ Dinfitem = 87
    MSK_DINF_SOL_ITR_DVIOLACC Dinfitem = 88
    MSK_DINF_SOL_ITR_DVIOLBARVAR Dinfitem = 89
    MSK_DINF_SOL_ITR_DVIOLCON Dinfitem = 90
    MSK_DINF_SOL_ITR_DVIOLCONES Dinfitem = 91
    MSK_DINF_SOL_ITR_DVIOLVAR Dinfitem = 92
    MSK_DINF_SOL_ITR_NRM_BARS Dinfitem = 93
    MSK_DINF_SOL_ITR_NRM_BARX Dinfitem = 94
    MSK_DINF_SOL_ITR_NRM_SLC Dinfitem = 95
    MSK_DINF_SOL_ITR_NRM_SLX Dinfitem = 96
    MSK_DINF_SOL_ITR_NRM_SNX Dinfitem = 97
    MSK_DINF_SOL_ITR_NRM_SUC Dinfitem = 98
    MSK_DINF_SOL_ITR_NRM_SUX Dinfitem = 99
    MSK_DINF_SOL_ITR_NRM_XC Dinfitem = 100
    MSK_DINF_SOL_ITR_NRM_XX Dinfitem = 101
    MSK_DINF_SOL_ITR_NRM_Y Dinfitem = 102
    MSK_DINF_SOL_ITR_PRIMAL_OBJ Dinfitem = 103
    MSK_DINF_SOL_ITR_PVIOLACC Dinfitem = 104
    MSK_DINF_SOL_ITR_PVIOLBARVAR Dinfitem = 105
    MSK_DINF_SOL_ITR_PVIOLCON Dinfitem = 106
    MSK_DINF_SOL_ITR_PVIOLCONES Dinfitem = 107
    MSK_DINF_SOL_ITR_PVIOLVAR Dinfitem = 108
    MSK_DINF_TO_CONIC_TIME Dinfitem = 109
    MSK_DINF_WRITE_DATA_TIME Dinfitem = 110
    MSK_DINF_END Dinfitem = 110
)
type Feature int32
const (
    MSK_FEATURE_PTS Feature = 0
    MSK_FEATURE_PTON Feature = 1
    MSK_FEATURE_END Feature = 1
)
type Dparam int32
const (
    MSK_DPAR_ANA_SOL_INFEAS_TOL Dparam = 0
    MSK_DPAR_BASIS_REL_TOL_S Dparam = 1
    MSK_DPAR_BASIS_TOL_S Dparam = 2
    MSK_DPAR_BASIS_TOL_X Dparam = 3
    MSK_DPAR_CHECK_CONVEXITY_REL_TOL Dparam = 4
    MSK_DPAR_DATA_SYM_MAT_TOL Dparam = 5
    MSK_DPAR_DATA_SYM_MAT_TOL_HUGE Dparam = 6
    MSK_DPAR_DATA_SYM_MAT_TOL_LARGE Dparam = 7
    MSK_DPAR_DATA_TOL_AIJ_HUGE Dparam = 8
    MSK_DPAR_DATA_TOL_AIJ_LARGE Dparam = 9
    MSK_DPAR_DATA_TOL_BOUND_INF Dparam = 10
    MSK_DPAR_DATA_TOL_BOUND_WRN Dparam = 11
    MSK_DPAR_DATA_TOL_C_HUGE Dparam = 12
    MSK_DPAR_DATA_TOL_CJ_LARGE Dparam = 13
    MSK_DPAR_DATA_TOL_QIJ Dparam = 14
    MSK_DPAR_DATA_TOL_X Dparam = 15
    MSK_DPAR_INTPNT_CO_TOL_DFEAS Dparam = 16
    MSK_DPAR_INTPNT_CO_TOL_INFEAS Dparam = 17
    MSK_DPAR_INTPNT_CO_TOL_MU_RED Dparam = 18
    MSK_DPAR_INTPNT_CO_TOL_NEAR_REL Dparam = 19
    MSK_DPAR_INTPNT_CO_TOL_PFEAS Dparam = 20
    MSK_DPAR_INTPNT_CO_TOL_REL_GAP Dparam = 21
    MSK_DPAR_INTPNT_QO_TOL_DFEAS Dparam = 22
    MSK_DPAR_INTPNT_QO_TOL_INFEAS Dparam = 23
    MSK_DPAR_INTPNT_QO_TOL_MU_RED Dparam = 24
    MSK_DPAR_INTPNT_QO_TOL_NEAR_REL Dparam = 25
    MSK_DPAR_INTPNT_QO_TOL_PFEAS Dparam = 26
    MSK_DPAR_INTPNT_QO_TOL_REL_GAP Dparam = 27
    MSK_DPAR_INTPNT_TOL_DFEAS Dparam = 28
    MSK_DPAR_INTPNT_TOL_DSAFE Dparam = 29
    MSK_DPAR_INTPNT_TOL_INFEAS Dparam = 30
    MSK_DPAR_INTPNT_TOL_MU_RED Dparam = 31
    MSK_DPAR_INTPNT_TOL_PATH Dparam = 32
    MSK_DPAR_INTPNT_TOL_PFEAS Dparam = 33
    MSK_DPAR_INTPNT_TOL_PSAFE Dparam = 34
    MSK_DPAR_INTPNT_TOL_REL_GAP Dparam = 35
    MSK_DPAR_INTPNT_TOL_REL_STEP Dparam = 36
    MSK_DPAR_INTPNT_TOL_STEP_SIZE Dparam = 37
    MSK_DPAR_LOWER_OBJ_CUT Dparam = 38
    MSK_DPAR_LOWER_OBJ_CUT_FINITE_TRH Dparam = 39
    MSK_DPAR_MIO_DJC_MAX_BIGM Dparam = 40
    MSK_DPAR_MIO_MAX_TIME Dparam = 41
    MSK_DPAR_MIO_REL_GAP_CONST Dparam = 42
    MSK_DPAR_MIO_TOL_ABS_GAP Dparam = 43
    MSK_DPAR_MIO_TOL_ABS_RELAX_INT Dparam = 44
    MSK_DPAR_MIO_TOL_FEAS Dparam = 45
    MSK_DPAR_MIO_TOL_REL_DUAL_BOUND_IMPROVEMENT Dparam = 46
    MSK_DPAR_MIO_TOL_REL_GAP Dparam = 47
    MSK_DPAR_OPTIMIZER_MAX_TICKS Dparam = 48
    MSK_DPAR_OPTIMIZER_MAX_TIME Dparam = 49
    MSK_DPAR_PRESOLVE_TOL_ABS_LINDEP Dparam = 50
    MSK_DPAR_PRESOLVE_TOL_AIJ Dparam = 51
    MSK_DPAR_PRESOLVE_TOL_PRIMAL_INFEAS_PERTURBATION Dparam = 52
    MSK_DPAR_PRESOLVE_TOL_REL_LINDEP Dparam = 53
    MSK_DPAR_PRESOLVE_TOL_S Dparam = 54
    MSK_DPAR_PRESOLVE_TOL_X Dparam = 55
    MSK_DPAR_QCQO_REFORMULATE_REL_DROP_TOL Dparam = 56
    MSK_DPAR_SEMIDEFINITE_TOL_APPROX Dparam = 57
    MSK_DPAR_SIM_LU_TOL_REL_PIV Dparam = 58
    MSK_DPAR_SIMPLEX_ABS_TOL_PIV Dparam = 59
    MSK_DPAR_UPPER_OBJ_CUT Dparam = 60
    MSK_DPAR_UPPER_OBJ_CUT_FINITE_TRH Dparam = 61
    MSK_DPAR_END Dparam = 61
)
type Liinfitem int32
const (
    MSK_LIINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_NUM_COLUMNS Liinfitem = 0
    MSK_LIINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_NUM_NZ Liinfitem = 1
    MSK_LIINF_ANA_PRO_SCALARIZED_CONSTRAINT_MATRIX_NUM_ROWS Liinfitem = 2
    MSK_LIINF_BI_CLEAN_DUAL_DEG_ITER Liinfitem = 3
    MSK_LIINF_BI_CLEAN_DUAL_ITER Liinfitem = 4
    MSK_LIINF_BI_CLEAN_PRIMAL_DEG_ITER Liinfitem = 5
    MSK_LIINF_BI_CLEAN_PRIMAL_ITER Liinfitem = 6
    MSK_LIINF_BI_DUAL_ITER Liinfitem = 7
    MSK_LIINF_BI_PRIMAL_ITER Liinfitem = 8
    MSK_LIINF_INTPNT_FACTOR_NUM_NZ Liinfitem = 9
    MSK_LIINF_MIO_ANZ Liinfitem = 10
    MSK_LIINF_MIO_INTPNT_ITER Liinfitem = 11
    MSK_LIINF_MIO_NUM_DUAL_ILLPOSED_CER Liinfitem = 12
    MSK_LIINF_MIO_NUM_PRIM_ILLPOSED_CER Liinfitem = 13
    MSK_LIINF_MIO_PRESOLVED_ANZ Liinfitem = 14
    MSK_LIINF_MIO_SIMPLEX_ITER Liinfitem = 15
    MSK_LIINF_RD_NUMACC Liinfitem = 16
    MSK_LIINF_RD_NUMANZ Liinfitem = 17
    MSK_LIINF_RD_NUMDJC Liinfitem = 18
    MSK_LIINF_RD_NUMQNZ Liinfitem = 19
    MSK_LIINF_SIMPLEX_ITER Liinfitem = 20
    MSK_LIINF_END Liinfitem = 20
)
type Iinfitem int32
const (
    MSK_IINF_ANA_PRO_NUM_CON Iinfitem = 0
    MSK_IINF_ANA_PRO_NUM_CON_EQ Iinfitem = 1
    MSK_IINF_ANA_PRO_NUM_CON_FR Iinfitem = 2
    MSK_IINF_ANA_PRO_NUM_CON_LO Iinfitem = 3
    MSK_IINF_ANA_PRO_NUM_CON_RA Iinfitem = 4
    MSK_IINF_ANA_PRO_NUM_CON_UP Iinfitem = 5
    MSK_IINF_ANA_PRO_NUM_VAR Iinfitem = 6
    MSK_IINF_ANA_PRO_NUM_VAR_BIN Iinfitem = 7
    MSK_IINF_ANA_PRO_NUM_VAR_CONT Iinfitem = 8
    MSK_IINF_ANA_PRO_NUM_VAR_EQ Iinfitem = 9
    MSK_IINF_ANA_PRO_NUM_VAR_FR Iinfitem = 10
    MSK_IINF_ANA_PRO_NUM_VAR_INT Iinfitem = 11
    MSK_IINF_ANA_PRO_NUM_VAR_LO Iinfitem = 12
    MSK_IINF_ANA_PRO_NUM_VAR_RA Iinfitem = 13
    MSK_IINF_ANA_PRO_NUM_VAR_UP Iinfitem = 14
    MSK_IINF_INTPNT_FACTOR_DIM_DENSE Iinfitem = 15
    MSK_IINF_INTPNT_ITER Iinfitem = 16
    MSK_IINF_INTPNT_NUM_THREADS Iinfitem = 17
    MSK_IINF_INTPNT_SOLVE_DUAL Iinfitem = 18
    MSK_IINF_MIO_ABSGAP_SATISFIED Iinfitem = 19
    MSK_IINF_MIO_CLIQUE_TABLE_SIZE Iinfitem = 20
    MSK_IINF_MIO_CONSTRUCT_SOLUTION Iinfitem = 21
    MSK_IINF_MIO_INITIAL_FEASIBLE_SOLUTION Iinfitem = 22
    MSK_IINF_MIO_NODE_DEPTH Iinfitem = 23
    MSK_IINF_MIO_NUM_ACTIVE_NODES Iinfitem = 24
    MSK_IINF_MIO_NUM_ACTIVE_ROOT_CUTS Iinfitem = 25
    MSK_IINF_MIO_NUM_BRANCH Iinfitem = 26
    MSK_IINF_MIO_NUM_INT_SOLUTIONS Iinfitem = 27
    MSK_IINF_MIO_NUM_RELAX Iinfitem = 28
    MSK_IINF_MIO_NUM_REPEATED_PRESOLVE Iinfitem = 29
    MSK_IINF_MIO_NUM_RESTARTS Iinfitem = 30
    MSK_IINF_MIO_NUM_ROOT_CUT_ROUNDS Iinfitem = 31
    MSK_IINF_MIO_NUM_SELECTED_CLIQUE_CUTS Iinfitem = 32
    MSK_IINF_MIO_NUM_SELECTED_CMIR_CUTS Iinfitem = 33
    MSK_IINF_MIO_NUM_SELECTED_GOMORY_CUTS Iinfitem = 34
    MSK_IINF_MIO_NUM_SELECTED_IMPLIED_BOUND_CUTS Iinfitem = 35
    MSK_IINF_MIO_NUM_SELECTED_KNAPSACK_COVER_CUTS Iinfitem = 36
    MSK_IINF_MIO_NUM_SELECTED_LIPRO_CUTS Iinfitem = 37
    MSK_IINF_MIO_NUM_SEPARATED_CLIQUE_CUTS Iinfitem = 38
    MSK_IINF_MIO_NUM_SEPARATED_CMIR_CUTS Iinfitem = 39
    MSK_IINF_MIO_NUM_SEPARATED_GOMORY_CUTS Iinfitem = 40
    MSK_IINF_MIO_NUM_SEPARATED_IMPLIED_BOUND_CUTS Iinfitem = 41
    MSK_IINF_MIO_NUM_SEPARATED_KNAPSACK_COVER_CUTS Iinfitem = 42
    MSK_IINF_MIO_NUM_SEPARATED_LIPRO_CUTS Iinfitem = 43
    MSK_IINF_MIO_NUM_SOLVED_NODES Iinfitem = 44
    MSK_IINF_MIO_NUMBIN Iinfitem = 45
    MSK_IINF_MIO_NUMBINCONEVAR Iinfitem = 46
    MSK_IINF_MIO_NUMCON Iinfitem = 47
    MSK_IINF_MIO_NUMCONE Iinfitem = 48
    MSK_IINF_MIO_NUMCONEVAR Iinfitem = 49
    MSK_IINF_MIO_NUMCONT Iinfitem = 50
    MSK_IINF_MIO_NUMCONTCONEVAR Iinfitem = 51
    MSK_IINF_MIO_NUMDEXPCONES Iinfitem = 52
    MSK_IINF_MIO_NUMDJC Iinfitem = 53
    MSK_IINF_MIO_NUMDPOWCONES Iinfitem = 54
    MSK_IINF_MIO_NUMINT Iinfitem = 55
    MSK_IINF_MIO_NUMINTCONEVAR Iinfitem = 56
    MSK_IINF_MIO_NUMPEXPCONES Iinfitem = 57
    MSK_IINF_MIO_NUMPPOWCONES Iinfitem = 58
    MSK_IINF_MIO_NUMQCONES Iinfitem = 59
    MSK_IINF_MIO_NUMRQCONES Iinfitem = 60
    MSK_IINF_MIO_NUMVAR Iinfitem = 61
    MSK_IINF_MIO_OBJ_BOUND_DEFINED Iinfitem = 62
    MSK_IINF_MIO_PRESOLVED_NUMBIN Iinfitem = 63
    MSK_IINF_MIO_PRESOLVED_NUMBINCONEVAR Iinfitem = 64
    MSK_IINF_MIO_PRESOLVED_NUMCON Iinfitem = 65
    MSK_IINF_MIO_PRESOLVED_NUMCONE Iinfitem = 66
    MSK_IINF_MIO_PRESOLVED_NUMCONEVAR Iinfitem = 67
    MSK_IINF_MIO_PRESOLVED_NUMCONT Iinfitem = 68
    MSK_IINF_MIO_PRESOLVED_NUMCONTCONEVAR Iinfitem = 69
    MSK_IINF_MIO_PRESOLVED_NUMDEXPCONES Iinfitem = 70
    MSK_IINF_MIO_PRESOLVED_NUMDJC Iinfitem = 71
    MSK_IINF_MIO_PRESOLVED_NUMDPOWCONES Iinfitem = 72
    MSK_IINF_MIO_PRESOLVED_NUMINT Iinfitem = 73
    MSK_IINF_MIO_PRESOLVED_NUMINTCONEVAR Iinfitem = 74
    MSK_IINF_MIO_PRESOLVED_NUMPEXPCONES Iinfitem = 75
    MSK_IINF_MIO_PRESOLVED_NUMPPOWCONES Iinfitem = 76
    MSK_IINF_MIO_PRESOLVED_NUMQCONES Iinfitem = 77
    MSK_IINF_MIO_PRESOLVED_NUMRQCONES Iinfitem = 78
    MSK_IINF_MIO_PRESOLVED_NUMVAR Iinfitem = 79
    MSK_IINF_MIO_RELGAP_SATISFIED Iinfitem = 80
    MSK_IINF_MIO_TOTAL_NUM_SELECTED_CUTS Iinfitem = 81
    MSK_IINF_MIO_TOTAL_NUM_SEPARATED_CUTS Iinfitem = 82
    MSK_IINF_MIO_USER_OBJ_CUT Iinfitem = 83
    MSK_IINF_OPT_NUMCON Iinfitem = 84
    MSK_IINF_OPT_NUMVAR Iinfitem = 85
    MSK_IINF_OPTIMIZE_RESPONSE Iinfitem = 86
    MSK_IINF_PRESOLVE_NUM_PRIMAL_PERTURBATIONS Iinfitem = 87
    MSK_IINF_PURIFY_DUAL_SUCCESS Iinfitem = 88
    MSK_IINF_PURIFY_PRIMAL_SUCCESS Iinfitem = 89
    MSK_IINF_RD_NUMBARVAR Iinfitem = 90
    MSK_IINF_RD_NUMCON Iinfitem = 91
    MSK_IINF_RD_NUMCONE Iinfitem = 92
    MSK_IINF_RD_NUMINTVAR Iinfitem = 93
    MSK_IINF_RD_NUMQ Iinfitem = 94
    MSK_IINF_RD_NUMVAR Iinfitem = 95
    MSK_IINF_RD_PROTYPE Iinfitem = 96
    MSK_IINF_SIM_DUAL_DEG_ITER Iinfitem = 97
    MSK_IINF_SIM_DUAL_HOTSTART Iinfitem = 98
    MSK_IINF_SIM_DUAL_HOTSTART_LU Iinfitem = 99
    MSK_IINF_SIM_DUAL_INF_ITER Iinfitem = 100
    MSK_IINF_SIM_DUAL_ITER Iinfitem = 101
    MSK_IINF_SIM_NUMCON Iinfitem = 102
    MSK_IINF_SIM_NUMVAR Iinfitem = 103
    MSK_IINF_SIM_PRIMAL_DEG_ITER Iinfitem = 104
    MSK_IINF_SIM_PRIMAL_HOTSTART Iinfitem = 105
    MSK_IINF_SIM_PRIMAL_HOTSTART_LU Iinfitem = 106
    MSK_IINF_SIM_PRIMAL_INF_ITER Iinfitem = 107
    MSK_IINF_SIM_PRIMAL_ITER Iinfitem = 108
    MSK_IINF_SIM_SOLVE_DUAL Iinfitem = 109
    MSK_IINF_SOL_BAS_PROSTA Iinfitem = 110
    MSK_IINF_SOL_BAS_SOLSTA Iinfitem = 111
    MSK_IINF_SOL_ITG_PROSTA Iinfitem = 112
    MSK_IINF_SOL_ITG_SOLSTA Iinfitem = 113
    MSK_IINF_SOL_ITR_PROSTA Iinfitem = 114
    MSK_IINF_SOL_ITR_SOLSTA Iinfitem = 115
    MSK_IINF_STO_NUM_A_REALLOC Iinfitem = 116
    MSK_IINF_END Iinfitem = 116
)
type Inftype int32
const (
    MSK_INF_DOU_TYPE Inftype = 0
    MSK_INF_INT_TYPE Inftype = 1
    MSK_INF_LINT_TYPE Inftype = 2
    MSK_INF_END Inftype = 2
)
type Iomode int32
const (
    MSK_IOMODE_READ Iomode = 0
    MSK_IOMODE_WRITE Iomode = 1
    MSK_IOMODE_READWRITE Iomode = 2
    MSK_IOMODE_END Iomode = 2
)
type Iparam int32
const (
    MSK_IPAR_ANA_SOL_BASIS Iparam = 0
    MSK_IPAR_ANA_SOL_PRINT_VIOLATED Iparam = 1
    MSK_IPAR_AUTO_SORT_A_BEFORE_OPT Iparam = 2
    MSK_IPAR_AUTO_UPDATE_SOL_INFO Iparam = 3
    MSK_IPAR_BASIS_SOLVE_USE_PLUS_ONE Iparam = 4
    MSK_IPAR_BI_CLEAN_OPTIMIZER Iparam = 5
    MSK_IPAR_BI_IGNORE_MAX_ITER Iparam = 6
    MSK_IPAR_BI_IGNORE_NUM_ERROR Iparam = 7
    MSK_IPAR_BI_MAX_ITERATIONS Iparam = 8
    MSK_IPAR_CACHE_LICENSE Iparam = 9
    MSK_IPAR_COMPRESS_STATFILE Iparam = 10
    MSK_IPAR_INFEAS_GENERIC_NAMES Iparam = 11
    MSK_IPAR_INFEAS_PREFER_PRIMAL Iparam = 12
    MSK_IPAR_INFEAS_REPORT_AUTO Iparam = 13
    MSK_IPAR_INFEAS_REPORT_LEVEL Iparam = 14
    MSK_IPAR_INTPNT_BASIS Iparam = 15
    MSK_IPAR_INTPNT_DIFF_STEP Iparam = 16
    MSK_IPAR_INTPNT_HOTSTART Iparam = 17
    MSK_IPAR_INTPNT_MAX_ITERATIONS Iparam = 18
    MSK_IPAR_INTPNT_MAX_NUM_COR Iparam = 19
    MSK_IPAR_INTPNT_MAX_NUM_REFINEMENT_STEPS Iparam = 20
    MSK_IPAR_INTPNT_OFF_COL_TRH Iparam = 21
    MSK_IPAR_INTPNT_ORDER_GP_NUM_SEEDS Iparam = 22
    MSK_IPAR_INTPNT_ORDER_METHOD Iparam = 23
    MSK_IPAR_INTPNT_PURIFY Iparam = 24
    MSK_IPAR_INTPNT_REGULARIZATION_USE Iparam = 25
    MSK_IPAR_INTPNT_SCALING Iparam = 26
    MSK_IPAR_INTPNT_SOLVE_FORM Iparam = 27
    MSK_IPAR_INTPNT_STARTING_POINT Iparam = 28
    MSK_IPAR_LICENSE_DEBUG Iparam = 29
    MSK_IPAR_LICENSE_PAUSE_TIME Iparam = 30
    MSK_IPAR_LICENSE_SUPPRESS_EXPIRE_WRNS Iparam = 31
    MSK_IPAR_LICENSE_TRH_EXPIRY_WRN Iparam = 32
    MSK_IPAR_LICENSE_WAIT Iparam = 33
    MSK_IPAR_LOG Iparam = 34
    MSK_IPAR_LOG_ANA_PRO Iparam = 35
    MSK_IPAR_LOG_BI Iparam = 36
    MSK_IPAR_LOG_BI_FREQ Iparam = 37
    MSK_IPAR_LOG_CUT_SECOND_OPT Iparam = 38
    MSK_IPAR_LOG_EXPAND Iparam = 39
    MSK_IPAR_LOG_FEAS_REPAIR Iparam = 40
    MSK_IPAR_LOG_FILE Iparam = 41
    MSK_IPAR_LOG_INCLUDE_SUMMARY Iparam = 42
    MSK_IPAR_LOG_INFEAS_ANA Iparam = 43
    MSK_IPAR_LOG_INTPNT Iparam = 44
    MSK_IPAR_LOG_LOCAL_INFO Iparam = 45
    MSK_IPAR_LOG_MIO Iparam = 46
    MSK_IPAR_LOG_MIO_FREQ Iparam = 47
    MSK_IPAR_LOG_ORDER Iparam = 48
    MSK_IPAR_LOG_PRESOLVE Iparam = 49
    MSK_IPAR_LOG_RESPONSE Iparam = 50
    MSK_IPAR_LOG_SENSITIVITY Iparam = 51
    MSK_IPAR_LOG_SENSITIVITY_OPT Iparam = 52
    MSK_IPAR_LOG_SIM Iparam = 53
    MSK_IPAR_LOG_SIM_FREQ Iparam = 54
    MSK_IPAR_LOG_SIM_MINOR Iparam = 55
    MSK_IPAR_LOG_STORAGE Iparam = 56
    MSK_IPAR_MAX_NUM_WARNINGS Iparam = 57
    MSK_IPAR_MIO_BRANCH_DIR Iparam = 58
    MSK_IPAR_MIO_CONIC_OUTER_APPROXIMATION Iparam = 59
    MSK_IPAR_MIO_CONSTRUCT_SOL Iparam = 60
    MSK_IPAR_MIO_CUT_CLIQUE Iparam = 61
    MSK_IPAR_MIO_CUT_CMIR Iparam = 62
    MSK_IPAR_MIO_CUT_GMI Iparam = 63
    MSK_IPAR_MIO_CUT_IMPLIED_BOUND Iparam = 64
    MSK_IPAR_MIO_CUT_KNAPSACK_COVER Iparam = 65
    MSK_IPAR_MIO_CUT_LIPRO Iparam = 66
    MSK_IPAR_MIO_CUT_SELECTION_LEVEL Iparam = 67
    MSK_IPAR_MIO_DATA_PERMUTATION_METHOD Iparam = 68
    MSK_IPAR_MIO_DUAL_RAY_ANALYSIS_LEVEL Iparam = 69
    MSK_IPAR_MIO_FEASPUMP_LEVEL Iparam = 70
    MSK_IPAR_MIO_HEURISTIC_LEVEL Iparam = 71
    MSK_IPAR_MIO_MAX_NUM_BRANCHES Iparam = 72
    MSK_IPAR_MIO_MAX_NUM_RELAXS Iparam = 73
    MSK_IPAR_MIO_MAX_NUM_RESTARTS Iparam = 74
    MSK_IPAR_MIO_MAX_NUM_ROOT_CUT_ROUNDS Iparam = 75
    MSK_IPAR_MIO_MAX_NUM_SOLUTIONS Iparam = 76
    MSK_IPAR_MIO_MEMORY_EMPHASIS_LEVEL Iparam = 77
    MSK_IPAR_MIO_MIN_REL Iparam = 78
    MSK_IPAR_MIO_MODE Iparam = 79
    MSK_IPAR_MIO_NODE_OPTIMIZER Iparam = 80
    MSK_IPAR_MIO_NODE_SELECTION Iparam = 81
    MSK_IPAR_MIO_NUMERICAL_EMPHASIS_LEVEL Iparam = 82
    MSK_IPAR_MIO_PERSPECTIVE_REFORMULATE Iparam = 83
    MSK_IPAR_MIO_PRESOLVE_AGGREGATOR_USE Iparam = 84
    MSK_IPAR_MIO_PROBING_LEVEL Iparam = 85
    MSK_IPAR_MIO_PROPAGATE_OBJECTIVE_CONSTRAINT Iparam = 86
    MSK_IPAR_MIO_QCQO_REFORMULATION_METHOD Iparam = 87
    MSK_IPAR_MIO_RINS_MAX_NODES Iparam = 88
    MSK_IPAR_MIO_ROOT_OPTIMIZER Iparam = 89
    MSK_IPAR_MIO_ROOT_REPEAT_PRESOLVE_LEVEL Iparam = 90
    MSK_IPAR_MIO_SEED Iparam = 91
    MSK_IPAR_MIO_SYMMETRY_LEVEL Iparam = 92
    MSK_IPAR_MIO_VAR_SELECTION Iparam = 93
    MSK_IPAR_MIO_VB_DETECTION_LEVEL Iparam = 94
    MSK_IPAR_MT_SPINCOUNT Iparam = 95
    MSK_IPAR_NG Iparam = 96
    MSK_IPAR_NUM_THREADS Iparam = 97
    MSK_IPAR_OPF_WRITE_HEADER Iparam = 98
    MSK_IPAR_OPF_WRITE_HINTS Iparam = 99
    MSK_IPAR_OPF_WRITE_LINE_LENGTH Iparam = 100
    MSK_IPAR_OPF_WRITE_PARAMETERS Iparam = 101
    MSK_IPAR_OPF_WRITE_PROBLEM Iparam = 102
    MSK_IPAR_OPF_WRITE_SOL_BAS Iparam = 103
    MSK_IPAR_OPF_WRITE_SOL_ITG Iparam = 104
    MSK_IPAR_OPF_WRITE_SOL_ITR Iparam = 105
    MSK_IPAR_OPF_WRITE_SOLUTIONS Iparam = 106
    MSK_IPAR_OPTIMIZER Iparam = 107
    MSK_IPAR_PARAM_READ_CASE_NAME Iparam = 108
    MSK_IPAR_PARAM_READ_IGN_ERROR Iparam = 109
    MSK_IPAR_PRESOLVE_ELIMINATOR_MAX_FILL Iparam = 110
    MSK_IPAR_PRESOLVE_ELIMINATOR_MAX_NUM_TRIES Iparam = 111
    MSK_IPAR_PRESOLVE_LEVEL Iparam = 112
    MSK_IPAR_PRESOLVE_LINDEP_ABS_WORK_TRH Iparam = 113
    MSK_IPAR_PRESOLVE_LINDEP_NEW Iparam = 114
    MSK_IPAR_PRESOLVE_LINDEP_REL_WORK_TRH Iparam = 115
    MSK_IPAR_PRESOLVE_LINDEP_USE Iparam = 116
    MSK_IPAR_PRESOLVE_MAX_NUM_PASS Iparam = 117
    MSK_IPAR_PRESOLVE_MAX_NUM_REDUCTIONS Iparam = 118
    MSK_IPAR_PRESOLVE_USE Iparam = 119
    MSK_IPAR_PRIMAL_REPAIR_OPTIMIZER Iparam = 120
    MSK_IPAR_PTF_WRITE_PARAMETERS Iparam = 121
    MSK_IPAR_PTF_WRITE_SOLUTIONS Iparam = 122
    MSK_IPAR_PTF_WRITE_TRANSFORM Iparam = 123
    MSK_IPAR_READ_DEBUG Iparam = 124
    MSK_IPAR_READ_KEEP_FREE_CON Iparam = 125
    MSK_IPAR_READ_MPS_FORMAT Iparam = 126
    MSK_IPAR_READ_MPS_WIDTH Iparam = 127
    MSK_IPAR_READ_TASK_IGNORE_PARAM Iparam = 128
    MSK_IPAR_REMOTE_USE_COMPRESSION Iparam = 129
    MSK_IPAR_REMOVE_UNUSED_SOLUTIONS Iparam = 130
    MSK_IPAR_SENSITIVITY_ALL Iparam = 131
    MSK_IPAR_SENSITIVITY_OPTIMIZER Iparam = 132
    MSK_IPAR_SENSITIVITY_TYPE Iparam = 133
    MSK_IPAR_SIM_BASIS_FACTOR_USE Iparam = 134
    MSK_IPAR_SIM_DEGEN Iparam = 135
    MSK_IPAR_SIM_DETECT_PWL Iparam = 136
    MSK_IPAR_SIM_DUAL_CRASH Iparam = 137
    MSK_IPAR_SIM_DUAL_PHASEONE_METHOD Iparam = 138
    MSK_IPAR_SIM_DUAL_RESTRICT_SELECTION Iparam = 139
    MSK_IPAR_SIM_DUAL_SELECTION Iparam = 140
    MSK_IPAR_SIM_EXPLOIT_DUPVEC Iparam = 141
    MSK_IPAR_SIM_HOTSTART Iparam = 142
    MSK_IPAR_SIM_HOTSTART_LU Iparam = 143
    MSK_IPAR_SIM_MAX_ITERATIONS Iparam = 144
    MSK_IPAR_SIM_MAX_NUM_SETBACKS Iparam = 145
    MSK_IPAR_SIM_NON_SINGULAR Iparam = 146
    MSK_IPAR_SIM_PRIMAL_CRASH Iparam = 147
    MSK_IPAR_SIM_PRIMAL_PHASEONE_METHOD Iparam = 148
    MSK_IPAR_SIM_PRIMAL_RESTRICT_SELECTION Iparam = 149
    MSK_IPAR_SIM_PRIMAL_SELECTION Iparam = 150
    MSK_IPAR_SIM_REFACTOR_FREQ Iparam = 151
    MSK_IPAR_SIM_REFORMULATION Iparam = 152
    MSK_IPAR_SIM_SAVE_LU Iparam = 153
    MSK_IPAR_SIM_SCALING Iparam = 154
    MSK_IPAR_SIM_SCALING_METHOD Iparam = 155
    MSK_IPAR_SIM_SEED Iparam = 156
    MSK_IPAR_SIM_SOLVE_FORM Iparam = 157
    MSK_IPAR_SIM_STABILITY_PRIORITY Iparam = 158
    MSK_IPAR_SIM_SWITCH_OPTIMIZER Iparam = 159
    MSK_IPAR_SOL_FILTER_KEEP_BASIC Iparam = 160
    MSK_IPAR_SOL_FILTER_KEEP_RANGED Iparam = 161
    MSK_IPAR_SOL_READ_NAME_WIDTH Iparam = 162
    MSK_IPAR_SOL_READ_WIDTH Iparam = 163
    MSK_IPAR_SOLUTION_CALLBACK Iparam = 164
    MSK_IPAR_TIMING_LEVEL Iparam = 165
    MSK_IPAR_WRITE_BAS_CONSTRAINTS Iparam = 166
    MSK_IPAR_WRITE_BAS_HEAD Iparam = 167
    MSK_IPAR_WRITE_BAS_VARIABLES Iparam = 168
    MSK_IPAR_WRITE_COMPRESSION Iparam = 169
    MSK_IPAR_WRITE_DATA_PARAM Iparam = 170
    MSK_IPAR_WRITE_FREE_CON Iparam = 171
    MSK_IPAR_WRITE_GENERIC_NAMES Iparam = 172
    MSK_IPAR_WRITE_GENERIC_NAMES_IO Iparam = 173
    MSK_IPAR_WRITE_IGNORE_INCOMPATIBLE_ITEMS Iparam = 174
    MSK_IPAR_WRITE_INT_CONSTRAINTS Iparam = 175
    MSK_IPAR_WRITE_INT_HEAD Iparam = 176
    MSK_IPAR_WRITE_INT_VARIABLES Iparam = 177
    MSK_IPAR_WRITE_JSON_INDENTATION Iparam = 178
    MSK_IPAR_WRITE_LP_FULL_OBJ Iparam = 179
    MSK_IPAR_WRITE_LP_LINE_WIDTH Iparam = 180
    MSK_IPAR_WRITE_MPS_FORMAT Iparam = 181
    MSK_IPAR_WRITE_MPS_INT Iparam = 182
    MSK_IPAR_WRITE_SOL_BARVARIABLES Iparam = 183
    MSK_IPAR_WRITE_SOL_CONSTRAINTS Iparam = 184
    MSK_IPAR_WRITE_SOL_HEAD Iparam = 185
    MSK_IPAR_WRITE_SOL_IGNORE_INVALID_NAMES Iparam = 186
    MSK_IPAR_WRITE_SOL_VARIABLES Iparam = 187
    MSK_IPAR_WRITE_TASK_INC_SOL Iparam = 188
    MSK_IPAR_WRITE_XML_MODE Iparam = 189
    MSK_IPAR_END Iparam = 189
)
type Branchdir int32
const (
    MSK_BRANCH_DIR_FREE Branchdir = 0
    MSK_BRANCH_DIR_UP Branchdir = 1
    MSK_BRANCH_DIR_DOWN Branchdir = 2
    MSK_BRANCH_DIR_NEAR Branchdir = 3
    MSK_BRANCH_DIR_FAR Branchdir = 4
    MSK_BRANCH_DIR_ROOT_LP Branchdir = 5
    MSK_BRANCH_DIR_GUIDED Branchdir = 6
    MSK_BRANCH_DIR_PSEUDOCOST Branchdir = 7
    MSK_BRANCH_DIR_END Branchdir = 7
)
type Miqcqoreformmethod int32
const (
    MSK_MIO_QCQO_REFORMULATION_METHOD_FREE Miqcqoreformmethod = 0
    MSK_MIO_QCQO_REFORMULATION_METHOD_NONE Miqcqoreformmethod = 1
    MSK_MIO_QCQO_REFORMULATION_METHOD_LINEARIZATION Miqcqoreformmethod = 2
    MSK_MIO_QCQO_REFORMULATION_METHOD_EIGEN_VAL_METHOD Miqcqoreformmethod = 3
    MSK_MIO_QCQO_REFORMULATION_METHOD_DIAG_SDP Miqcqoreformmethod = 4
    MSK_MIO_QCQO_REFORMULATION_METHOD_RELAX_SDP Miqcqoreformmethod = 5
    MSK_MIO_QCQO_REFORMULATION_METHOD_END Miqcqoreformmethod = 5
)
type Miodatapermmethod int32
const (
    MSK_MIO_DATA_PERMUTATION_METHOD_NONE Miodatapermmethod = 0
    MSK_MIO_DATA_PERMUTATION_METHOD_CYCLIC_SHIFT Miodatapermmethod = 1
    MSK_MIO_DATA_PERMUTATION_METHOD_RANDOM Miodatapermmethod = 2
    MSK_MIO_DATA_PERMUTATION_METHOD_END Miodatapermmethod = 2
)
type Miocontsoltype int32
const (
    MSK_MIO_CONT_SOL_NONE Miocontsoltype = 0
    MSK_MIO_CONT_SOL_ROOT Miocontsoltype = 1
    MSK_MIO_CONT_SOL_ITG Miocontsoltype = 2
    MSK_MIO_CONT_SOL_ITG_REL Miocontsoltype = 3
    MSK_MIO_CONT_SOL_END Miocontsoltype = 3
)
type Miomode int32
const (
    MSK_MIO_MODE_IGNORED Miomode = 0
    MSK_MIO_MODE_SATISFIED Miomode = 1
    MSK_MIO_MODE_END Miomode = 1
)
type Mionodeseltype int32
const (
    MSK_MIO_NODE_SELECTION_FREE Mionodeseltype = 0
    MSK_MIO_NODE_SELECTION_FIRST Mionodeseltype = 1
    MSK_MIO_NODE_SELECTION_BEST Mionodeseltype = 2
    MSK_MIO_NODE_SELECTION_PSEUDO Mionodeseltype = 3
    MSK_MIO_NODE_SELECTION_END Mionodeseltype = 3
)
type Miovarseltype int32
const (
    MSK_MIO_VAR_SELECTION_FREE Miovarseltype = 0
    MSK_MIO_VAR_SELECTION_PSEUDOCOST Miovarseltype = 1
    MSK_MIO_VAR_SELECTION_STRONG Miovarseltype = 2
    MSK_MIO_VAR_SELECTION_END Miovarseltype = 2
)
type Mpsformat int32
const (
    MSK_MPS_FORMAT_STRICT Mpsformat = 0
    MSK_MPS_FORMAT_RELAXED Mpsformat = 1
    MSK_MPS_FORMAT_FREE Mpsformat = 2
    MSK_MPS_FORMAT_CPLEX Mpsformat = 3
    MSK_MPS_FORMAT_END Mpsformat = 3
)
type Objsense int32
const (
    MSK_OBJECTIVE_SENSE_MINIMIZE Objsense = 0
    MSK_OBJECTIVE_SENSE_MAXIMIZE Objsense = 1
    MSK_OBJECTIVE_SENSE_END Objsense = 1
)
type Onoffkey int32
const (
    MSK_OFF Onoffkey = 0
    MSK_ON Onoffkey = 1
)
type Optimizertype int32
const (
    MSK_OPTIMIZER_CONIC Optimizertype = 0
    MSK_OPTIMIZER_DUAL_SIMPLEX Optimizertype = 1
    MSK_OPTIMIZER_FREE Optimizertype = 2
    MSK_OPTIMIZER_FREE_SIMPLEX Optimizertype = 3
    MSK_OPTIMIZER_INTPNT Optimizertype = 4
    MSK_OPTIMIZER_MIXED_INT Optimizertype = 5
    MSK_OPTIMIZER_PRIMAL_SIMPLEX Optimizertype = 6
    MSK_OPTIMIZER_END Optimizertype = 6
)
type Orderingtype int32
const (
    MSK_ORDER_METHOD_FREE Orderingtype = 0
    MSK_ORDER_METHOD_APPMINLOC Orderingtype = 1
    MSK_ORDER_METHOD_EXPERIMENTAL Orderingtype = 2
    MSK_ORDER_METHOD_TRY_GRAPHPAR Orderingtype = 3
    MSK_ORDER_METHOD_FORCE_GRAPHPAR Orderingtype = 4
    MSK_ORDER_METHOD_NONE Orderingtype = 5
    MSK_ORDER_METHOD_END Orderingtype = 5
)
type Presolvemode int32
const (
    MSK_PRESOLVE_MODE_OFF Presolvemode = 0
    MSK_PRESOLVE_MODE_ON Presolvemode = 1
    MSK_PRESOLVE_MODE_FREE Presolvemode = 2
    MSK_PRESOLVE_MODE_END Presolvemode = 2
)
type Parametertype int32
const (
    MSK_PAR_INVALID_TYPE Parametertype = 0
    MSK_PAR_DOU_TYPE Parametertype = 1
    MSK_PAR_INT_TYPE Parametertype = 2
    MSK_PAR_STR_TYPE Parametertype = 3
    MSK_PAR_END Parametertype = 3
)
type Problemitem int32
const (
    MSK_PI_VAR Problemitem = 0
    MSK_PI_CON Problemitem = 1
    MSK_PI_CONE Problemitem = 2
    MSK_PI_END Problemitem = 2
)
type Problemtype int32
const (
    MSK_PROBTYPE_LO Problemtype = 0
    MSK_PROBTYPE_QO Problemtype = 1
    MSK_PROBTYPE_QCQO Problemtype = 2
    MSK_PROBTYPE_CONIC Problemtype = 3
    MSK_PROBTYPE_MIXED Problemtype = 4
    MSK_PROBTYPE_END Problemtype = 4
)
type Prosta int32
const (
    MSK_PRO_STA_UNKNOWN Prosta = 0
    MSK_PRO_STA_PRIM_AND_DUAL_FEAS Prosta = 1
    MSK_PRO_STA_PRIM_FEAS Prosta = 2
    MSK_PRO_STA_DUAL_FEAS Prosta = 3
    MSK_PRO_STA_PRIM_INFEAS Prosta = 4
    MSK_PRO_STA_DUAL_INFEAS Prosta = 5
    MSK_PRO_STA_PRIM_AND_DUAL_INFEAS Prosta = 6
    MSK_PRO_STA_ILL_POSED Prosta = 7
    MSK_PRO_STA_PRIM_INFEAS_OR_UNBOUNDED Prosta = 8
    MSK_PRO_STA_END Prosta = 8
)
type Xmlwriteroutputtype int32
const (
    MSK_WRITE_XML_MODE_ROW Xmlwriteroutputtype = 0
    MSK_WRITE_XML_MODE_COL Xmlwriteroutputtype = 1
    MSK_WRITE_XML_MODE_END Xmlwriteroutputtype = 1
)
type Rescode int32
const (
    MSK_RES_OK Rescode = 0
    MSK_RES_WRN_OPEN_PARAM_FILE Rescode = 50
    MSK_RES_WRN_LARGE_BOUND Rescode = 51
    MSK_RES_WRN_LARGE_LO_BOUND Rescode = 52
    MSK_RES_WRN_LARGE_UP_BOUND Rescode = 53
    MSK_RES_WRN_LARGE_CON_FX Rescode = 54
    MSK_RES_WRN_LARGE_CJ Rescode = 57
    MSK_RES_WRN_LARGE_AIJ Rescode = 62
    MSK_RES_WRN_ZERO_AIJ Rescode = 63
    MSK_RES_WRN_NAME_MAX_LEN Rescode = 65
    MSK_RES_WRN_SPAR_MAX_LEN Rescode = 66
    MSK_RES_WRN_MPS_SPLIT_RHS_VECTOR Rescode = 70
    MSK_RES_WRN_MPS_SPLIT_RAN_VECTOR Rescode = 71
    MSK_RES_WRN_MPS_SPLIT_BOU_VECTOR Rescode = 72
    MSK_RES_WRN_LP_OLD_QUAD_FORMAT Rescode = 80
    MSK_RES_WRN_LP_DROP_VARIABLE Rescode = 85
    MSK_RES_WRN_NZ_IN_UPR_TRI Rescode = 200
    MSK_RES_WRN_DROPPED_NZ_QOBJ Rescode = 201
    MSK_RES_WRN_IGNORE_INTEGER Rescode = 250
    MSK_RES_WRN_NO_GLOBAL_OPTIMIZER Rescode = 251
    MSK_RES_WRN_MIO_INFEASIBLE_FINAL Rescode = 270
    MSK_RES_WRN_SOL_FILTER Rescode = 300
    MSK_RES_WRN_UNDEF_SOL_FILE_NAME Rescode = 350
    MSK_RES_WRN_SOL_FILE_IGNORED_CON Rescode = 351
    MSK_RES_WRN_SOL_FILE_IGNORED_VAR Rescode = 352
    MSK_RES_WRN_TOO_FEW_BASIS_VARS Rescode = 400
    MSK_RES_WRN_TOO_MANY_BASIS_VARS Rescode = 405
    MSK_RES_WRN_LICENSE_EXPIRE Rescode = 500
    MSK_RES_WRN_LICENSE_SERVER Rescode = 501
    MSK_RES_WRN_EMPTY_NAME Rescode = 502
    MSK_RES_WRN_USING_GENERIC_NAMES Rescode = 503
    MSK_RES_WRN_INVALID_MPS_NAME Rescode = 504
    MSK_RES_WRN_INVALID_MPS_OBJ_NAME Rescode = 505
    MSK_RES_WRN_LICENSE_FEATURE_EXPIRE Rescode = 509
    MSK_RES_WRN_PARAM_NAME_DOU Rescode = 510
    MSK_RES_WRN_PARAM_NAME_INT Rescode = 511
    MSK_RES_WRN_PARAM_NAME_STR Rescode = 512
    MSK_RES_WRN_PARAM_STR_VALUE Rescode = 515
    MSK_RES_WRN_PARAM_IGNORED_CMIO Rescode = 516
    MSK_RES_WRN_ZEROS_IN_SPARSE_ROW Rescode = 705
    MSK_RES_WRN_ZEROS_IN_SPARSE_COL Rescode = 710
    MSK_RES_WRN_INCOMPLETE_LINEAR_DEPENDENCY_CHECK Rescode = 800
    MSK_RES_WRN_ELIMINATOR_SPACE Rescode = 801
    MSK_RES_WRN_PRESOLVE_OUTOFSPACE Rescode = 802
    MSK_RES_WRN_PRESOLVE_PRIMAL_PERTUBATIONS Rescode = 803
    MSK_RES_WRN_WRITE_CHANGED_NAMES Rescode = 830
    MSK_RES_WRN_WRITE_DISCARDED_CFIX Rescode = 831
    MSK_RES_WRN_DUPLICATE_CONSTRAINT_NAMES Rescode = 850
    MSK_RES_WRN_DUPLICATE_VARIABLE_NAMES Rescode = 851
    MSK_RES_WRN_DUPLICATE_BARVARIABLE_NAMES Rescode = 852
    MSK_RES_WRN_DUPLICATE_CONE_NAMES Rescode = 853
    MSK_RES_WRN_WRITE_LP_INVALID_VAR_NAMES Rescode = 854
    MSK_RES_WRN_WRITE_LP_DUPLICATE_VAR_NAMES Rescode = 855
    MSK_RES_WRN_WRITE_LP_INVALID_CON_NAMES Rescode = 856
    MSK_RES_WRN_WRITE_LP_DUPLICATE_CON_NAMES Rescode = 857
    MSK_RES_WRN_ANA_LARGE_BOUNDS Rescode = 900
    MSK_RES_WRN_ANA_C_ZERO Rescode = 901
    MSK_RES_WRN_ANA_EMPTY_COLS Rescode = 902
    MSK_RES_WRN_ANA_CLOSE_BOUNDS Rescode = 903
    MSK_RES_WRN_ANA_ALMOST_INT_BOUNDS Rescode = 904
    MSK_RES_WRN_NO_INFEASIBILITY_REPORT_WHEN_MATRIX_VARIABLES Rescode = 930
    MSK_RES_WRN_NO_DUALIZER Rescode = 950
    MSK_RES_WRN_SYM_MAT_LARGE Rescode = 960
    MSK_RES_WRN_MODIFIED_DOUBLE_PARAMETER Rescode = 970
    MSK_RES_WRN_LARGE_FIJ Rescode = 980
    MSK_RES_ERR_LICENSE Rescode = 1000
    MSK_RES_ERR_LICENSE_EXPIRED Rescode = 1001
    MSK_RES_ERR_LICENSE_VERSION Rescode = 1002
    MSK_RES_ERR_LICENSE_OLD_SERVER_VERSION Rescode = 1003
    MSK_RES_ERR_SIZE_LICENSE Rescode = 1005
    MSK_RES_ERR_PROB_LICENSE Rescode = 1006
    MSK_RES_ERR_FILE_LICENSE Rescode = 1007
    MSK_RES_ERR_MISSING_LICENSE_FILE Rescode = 1008
    MSK_RES_ERR_SIZE_LICENSE_CON Rescode = 1010
    MSK_RES_ERR_SIZE_LICENSE_VAR Rescode = 1011
    MSK_RES_ERR_SIZE_LICENSE_INTVAR Rescode = 1012
    MSK_RES_ERR_OPTIMIZER_LICENSE Rescode = 1013
    MSK_RES_ERR_FLEXLM Rescode = 1014
    MSK_RES_ERR_LICENSE_SERVER Rescode = 1015
    MSK_RES_ERR_LICENSE_MAX Rescode = 1016
    MSK_RES_ERR_LICENSE_MOSEKLM_DAEMON Rescode = 1017
    MSK_RES_ERR_LICENSE_FEATURE Rescode = 1018
    MSK_RES_ERR_PLATFORM_NOT_LICENSED Rescode = 1019
    MSK_RES_ERR_LICENSE_CANNOT_ALLOCATE Rescode = 1020
    MSK_RES_ERR_LICENSE_CANNOT_CONNECT Rescode = 1021
    MSK_RES_ERR_LICENSE_INVALID_HOSTID Rescode = 1025
    MSK_RES_ERR_LICENSE_SERVER_VERSION Rescode = 1026
    MSK_RES_ERR_LICENSE_NO_SERVER_SUPPORT Rescode = 1027
    MSK_RES_ERR_LICENSE_NO_SERVER_LINE Rescode = 1028
    MSK_RES_ERR_OLDER_DLL Rescode = 1035
    MSK_RES_ERR_NEWER_DLL Rescode = 1036
    MSK_RES_ERR_LINK_FILE_DLL Rescode = 1040
    MSK_RES_ERR_THREAD_MUTEX_INIT Rescode = 1045
    MSK_RES_ERR_THREAD_MUTEX_LOCK Rescode = 1046
    MSK_RES_ERR_THREAD_MUTEX_UNLOCK Rescode = 1047
    MSK_RES_ERR_THREAD_CREATE Rescode = 1048
    MSK_RES_ERR_THREAD_COND_INIT Rescode = 1049
    MSK_RES_ERR_UNKNOWN Rescode = 1050
    MSK_RES_ERR_SPACE Rescode = 1051
    MSK_RES_ERR_FILE_OPEN Rescode = 1052
    MSK_RES_ERR_FILE_READ Rescode = 1053
    MSK_RES_ERR_FILE_WRITE Rescode = 1054
    MSK_RES_ERR_DATA_FILE_EXT Rescode = 1055
    MSK_RES_ERR_INVALID_FILE_NAME Rescode = 1056
    MSK_RES_ERR_INVALID_SOL_FILE_NAME Rescode = 1057
    MSK_RES_ERR_END_OF_FILE Rescode = 1059
    MSK_RES_ERR_NULL_ENV Rescode = 1060
    MSK_RES_ERR_NULL_TASK Rescode = 1061
    MSK_RES_ERR_INVALID_STREAM Rescode = 1062
    MSK_RES_ERR_NO_INIT_ENV Rescode = 1063
    MSK_RES_ERR_INVALID_TASK Rescode = 1064
    MSK_RES_ERR_NULL_POINTER Rescode = 1065
    MSK_RES_ERR_LIVING_TASKS Rescode = 1066
    MSK_RES_ERR_READ_GZIP Rescode = 1067
    MSK_RES_ERR_READ_ZSTD Rescode = 1068
    MSK_RES_ERR_BLANK_NAME Rescode = 1070
    MSK_RES_ERR_DUP_NAME Rescode = 1071
    MSK_RES_ERR_FORMAT_STRING Rescode = 1072
    MSK_RES_ERR_SPARSITY_SPECIFICATION Rescode = 1073
    MSK_RES_ERR_MISMATCHING_DIMENSION Rescode = 1074
    MSK_RES_ERR_INVALID_OBJ_NAME Rescode = 1075
    MSK_RES_ERR_INVALID_CON_NAME Rescode = 1076
    MSK_RES_ERR_INVALID_VAR_NAME Rescode = 1077
    MSK_RES_ERR_INVALID_CONE_NAME Rescode = 1078
    MSK_RES_ERR_INVALID_BARVAR_NAME Rescode = 1079
    MSK_RES_ERR_SPACE_LEAKING Rescode = 1080
    MSK_RES_ERR_SPACE_NO_INFO Rescode = 1081
    MSK_RES_ERR_DIMENSION_SPECIFICATION Rescode = 1082
    MSK_RES_ERR_AXIS_NAME_SPECIFICATION Rescode = 1083
    MSK_RES_ERR_READ_FORMAT Rescode = 1090
    MSK_RES_ERR_MPS_FILE Rescode = 1100
    MSK_RES_ERR_MPS_INV_FIELD Rescode = 1101
    MSK_RES_ERR_MPS_INV_MARKER Rescode = 1102
    MSK_RES_ERR_MPS_NULL_CON_NAME Rescode = 1103
    MSK_RES_ERR_MPS_NULL_VAR_NAME Rescode = 1104
    MSK_RES_ERR_MPS_UNDEF_CON_NAME Rescode = 1105
    MSK_RES_ERR_MPS_UNDEF_VAR_NAME Rescode = 1106
    MSK_RES_ERR_MPS_INVALID_CON_KEY Rescode = 1107
    MSK_RES_ERR_MPS_INVALID_BOUND_KEY Rescode = 1108
    MSK_RES_ERR_MPS_INVALID_SEC_NAME Rescode = 1109
    MSK_RES_ERR_MPS_NO_OBJECTIVE Rescode = 1110
    MSK_RES_ERR_MPS_SPLITTED_VAR Rescode = 1111
    MSK_RES_ERR_MPS_MUL_CON_NAME Rescode = 1112
    MSK_RES_ERR_MPS_MUL_QSEC Rescode = 1113
    MSK_RES_ERR_MPS_MUL_QOBJ Rescode = 1114
    MSK_RES_ERR_MPS_INV_SEC_ORDER Rescode = 1115
    MSK_RES_ERR_MPS_MUL_CSEC Rescode = 1116
    MSK_RES_ERR_MPS_CONE_TYPE Rescode = 1117
    MSK_RES_ERR_MPS_CONE_OVERLAP Rescode = 1118
    MSK_RES_ERR_MPS_CONE_REPEAT Rescode = 1119
    MSK_RES_ERR_MPS_NON_SYMMETRIC_Q Rescode = 1120
    MSK_RES_ERR_MPS_DUPLICATE_Q_ELEMENT Rescode = 1121
    MSK_RES_ERR_MPS_INVALID_OBJSENSE Rescode = 1122
    MSK_RES_ERR_MPS_TAB_IN_FIELD2 Rescode = 1125
    MSK_RES_ERR_MPS_TAB_IN_FIELD3 Rescode = 1126
    MSK_RES_ERR_MPS_TAB_IN_FIELD5 Rescode = 1127
    MSK_RES_ERR_MPS_INVALID_OBJ_NAME Rescode = 1128
    MSK_RES_ERR_MPS_INVALID_KEY Rescode = 1129
    MSK_RES_ERR_MPS_INVALID_INDICATOR_CONSTRAINT Rescode = 1130
    MSK_RES_ERR_MPS_INVALID_INDICATOR_VARIABLE Rescode = 1131
    MSK_RES_ERR_MPS_INVALID_INDICATOR_VALUE Rescode = 1132
    MSK_RES_ERR_MPS_INVALID_INDICATOR_QUADRATIC_CONSTRAINT Rescode = 1133
    MSK_RES_ERR_OPF_SYNTAX Rescode = 1134
    MSK_RES_ERR_OPF_PREMATURE_EOF Rescode = 1136
    MSK_RES_ERR_OPF_MISMATCHED_TAG Rescode = 1137
    MSK_RES_ERR_OPF_DUPLICATE_BOUND Rescode = 1138
    MSK_RES_ERR_OPF_DUPLICATE_CONSTRAINT_NAME Rescode = 1139
    MSK_RES_ERR_OPF_INVALID_CONE_TYPE Rescode = 1140
    MSK_RES_ERR_OPF_INCORRECT_TAG_PARAM Rescode = 1141
    MSK_RES_ERR_OPF_INVALID_TAG Rescode = 1142
    MSK_RES_ERR_OPF_DUPLICATE_CONE_ENTRY Rescode = 1143
    MSK_RES_ERR_OPF_TOO_LARGE Rescode = 1144
    MSK_RES_ERR_OPF_DUAL_INTEGER_SOLUTION Rescode = 1146
    MSK_RES_ERR_LP_EMPTY Rescode = 1151
    MSK_RES_ERR_WRITE_MPS_INVALID_NAME Rescode = 1153
    MSK_RES_ERR_LP_INVALID_VAR_NAME Rescode = 1154
    MSK_RES_ERR_WRITE_OPF_INVALID_VAR_NAME Rescode = 1156
    MSK_RES_ERR_LP_FILE_FORMAT Rescode = 1157
    MSK_RES_ERR_LP_EXPECTED_NUMBER Rescode = 1158
    MSK_RES_ERR_READ_LP_MISSING_END_TAG Rescode = 1159
    MSK_RES_ERR_LP_INDICATOR_VAR Rescode = 1160
    MSK_RES_ERR_LP_EXPECTED_OBJECTIVE Rescode = 1161
    MSK_RES_ERR_LP_EXPECTED_CONSTRAINT_RELATION Rescode = 1162
    MSK_RES_ERR_LP_AMBIGUOUS_CONSTRAINT_BOUND Rescode = 1163
    MSK_RES_ERR_LP_DUPLICATE_SECTION Rescode = 1164
    MSK_RES_ERR_READ_LP_DELAYED_ROWS_NOT_SUPPORTED Rescode = 1165
    MSK_RES_ERR_WRITING_FILE Rescode = 1166
    MSK_RES_ERR_INVALID_NAME_IN_SOL_FILE Rescode = 1170
    MSK_RES_ERR_JSON_SYNTAX Rescode = 1175
    MSK_RES_ERR_JSON_STRING Rescode = 1176
    MSK_RES_ERR_JSON_NUMBER_OVERFLOW Rescode = 1177
    MSK_RES_ERR_JSON_FORMAT Rescode = 1178
    MSK_RES_ERR_JSON_DATA Rescode = 1179
    MSK_RES_ERR_JSON_MISSING_DATA Rescode = 1180
    MSK_RES_ERR_PTF_INCOMPATIBILITY Rescode = 1181
    MSK_RES_ERR_PTF_UNDEFINED_ITEM Rescode = 1182
    MSK_RES_ERR_PTF_INCONSISTENCY Rescode = 1183
    MSK_RES_ERR_PTF_FORMAT Rescode = 1184
    MSK_RES_ERR_ARGUMENT_LENNEQ Rescode = 1197
    MSK_RES_ERR_ARGUMENT_TYPE Rescode = 1198
    MSK_RES_ERR_NUM_ARGUMENTS Rescode = 1199
    MSK_RES_ERR_IN_ARGUMENT Rescode = 1200
    MSK_RES_ERR_ARGUMENT_DIMENSION Rescode = 1201
    MSK_RES_ERR_SHAPE_IS_TOO_LARGE Rescode = 1202
    MSK_RES_ERR_INDEX_IS_TOO_SMALL Rescode = 1203
    MSK_RES_ERR_INDEX_IS_TOO_LARGE Rescode = 1204
    MSK_RES_ERR_INDEX_IS_NOT_UNIQUE Rescode = 1205
    MSK_RES_ERR_PARAM_NAME Rescode = 1206
    MSK_RES_ERR_PARAM_NAME_DOU Rescode = 1207
    MSK_RES_ERR_PARAM_NAME_INT Rescode = 1208
    MSK_RES_ERR_PARAM_NAME_STR Rescode = 1209
    MSK_RES_ERR_PARAM_INDEX Rescode = 1210
    MSK_RES_ERR_PARAM_IS_TOO_LARGE Rescode = 1215
    MSK_RES_ERR_PARAM_IS_TOO_SMALL Rescode = 1216
    MSK_RES_ERR_PARAM_VALUE_STR Rescode = 1217
    MSK_RES_ERR_PARAM_TYPE Rescode = 1218
    MSK_RES_ERR_INF_DOU_INDEX Rescode = 1219
    MSK_RES_ERR_INF_INT_INDEX Rescode = 1220
    MSK_RES_ERR_INDEX_ARR_IS_TOO_SMALL Rescode = 1221
    MSK_RES_ERR_INDEX_ARR_IS_TOO_LARGE Rescode = 1222
    MSK_RES_ERR_INF_LINT_INDEX Rescode = 1225
    MSK_RES_ERR_ARG_IS_TOO_SMALL Rescode = 1226
    MSK_RES_ERR_ARG_IS_TOO_LARGE Rescode = 1227
    MSK_RES_ERR_INVALID_WHICHSOL Rescode = 1228
    MSK_RES_ERR_INF_DOU_NAME Rescode = 1230
    MSK_RES_ERR_INF_INT_NAME Rescode = 1231
    MSK_RES_ERR_INF_TYPE Rescode = 1232
    MSK_RES_ERR_INF_LINT_NAME Rescode = 1234
    MSK_RES_ERR_INDEX Rescode = 1235
    MSK_RES_ERR_WHICHSOL Rescode = 1236
    MSK_RES_ERR_SOLITEM Rescode = 1237
    MSK_RES_ERR_WHICHITEM_NOT_ALLOWED Rescode = 1238
    MSK_RES_ERR_MAXNUMCON Rescode = 1240
    MSK_RES_ERR_MAXNUMVAR Rescode = 1241
    MSK_RES_ERR_MAXNUMBARVAR Rescode = 1242
    MSK_RES_ERR_MAXNUMQNZ Rescode = 1243
    MSK_RES_ERR_TOO_SMALL_MAX_NUM_NZ Rescode = 1245
    MSK_RES_ERR_INVALID_IDX Rescode = 1246
    MSK_RES_ERR_INVALID_MAX_NUM Rescode = 1247
    MSK_RES_ERR_UNALLOWED_WHICHSOL Rescode = 1248
    MSK_RES_ERR_NUMCONLIM Rescode = 1250
    MSK_RES_ERR_NUMVARLIM Rescode = 1251
    MSK_RES_ERR_TOO_SMALL_MAXNUMANZ Rescode = 1252
    MSK_RES_ERR_INV_APTRE Rescode = 1253
    MSK_RES_ERR_MUL_A_ELEMENT Rescode = 1254
    MSK_RES_ERR_INV_BK Rescode = 1255
    MSK_RES_ERR_INV_BKC Rescode = 1256
    MSK_RES_ERR_INV_BKX Rescode = 1257
    MSK_RES_ERR_INV_VAR_TYPE Rescode = 1258
    MSK_RES_ERR_SOLVER_PROBTYPE Rescode = 1259
    MSK_RES_ERR_OBJECTIVE_RANGE Rescode = 1260
    MSK_RES_ERR_INV_RESCODE Rescode = 1261
    MSK_RES_ERR_INV_IINF Rescode = 1262
    MSK_RES_ERR_INV_LIINF Rescode = 1263
    MSK_RES_ERR_INV_DINF Rescode = 1264
    MSK_RES_ERR_BASIS Rescode = 1266
    MSK_RES_ERR_INV_SKC Rescode = 1267
    MSK_RES_ERR_INV_SKX Rescode = 1268
    MSK_RES_ERR_INV_SK_STR Rescode = 1269
    MSK_RES_ERR_INV_SK Rescode = 1270
    MSK_RES_ERR_INV_CONE_TYPE_STR Rescode = 1271
    MSK_RES_ERR_INV_CONE_TYPE Rescode = 1272
    MSK_RES_ERR_INV_SKN Rescode = 1274
    MSK_RES_ERR_INVALID_SURPLUS Rescode = 1275
    MSK_RES_ERR_INV_NAME_ITEM Rescode = 1280
    MSK_RES_ERR_PRO_ITEM Rescode = 1281
    MSK_RES_ERR_INVALID_FORMAT_TYPE Rescode = 1283
    MSK_RES_ERR_FIRSTI Rescode = 1285
    MSK_RES_ERR_LASTI Rescode = 1286
    MSK_RES_ERR_FIRSTJ Rescode = 1287
    MSK_RES_ERR_LASTJ Rescode = 1288
    MSK_RES_ERR_MAX_LEN_IS_TOO_SMALL Rescode = 1289
    MSK_RES_ERR_NONLINEAR_EQUALITY Rescode = 1290
    MSK_RES_ERR_NONCONVEX Rescode = 1291
    MSK_RES_ERR_NONLINEAR_RANGED Rescode = 1292
    MSK_RES_ERR_CON_Q_NOT_PSD Rescode = 1293
    MSK_RES_ERR_CON_Q_NOT_NSD Rescode = 1294
    MSK_RES_ERR_OBJ_Q_NOT_PSD Rescode = 1295
    MSK_RES_ERR_OBJ_Q_NOT_NSD Rescode = 1296
    MSK_RES_ERR_ARGUMENT_PERM_ARRAY Rescode = 1299
    MSK_RES_ERR_CONE_INDEX Rescode = 1300
    MSK_RES_ERR_CONE_SIZE Rescode = 1301
    MSK_RES_ERR_CONE_OVERLAP Rescode = 1302
    MSK_RES_ERR_CONE_REP_VAR Rescode = 1303
    MSK_RES_ERR_MAXNUMCONE Rescode = 1304
    MSK_RES_ERR_CONE_TYPE Rescode = 1305
    MSK_RES_ERR_CONE_TYPE_STR Rescode = 1306
    MSK_RES_ERR_CONE_OVERLAP_APPEND Rescode = 1307
    MSK_RES_ERR_REMOVE_CONE_VARIABLE Rescode = 1310
    MSK_RES_ERR_APPENDING_TOO_BIG_CONE Rescode = 1311
    MSK_RES_ERR_CONE_PARAMETER Rescode = 1320
    MSK_RES_ERR_SOL_FILE_INVALID_NUMBER Rescode = 1350
    MSK_RES_ERR_HUGE_C Rescode = 1375
    MSK_RES_ERR_HUGE_AIJ Rescode = 1380
    MSK_RES_ERR_DUPLICATE_AIJ Rescode = 1385
    MSK_RES_ERR_LOWER_BOUND_IS_A_NAN Rescode = 1390
    MSK_RES_ERR_UPPER_BOUND_IS_A_NAN Rescode = 1391
    MSK_RES_ERR_INFINITE_BOUND Rescode = 1400
    MSK_RES_ERR_INV_QOBJ_SUBI Rescode = 1401
    MSK_RES_ERR_INV_QOBJ_SUBJ Rescode = 1402
    MSK_RES_ERR_INV_QOBJ_VAL Rescode = 1403
    MSK_RES_ERR_INV_QCON_SUBK Rescode = 1404
    MSK_RES_ERR_INV_QCON_SUBI Rescode = 1405
    MSK_RES_ERR_INV_QCON_SUBJ Rescode = 1406
    MSK_RES_ERR_INV_QCON_VAL Rescode = 1407
    MSK_RES_ERR_QCON_SUBI_TOO_SMALL Rescode = 1408
    MSK_RES_ERR_QCON_SUBI_TOO_LARGE Rescode = 1409
    MSK_RES_ERR_QOBJ_UPPER_TRIANGLE Rescode = 1415
    MSK_RES_ERR_QCON_UPPER_TRIANGLE Rescode = 1417
    MSK_RES_ERR_FIXED_BOUND_VALUES Rescode = 1420
    MSK_RES_ERR_TOO_SMALL_A_TRUNCATION_VALUE Rescode = 1421
    MSK_RES_ERR_INVALID_OBJECTIVE_SENSE Rescode = 1445
    MSK_RES_ERR_UNDEFINED_OBJECTIVE_SENSE Rescode = 1446
    MSK_RES_ERR_Y_IS_UNDEFINED Rescode = 1449
    MSK_RES_ERR_NAN_IN_DOUBLE_DATA Rescode = 1450
    MSK_RES_ERR_INF_IN_DOUBLE_DATA Rescode = 1451
    MSK_RES_ERR_NAN_IN_BLC Rescode = 1461
    MSK_RES_ERR_NAN_IN_BUC Rescode = 1462
    MSK_RES_ERR_INVALID_CFIX Rescode = 1469
    MSK_RES_ERR_NAN_IN_C Rescode = 1470
    MSK_RES_ERR_NAN_IN_BLX Rescode = 1471
    MSK_RES_ERR_NAN_IN_BUX Rescode = 1472
    MSK_RES_ERR_INVALID_AIJ Rescode = 1473
    MSK_RES_ERR_INVALID_CJ Rescode = 1474
    MSK_RES_ERR_SYM_MAT_INVALID Rescode = 1480
    MSK_RES_ERR_SYM_MAT_HUGE Rescode = 1482
    MSK_RES_ERR_INV_PROBLEM Rescode = 1500
    MSK_RES_ERR_MIXED_CONIC_AND_NL Rescode = 1501
    MSK_RES_ERR_GLOBAL_INV_CONIC_PROBLEM Rescode = 1503
    MSK_RES_ERR_INV_OPTIMIZER Rescode = 1550
    MSK_RES_ERR_MIO_NO_OPTIMIZER Rescode = 1551
    MSK_RES_ERR_NO_OPTIMIZER_VAR_TYPE Rescode = 1552
    MSK_RES_ERR_FINAL_SOLUTION Rescode = 1560
    MSK_RES_ERR_FIRST Rescode = 1570
    MSK_RES_ERR_LAST Rescode = 1571
    MSK_RES_ERR_SLICE_SIZE Rescode = 1572
    MSK_RES_ERR_NEGATIVE_SURPLUS Rescode = 1573
    MSK_RES_ERR_NEGATIVE_APPEND Rescode = 1578
    MSK_RES_ERR_POSTSOLVE Rescode = 1580
    MSK_RES_ERR_OVERFLOW Rescode = 1590
    MSK_RES_ERR_NO_BASIS_SOL Rescode = 1600
    MSK_RES_ERR_BASIS_FACTOR Rescode = 1610
    MSK_RES_ERR_BASIS_SINGULAR Rescode = 1615
    MSK_RES_ERR_FACTOR Rescode = 1650
    MSK_RES_ERR_FEASREPAIR_CANNOT_RELAX Rescode = 1700
    MSK_RES_ERR_FEASREPAIR_SOLVING_RELAXED Rescode = 1701
    MSK_RES_ERR_FEASREPAIR_INCONSISTENT_BOUND Rescode = 1702
    MSK_RES_ERR_REPAIR_INVALID_PROBLEM Rescode = 1710
    MSK_RES_ERR_REPAIR_OPTIMIZATION_FAILED Rescode = 1711
    MSK_RES_ERR_NAME_MAX_LEN Rescode = 1750
    MSK_RES_ERR_NAME_IS_NULL Rescode = 1760
    MSK_RES_ERR_INVALID_COMPRESSION Rescode = 1800
    MSK_RES_ERR_INVALID_IOMODE Rescode = 1801
    MSK_RES_ERR_NO_PRIMAL_INFEAS_CER Rescode = 2000
    MSK_RES_ERR_NO_DUAL_INFEAS_CER Rescode = 2001
    MSK_RES_ERR_NO_SOLUTION_IN_CALLBACK Rescode = 2500
    MSK_RES_ERR_INV_MARKI Rescode = 2501
    MSK_RES_ERR_INV_MARKJ Rescode = 2502
    MSK_RES_ERR_INV_NUMI Rescode = 2503
    MSK_RES_ERR_INV_NUMJ Rescode = 2504
    MSK_RES_ERR_TASK_INCOMPATIBLE Rescode = 2560
    MSK_RES_ERR_TASK_INVALID Rescode = 2561
    MSK_RES_ERR_TASK_WRITE Rescode = 2562
    MSK_RES_ERR_LU_MAX_NUM_TRIES Rescode = 2800
    MSK_RES_ERR_INVALID_UTF8 Rescode = 2900
    MSK_RES_ERR_INVALID_WCHAR Rescode = 2901
    MSK_RES_ERR_NO_DUAL_FOR_ITG_SOL Rescode = 2950
    MSK_RES_ERR_NO_SNX_FOR_BAS_SOL Rescode = 2953
    MSK_RES_ERR_INTERNAL Rescode = 3000
    MSK_RES_ERR_API_ARRAY_TOO_SMALL Rescode = 3001
    MSK_RES_ERR_API_CB_CONNECT Rescode = 3002
    MSK_RES_ERR_API_FATAL_ERROR Rescode = 3005
    MSK_RES_ERR_SEN_FORMAT Rescode = 3050
    MSK_RES_ERR_SEN_UNDEF_NAME Rescode = 3051
    MSK_RES_ERR_SEN_INDEX_RANGE Rescode = 3052
    MSK_RES_ERR_SEN_BOUND_INVALID_UP Rescode = 3053
    MSK_RES_ERR_SEN_BOUND_INVALID_LO Rescode = 3054
    MSK_RES_ERR_SEN_INDEX_INVALID Rescode = 3055
    MSK_RES_ERR_SEN_INVALID_REGEXP Rescode = 3056
    MSK_RES_ERR_SEN_SOLUTION_STATUS Rescode = 3057
    MSK_RES_ERR_SEN_NUMERICAL Rescode = 3058
    MSK_RES_ERR_SEN_UNHANDLED_PROBLEM_TYPE Rescode = 3080
    MSK_RES_ERR_UNB_STEP_SIZE Rescode = 3100
    MSK_RES_ERR_IDENTICAL_TASKS Rescode = 3101
    MSK_RES_ERR_AD_INVALID_CODELIST Rescode = 3102
    MSK_RES_ERR_INTERNAL_TEST_FAILED Rescode = 3500
    MSK_RES_ERR_XML_INVALID_PROBLEM_TYPE Rescode = 3600
    MSK_RES_ERR_INVALID_AMPL_STUB Rescode = 3700
    MSK_RES_ERR_INT64_TO_INT32_CAST Rescode = 3800
    MSK_RES_ERR_SIZE_LICENSE_NUMCORES Rescode = 3900
    MSK_RES_ERR_INFEAS_UNDEFINED Rescode = 3910
    MSK_RES_ERR_NO_BARX_FOR_SOLUTION Rescode = 3915
    MSK_RES_ERR_NO_BARS_FOR_SOLUTION Rescode = 3916
    MSK_RES_ERR_BAR_VAR_DIM Rescode = 3920
    MSK_RES_ERR_SYM_MAT_INVALID_ROW_INDEX Rescode = 3940
    MSK_RES_ERR_SYM_MAT_INVALID_COL_INDEX Rescode = 3941
    MSK_RES_ERR_SYM_MAT_NOT_LOWER_TRINGULAR Rescode = 3942
    MSK_RES_ERR_SYM_MAT_INVALID_VALUE Rescode = 3943
    MSK_RES_ERR_SYM_MAT_DUPLICATE Rescode = 3944
    MSK_RES_ERR_INVALID_SYM_MAT_DIM Rescode = 3950
    MSK_RES_ERR_API_INTERNAL Rescode = 3999
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_SYM_MAT Rescode = 4000
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_CFIX Rescode = 4001
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_RANGED_CONSTRAINTS Rescode = 4002
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_FREE_CONSTRAINTS Rescode = 4003
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_CONES Rescode = 4005
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_QUADRATIC_TERMS Rescode = 4006
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_NONLINEAR Rescode = 4010
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_DISJUNCTIVE_CONSTRAINTS Rescode = 4011
    MSK_RES_ERR_INVALID_FILE_FORMAT_FOR_AFFINE_CONIC_CONSTRAINTS Rescode = 4012
    MSK_RES_ERR_DUPLICATE_CONSTRAINT_NAMES Rescode = 4500
    MSK_RES_ERR_DUPLICATE_VARIABLE_NAMES Rescode = 4501
    MSK_RES_ERR_DUPLICATE_BARVARIABLE_NAMES Rescode = 4502
    MSK_RES_ERR_DUPLICATE_CONE_NAMES Rescode = 4503
    MSK_RES_ERR_DUPLICATE_DOMAIN_NAMES Rescode = 4504
    MSK_RES_ERR_DUPLICATE_DJC_NAMES Rescode = 4505
    MSK_RES_ERR_NON_UNIQUE_ARRAY Rescode = 5000
    MSK_RES_ERR_ARGUMENT_IS_TOO_SMALL Rescode = 5004
    MSK_RES_ERR_ARGUMENT_IS_TOO_LARGE Rescode = 5005
    MSK_RES_ERR_MIO_INTERNAL Rescode = 5010
    MSK_RES_ERR_INVALID_PROBLEM_TYPE Rescode = 6000
    MSK_RES_ERR_UNHANDLED_SOLUTION_STATUS Rescode = 6010
    MSK_RES_ERR_UPPER_TRIANGLE Rescode = 6020
    MSK_RES_ERR_LAU_SINGULAR_MATRIX Rescode = 7000
    MSK_RES_ERR_LAU_NOT_POSITIVE_DEFINITE Rescode = 7001
    MSK_RES_ERR_LAU_INVALID_LOWER_TRIANGULAR_MATRIX Rescode = 7002
    MSK_RES_ERR_LAU_UNKNOWN Rescode = 7005
    MSK_RES_ERR_LAU_ARG_M Rescode = 7010
    MSK_RES_ERR_LAU_ARG_N Rescode = 7011
    MSK_RES_ERR_LAU_ARG_K Rescode = 7012
    MSK_RES_ERR_LAU_ARG_TRANSA Rescode = 7015
    MSK_RES_ERR_LAU_ARG_TRANSB Rescode = 7016
    MSK_RES_ERR_LAU_ARG_UPLO Rescode = 7017
    MSK_RES_ERR_LAU_ARG_TRANS Rescode = 7018
    MSK_RES_ERR_LAU_INVALID_SPARSE_SYMMETRIC_MATRIX Rescode = 7019
    MSK_RES_ERR_CBF_PARSE Rescode = 7100
    MSK_RES_ERR_CBF_OBJ_SENSE Rescode = 7101
    MSK_RES_ERR_CBF_NO_VARIABLES Rescode = 7102
    MSK_RES_ERR_CBF_TOO_MANY_CONSTRAINTS Rescode = 7103
    MSK_RES_ERR_CBF_TOO_MANY_VARIABLES Rescode = 7104
    MSK_RES_ERR_CBF_NO_VERSION_SPECIFIED Rescode = 7105
    MSK_RES_ERR_CBF_SYNTAX Rescode = 7106
    MSK_RES_ERR_CBF_DUPLICATE_OBJ Rescode = 7107
    MSK_RES_ERR_CBF_DUPLICATE_CON Rescode = 7108
    MSK_RES_ERR_CBF_DUPLICATE_VAR Rescode = 7110
    MSK_RES_ERR_CBF_DUPLICATE_INT Rescode = 7111
    MSK_RES_ERR_CBF_INVALID_VAR_TYPE Rescode = 7112
    MSK_RES_ERR_CBF_INVALID_CON_TYPE Rescode = 7113
    MSK_RES_ERR_CBF_INVALID_DOMAIN_DIMENSION Rescode = 7114
    MSK_RES_ERR_CBF_DUPLICATE_OBJACOORD Rescode = 7115
    MSK_RES_ERR_CBF_DUPLICATE_BCOORD Rescode = 7116
    MSK_RES_ERR_CBF_DUPLICATE_ACOORD Rescode = 7117
    MSK_RES_ERR_CBF_TOO_FEW_VARIABLES Rescode = 7118
    MSK_RES_ERR_CBF_TOO_FEW_CONSTRAINTS Rescode = 7119
    MSK_RES_ERR_CBF_TOO_FEW_INTS Rescode = 7120
    MSK_RES_ERR_CBF_TOO_MANY_INTS Rescode = 7121
    MSK_RES_ERR_CBF_INVALID_INT_INDEX Rescode = 7122
    MSK_RES_ERR_CBF_UNSUPPORTED Rescode = 7123
    MSK_RES_ERR_CBF_DUPLICATE_PSDVAR Rescode = 7124
    MSK_RES_ERR_CBF_INVALID_PSDVAR_DIMENSION Rescode = 7125
    MSK_RES_ERR_CBF_TOO_FEW_PSDVAR Rescode = 7126
    MSK_RES_ERR_CBF_INVALID_EXP_DIMENSION Rescode = 7127
    MSK_RES_ERR_CBF_DUPLICATE_POW_CONES Rescode = 7130
    MSK_RES_ERR_CBF_DUPLICATE_POW_STAR_CONES Rescode = 7131
    MSK_RES_ERR_CBF_INVALID_POWER Rescode = 7132
    MSK_RES_ERR_CBF_POWER_CONE_IS_TOO_LONG Rescode = 7133
    MSK_RES_ERR_CBF_INVALID_POWER_CONE_INDEX Rescode = 7134
    MSK_RES_ERR_CBF_INVALID_POWER_STAR_CONE_INDEX Rescode = 7135
    MSK_RES_ERR_CBF_UNHANDLED_POWER_CONE_TYPE Rescode = 7136
    MSK_RES_ERR_CBF_UNHANDLED_POWER_STAR_CONE_TYPE Rescode = 7137
    MSK_RES_ERR_CBF_POWER_CONE_MISMATCH Rescode = 7138
    MSK_RES_ERR_CBF_POWER_STAR_CONE_MISMATCH Rescode = 7139
    MSK_RES_ERR_CBF_INVALID_NUMBER_OF_CONES Rescode = 7140
    MSK_RES_ERR_CBF_INVALID_DIMENSION_OF_CONES Rescode = 7141
    MSK_RES_ERR_CBF_INVALID_NUM_OBJACOORD Rescode = 7150
    MSK_RES_ERR_CBF_INVALID_NUM_OBJFCOORD Rescode = 7151
    MSK_RES_ERR_CBF_INVALID_NUM_ACOORD Rescode = 7152
    MSK_RES_ERR_CBF_INVALID_NUM_BCOORD Rescode = 7153
    MSK_RES_ERR_CBF_INVALID_NUM_FCOORD Rescode = 7155
    MSK_RES_ERR_CBF_INVALID_NUM_HCOORD Rescode = 7156
    MSK_RES_ERR_CBF_INVALID_NUM_DCOORD Rescode = 7157
    MSK_RES_ERR_CBF_EXPECTED_A_KEYWORD Rescode = 7158
    MSK_RES_ERR_CBF_INVALID_NUM_PSDCON Rescode = 7200
    MSK_RES_ERR_CBF_DUPLICATE_PSDCON Rescode = 7201
    MSK_RES_ERR_CBF_INVALID_DIMENSION_OF_PSDCON Rescode = 7202
    MSK_RES_ERR_CBF_INVALID_PSDCON_INDEX Rescode = 7203
    MSK_RES_ERR_CBF_INVALID_PSDCON_VARIABLE_INDEX Rescode = 7204
    MSK_RES_ERR_CBF_INVALID_PSDCON_BLOCK_INDEX Rescode = 7205
    MSK_RES_ERR_CBF_UNSUPPORTED_CHANGE Rescode = 7210
    MSK_RES_ERR_MIO_INVALID_ROOT_OPTIMIZER Rescode = 7700
    MSK_RES_ERR_MIO_INVALID_NODE_OPTIMIZER Rescode = 7701
    MSK_RES_ERR_MPS_WRITE_CPLEX_INVALID_CONE_TYPE Rescode = 7750
    MSK_RES_ERR_TOCONIC_CONSTR_Q_NOT_PSD Rescode = 7800
    MSK_RES_ERR_TOCONIC_CONSTRAINT_FX Rescode = 7801
    MSK_RES_ERR_TOCONIC_CONSTRAINT_RA Rescode = 7802
    MSK_RES_ERR_TOCONIC_CONSTR_NOT_CONIC Rescode = 7803
    MSK_RES_ERR_TOCONIC_OBJECTIVE_NOT_PSD Rescode = 7804
    MSK_RES_ERR_SERVER_CONNECT Rescode = 8000
    MSK_RES_ERR_SERVER_PROTOCOL Rescode = 8001
    MSK_RES_ERR_SERVER_STATUS Rescode = 8002
    MSK_RES_ERR_SERVER_TOKEN Rescode = 8003
    MSK_RES_ERR_SERVER_ADDRESS Rescode = 8004
    MSK_RES_ERR_SERVER_CERTIFICATE Rescode = 8005
    MSK_RES_ERR_SERVER_TLS_CLIENT Rescode = 8006
    MSK_RES_ERR_SERVER_ACCESS_TOKEN Rescode = 8007
    MSK_RES_ERR_SERVER_PROBLEM_SIZE Rescode = 8008
    MSK_RES_ERR_DUPLICATE_INDEX_IN_A_SPARSE_MATRIX Rescode = 20050
    MSK_RES_ERR_DUPLICATE_INDEX_IN_AFEIDX_LIST Rescode = 20060
    MSK_RES_ERR_DUPLICATE_FIJ Rescode = 20100
    MSK_RES_ERR_INVALID_FIJ Rescode = 20101
    MSK_RES_ERR_HUGE_FIJ Rescode = 20102
    MSK_RES_ERR_INVALID_G Rescode = 20103
    MSK_RES_ERR_INVALID_B Rescode = 20150
    MSK_RES_ERR_DOMAIN_INVALID_INDEX Rescode = 20400
    MSK_RES_ERR_DOMAIN_DIMENSION Rescode = 20401
    MSK_RES_ERR_DOMAIN_DIMENSION_PSD Rescode = 20402
    MSK_RES_ERR_NOT_POWER_DOMAIN Rescode = 20403
    MSK_RES_ERR_DOMAIN_POWER_INVALID_ALPHA Rescode = 20404
    MSK_RES_ERR_DOMAIN_POWER_NEGATIVE_ALPHA Rescode = 20405
    MSK_RES_ERR_DOMAIN_POWER_NLEFT Rescode = 20406
    MSK_RES_ERR_AFE_INVALID_INDEX Rescode = 20500
    MSK_RES_ERR_ACC_INVALID_INDEX Rescode = 20600
    MSK_RES_ERR_ACC_INVALID_ENTRY_INDEX Rescode = 20601
    MSK_RES_ERR_ACC_AFE_DOMAIN_MISMATCH Rescode = 20602
    MSK_RES_ERR_DJC_INVALID_INDEX Rescode = 20700
    MSK_RES_ERR_DJC_UNSUPPORTED_DOMAIN_TYPE Rescode = 20701
    MSK_RES_ERR_DJC_AFE_DOMAIN_MISMATCH Rescode = 20702
    MSK_RES_ERR_DJC_INVALID_TERM_SIZE Rescode = 20703
    MSK_RES_ERR_DJC_DOMAIN_TERMSIZE_MISMATCH Rescode = 20704
    MSK_RES_ERR_DJC_TOTAL_NUM_TERMS_MISMATCH Rescode = 20705
    MSK_RES_ERR_UNDEF_SOLUTION Rescode = 22000
    MSK_RES_ERR_NO_DOTY Rescode = 22010
    MSK_RES_TRM_MAX_ITERATIONS Rescode = 100000
    MSK_RES_TRM_MAX_TIME Rescode = 100001
    MSK_RES_TRM_OBJECTIVE_RANGE Rescode = 100002
    MSK_RES_TRM_STALL Rescode = 100006
    MSK_RES_TRM_USER_CALLBACK Rescode = 100007
    MSK_RES_TRM_MIO_NUM_RELAXS Rescode = 100008
    MSK_RES_TRM_MIO_NUM_BRANCHES Rescode = 100009
    MSK_RES_TRM_NUM_MAX_NUM_INT_SOLUTIONS Rescode = 100015
    MSK_RES_TRM_MAX_NUM_SETBACKS Rescode = 100020
    MSK_RES_TRM_NUMERICAL_PROBLEM Rescode = 100025
    MSK_RES_TRM_LOST_RACE Rescode = 100027
    MSK_RES_TRM_INTERNAL Rescode = 100030
    MSK_RES_TRM_INTERNAL_STOP Rescode = 100031
    MSK_RES_END Rescode = 100031
)
type Rescodetype int32
const (
    MSK_RESPONSE_OK Rescodetype = 0
    MSK_RESPONSE_WRN Rescodetype = 1
    MSK_RESPONSE_TRM Rescodetype = 2
    MSK_RESPONSE_ERR Rescodetype = 3
    MSK_RESPONSE_UNK Rescodetype = 4
    MSK_RESPONSE_END Rescodetype = 4
)
type Scalingtype int32
const (
    MSK_SCALING_FREE Scalingtype = 0
    MSK_SCALING_NONE Scalingtype = 1
    MSK_SCALING_END Scalingtype = 1
)
type Scalingmethod int32
const (
    MSK_SCALING_METHOD_POW2 Scalingmethod = 0
    MSK_SCALING_METHOD_FREE Scalingmethod = 1
    MSK_SCALING_METHOD_END Scalingmethod = 1
)
type Sensitivitytype int32
const (
    MSK_SENSITIVITY_TYPE_BASIS Sensitivitytype = 0
    MSK_SENSITIVITY_TYPE_END Sensitivitytype = 0
)
type Simseltype int32
const (
    MSK_SIM_SELECTION_FREE Simseltype = 0
    MSK_SIM_SELECTION_FULL Simseltype = 1
    MSK_SIM_SELECTION_ASE Simseltype = 2
    MSK_SIM_SELECTION_DEVEX Simseltype = 3
    MSK_SIM_SELECTION_SE Simseltype = 4
    MSK_SIM_SELECTION_PARTIAL Simseltype = 5
    MSK_SIM_SELECTION_END Simseltype = 5
)
type Solitem int32
const (
    MSK_SOL_ITEM_XC Solitem = 0
    MSK_SOL_ITEM_XX Solitem = 1
    MSK_SOL_ITEM_Y Solitem = 2
    MSK_SOL_ITEM_SLC Solitem = 3
    MSK_SOL_ITEM_SUC Solitem = 4
    MSK_SOL_ITEM_SLX Solitem = 5
    MSK_SOL_ITEM_SUX Solitem = 6
    MSK_SOL_ITEM_SNX Solitem = 7
    MSK_SOL_ITEM_END Solitem = 7
)
type Solsta int32
const (
    MSK_SOL_STA_UNKNOWN Solsta = 0
    MSK_SOL_STA_OPTIMAL Solsta = 1
    MSK_SOL_STA_PRIM_FEAS Solsta = 2
    MSK_SOL_STA_DUAL_FEAS Solsta = 3
    MSK_SOL_STA_PRIM_AND_DUAL_FEAS Solsta = 4
    MSK_SOL_STA_PRIM_INFEAS_CER Solsta = 5
    MSK_SOL_STA_DUAL_INFEAS_CER Solsta = 6
    MSK_SOL_STA_PRIM_ILLPOSED_CER Solsta = 7
    MSK_SOL_STA_DUAL_ILLPOSED_CER Solsta = 8
    MSK_SOL_STA_INTEGER_OPTIMAL Solsta = 9
    MSK_SOL_STA_END Solsta = 9
)
type Soltype int32
const (
    MSK_SOL_ITR Soltype = 0
    MSK_SOL_BAS Soltype = 1
    MSK_SOL_ITG Soltype = 2
    MSK_SOL_END Soltype = 2
)
type Solveform int32
const (
    MSK_SOLVE_FREE Solveform = 0
    MSK_SOLVE_PRIMAL Solveform = 1
    MSK_SOLVE_DUAL Solveform = 2
    MSK_SOLVE_END Solveform = 2
)
type Sparam int32
const (
    MSK_SPAR_BAS_SOL_FILE_NAME Sparam = 0
    MSK_SPAR_DATA_FILE_NAME Sparam = 1
    MSK_SPAR_DEBUG_FILE_NAME Sparam = 2
    MSK_SPAR_INT_SOL_FILE_NAME Sparam = 3
    MSK_SPAR_ITR_SOL_FILE_NAME Sparam = 4
    MSK_SPAR_MIO_DEBUG_STRING Sparam = 5
    MSK_SPAR_PARAM_COMMENT_SIGN Sparam = 6
    MSK_SPAR_PARAM_READ_FILE_NAME Sparam = 7
    MSK_SPAR_PARAM_WRITE_FILE_NAME Sparam = 8
    MSK_SPAR_READ_MPS_BOU_NAME Sparam = 9
    MSK_SPAR_READ_MPS_OBJ_NAME Sparam = 10
    MSK_SPAR_READ_MPS_RAN_NAME Sparam = 11
    MSK_SPAR_READ_MPS_RHS_NAME Sparam = 12
    MSK_SPAR_REMOTE_OPTSERVER_HOST Sparam = 13
    MSK_SPAR_REMOTE_TLS_CERT Sparam = 14
    MSK_SPAR_REMOTE_TLS_CERT_PATH Sparam = 15
    MSK_SPAR_SENSITIVITY_FILE_NAME Sparam = 16
    MSK_SPAR_SENSITIVITY_RES_FILE_NAME Sparam = 17
    MSK_SPAR_SOL_FILTER_XC_LOW Sparam = 18
    MSK_SPAR_SOL_FILTER_XC_UPR Sparam = 19
    MSK_SPAR_SOL_FILTER_XX_LOW Sparam = 20
    MSK_SPAR_SOL_FILTER_XX_UPR Sparam = 21
    MSK_SPAR_STAT_KEY Sparam = 22
    MSK_SPAR_STAT_NAME Sparam = 23
    MSK_SPAR_WRITE_LP_GEN_VAR_NAME Sparam = 24
    MSK_SPAR_END Sparam = 24
)
type Stakey int32
const (
    MSK_SK_UNK Stakey = 0
    MSK_SK_BAS Stakey = 1
    MSK_SK_SUPBAS Stakey = 2
    MSK_SK_LOW Stakey = 3
    MSK_SK_UPR Stakey = 4
    MSK_SK_FIX Stakey = 5
    MSK_SK_INF Stakey = 6
    MSK_SK_END Stakey = 6
)
type Startpointtype int32
const (
    MSK_STARTING_POINT_FREE Startpointtype = 0
    MSK_STARTING_POINT_GUESS Startpointtype = 1
    MSK_STARTING_POINT_CONSTANT Startpointtype = 2
    MSK_STARTING_POINT_END Startpointtype = 2
)
type Streamtype int32
const (
    MSK_STREAM_LOG Streamtype = 0
    MSK_STREAM_MSG Streamtype = 1
    MSK_STREAM_ERR Streamtype = 2
    MSK_STREAM_WRN Streamtype = 3
    MSK_STREAM_END Streamtype = 3
)
type Value int32
const (
    MSK_LICENSE_BUFFER_LENGTH Value = 21
    MSK_MAX_STR_LEN Value = 1024
)
type Variabletype int32
const (
    MSK_VAR_TYPE_CONT Variabletype = 0
    MSK_VAR_TYPE_INT Variabletype = 1
    MSK_VAR_END Variabletype = 1
)


type MosekError struct {
    code Rescode
    msg string
}
func (self*MosekError) Error() string {
    return fmt.Sprintf("%d: %s",self.code,self.msg)
}


type ArrayLengthError struct {
    fun string
    arg string
}
func (self*ArrayLengthError) Error() string {
    return fmt.Sprintf("%s: Argument %s is too short",self.fun,self.arg)   
}


type Env struct {
        r    Rescode
        cptr C.MSKenv_t
}

type Task struct {
    cptr            C.MSKtask_t
    streamfunc      [4]func(string)
    callbackfunc    func(int32)int
    infcallbackfunc func(int32,[]float64,[]int32,[]int64)int
}

func (t * Task) ptr() C.MSKtask_t { return t.cptr }
func (e * Env)  ptr() C.MSKenv_t  { return e.cptr }

func (self * Env) getlasterror(res C.int32_t) (Rescode,string) {
    return Rescode(res),""
}
func (self * Task) getlasterror(res C.int32_t) (Rescode,string) {
    var lastcode C.int32_t = res
    var lastmsglen C.long
    if 0 != C.MSK_getlasterror64(self.ptr(),&lastcode,0, &lastmsglen, nil) {
        return Rescode(lastcode),""
    } else {
        lastmsgbytes := make([]C.char,lastmsglen+1)
        if 0 != C.MSK_getlasterror64(self.ptr(),&lastcode,lastmsglen,&lastmsglen,(*C.char)(&lastmsgbytes[0])) {
            return Rescode(lastcode),""
        } else {
            return Rescode(lastcode),C.GoString(&lastmsgbytes[0])
        }
    }
}


//export streamfunc_log
func streamfunc_log(handle unsafe.Pointer, msg *C.char) {
	task := (*Task)(handle)
	if task.streamfunc[MSK_STREAM_LOG] != nil { task.streamfunc[MSK_STREAM_LOG](C.GoString(msg)) }
}

//export streamfunc_msg
func streamfunc_msg(handle unsafe.Pointer, msg *C.char) {
	task := (*Task)(handle)
	if task.streamfunc[MSK_STREAM_MSG] != nil { task.streamfunc[MSK_STREAM_MSG](C.GoString(msg)) }
}

//export streamfunc_wrn
func streamfunc_wrn(handle unsafe.Pointer, msg *C.char) {
	task := (*Task)(handle)
	if task.streamfunc[MSK_STREAM_WRN] != nil { task.streamfunc[MSK_STREAM_WRN](C.GoString(msg)) }
}

//export streamfunc_err
func streamfunc_err(handle unsafe.Pointer, msg *C.char) {
	task := (*Task)(handle)
	if task.streamfunc[MSK_STREAM_ERR] != nil { task.streamfunc[MSK_STREAM_ERR](C.GoString(msg)) }
}

//export callbackfunc
func callbackfunc(
	nativetask unsafe.Pointer,
	handle  unsafe.Pointer,
	code    C.int,
	dinf  * C.double,
	iinf  * C.int32_t,
	liinf * C.int64_t) (C.int) {

	task := (*Task)(handle)

	var r int = 0

	if task.infcallbackfunc != nil {
		_dinf  := (*[int(MSK_DINF_END)]float64)(unsafe.Pointer(dinf))[0:MSK_DINF_END]
		_iinf  := (*[int(MSK_IINF_END)]int32)  (unsafe.Pointer(iinf))[0:MSK_IINF_END]
		_liinf := (*[int(MSK_LIINF_END)]int64) (unsafe.Pointer(liinf))[0:MSK_LIINF_END]

		r = task.infcallbackfunc(int32(code),_dinf,_iinf,_liinf)
	} else if task.callbackfunc != nil {
		r = task.callbackfunc(int32(code))
	}
	return C.int(r)
}


func MakeEnv() (env Env, err error) {
        if res := C.MSK_makeenv(&env.cptr,nil); res != 0 {
            err = &MosekError{code : Rescode(res) }
        }
        return
}

func (env *Env) MakeTask() (task Task, err error) {
    if res := C.MSK_makeemptytask(env.ptr(), &task.cptr); res != 0 {
        err = &MosekError{ code:Rescode(res) }
    } else { 
	task.streamfunc[0]   = nil
	task.streamfunc[1]   = nil
	task.streamfunc[2]   = nil
	task.streamfunc[3]   = nil
	task.callbackfunc    = nil
	task.infcallbackfunc = nil
    }
    return
}

func NewTask() (task Task, err error) {
    if res := C.MSK_makeemptytask(nil, &task.cptr); res != 0 {
        err = &MosekError{ code:Rescode(res) }
    } else { 
	task.streamfunc[0]   = nil
	task.streamfunc[1]   = nil
	task.streamfunc[2]   = nil
	task.streamfunc[3]   = nil
	task.callbackfunc    = nil
	task.infcallbackfunc = nil
    }
    return
}

func (self *Env) Delete() (err error) {
    if r := C.MSK_deleteenv(&self.cptr); r != 0 {
        err = &MosekError{code : Rescode(r) }
    }
    return 
}

func (self *Task) Delete() (err error) {
    if r := C.MSK_deletetask(&self.cptr); r != 0 {
        err = &MosekError{code : Rescode(r) }
    }
    return
}

func (self *Task) PutStreamFunc(whichstream Streamtype, fun func(string)) {
	self.streamfunc[whichstream] = fun

	if fun == nil {
		C.MSK_linkfunctotaskstream(
			self.ptr(),
			C.int32_t(whichstream),
			nil,
			nil)
	} else {
		var strmfun (*[0]byte)
		switch whichstream {
		case MSK_STREAM_MSG: strmfun = (*[0]byte)(C.streamfunc_msg)
		case MSK_STREAM_LOG: strmfun = (*[0]byte)(C.streamfunc_log)
		case MSK_STREAM_ERR: strmfun = (*[0]byte)(C.streamfunc_err)
		case MSK_STREAM_WRN: strmfun = (*[0]byte)(C.streamfunc_wrn)
		}

		C.MSK_linkfunctotaskstream(
			self.ptr(),
			C.int32_t(whichstream),
			C.MSKuserhandle_t(unsafe.Pointer(self)),
			strmfun) // ?!?
	}
}

func (self *Task) PutCallbackFunc(fun func(int32) int) {
	self.callbackfunc = fun
	if fun == nil {
		C.MSK_putcallbackfunc(self.ptr(), nil, nil)
	} else {
		C.MSK_putcallbackfunc(self.ptr(), (*[0]byte)(C.callbackfunc), C.MSKuserhandle_t(unsafe.Pointer(self)))
	}
}

func (self *Task) PutInfoCallbackFunc(fun func(int32,[]float64,[]int32,[]int64) int) {
	self.infcallbackfunc = fun
	if fun == nil {
		C.MSK_putcallbackfunc(self.ptr(), nil, nil)
	} else {
		C.MSK_putcallbackfunc(self.ptr(), (*[0]byte)(C.callbackfunc), C.MSKuserhandle_t(unsafe.Pointer(self)))
	}
}

func minint(a []int) (r int) {
        if len(a) == 0 { panic("Minimum of empty array") }
        r = a[0]
        for i := 1; i < len(a); i++ {
                if a[i] < r { r = a[i] }
        }
        return
}

func sum[T int32|int64|float64](data []T) T {
    var r T
    for _,v := range data { r += v }
    return r
}


// Methods
func (self *Task) AnalyzeNames(whichstream Streamtype,nametype Nametype) (err error) {
  if _tmp0 := C.MSK_analyzenames(self.ptr(),C.int32_t(whichstream),C.int32_t(nametype)); _tmp0 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp0)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AnalyzeProblem(whichstream Streamtype) (err error) {
  if _tmp1 := C.MSK_analyzeproblem(self.ptr(),C.int32_t(whichstream)); _tmp1 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AnalyzeSolution(whichstream Streamtype,whichsol Soltype) (err error) {
  if _tmp2 := C.MSK_analyzesolution(self.ptr(),C.int32_t(whichstream),C.int32_t(whichsol)); _tmp2 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp2)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendAcc(domidx int64,afeidxlist []int64,b []float64) (err error) {
  _tmp3 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp3)
  var _tmp4 *int64
  if afeidxlist != nil { _tmp4 = (*int64)(&afeidxlist[0]) }
  if int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"AppendAcc",arg:"b"}
    return
  }
  var _tmp5 *float64
  if b != nil { _tmp5 = (*float64)(&b[0]) }
  if _tmp6 := C.MSK_appendacc(self.ptr(),C.int64_t(domidx),C.int64_t(numafeidx),(*C.int64_t)(_tmp4),(*C.double)(_tmp5)); _tmp6 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp6)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendAccs(domidxs []int64,afeidxlist []int64,b []float64) (err error) {
  _tmp7 := len(domidxs)
  var numaccs int64 = int64(_tmp7)
  var _tmp8 *int64
  if domidxs != nil { _tmp8 = (*int64)(&domidxs[0]) }
  _tmp9 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp9)
  var _tmp10 *int64
  if afeidxlist != nil { _tmp10 = (*int64)(&afeidxlist[0]) }
  if int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"AppendAccs",arg:"b"}
    return
  }
  var _tmp11 *float64
  if b != nil { _tmp11 = (*float64)(&b[0]) }
  if _tmp12 := C.MSK_appendaccs(self.ptr(),C.int64_t(numaccs),(*C.int64_t)(_tmp8),C.int64_t(numafeidx),(*C.int64_t)(_tmp10),(*C.double)(_tmp11)); _tmp12 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp12)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendAccSeq(domidx int64,afeidxfirst int64,b []float64) (err error) {
  var _tmp13 C.int64_t
  if _tmp14 := C.MSK_getdomainn(self.ptr(),(C.int64_t)(domidx),(&_tmp13)); _tmp14 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp14)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var numafeidx int64 = int64(_tmp13)
  if int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"AppendAccSeq",arg:"b"}
    return
  }
  var _tmp15 *float64
  if b != nil { _tmp15 = (*float64)(&b[0]) }
  if _tmp16 := C.MSK_appendaccseq(self.ptr(),C.int64_t(domidx),C.int64_t(numafeidx),C.int64_t(afeidxfirst),(*C.double)(_tmp15)); _tmp16 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp16)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendAccsSeq(domidxs []int64,numafeidx int64,afeidxfirst int64,b []float64) (err error) {
  _tmp17 := len(domidxs)
  var numaccs int64 = int64(_tmp17)
  var _tmp18 *int64
  if domidxs != nil { _tmp18 = (*int64)(&domidxs[0]) }
  if int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"AppendAccsSeq",arg:"b"}
    return
  }
  var _tmp19 *float64
  if b != nil { _tmp19 = (*float64)(&b[0]) }
  if _tmp20 := C.MSK_appendaccsseq(self.ptr(),C.int64_t(numaccs),(*C.int64_t)(_tmp18),C.int64_t(numafeidx),C.int64_t(afeidxfirst),(*C.double)(_tmp19)); _tmp20 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp20)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendAfes(num int64) (err error) {
  if _tmp21 := C.MSK_appendafes(self.ptr(),C.int64_t(num)); _tmp21 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp21)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendBarvars(dim []int32) (err error) {
  _tmp22 := len(dim)
  var num int32 = int32(_tmp22)
  var _tmp23 *int32
  if dim != nil { _tmp23 = (*int32)(&dim[0]) }
  if _tmp24 := C.MSK_appendbarvars(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp23)); _tmp24 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp24)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendCone(ct Conetype,conepar float64,submem []int32) (err error) {
  _tmp25 := len(submem)
  var nummem int32 = int32(_tmp25)
  var _tmp26 *int32
  if submem != nil { _tmp26 = (*int32)(&submem[0]) }
  if _tmp27 := C.MSK_appendcone(self.ptr(),C.int32_t(ct),C.double(conepar),C.int32_t(nummem),(*C.int32_t)(_tmp26)); _tmp27 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp27)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendConeSeq(ct Conetype,conepar float64,nummem int32,j int32) (err error) {
  if _tmp28 := C.MSK_appendconeseq(self.ptr(),C.int32_t(ct),C.double(conepar),C.int32_t(nummem),C.int32_t(j)); _tmp28 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp28)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendConesSeq(ct []Conetype,conepar []float64,nummem []int32,j int32) (err error) {
  _tmp29 := len(conepar)
  if _tmp29 < len(ct) { _tmp29 = len(ct) }
  if _tmp29 < len(nummem) { _tmp29 = len(nummem) }
  var num int32 = int32(_tmp29)
  var _tmp30 *Conetype
  if ct != nil { _tmp30 = (*Conetype)(&ct[0]) }
  var _tmp31 *float64
  if conepar != nil { _tmp31 = (*float64)(&conepar[0]) }
  var _tmp32 *int32
  if nummem != nil { _tmp32 = (*int32)(&nummem[0]) }
  if _tmp33 := C.MSK_appendconesseq(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp30),(*C.double)(_tmp31),(*C.int32_t)(_tmp32),C.int32_t(j)); _tmp33 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp33)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendCons(num int32) (err error) {
  if _tmp34 := C.MSK_appendcons(self.ptr(),C.int32_t(num)); _tmp34 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp34)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendDjcs(num int64) (err error) {
  if _tmp35 := C.MSK_appenddjcs(self.ptr(),C.int64_t(num)); _tmp35 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp35)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendDualExpConeDomain() (domidx int64,err error) {
  if _tmp36 := C.MSK_appenddualexpconedomain(self.ptr(),(*C.int64_t)(&domidx)); _tmp36 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp36)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendDualGeoMeanConeDomain(n int64) (domidx int64,err error) {
  if _tmp37 := C.MSK_appenddualgeomeanconedomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp37 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp37)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendDualPowerConeDomain(n int64,alpha []float64) (domidx int64,err error) {
  _tmp38 := len(alpha)
  var nleft int64 = int64(_tmp38)
  var _tmp39 *float64
  if alpha != nil { _tmp39 = (*float64)(&alpha[0]) }
  if _tmp40 := C.MSK_appenddualpowerconedomain(self.ptr(),C.int64_t(n),C.int64_t(nleft),(*C.double)(_tmp39),(*C.int64_t)(&domidx)); _tmp40 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp40)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendPrimalExpConeDomain() (domidx int64,err error) {
  if _tmp41 := C.MSK_appendprimalexpconedomain(self.ptr(),(*C.int64_t)(&domidx)); _tmp41 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp41)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendPrimalGeoMeanConeDomain(n int64) (domidx int64,err error) {
  if _tmp42 := C.MSK_appendprimalgeomeanconedomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp42 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp42)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendPrimalPowerConeDomain(n int64,alpha []float64) (domidx int64,err error) {
  _tmp43 := len(alpha)
  var nleft int64 = int64(_tmp43)
  var _tmp44 *float64
  if alpha != nil { _tmp44 = (*float64)(&alpha[0]) }
  if _tmp45 := C.MSK_appendprimalpowerconedomain(self.ptr(),C.int64_t(n),C.int64_t(nleft),(*C.double)(_tmp44),(*C.int64_t)(&domidx)); _tmp45 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp45)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendQuadraticConeDomain(n int64) (domidx int64,err error) {
  if _tmp46 := C.MSK_appendquadraticconedomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp46 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp46)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendRDomain(n int64) (domidx int64,err error) {
  if _tmp47 := C.MSK_appendrdomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp47 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp47)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendRminusDomain(n int64) (domidx int64,err error) {
  if _tmp48 := C.MSK_appendrminusdomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp48 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp48)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendRplusDomain(n int64) (domidx int64,err error) {
  if _tmp49 := C.MSK_appendrplusdomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp49 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp49)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendRQuadraticConeDomain(n int64) (domidx int64,err error) {
  if _tmp50 := C.MSK_appendrquadraticconedomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp50 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp50)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendRzeroDomain(n int64) (domidx int64,err error) {
  if _tmp51 := C.MSK_appendrzerodomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp51 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp51)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendSparseSymMat(dim int32,subi []int32,subj []int32,valij []float64) (idx int64,err error) {
  _tmp52 := len(subi)
  if _tmp52 < len(valij) { _tmp52 = len(valij) }
  if _tmp52 < len(subj) { _tmp52 = len(subj) }
  var nz int64 = int64(_tmp52)
  var _tmp53 *int32
  if subi != nil { _tmp53 = (*int32)(&subi[0]) }
  var _tmp54 *int32
  if subj != nil { _tmp54 = (*int32)(&subj[0]) }
  var _tmp55 *float64
  if valij != nil { _tmp55 = (*float64)(&valij[0]) }
  if _tmp56 := C.MSK_appendsparsesymmat(self.ptr(),C.int32_t(dim),C.int64_t(nz),(*C.int32_t)(_tmp53),(*C.int32_t)(_tmp54),(*C.double)(_tmp55),(*C.int64_t)(&idx)); _tmp56 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp56)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendSparseSymMatList(dims []int32,nz []int64,subi []int32,subj []int32,valij []float64) (idx []int64,err error) {
  _tmp57 := len(dims)
  if _tmp57 < len(nz) { _tmp57 = len(nz) }
  var num int32 = int32(_tmp57)
  var _tmp58 *int32
  if dims != nil { _tmp58 = (*int32)(&dims[0]) }
  var _tmp59 *int64
  if nz != nil { _tmp59 = (*int64)(&nz[0]) }
  if int64(len(subi)) < int64(sum(nz)) {
    err = &ArrayLengthError{fun:"AppendSparseSymMatList",arg:"subi"}
    return
  }
  var _tmp60 *int32
  if subi != nil { _tmp60 = (*int32)(&subi[0]) }
  if int64(len(subj)) < int64(sum(nz)) {
    err = &ArrayLengthError{fun:"AppendSparseSymMatList",arg:"subj"}
    return
  }
  var _tmp61 *int32
  if subj != nil { _tmp61 = (*int32)(&subj[0]) }
  if int64(len(valij)) < int64(sum(nz)) {
    err = &ArrayLengthError{fun:"AppendSparseSymMatList",arg:"valij"}
    return
  }
  var _tmp62 *float64
  if valij != nil { _tmp62 = (*float64)(&valij[0]) }
  var _tmp63 *int64
  idx = make([]int64,num)
  if len(idx) > 0 { _tmp63 = (*int64)(&idx[0]) }
  if _tmp64 := C.MSK_appendsparsesymmatlist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp58),(*C.int64_t)(_tmp59),(*C.int32_t)(_tmp60),(*C.int32_t)(_tmp61),(*C.double)(_tmp62),(*C.int64_t)(_tmp63)); _tmp64 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp64)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendSvecPsdConeDomain(n int64) (domidx int64,err error) {
  if _tmp65 := C.MSK_appendsvecpsdconedomain(self.ptr(),C.int64_t(n),(*C.int64_t)(&domidx)); _tmp65 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp65)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendVars(num int32) (err error) {
  if _tmp66 := C.MSK_appendvars(self.ptr(),C.int32_t(num)); _tmp66 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp66)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) BasisCond() (nrmbasis float64,nrminvbasis float64,err error) {
  if _tmp67 := C.MSK_basiscond(self.ptr(),(*C.double)(&nrmbasis),(*C.double)(&nrminvbasis)); _tmp67 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp67)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) CheckMem(file string,line int32) (err error) {
  _tmp68 := C.CString(file)
  if _tmp69 := C.MSK_checkmemtask(self.ptr(),_tmp68,C.int32_t(line)); _tmp69 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp69)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) ChgConBound(i int32,lower int32,finite int32,value float64) (err error) {
  if _tmp70 := C.MSK_chgconbound(self.ptr(),C.int32_t(i),C.int32_t(lower),C.int32_t(finite),C.double(value)); _tmp70 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp70)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) ChgVarBound(j int32,lower int32,finite int32,value float64) (err error) {
  if _tmp71 := C.MSK_chgvarbound(self.ptr(),C.int32_t(j),C.int32_t(lower),C.int32_t(finite),C.double(value)); _tmp71 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp71)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) CommitChanges() (err error) {
  if _tmp72 := C.MSK_commitchanges(self.ptr()); _tmp72 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp72)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) DeleteSolution(whichsol Soltype) (err error) {
  if _tmp73 := C.MSK_deletesolution(self.ptr(),C.int32_t(whichsol)); _tmp73 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp73)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) DualSensitivity(subj []int32) (leftpricej []float64,rightpricej []float64,leftrangej []float64,rightrangej []float64,err error) {
  _tmp74 := len(subj)
  var numj int32 = int32(_tmp74)
  var _tmp75 *int32
  if subj != nil { _tmp75 = (*int32)(&subj[0]) }
  var _tmp76 *float64
  leftpricej = make([]float64,numj)
  if len(leftpricej) > 0 { _tmp76 = (*float64)(&leftpricej[0]) }
  var _tmp77 *float64
  rightpricej = make([]float64,numj)
  if len(rightpricej) > 0 { _tmp77 = (*float64)(&rightpricej[0]) }
  var _tmp78 *float64
  leftrangej = make([]float64,numj)
  if len(leftrangej) > 0 { _tmp78 = (*float64)(&leftrangej[0]) }
  var _tmp79 *float64
  rightrangej = make([]float64,numj)
  if len(rightrangej) > 0 { _tmp79 = (*float64)(&rightrangej[0]) }
  if _tmp80 := C.MSK_dualsensitivity(self.ptr(),C.int32_t(numj),(*C.int32_t)(_tmp75),(*C.double)(_tmp76),(*C.double)(_tmp77),(*C.double)(_tmp78),(*C.double)(_tmp79)); _tmp80 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp80)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) EmptyAfeBarfRow(afeidx int64) (err error) {
  if _tmp81 := C.MSK_emptyafebarfrow(self.ptr(),C.int64_t(afeidx)); _tmp81 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp81)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) EmptyAfeBarfRowList(afeidxlist []int64) (err error) {
  _tmp82 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp82)
  var _tmp83 *int64
  if afeidxlist != nil { _tmp83 = (*int64)(&afeidxlist[0]) }
  if _tmp84 := C.MSK_emptyafebarfrowlist(self.ptr(),C.int64_t(numafeidx),(*C.int64_t)(_tmp83)); _tmp84 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp84)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) EmptyAfeFCol(varidx int32) (err error) {
  if _tmp85 := C.MSK_emptyafefcol(self.ptr(),C.int32_t(varidx)); _tmp85 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp85)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) EmptyAfeFColList(varidx []int32) (err error) {
  _tmp86 := len(varidx)
  var numvaridx int64 = int64(_tmp86)
  var _tmp87 *int32
  if varidx != nil { _tmp87 = (*int32)(&varidx[0]) }
  if _tmp88 := C.MSK_emptyafefcollist(self.ptr(),C.int64_t(numvaridx),(*C.int32_t)(_tmp87)); _tmp88 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp88)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) EmptyAfeFRow(afeidx int64) (err error) {
  if _tmp89 := C.MSK_emptyafefrow(self.ptr(),C.int64_t(afeidx)); _tmp89 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp89)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) EmptyAfeFRowList(afeidx []int64) (err error) {
  _tmp90 := len(afeidx)
  var numafeidx int64 = int64(_tmp90)
  var _tmp91 *int64
  if afeidx != nil { _tmp91 = (*int64)(&afeidx[0]) }
  if _tmp92 := C.MSK_emptyafefrowlist(self.ptr(),C.int64_t(numafeidx),(*C.int64_t)(_tmp91)); _tmp92 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp92)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) EvaluateAcc(whichsol Soltype,accidx int64) (activity []float64,err error) {
  var _tmp93 C.int64_t
  if _tmp94 := C.MSK_getaccn(self.ptr(),(C.int64_t)(accidx),(&_tmp93)); _tmp94 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp94)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp95 *float64
  activity = make([]float64,_tmp93)
  if len(activity) > 0 { _tmp95 = (*float64)(&activity[0]) }
  if _tmp96 := C.MSK_evaluateacc(self.ptr(),C.int32_t(whichsol),C.int64_t(accidx),(*C.double)(_tmp95)); _tmp96 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp96)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) EvaluateAccs(whichsol Soltype) (activity []float64,err error) {
  var _tmp97 C.int64_t
  if _tmp98 := C.MSK_getaccntot(self.ptr(),(&_tmp97)); _tmp98 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp98)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp99 *float64
  activity = make([]float64,_tmp97)
  if len(activity) > 0 { _tmp99 = (*float64)(&activity[0]) }
  if _tmp100 := C.MSK_evaluateaccs(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp99)); _tmp100 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp100)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GenerateAccNames(sub []int64,fmt string,dims []int32,sp []int64,namedaxisidxs []int32,names []string) (err error) {
  _tmp101 := len(sub)
  var num int64 = int64(_tmp101)
  var _tmp102 *int64
  if sub != nil { _tmp102 = (*int64)(&sub[0]) }
  _tmp103 := C.CString(fmt)
  _tmp104 := len(dims)
  var ndims int32 = int32(_tmp104)
  var _tmp105 *int32
  if dims != nil { _tmp105 = (*int32)(&dims[0]) }
  if int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateAccNames",arg:"sp"}
    return
  }
  var _tmp106 *int64
  if sp != nil { _tmp106 = (*int64)(&sp[0]) }
  _tmp107 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp107)
  var _tmp108 *int32
  if namedaxisidxs != nil { _tmp108 = (*int32)(&namedaxisidxs[0]) }
  _tmp109 := len(names)
  var numnames int64 = int64(_tmp109)
  _tmp111 := make([]*C.char,len(names))
  for _tmp112,_tmp113 := range names {
    _tmp111[_tmp112] = C.CString(_tmp113)
  }
  var _tmp110 **C.char;
  if len(_tmp111) > 0 { _tmp110 = &_tmp111[0] }
  if _tmp114 := C.MSK_generateaccnames(self.ptr(),C.int64_t(num),(*C.int64_t)(_tmp102),_tmp103,C.int32_t(ndims),(*C.int32_t)(_tmp105),(*C.int64_t)(_tmp106),C.int32_t(numnamedaxis),(*C.int32_t)(_tmp108),C.int64_t(numnames),((**C.char))(_tmp110)); _tmp114 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp114)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GenerateBarvarNames(subj []int32,fmt string,dims []int32,sp []int64,namedaxisidxs []int32,names []string) (err error) {
  _tmp115 := len(subj)
  var num int32 = int32(_tmp115)
  var _tmp116 *int32
  if subj != nil { _tmp116 = (*int32)(&subj[0]) }
  _tmp117 := C.CString(fmt)
  _tmp118 := len(dims)
  var ndims int32 = int32(_tmp118)
  var _tmp119 *int32
  if dims != nil { _tmp119 = (*int32)(&dims[0]) }
  if int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateBarvarNames",arg:"sp"}
    return
  }
  var _tmp120 *int64
  if sp != nil { _tmp120 = (*int64)(&sp[0]) }
  _tmp121 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp121)
  var _tmp122 *int32
  if namedaxisidxs != nil { _tmp122 = (*int32)(&namedaxisidxs[0]) }
  _tmp123 := len(names)
  var numnames int64 = int64(_tmp123)
  _tmp125 := make([]*C.char,len(names))
  for _tmp126,_tmp127 := range names {
    _tmp125[_tmp126] = C.CString(_tmp127)
  }
  var _tmp124 **C.char;
  if len(_tmp125) > 0 { _tmp124 = &_tmp125[0] }
  if _tmp128 := C.MSK_generatebarvarnames(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp116),_tmp117,C.int32_t(ndims),(*C.int32_t)(_tmp119),(*C.int64_t)(_tmp120),C.int32_t(numnamedaxis),(*C.int32_t)(_tmp122),C.int64_t(numnames),((**C.char))(_tmp124)); _tmp128 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp128)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GenerateConeNames(subk []int32,fmt string,dims []int32,sp []int64,namedaxisidxs []int32,names []string) (err error) {
  _tmp129 := len(subk)
  var num int32 = int32(_tmp129)
  var _tmp130 *int32
  if subk != nil { _tmp130 = (*int32)(&subk[0]) }
  _tmp131 := C.CString(fmt)
  _tmp132 := len(dims)
  var ndims int32 = int32(_tmp132)
  var _tmp133 *int32
  if dims != nil { _tmp133 = (*int32)(&dims[0]) }
  if int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateConeNames",arg:"sp"}
    return
  }
  var _tmp134 *int64
  if sp != nil { _tmp134 = (*int64)(&sp[0]) }
  _tmp135 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp135)
  var _tmp136 *int32
  if namedaxisidxs != nil { _tmp136 = (*int32)(&namedaxisidxs[0]) }
  _tmp137 := len(names)
  var numnames int64 = int64(_tmp137)
  _tmp139 := make([]*C.char,len(names))
  for _tmp140,_tmp141 := range names {
    _tmp139[_tmp140] = C.CString(_tmp141)
  }
  var _tmp138 **C.char;
  if len(_tmp139) > 0 { _tmp138 = &_tmp139[0] }
  if _tmp142 := C.MSK_generateconenames(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp130),_tmp131,C.int32_t(ndims),(*C.int32_t)(_tmp133),(*C.int64_t)(_tmp134),C.int32_t(numnamedaxis),(*C.int32_t)(_tmp136),C.int64_t(numnames),((**C.char))(_tmp138)); _tmp142 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp142)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GenerateConNames(subi []int32,fmt string,dims []int32,sp []int64,namedaxisidxs []int32,names []string) (err error) {
  _tmp143 := len(subi)
  var num int32 = int32(_tmp143)
  var _tmp144 *int32
  if subi != nil { _tmp144 = (*int32)(&subi[0]) }
  _tmp145 := C.CString(fmt)
  _tmp146 := len(dims)
  var ndims int32 = int32(_tmp146)
  var _tmp147 *int32
  if dims != nil { _tmp147 = (*int32)(&dims[0]) }
  if int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateConNames",arg:"sp"}
    return
  }
  var _tmp148 *int64
  if sp != nil { _tmp148 = (*int64)(&sp[0]) }
  _tmp149 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp149)
  var _tmp150 *int32
  if namedaxisidxs != nil { _tmp150 = (*int32)(&namedaxisidxs[0]) }
  _tmp151 := len(names)
  var numnames int64 = int64(_tmp151)
  _tmp153 := make([]*C.char,len(names))
  for _tmp154,_tmp155 := range names {
    _tmp153[_tmp154] = C.CString(_tmp155)
  }
  var _tmp152 **C.char;
  if len(_tmp153) > 0 { _tmp152 = &_tmp153[0] }
  if _tmp156 := C.MSK_generateconnames(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp144),_tmp145,C.int32_t(ndims),(*C.int32_t)(_tmp147),(*C.int64_t)(_tmp148),C.int32_t(numnamedaxis),(*C.int32_t)(_tmp150),C.int64_t(numnames),((**C.char))(_tmp152)); _tmp156 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp156)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GenerateDjcNames(sub []int64,fmt string,dims []int32,sp []int64,namedaxisidxs []int32,names []string) (err error) {
  _tmp157 := len(sub)
  var num int64 = int64(_tmp157)
  var _tmp158 *int64
  if sub != nil { _tmp158 = (*int64)(&sub[0]) }
  _tmp159 := C.CString(fmt)
  _tmp160 := len(dims)
  var ndims int32 = int32(_tmp160)
  var _tmp161 *int32
  if dims != nil { _tmp161 = (*int32)(&dims[0]) }
  if int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateDjcNames",arg:"sp"}
    return
  }
  var _tmp162 *int64
  if sp != nil { _tmp162 = (*int64)(&sp[0]) }
  _tmp163 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp163)
  var _tmp164 *int32
  if namedaxisidxs != nil { _tmp164 = (*int32)(&namedaxisidxs[0]) }
  _tmp165 := len(names)
  var numnames int64 = int64(_tmp165)
  _tmp167 := make([]*C.char,len(names))
  for _tmp168,_tmp169 := range names {
    _tmp167[_tmp168] = C.CString(_tmp169)
  }
  var _tmp166 **C.char;
  if len(_tmp167) > 0 { _tmp166 = &_tmp167[0] }
  if _tmp170 := C.MSK_generatedjcnames(self.ptr(),C.int64_t(num),(*C.int64_t)(_tmp158),_tmp159,C.int32_t(ndims),(*C.int32_t)(_tmp161),(*C.int64_t)(_tmp162),C.int32_t(numnamedaxis),(*C.int32_t)(_tmp164),C.int64_t(numnames),((**C.char))(_tmp166)); _tmp170 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp170)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GenerateVarNames(subj []int32,fmt string,dims []int32,sp []int64,namedaxisidxs []int32,names []string) (err error) {
  _tmp171 := len(subj)
  var num int32 = int32(_tmp171)
  var _tmp172 *int32
  if subj != nil { _tmp172 = (*int32)(&subj[0]) }
  _tmp173 := C.CString(fmt)
  _tmp174 := len(dims)
  var ndims int32 = int32(_tmp174)
  var _tmp175 *int32
  if dims != nil { _tmp175 = (*int32)(&dims[0]) }
  if int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateVarNames",arg:"sp"}
    return
  }
  var _tmp176 *int64
  if sp != nil { _tmp176 = (*int64)(&sp[0]) }
  _tmp177 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp177)
  var _tmp178 *int32
  if namedaxisidxs != nil { _tmp178 = (*int32)(&namedaxisidxs[0]) }
  _tmp179 := len(names)
  var numnames int64 = int64(_tmp179)
  _tmp181 := make([]*C.char,len(names))
  for _tmp182,_tmp183 := range names {
    _tmp181[_tmp182] = C.CString(_tmp183)
  }
  var _tmp180 **C.char;
  if len(_tmp181) > 0 { _tmp180 = &_tmp181[0] }
  if _tmp184 := C.MSK_generatevarnames(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp172),_tmp173,C.int32_t(ndims),(*C.int32_t)(_tmp175),(*C.int64_t)(_tmp176),C.int32_t(numnamedaxis),(*C.int32_t)(_tmp178),C.int64_t(numnames),((**C.char))(_tmp180)); _tmp184 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp184)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccAfeIdxList(accidx int64) (afeidxlist []int64,err error) {
  var _tmp185 C.int64_t
  if _tmp186 := C.MSK_getaccn(self.ptr(),(C.int64_t)(accidx),(&_tmp185)); _tmp186 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp186)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp187 *int64
  afeidxlist = make([]int64,_tmp185)
  if len(afeidxlist) > 0 { _tmp187 = (*int64)(&afeidxlist[0]) }
  if _tmp188 := C.MSK_getaccafeidxlist(self.ptr(),C.int64_t(accidx),(*C.int64_t)(_tmp187)); _tmp188 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp188)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccB(accidx int64) (b []float64,err error) {
  var _tmp189 C.int64_t
  if _tmp190 := C.MSK_getaccn(self.ptr(),(C.int64_t)(accidx),(&_tmp189)); _tmp190 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp190)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp191 *float64
  b = make([]float64,_tmp189)
  if len(b) > 0 { _tmp191 = (*float64)(&b[0]) }
  if _tmp192 := C.MSK_getaccb(self.ptr(),C.int64_t(accidx),(*C.double)(_tmp191)); _tmp192 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp192)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccBarfBlockTriplet() (numtrip int64,acc_afe []int64,bar_var []int32,blk_row []int32,blk_col []int32,blk_val []float64,err error) {
  var _tmp193 C.int64_t
  if _tmp194 := C.MSK_getaccbarfnumblocktriplets(self.ptr(),(&_tmp193)); _tmp194 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp194)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumtrip int64 = int64(_tmp193)
  var _tmp195 *int64
  acc_afe = make([]int64,maxnumtrip)
  if len(acc_afe) > 0 { _tmp195 = (*int64)(&acc_afe[0]) }
  var _tmp196 *int32
  bar_var = make([]int32,maxnumtrip)
  if len(bar_var) > 0 { _tmp196 = (*int32)(&bar_var[0]) }
  var _tmp197 *int32
  blk_row = make([]int32,maxnumtrip)
  if len(blk_row) > 0 { _tmp197 = (*int32)(&blk_row[0]) }
  var _tmp198 *int32
  blk_col = make([]int32,maxnumtrip)
  if len(blk_col) > 0 { _tmp198 = (*int32)(&blk_col[0]) }
  var _tmp199 *float64
  blk_val = make([]float64,maxnumtrip)
  if len(blk_val) > 0 { _tmp199 = (*float64)(&blk_val[0]) }
  if _tmp200 := C.MSK_getaccbarfblocktriplet(self.ptr(),C.int64_t(maxnumtrip),(*C.int64_t)(&numtrip),(*C.int64_t)(_tmp195),(*C.int32_t)(_tmp196),(*C.int32_t)(_tmp197),(*C.int32_t)(_tmp198),(*C.double)(_tmp199)); _tmp200 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp200)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccBarfNumBlockTriplets() (numtrip int64,err error) {
  if _tmp201 := C.MSK_getaccbarfnumblocktriplets(self.ptr(),(*C.int64_t)(&numtrip)); _tmp201 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp201)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccDomain(accidx int64) (domidx int64,err error) {
  if _tmp202 := C.MSK_getaccdomain(self.ptr(),C.int64_t(accidx),(*C.int64_t)(&domidx)); _tmp202 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp202)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccDotY(whichsol Soltype,accidx int64) (doty []float64,err error) {
  var _tmp203 C.int64_t
  if _tmp204 := C.MSK_getaccn(self.ptr(),(C.int64_t)(accidx),(&_tmp203)); _tmp204 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp204)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp205 *float64
  doty = make([]float64,_tmp203)
  if len(doty) > 0 { _tmp205 = (*float64)(&doty[0]) }
  if _tmp206 := C.MSK_getaccdoty(self.ptr(),C.int32_t(whichsol),C.int64_t(accidx),(*C.double)(_tmp205)); _tmp206 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp206)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccDotYS(whichsol Soltype) (doty []float64,err error) {
  var _tmp207 C.int64_t
  if _tmp208 := C.MSK_getaccntot(self.ptr(),(&_tmp207)); _tmp208 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp208)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp209 *float64
  doty = make([]float64,_tmp207)
  if len(doty) > 0 { _tmp209 = (*float64)(&doty[0]) }
  if _tmp210 := C.MSK_getaccdotys(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp209)); _tmp210 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp210)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccFNumnz() (accfnnz int64,err error) {
  if _tmp211 := C.MSK_getaccfnumnz(self.ptr(),(*C.int64_t)(&accfnnz)); _tmp211 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp211)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccFTrip() (frow []int64,fcol []int32,fval []float64,err error) {
  var _tmp212 C.int64_t
  if _tmp213 := C.MSK_getaccfnumnz(self.ptr(),(&_tmp212)); _tmp213 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp213)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp214 *int64
  frow = make([]int64,_tmp212)
  if len(frow) > 0 { _tmp214 = (*int64)(&frow[0]) }
  var _tmp215 C.int64_t
  if _tmp216 := C.MSK_getaccfnumnz(self.ptr(),(&_tmp215)); _tmp216 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp216)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp217 *int32
  fcol = make([]int32,_tmp215)
  if len(fcol) > 0 { _tmp217 = (*int32)(&fcol[0]) }
  var _tmp218 C.int64_t
  if _tmp219 := C.MSK_getaccfnumnz(self.ptr(),(&_tmp218)); _tmp219 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp219)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp220 *float64
  fval = make([]float64,_tmp218)
  if len(fval) > 0 { _tmp220 = (*float64)(&fval[0]) }
  if _tmp221 := C.MSK_getaccftrip(self.ptr(),(*C.int64_t)(_tmp214),(*C.int32_t)(_tmp217),(*C.double)(_tmp220)); _tmp221 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp221)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccGVector() (g []float64,err error) {
  var _tmp222 C.int64_t
  if _tmp223 := C.MSK_getaccntot(self.ptr(),(&_tmp222)); _tmp223 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp223)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp224 *float64
  g = make([]float64,_tmp222)
  if len(g) > 0 { _tmp224 = (*float64)(&g[0]) }
  if _tmp225 := C.MSK_getaccgvector(self.ptr(),(*C.double)(_tmp224)); _tmp225 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp225)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccN(accidx int64) (n int64,err error) {
  if _tmp226 := C.MSK_getaccn(self.ptr(),C.int64_t(accidx),(*C.int64_t)(&n)); _tmp226 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp226)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccName(accidx int64) (name string,err error) {
  var _tmp227 C.int32_t
  if _tmp228 := C.MSK_getaccnamelen(self.ptr(),(C.int64_t)(accidx),(&_tmp227)); _tmp228 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp228)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp227))
  _tmp229 := make([]byte,sizename)
  if _tmp230 := C.MSK_getaccname(self.ptr(),C.int64_t(accidx),C.int32_t(sizename),(*C.uchar)(&_tmp229[0])); _tmp230 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp230)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp229,byte(0)); p < 0 {
    name = string(_tmp229)
  } else {
    name = string(_tmp229[:p])
  }
  return
}
func (self *Task) GetAccNameLen(accidx int64) (len int32,err error) {
  if _tmp231 := C.MSK_getaccnamelen(self.ptr(),C.int64_t(accidx),(*C.int32_t)(&len)); _tmp231 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp231)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccNTot() (n int64,err error) {
  if _tmp232 := C.MSK_getaccntot(self.ptr(),(*C.int64_t)(&n)); _tmp232 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp232)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccs() (domidxlist []int64,afeidxlist []int64,b []float64,err error) {
  var _tmp233 C.int64_t
  if _tmp234 := C.MSK_getnumacc(self.ptr(),(&_tmp233)); _tmp234 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp234)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp235 *int64
  domidxlist = make([]int64,_tmp233)
  if len(domidxlist) > 0 { _tmp235 = (*int64)(&domidxlist[0]) }
  var _tmp236 C.int64_t
  if _tmp237 := C.MSK_getaccntot(self.ptr(),(&_tmp236)); _tmp237 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp237)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp238 *int64
  afeidxlist = make([]int64,_tmp236)
  if len(afeidxlist) > 0 { _tmp238 = (*int64)(&afeidxlist[0]) }
  var _tmp239 C.int64_t
  if _tmp240 := C.MSK_getaccntot(self.ptr(),(&_tmp239)); _tmp240 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp240)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp241 *float64
  b = make([]float64,_tmp239)
  if len(b) > 0 { _tmp241 = (*float64)(&b[0]) }
  if _tmp242 := C.MSK_getaccs(self.ptr(),(*C.int64_t)(_tmp235),(*C.int64_t)(_tmp238),(*C.double)(_tmp241)); _tmp242 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp242)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetACol(j int32) (nzj int32,subj []int32,valj []float64,err error) {
  var _tmp243 C.int32_t
  if _tmp244 := C.MSK_getacolnumnz(self.ptr(),(C.int32_t)(j),(&_tmp243)); _tmp244 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp244)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp245 *int32
  subj = make([]int32,_tmp243)
  if len(subj) > 0 { _tmp245 = (*int32)(&subj[0]) }
  var _tmp246 C.int32_t
  if _tmp247 := C.MSK_getacolnumnz(self.ptr(),(C.int32_t)(j),(&_tmp246)); _tmp247 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp247)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp248 *float64
  valj = make([]float64,_tmp246)
  if len(valj) > 0 { _tmp248 = (*float64)(&valj[0]) }
  if _tmp249 := C.MSK_getacol(self.ptr(),C.int32_t(j),(*C.int32_t)(&nzj),(*C.int32_t)(_tmp245),(*C.double)(_tmp248)); _tmp249 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp249)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAColNumNz(i int32) (nzj int32,err error) {
  if _tmp250 := C.MSK_getacolnumnz(self.ptr(),C.int32_t(i),(*C.int32_t)(&nzj)); _tmp250 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp250)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAColSlice(first int32,last int32) (ptrb []int64,ptre []int64,sub []int32,val []float64,err error) {
  var _tmp251 C.int64_t
  if _tmp252 := C.MSK_getacolslicenumnz64(self.ptr(),(C.int32_t)(first),(C.int32_t)(last),(&_tmp251)); _tmp252 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp252)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp251)
  var _tmp253 *int64
  ptrb = make([]int64,(last - first))
  if len(ptrb) > 0 { _tmp253 = (*int64)(&ptrb[0]) }
  var _tmp254 *int64
  ptre = make([]int64,(last - first))
  if len(ptre) > 0 { _tmp254 = (*int64)(&ptre[0]) }
  var _tmp255 *int32
  sub = make([]int32,maxnumnz)
  if len(sub) > 0 { _tmp255 = (*int32)(&sub[0]) }
  var _tmp256 *float64
  val = make([]float64,maxnumnz)
  if len(val) > 0 { _tmp256 = (*float64)(&val[0]) }
  if _tmp257 := C.MSK_getacolslice64(self.ptr(),C.int32_t(first),C.int32_t(last),C.int64_t(maxnumnz),(*C.int64_t)(_tmp253),(*C.int64_t)(_tmp254),(*C.int32_t)(_tmp255),(*C.double)(_tmp256)); _tmp257 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp257)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAColSliceNumNz(first int32,last int32) (numnz int64,err error) {
  if _tmp258 := C.MSK_getacolslicenumnz64(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.int64_t)(&numnz)); _tmp258 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp258)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAColSliceTrip(first int32,last int32) (subi []int32,subj []int32,val []float64,err error) {
  var _tmp259 C.int64_t
  if _tmp260 := C.MSK_getacolslicenumnz64(self.ptr(),(C.int32_t)(first),(C.int32_t)(last),(&_tmp259)); _tmp260 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp260)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp259)
  var _tmp261 *int32
  subi = make([]int32,maxnumnz)
  if len(subi) > 0 { _tmp261 = (*int32)(&subi[0]) }
  var _tmp262 *int32
  subj = make([]int32,maxnumnz)
  if len(subj) > 0 { _tmp262 = (*int32)(&subj[0]) }
  var _tmp263 *float64
  val = make([]float64,maxnumnz)
  if len(val) > 0 { _tmp263 = (*float64)(&val[0]) }
  if _tmp264 := C.MSK_getacolslicetrip(self.ptr(),C.int32_t(first),C.int32_t(last),C.int64_t(maxnumnz),(*C.int32_t)(_tmp261),(*C.int32_t)(_tmp262),(*C.double)(_tmp263)); _tmp264 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp264)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeBarfBlockTriplet() (numtrip int64,afeidx []int64,barvaridx []int32,subk []int32,subl []int32,valkl []float64,err error) {
  var _tmp265 C.int64_t
  if _tmp266 := C.MSK_getafebarfnumblocktriplets(self.ptr(),(&_tmp265)); _tmp266 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp266)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumtrip int64 = int64(_tmp265)
  var _tmp267 *int64
  afeidx = make([]int64,maxnumtrip)
  if len(afeidx) > 0 { _tmp267 = (*int64)(&afeidx[0]) }
  var _tmp268 *int32
  barvaridx = make([]int32,maxnumtrip)
  if len(barvaridx) > 0 { _tmp268 = (*int32)(&barvaridx[0]) }
  var _tmp269 *int32
  subk = make([]int32,maxnumtrip)
  if len(subk) > 0 { _tmp269 = (*int32)(&subk[0]) }
  var _tmp270 *int32
  subl = make([]int32,maxnumtrip)
  if len(subl) > 0 { _tmp270 = (*int32)(&subl[0]) }
  var _tmp271 *float64
  valkl = make([]float64,maxnumtrip)
  if len(valkl) > 0 { _tmp271 = (*float64)(&valkl[0]) }
  if _tmp272 := C.MSK_getafebarfblocktriplet(self.ptr(),C.int64_t(maxnumtrip),(*C.int64_t)(&numtrip),(*C.int64_t)(_tmp267),(*C.int32_t)(_tmp268),(*C.int32_t)(_tmp269),(*C.int32_t)(_tmp270),(*C.double)(_tmp271)); _tmp272 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp272)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeBarfNumBlockTriplets() (numtrip int64,err error) {
  if _tmp273 := C.MSK_getafebarfnumblocktriplets(self.ptr(),(*C.int64_t)(&numtrip)); _tmp273 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp273)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeBarfNumRowEntries(afeidx int64) (numentr int32,err error) {
  if _tmp274 := C.MSK_getafebarfnumrowentries(self.ptr(),C.int64_t(afeidx),(*C.int32_t)(&numentr)); _tmp274 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp274)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeBarfRow(afeidx int64) (barvaridx []int32,ptrterm []int64,numterm []int64,termidx []int64,termweight []float64,err error) {
  var _tmp275 C.int32_t
  var _tmp276 C.int64_t
  if _tmp277 := C.MSK_getafebarfrowinfo(self.ptr(),(C.int64_t)(afeidx),(&_tmp275),(&_tmp276)); _tmp277 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp277)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp278 *int32
  barvaridx = make([]int32,_tmp275)
  if len(barvaridx) > 0 { _tmp278 = (*int32)(&barvaridx[0]) }
  var _tmp279 C.int32_t
  var _tmp280 C.int64_t
  if _tmp281 := C.MSK_getafebarfrowinfo(self.ptr(),(C.int64_t)(afeidx),(&_tmp279),(&_tmp280)); _tmp281 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp281)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp282 *int64
  ptrterm = make([]int64,_tmp279)
  if len(ptrterm) > 0 { _tmp282 = (*int64)(&ptrterm[0]) }
  var _tmp283 C.int32_t
  var _tmp284 C.int64_t
  if _tmp285 := C.MSK_getafebarfrowinfo(self.ptr(),(C.int64_t)(afeidx),(&_tmp283),(&_tmp284)); _tmp285 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp285)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp286 *int64
  numterm = make([]int64,_tmp283)
  if len(numterm) > 0 { _tmp286 = (*int64)(&numterm[0]) }
  var _tmp287 C.int32_t
  var _tmp288 C.int64_t
  if _tmp289 := C.MSK_getafebarfrowinfo(self.ptr(),(C.int64_t)(afeidx),(&_tmp287),(&_tmp288)); _tmp289 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp289)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp290 *int64
  termidx = make([]int64,_tmp288)
  if len(termidx) > 0 { _tmp290 = (*int64)(&termidx[0]) }
  var _tmp291 C.int32_t
  var _tmp292 C.int64_t
  if _tmp293 := C.MSK_getafebarfrowinfo(self.ptr(),(C.int64_t)(afeidx),(&_tmp291),(&_tmp292)); _tmp293 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp293)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp294 *float64
  termweight = make([]float64,_tmp292)
  if len(termweight) > 0 { _tmp294 = (*float64)(&termweight[0]) }
  if _tmp295 := C.MSK_getafebarfrow(self.ptr(),C.int64_t(afeidx),(*C.int32_t)(_tmp278),(*C.int64_t)(_tmp282),(*C.int64_t)(_tmp286),(*C.int64_t)(_tmp290),(*C.double)(_tmp294)); _tmp295 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp295)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeBarfRowInfo(afeidx int64) (numentr int32,numterm int64,err error) {
  if _tmp296 := C.MSK_getafebarfrowinfo(self.ptr(),C.int64_t(afeidx),(*C.int32_t)(&numentr),(*C.int64_t)(&numterm)); _tmp296 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp296)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeFNumNz() (numnz int64,err error) {
  if _tmp297 := C.MSK_getafefnumnz(self.ptr(),(*C.int64_t)(&numnz)); _tmp297 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp297)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeFRow(afeidx int64) (numnz int32,varidx []int32,val []float64,err error) {
  var _tmp298 C.int32_t
  if _tmp299 := C.MSK_getafefrownumnz(self.ptr(),(C.int64_t)(afeidx),(&_tmp298)); _tmp299 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp299)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp300 *int32
  varidx = make([]int32,_tmp298)
  if len(varidx) > 0 { _tmp300 = (*int32)(&varidx[0]) }
  var _tmp301 C.int32_t
  if _tmp302 := C.MSK_getafefrownumnz(self.ptr(),(C.int64_t)(afeidx),(&_tmp301)); _tmp302 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp302)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp303 *float64
  val = make([]float64,_tmp301)
  if len(val) > 0 { _tmp303 = (*float64)(&val[0]) }
  if _tmp304 := C.MSK_getafefrow(self.ptr(),C.int64_t(afeidx),(*C.int32_t)(&numnz),(*C.int32_t)(_tmp300),(*C.double)(_tmp303)); _tmp304 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp304)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeFRowNumNz(afeidx int64) (numnz int32,err error) {
  if _tmp305 := C.MSK_getafefrownumnz(self.ptr(),C.int64_t(afeidx),(*C.int32_t)(&numnz)); _tmp305 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp305)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeFTrip() (afeidx []int64,varidx []int32,val []float64,err error) {
  var _tmp306 C.int64_t
  if _tmp307 := C.MSK_getafefnumnz(self.ptr(),(&_tmp306)); _tmp307 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp307)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp308 *int64
  afeidx = make([]int64,_tmp306)
  if len(afeidx) > 0 { _tmp308 = (*int64)(&afeidx[0]) }
  var _tmp309 C.int64_t
  if _tmp310 := C.MSK_getafefnumnz(self.ptr(),(&_tmp309)); _tmp310 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp310)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp311 *int32
  varidx = make([]int32,_tmp309)
  if len(varidx) > 0 { _tmp311 = (*int32)(&varidx[0]) }
  var _tmp312 C.int64_t
  if _tmp313 := C.MSK_getafefnumnz(self.ptr(),(&_tmp312)); _tmp313 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp313)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp314 *float64
  val = make([]float64,_tmp312)
  if len(val) > 0 { _tmp314 = (*float64)(&val[0]) }
  if _tmp315 := C.MSK_getafeftrip(self.ptr(),(*C.int64_t)(_tmp308),(*C.int32_t)(_tmp311),(*C.double)(_tmp314)); _tmp315 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp315)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeG(afeidx int64) (g float64,err error) {
  if _tmp316 := C.MSK_getafeg(self.ptr(),C.int64_t(afeidx),(*C.double)(&g)); _tmp316 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp316)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeGSlice(first int64,last int64) (g []float64,err error) {
  var _tmp317 *float64
  g = make([]float64,(last - first))
  if len(g) > 0 { _tmp317 = (*float64)(&g[0]) }
  if _tmp318 := C.MSK_getafegslice(self.ptr(),C.int64_t(first),C.int64_t(last),(*C.double)(_tmp317)); _tmp318 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp318)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAij(i int32,j int32) (aij float64,err error) {
  if _tmp319 := C.MSK_getaij(self.ptr(),C.int32_t(i),C.int32_t(j),(*C.double)(&aij)); _tmp319 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp319)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAPieceNumNz(firsti int32,lasti int32,firstj int32,lastj int32) (numnz int32,err error) {
  if _tmp320 := C.MSK_getapiecenumnz(self.ptr(),C.int32_t(firsti),C.int32_t(lasti),C.int32_t(firstj),C.int32_t(lastj),(*C.int32_t)(&numnz)); _tmp320 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp320)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetARow(i int32) (nzi int32,subi []int32,vali []float64,err error) {
  var _tmp321 C.int32_t
  if _tmp322 := C.MSK_getarownumnz(self.ptr(),(C.int32_t)(i),(&_tmp321)); _tmp322 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp322)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp323 *int32
  subi = make([]int32,_tmp321)
  if len(subi) > 0 { _tmp323 = (*int32)(&subi[0]) }
  var _tmp324 C.int32_t
  if _tmp325 := C.MSK_getarownumnz(self.ptr(),(C.int32_t)(i),(&_tmp324)); _tmp325 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp325)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp326 *float64
  vali = make([]float64,_tmp324)
  if len(vali) > 0 { _tmp326 = (*float64)(&vali[0]) }
  if _tmp327 := C.MSK_getarow(self.ptr(),C.int32_t(i),(*C.int32_t)(&nzi),(*C.int32_t)(_tmp323),(*C.double)(_tmp326)); _tmp327 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp327)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetARowNumNz(i int32) (nzi int32,err error) {
  if _tmp328 := C.MSK_getarownumnz(self.ptr(),C.int32_t(i),(*C.int32_t)(&nzi)); _tmp328 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp328)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetARowSlice(first int32,last int32) (ptrb []int64,ptre []int64,sub []int32,val []float64,err error) {
  var _tmp329 C.int64_t
  if _tmp330 := C.MSK_getarowslicenumnz64(self.ptr(),(C.int32_t)(first),(C.int32_t)(last),(&_tmp329)); _tmp330 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp330)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp329)
  var _tmp331 *int64
  ptrb = make([]int64,(last - first))
  if len(ptrb) > 0 { _tmp331 = (*int64)(&ptrb[0]) }
  var _tmp332 *int64
  ptre = make([]int64,(last - first))
  if len(ptre) > 0 { _tmp332 = (*int64)(&ptre[0]) }
  var _tmp333 *int32
  sub = make([]int32,maxnumnz)
  if len(sub) > 0 { _tmp333 = (*int32)(&sub[0]) }
  var _tmp334 *float64
  val = make([]float64,maxnumnz)
  if len(val) > 0 { _tmp334 = (*float64)(&val[0]) }
  if _tmp335 := C.MSK_getarowslice64(self.ptr(),C.int32_t(first),C.int32_t(last),C.int64_t(maxnumnz),(*C.int64_t)(_tmp331),(*C.int64_t)(_tmp332),(*C.int32_t)(_tmp333),(*C.double)(_tmp334)); _tmp335 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp335)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetARowSliceNumNz(first int32,last int32) (numnz int64,err error) {
  if _tmp336 := C.MSK_getarowslicenumnz64(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.int64_t)(&numnz)); _tmp336 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp336)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetARowSliceTrip(first int32,last int32) (subi []int32,subj []int32,val []float64,err error) {
  var _tmp337 C.int64_t
  if _tmp338 := C.MSK_getarowslicenumnz64(self.ptr(),(C.int32_t)(first),(C.int32_t)(last),(&_tmp337)); _tmp338 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp338)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp337)
  var _tmp339 *int32
  subi = make([]int32,maxnumnz)
  if len(subi) > 0 { _tmp339 = (*int32)(&subi[0]) }
  var _tmp340 *int32
  subj = make([]int32,maxnumnz)
  if len(subj) > 0 { _tmp340 = (*int32)(&subj[0]) }
  var _tmp341 *float64
  val = make([]float64,maxnumnz)
  if len(val) > 0 { _tmp341 = (*float64)(&val[0]) }
  if _tmp342 := C.MSK_getarowslicetrip(self.ptr(),C.int32_t(first),C.int32_t(last),C.int64_t(maxnumnz),(*C.int32_t)(_tmp339),(*C.int32_t)(_tmp340),(*C.double)(_tmp341)); _tmp342 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp342)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetATrip() (subi []int32,subj []int32,val []float64,err error) {
  var _tmp343 C.int64_t
  if _tmp344 := C.MSK_getnumanz64(self.ptr(),(&_tmp343)); _tmp344 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp344)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp343)
  var _tmp345 *int32
  subi = make([]int32,maxnumnz)
  if len(subi) > 0 { _tmp345 = (*int32)(&subi[0]) }
  var _tmp346 *int32
  subj = make([]int32,maxnumnz)
  if len(subj) > 0 { _tmp346 = (*int32)(&subj[0]) }
  var _tmp347 *float64
  val = make([]float64,maxnumnz)
  if len(val) > 0 { _tmp347 = (*float64)(&val[0]) }
  if _tmp348 := C.MSK_getatrip(self.ptr(),C.int64_t(maxnumnz),(*C.int32_t)(_tmp345),(*C.int32_t)(_tmp346),(*C.double)(_tmp347)); _tmp348 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp348)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetATruncateTol() (tolzero []float64,err error) {
  var _tmp349 *float64
  tolzero = make([]float64,1)
  if len(tolzero) > 0 { _tmp349 = (*float64)(&tolzero[0]) }
  if _tmp350 := C.MSK_getatruncatetol(self.ptr(),(*C.double)(_tmp349)); _tmp350 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp350)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBaraBlockTriplet() (num int64,subi []int32,subj []int32,subk []int32,subl []int32,valijkl []float64,err error) {
  var _tmp351 C.int64_t
  if _tmp352 := C.MSK_getnumbarablocktriplets(self.ptr(),(&_tmp351)); _tmp352 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp352)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnum int64 = int64(_tmp351)
  var _tmp353 *int32
  subi = make([]int32,maxnum)
  if len(subi) > 0 { _tmp353 = (*int32)(&subi[0]) }
  var _tmp354 *int32
  subj = make([]int32,maxnum)
  if len(subj) > 0 { _tmp354 = (*int32)(&subj[0]) }
  var _tmp355 *int32
  subk = make([]int32,maxnum)
  if len(subk) > 0 { _tmp355 = (*int32)(&subk[0]) }
  var _tmp356 *int32
  subl = make([]int32,maxnum)
  if len(subl) > 0 { _tmp356 = (*int32)(&subl[0]) }
  var _tmp357 *float64
  valijkl = make([]float64,maxnum)
  if len(valijkl) > 0 { _tmp357 = (*float64)(&valijkl[0]) }
  if _tmp358 := C.MSK_getbarablocktriplet(self.ptr(),C.int64_t(maxnum),(*C.int64_t)(&num),(*C.int32_t)(_tmp353),(*C.int32_t)(_tmp354),(*C.int32_t)(_tmp355),(*C.int32_t)(_tmp356),(*C.double)(_tmp357)); _tmp358 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp358)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBaraIdx(idx int64) (i int32,j int32,num int64,sub []int64,weights []float64,err error) {
  var _tmp359 C.int64_t
  if _tmp360 := C.MSK_getbaraidxinfo(self.ptr(),(C.int64_t)(idx),(&_tmp359)); _tmp360 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp360)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnum int64 = int64(_tmp359)
  var _tmp361 *int64
  sub = make([]int64,maxnum)
  if len(sub) > 0 { _tmp361 = (*int64)(&sub[0]) }
  var _tmp362 *float64
  weights = make([]float64,maxnum)
  if len(weights) > 0 { _tmp362 = (*float64)(&weights[0]) }
  if _tmp363 := C.MSK_getbaraidx(self.ptr(),C.int64_t(idx),C.int64_t(maxnum),(*C.int32_t)(&i),(*C.int32_t)(&j),(*C.int64_t)(&num),(*C.int64_t)(_tmp361),(*C.double)(_tmp362)); _tmp363 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp363)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBaraIdxIJ(idx int64) (i int32,j int32,err error) {
  if _tmp364 := C.MSK_getbaraidxij(self.ptr(),C.int64_t(idx),(*C.int32_t)(&i),(*C.int32_t)(&j)); _tmp364 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp364)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBaraIdxInfo(idx int64) (num int64,err error) {
  if _tmp365 := C.MSK_getbaraidxinfo(self.ptr(),C.int64_t(idx),(*C.int64_t)(&num)); _tmp365 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp365)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBaraSparsity() (numnz int64,idxij []int64,err error) {
  var _tmp366 C.int64_t
  if _tmp367 := C.MSK_getnumbaranz(self.ptr(),(&_tmp366)); _tmp367 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp367)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp366)
  var _tmp368 *int64
  idxij = make([]int64,maxnumnz)
  if len(idxij) > 0 { _tmp368 = (*int64)(&idxij[0]) }
  if _tmp369 := C.MSK_getbarasparsity(self.ptr(),C.int64_t(maxnumnz),(*C.int64_t)(&numnz),(*C.int64_t)(_tmp368)); _tmp369 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp369)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarcBlockTriplet() (num int64,subj []int32,subk []int32,subl []int32,valjkl []float64,err error) {
  var _tmp370 C.int64_t
  if _tmp371 := C.MSK_getnumbarcblocktriplets(self.ptr(),(&_tmp370)); _tmp371 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp371)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnum int64 = int64(_tmp370)
  var _tmp372 *int32
  subj = make([]int32,maxnum)
  if len(subj) > 0 { _tmp372 = (*int32)(&subj[0]) }
  var _tmp373 *int32
  subk = make([]int32,maxnum)
  if len(subk) > 0 { _tmp373 = (*int32)(&subk[0]) }
  var _tmp374 *int32
  subl = make([]int32,maxnum)
  if len(subl) > 0 { _tmp374 = (*int32)(&subl[0]) }
  var _tmp375 *float64
  valjkl = make([]float64,maxnum)
  if len(valjkl) > 0 { _tmp375 = (*float64)(&valjkl[0]) }
  if _tmp376 := C.MSK_getbarcblocktriplet(self.ptr(),C.int64_t(maxnum),(*C.int64_t)(&num),(*C.int32_t)(_tmp372),(*C.int32_t)(_tmp373),(*C.int32_t)(_tmp374),(*C.double)(_tmp375)); _tmp376 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp376)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarcIdx(idx int64) (j int32,num int64,sub []int64,weights []float64,err error) {
  var _tmp377 C.int64_t
  if _tmp378 := C.MSK_getbarcidxinfo(self.ptr(),(C.int64_t)(idx),(&_tmp377)); _tmp378 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp378)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnum int64 = int64(_tmp377)
  var _tmp379 *int64
  sub = make([]int64,maxnum)
  if len(sub) > 0 { _tmp379 = (*int64)(&sub[0]) }
  var _tmp380 *float64
  weights = make([]float64,maxnum)
  if len(weights) > 0 { _tmp380 = (*float64)(&weights[0]) }
  if _tmp381 := C.MSK_getbarcidx(self.ptr(),C.int64_t(idx),C.int64_t(maxnum),(*C.int32_t)(&j),(*C.int64_t)(&num),(*C.int64_t)(_tmp379),(*C.double)(_tmp380)); _tmp381 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp381)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarcIdxInfo(idx int64) (num int64,err error) {
  if _tmp382 := C.MSK_getbarcidxinfo(self.ptr(),C.int64_t(idx),(*C.int64_t)(&num)); _tmp382 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp382)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarcIdxJ(idx int64) (j int32,err error) {
  if _tmp383 := C.MSK_getbarcidxj(self.ptr(),C.int64_t(idx),(*C.int32_t)(&j)); _tmp383 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp383)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarcSparsity() (numnz int64,idxj []int64,err error) {
  var _tmp384 C.int64_t
  if _tmp385 := C.MSK_getnumbarcnz(self.ptr(),(&_tmp384)); _tmp385 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp385)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp384)
  var _tmp386 *int64
  idxj = make([]int64,maxnumnz)
  if len(idxj) > 0 { _tmp386 = (*int64)(&idxj[0]) }
  if _tmp387 := C.MSK_getbarcsparsity(self.ptr(),C.int64_t(maxnumnz),(*C.int64_t)(&numnz),(*C.int64_t)(_tmp386)); _tmp387 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp387)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarsJ(whichsol Soltype,j int32) (barsj []float64,err error) {
  var _tmp388 C.int64_t
  if _tmp389 := C.MSK_getlenbarvarj(self.ptr(),(C.int32_t)(j),(&_tmp388)); _tmp389 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp389)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp390 *float64
  barsj = make([]float64,_tmp388)
  if len(barsj) > 0 { _tmp390 = (*float64)(&barsj[0]) }
  if _tmp391 := C.MSK_getbarsj(self.ptr(),C.int32_t(whichsol),C.int32_t(j),(*C.double)(_tmp390)); _tmp391 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp391)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarsSlice(whichsol Soltype,first int32,last int32,slicesize int64) (barsslice []float64,err error) {
  var _tmp392 *float64
  barsslice = make([]float64,slicesize)
  if len(barsslice) > 0 { _tmp392 = (*float64)(&barsslice[0]) }
  if _tmp393 := C.MSK_getbarsslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),C.int64_t(slicesize),(*C.double)(_tmp392)); _tmp393 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp393)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarvarName(i int32) (name string,err error) {
  var _tmp394 C.int32_t
  if _tmp395 := C.MSK_getbarvarnamelen(self.ptr(),(C.int32_t)(i),(&_tmp394)); _tmp395 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp395)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp394))
  _tmp396 := make([]byte,sizename)
  if _tmp397 := C.MSK_getbarvarname(self.ptr(),C.int32_t(i),C.int32_t(sizename),(*C.uchar)(&_tmp396[0])); _tmp397 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp397)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp396,byte(0)); p < 0 {
    name = string(_tmp396)
  } else {
    name = string(_tmp396[:p])
  }
  return
}
func (self *Task) GetBarvarNameIndex(somename string) (asgn int32,index int32,err error) {
  _tmp398 := C.CString(somename)
  if _tmp399 := C.MSK_getbarvarnameindex(self.ptr(),_tmp398,(*C.int32_t)(&asgn),(*C.int32_t)(&index)); _tmp399 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp399)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarvarNameLen(i int32) (len int32,err error) {
  if _tmp400 := C.MSK_getbarvarnamelen(self.ptr(),C.int32_t(i),(*C.int32_t)(&len)); _tmp400 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp400)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarxJ(whichsol Soltype,j int32) (barxj []float64,err error) {
  var _tmp401 C.int64_t
  if _tmp402 := C.MSK_getlenbarvarj(self.ptr(),(C.int32_t)(j),(&_tmp401)); _tmp402 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp402)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp403 *float64
  barxj = make([]float64,_tmp401)
  if len(barxj) > 0 { _tmp403 = (*float64)(&barxj[0]) }
  if _tmp404 := C.MSK_getbarxj(self.ptr(),C.int32_t(whichsol),C.int32_t(j),(*C.double)(_tmp403)); _tmp404 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp404)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarxSlice(whichsol Soltype,first int32,last int32,slicesize int64) (barxslice []float64,err error) {
  var _tmp405 *float64
  barxslice = make([]float64,slicesize)
  if len(barxslice) > 0 { _tmp405 = (*float64)(&barxslice[0]) }
  if _tmp406 := C.MSK_getbarxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),C.int64_t(slicesize),(*C.double)(_tmp405)); _tmp406 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp406)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetC() (c []float64,err error) {
  var _tmp407 C.int32_t
  if _tmp408 := C.MSK_getnumvar(self.ptr(),(&_tmp407)); _tmp408 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp408)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp409 *float64
  c = make([]float64,_tmp407)
  if len(c) > 0 { _tmp409 = (*float64)(&c[0]) }
  if _tmp410 := C.MSK_getc(self.ptr(),(*C.double)(_tmp409)); _tmp410 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp410)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetCfix() (cfix float64,err error) {
  if _tmp411 := C.MSK_getcfix(self.ptr(),(*C.double)(&cfix)); _tmp411 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp411)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetCJ(j int32) (cj float64,err error) {
  if _tmp412 := C.MSK_getcj(self.ptr(),C.int32_t(j),(*C.double)(&cj)); _tmp412 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp412)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetCList(subj []int32) (c []float64,err error) {
  _tmp413 := len(subj)
  var num int32 = int32(_tmp413)
  var _tmp414 *int32
  if subj != nil { _tmp414 = (*int32)(&subj[0]) }
  var _tmp415 *float64
  c = make([]float64,num)
  if len(c) > 0 { _tmp415 = (*float64)(&c[0]) }
  if _tmp416 := C.MSK_getclist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp414),(*C.double)(_tmp415)); _tmp416 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp416)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetConBound(i int32) (bk Boundkey,bl float64,bu float64,err error) {
  var _tmp417 int32
  if _tmp418 := C.MSK_getconbound(self.ptr(),C.int32_t(i),(*C.int32_t)(&_tmp417),(*C.double)(&bl),(*C.double)(&bu)); _tmp418 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp418)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  bk = Boundkey(_tmp417)
  return
}
func (self *Task) GetConBoundSlice(first int32,last int32) (bk []Boundkey,bl []float64,bu []float64,err error) {
  var _tmp419 *Boundkey
  bk = make([]Boundkey,(last - first))
  if len(bk) > 0 { _tmp419 = (*Boundkey)(&bk[0]) }
  var _tmp420 *float64
  bl = make([]float64,(last - first))
  if len(bl) > 0 { _tmp420 = (*float64)(&bl[0]) }
  var _tmp421 *float64
  bu = make([]float64,(last - first))
  if len(bu) > 0 { _tmp421 = (*float64)(&bu[0]) }
  if _tmp422 := C.MSK_getconboundslice(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.int32_t)(_tmp419),(*C.double)(_tmp420),(*C.double)(_tmp421)); _tmp422 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp422)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetCone(k int32) (ct Conetype,conepar float64,nummem int32,submem []int32,err error) {
  var _tmp423 int32
  var _tmp424 C.int32_t
  var _tmp425 C.double
  var _tmp426 C.int32_t
  if _tmp427 := C.MSK_getconeinfo(self.ptr(),(C.int32_t)(k),(&_tmp424),(&_tmp425),(&_tmp426)); _tmp427 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp427)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp428 *int32
  submem = make([]int32,_tmp426)
  if len(submem) > 0 { _tmp428 = (*int32)(&submem[0]) }
  if _tmp429 := C.MSK_getcone(self.ptr(),C.int32_t(k),(*C.int32_t)(&_tmp423),(*C.double)(&conepar),(*C.int32_t)(&nummem),(*C.int32_t)(_tmp428)); _tmp429 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp429)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  ct = Conetype(_tmp423)
  return
}
func (self *Task) GetConeInfo(k int32) (ct Conetype,conepar float64,nummem int32,err error) {
  var _tmp430 int32
  if _tmp431 := C.MSK_getconeinfo(self.ptr(),C.int32_t(k),(*C.int32_t)(&_tmp430),(*C.double)(&conepar),(*C.int32_t)(&nummem)); _tmp431 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp431)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  ct = Conetype(_tmp430)
  return
}
func (self *Task) GetConeName(i int32) (name string,err error) {
  var _tmp432 C.int32_t
  if _tmp433 := C.MSK_getconenamelen(self.ptr(),(C.int32_t)(i),(&_tmp432)); _tmp433 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp433)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp432))
  _tmp434 := make([]byte,sizename)
  if _tmp435 := C.MSK_getconename(self.ptr(),C.int32_t(i),C.int32_t(sizename),(*C.uchar)(&_tmp434[0])); _tmp435 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp435)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp434,byte(0)); p < 0 {
    name = string(_tmp434)
  } else {
    name = string(_tmp434[:p])
  }
  return
}
func (self *Task) GetConeNameIndex(somename string) (asgn int32,index int32,err error) {
  _tmp436 := C.CString(somename)
  if _tmp437 := C.MSK_getconenameindex(self.ptr(),_tmp436,(*C.int32_t)(&asgn),(*C.int32_t)(&index)); _tmp437 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp437)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetConeNameLen(i int32) (len int32,err error) {
  if _tmp438 := C.MSK_getconenamelen(self.ptr(),C.int32_t(i),(*C.int32_t)(&len)); _tmp438 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp438)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetConName(i int32) (name string,err error) {
  var _tmp439 C.int32_t
  if _tmp440 := C.MSK_getconnamelen(self.ptr(),(C.int32_t)(i),(&_tmp439)); _tmp440 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp440)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp439))
  _tmp441 := make([]byte,sizename)
  if _tmp442 := C.MSK_getconname(self.ptr(),C.int32_t(i),C.int32_t(sizename),(*C.uchar)(&_tmp441[0])); _tmp442 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp442)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp441,byte(0)); p < 0 {
    name = string(_tmp441)
  } else {
    name = string(_tmp441[:p])
  }
  return
}
func (self *Task) GetConNameIndex(somename string) (asgn int32,index int32,err error) {
  _tmp443 := C.CString(somename)
  if _tmp444 := C.MSK_getconnameindex(self.ptr(),_tmp443,(*C.int32_t)(&asgn),(*C.int32_t)(&index)); _tmp444 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp444)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetConNameLen(i int32) (len int32,err error) {
  if _tmp445 := C.MSK_getconnamelen(self.ptr(),C.int32_t(i),(*C.int32_t)(&len)); _tmp445 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp445)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetCSlice(first int32,last int32) (c []float64,err error) {
  var _tmp446 *float64
  c = make([]float64,(last - first))
  if len(c) > 0 { _tmp446 = (*float64)(&c[0]) }
  if _tmp447 := C.MSK_getcslice(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp446)); _tmp447 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp447)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDimBarvarJ(j int32) (dimbarvarj int32,err error) {
  if _tmp448 := C.MSK_getdimbarvarj(self.ptr(),C.int32_t(j),(*C.int32_t)(&dimbarvarj)); _tmp448 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp448)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcAfeIdxList(djcidx int64) (afeidxlist []int64,err error) {
  var _tmp449 C.int64_t
  if _tmp450 := C.MSK_getdjcnumafe(self.ptr(),(C.int64_t)(djcidx),(&_tmp449)); _tmp450 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp450)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp451 *int64
  afeidxlist = make([]int64,_tmp449)
  if len(afeidxlist) > 0 { _tmp451 = (*int64)(&afeidxlist[0]) }
  if _tmp452 := C.MSK_getdjcafeidxlist(self.ptr(),C.int64_t(djcidx),(*C.int64_t)(_tmp451)); _tmp452 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp452)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcB(djcidx int64) (b []float64,err error) {
  var _tmp453 C.int64_t
  if _tmp454 := C.MSK_getdjcnumafe(self.ptr(),(C.int64_t)(djcidx),(&_tmp453)); _tmp454 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp454)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp455 *float64
  b = make([]float64,_tmp453)
  if len(b) > 0 { _tmp455 = (*float64)(&b[0]) }
  if _tmp456 := C.MSK_getdjcb(self.ptr(),C.int64_t(djcidx),(*C.double)(_tmp455)); _tmp456 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp456)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcDomainIdxList(djcidx int64) (domidxlist []int64,err error) {
  var _tmp457 C.int64_t
  if _tmp458 := C.MSK_getdjcnumdomain(self.ptr(),(C.int64_t)(djcidx),(&_tmp457)); _tmp458 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp458)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp459 *int64
  domidxlist = make([]int64,_tmp457)
  if len(domidxlist) > 0 { _tmp459 = (*int64)(&domidxlist[0]) }
  if _tmp460 := C.MSK_getdjcdomainidxlist(self.ptr(),C.int64_t(djcidx),(*C.int64_t)(_tmp459)); _tmp460 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp460)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcName(djcidx int64) (name string,err error) {
  var _tmp461 C.int32_t
  if _tmp462 := C.MSK_getdjcnamelen(self.ptr(),(C.int64_t)(djcidx),(&_tmp461)); _tmp462 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp462)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp461))
  _tmp463 := make([]byte,sizename)
  if _tmp464 := C.MSK_getdjcname(self.ptr(),C.int64_t(djcidx),C.int32_t(sizename),(*C.uchar)(&_tmp463[0])); _tmp464 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp464)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp463,byte(0)); p < 0 {
    name = string(_tmp463)
  } else {
    name = string(_tmp463[:p])
  }
  return
}
func (self *Task) GetDjcNameLen(djcidx int64) (len int32,err error) {
  if _tmp465 := C.MSK_getdjcnamelen(self.ptr(),C.int64_t(djcidx),(*C.int32_t)(&len)); _tmp465 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp465)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcNumAfe(djcidx int64) (numafe int64,err error) {
  if _tmp466 := C.MSK_getdjcnumafe(self.ptr(),C.int64_t(djcidx),(*C.int64_t)(&numafe)); _tmp466 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp466)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcNumAfeTot() (numafetot int64,err error) {
  if _tmp467 := C.MSK_getdjcnumafetot(self.ptr(),(*C.int64_t)(&numafetot)); _tmp467 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp467)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcNumDomain(djcidx int64) (numdomain int64,err error) {
  if _tmp468 := C.MSK_getdjcnumdomain(self.ptr(),C.int64_t(djcidx),(*C.int64_t)(&numdomain)); _tmp468 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp468)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcNumDomainTot() (numdomaintot int64,err error) {
  if _tmp469 := C.MSK_getdjcnumdomaintot(self.ptr(),(*C.int64_t)(&numdomaintot)); _tmp469 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp469)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcNumTerm(djcidx int64) (numterm int64,err error) {
  if _tmp470 := C.MSK_getdjcnumterm(self.ptr(),C.int64_t(djcidx),(*C.int64_t)(&numterm)); _tmp470 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp470)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcNumTermTot() (numtermtot int64,err error) {
  if _tmp471 := C.MSK_getdjcnumtermtot(self.ptr(),(*C.int64_t)(&numtermtot)); _tmp471 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp471)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcs() (domidxlist []int64,afeidxlist []int64,b []float64,termsizelist []int64,numterms []int64,err error) {
  var _tmp472 C.int64_t
  if _tmp473 := C.MSK_getdjcnumdomaintot(self.ptr(),(&_tmp472)); _tmp473 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp473)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp474 *int64
  domidxlist = make([]int64,_tmp472)
  if len(domidxlist) > 0 { _tmp474 = (*int64)(&domidxlist[0]) }
  var _tmp475 C.int64_t
  if _tmp476 := C.MSK_getdjcnumafetot(self.ptr(),(&_tmp475)); _tmp476 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp476)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp477 *int64
  afeidxlist = make([]int64,_tmp475)
  if len(afeidxlist) > 0 { _tmp477 = (*int64)(&afeidxlist[0]) }
  var _tmp478 C.int64_t
  if _tmp479 := C.MSK_getdjcnumafetot(self.ptr(),(&_tmp478)); _tmp479 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp479)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp480 *float64
  b = make([]float64,_tmp478)
  if len(b) > 0 { _tmp480 = (*float64)(&b[0]) }
  var _tmp481 C.int64_t
  if _tmp482 := C.MSK_getdjcnumtermtot(self.ptr(),(&_tmp481)); _tmp482 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp482)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp483 *int64
  termsizelist = make([]int64,_tmp481)
  if len(termsizelist) > 0 { _tmp483 = (*int64)(&termsizelist[0]) }
  var _tmp484 C.int64_t
  if _tmp485 := C.MSK_getnumdjc(self.ptr(),(&_tmp484)); _tmp485 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp485)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp486 *int64
  numterms = make([]int64,_tmp484)
  if len(numterms) > 0 { _tmp486 = (*int64)(&numterms[0]) }
  if _tmp487 := C.MSK_getdjcs(self.ptr(),(*C.int64_t)(_tmp474),(*C.int64_t)(_tmp477),(*C.double)(_tmp480),(*C.int64_t)(_tmp483),(*C.int64_t)(_tmp486)); _tmp487 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp487)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcTermSizeList(djcidx int64) (termsizelist []int64,err error) {
  var _tmp488 C.int64_t
  if _tmp489 := C.MSK_getdjcnumterm(self.ptr(),(C.int64_t)(djcidx),(&_tmp488)); _tmp489 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp489)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp490 *int64
  termsizelist = make([]int64,_tmp488)
  if len(termsizelist) > 0 { _tmp490 = (*int64)(&termsizelist[0]) }
  if _tmp491 := C.MSK_getdjctermsizelist(self.ptr(),C.int64_t(djcidx),(*C.int64_t)(_tmp490)); _tmp491 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp491)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDomainN(domidx int64) (n int64,err error) {
  if _tmp492 := C.MSK_getdomainn(self.ptr(),C.int64_t(domidx),(*C.int64_t)(&n)); _tmp492 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp492)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDomainName(domidx int64) (name string,err error) {
  var _tmp493 C.int32_t
  if _tmp494 := C.MSK_getdomainnamelen(self.ptr(),(C.int64_t)(domidx),(&_tmp493)); _tmp494 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp494)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp493))
  _tmp495 := make([]byte,sizename)
  if _tmp496 := C.MSK_getdomainname(self.ptr(),C.int64_t(domidx),C.int32_t(sizename),(*C.uchar)(&_tmp495[0])); _tmp496 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp496)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp495,byte(0)); p < 0 {
    name = string(_tmp495)
  } else {
    name = string(_tmp495[:p])
  }
  return
}
func (self *Task) GetDomainNameLen(domidx int64) (len int32,err error) {
  if _tmp497 := C.MSK_getdomainnamelen(self.ptr(),C.int64_t(domidx),(*C.int32_t)(&len)); _tmp497 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp497)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDomainType(domidx int64) (domtype Domaintype,err error) {
  var _tmp498 int32
  if _tmp499 := C.MSK_getdomaintype(self.ptr(),C.int64_t(domidx),(*C.int32_t)(&_tmp498)); _tmp499 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp499)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  domtype = Domaintype(_tmp498)
  return
}
func (self *Task) GetDouInf(whichdinf Dinfitem) (dvalue float64,err error) {
  if _tmp500 := C.MSK_getdouinf(self.ptr(),C.int32_t(whichdinf),(*C.double)(&dvalue)); _tmp500 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp500)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDouParam(param Dparam) (parvalue float64,err error) {
  if _tmp501 := C.MSK_getdouparam(self.ptr(),C.int32_t(param),(*C.double)(&parvalue)); _tmp501 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp501)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDualObj(whichsol Soltype) (dualobj float64,err error) {
  if _tmp502 := C.MSK_getdualobj(self.ptr(),C.int32_t(whichsol),(*C.double)(&dualobj)); _tmp502 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp502)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDualSolutionNorms(whichsol Soltype) (nrmy float64,nrmslc float64,nrmsuc float64,nrmslx float64,nrmsux float64,nrmsnx float64,nrmbars float64,err error) {
  if _tmp503 := C.MSK_getdualsolutionnorms(self.ptr(),C.int32_t(whichsol),(*C.double)(&nrmy),(*C.double)(&nrmslc),(*C.double)(&nrmsuc),(*C.double)(&nrmslx),(*C.double)(&nrmsux),(*C.double)(&nrmsnx),(*C.double)(&nrmbars)); _tmp503 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp503)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDviolAcc(whichsol Soltype,accidxlist []int64) (viol []float64,err error) {
  _tmp504 := len(accidxlist)
  var numaccidx int64 = int64(_tmp504)
  var _tmp505 *int64
  if accidxlist != nil { _tmp505 = (*int64)(&accidxlist[0]) }
  var _tmp506 *float64
  viol = make([]float64,numaccidx)
  if len(viol) > 0 { _tmp506 = (*float64)(&viol[0]) }
  if _tmp507 := C.MSK_getdviolacc(self.ptr(),C.int32_t(whichsol),C.int64_t(numaccidx),(*C.int64_t)(_tmp505),(*C.double)(_tmp506)); _tmp507 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp507)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDviolBarvar(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp508 := len(sub)
  var num int32 = int32(_tmp508)
  var _tmp509 *int32
  if sub != nil { _tmp509 = (*int32)(&sub[0]) }
  var _tmp510 *float64
  viol = make([]float64,num)
  if len(viol) > 0 { _tmp510 = (*float64)(&viol[0]) }
  if _tmp511 := C.MSK_getdviolbarvar(self.ptr(),C.int32_t(whichsol),C.int32_t(num),(*C.int32_t)(_tmp509),(*C.double)(_tmp510)); _tmp511 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp511)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDviolCon(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp512 := len(sub)
  var num int32 = int32(_tmp512)
  var _tmp513 *int32
  if sub != nil { _tmp513 = (*int32)(&sub[0]) }
  var _tmp514 *float64
  viol = make([]float64,num)
  if len(viol) > 0 { _tmp514 = (*float64)(&viol[0]) }
  if _tmp515 := C.MSK_getdviolcon(self.ptr(),C.int32_t(whichsol),C.int32_t(num),(*C.int32_t)(_tmp513),(*C.double)(_tmp514)); _tmp515 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp515)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDviolCones(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp516 := len(sub)
  var num int32 = int32(_tmp516)
  var _tmp517 *int32
  if sub != nil { _tmp517 = (*int32)(&sub[0]) }
  var _tmp518 *float64
  viol = make([]float64,num)
  if len(viol) > 0 { _tmp518 = (*float64)(&viol[0]) }
  if _tmp519 := C.MSK_getdviolcones(self.ptr(),C.int32_t(whichsol),C.int32_t(num),(*C.int32_t)(_tmp517),(*C.double)(_tmp518)); _tmp519 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp519)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDviolVar(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp520 := len(sub)
  var num int32 = int32(_tmp520)
  var _tmp521 *int32
  if sub != nil { _tmp521 = (*int32)(&sub[0]) }
  var _tmp522 *float64
  viol = make([]float64,num)
  if len(viol) > 0 { _tmp522 = (*float64)(&viol[0]) }
  if _tmp523 := C.MSK_getdviolvar(self.ptr(),C.int32_t(whichsol),C.int32_t(num),(*C.int32_t)(_tmp521),(*C.double)(_tmp522)); _tmp523 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp523)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetInfIndex(inftype Inftype,infname string) (infindex int32,err error) {
  _tmp524 := C.CString(infname)
  if _tmp525 := C.MSK_getinfindex(self.ptr(),C.int32_t(inftype),_tmp524,(*C.int32_t)(&infindex)); _tmp525 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp525)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetInfMax(inftype Inftype) (infmax []int32,err error) {
  var _tmp526 *int32
  infmax = make([]int32,MSK_MAX_STR_LEN)
  if len(infmax) > 0 { _tmp526 = (*int32)(&infmax[0]) }
  if _tmp527 := C.MSK_getinfmax(self.ptr(),C.int32_t(inftype),(*C.int32_t)(_tmp526)); _tmp527 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp527)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetInfName(inftype Inftype,whichinf int32) (infname string,err error) {
  _tmp528 := make([]byte,MSK_MAX_STR_LEN)
  if _tmp529 := C.MSK_getinfname(self.ptr(),C.int32_t(inftype),C.int32_t(whichinf),(*C.uchar)(&_tmp528[0])); _tmp529 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp529)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp528,byte(0)); p < 0 {
    infname = string(_tmp528)
  } else {
    infname = string(_tmp528[:p])
  }
  return
}
func (self *Task) GetIntInf(whichiinf Iinfitem) (ivalue int32,err error) {
  if _tmp530 := C.MSK_getintinf(self.ptr(),C.int32_t(whichiinf),(*C.int32_t)(&ivalue)); _tmp530 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp530)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetIntParam(param Iparam) (parvalue int32,err error) {
  if _tmp531 := C.MSK_getintparam(self.ptr(),C.int32_t(param),(*C.int32_t)(&parvalue)); _tmp531 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp531)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetLenBarvarJ(j int32) (lenbarvarj int64,err error) {
  if _tmp532 := C.MSK_getlenbarvarj(self.ptr(),C.int32_t(j),(*C.int64_t)(&lenbarvarj)); _tmp532 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp532)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetLintInf(whichliinf Liinfitem) (ivalue int64,err error) {
  if _tmp533 := C.MSK_getlintinf(self.ptr(),C.int32_t(whichliinf),(*C.int64_t)(&ivalue)); _tmp533 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp533)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetMaxNumANz() (maxnumanz int64,err error) {
  if _tmp534 := C.MSK_getmaxnumanz64(self.ptr(),(*C.int64_t)(&maxnumanz)); _tmp534 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp534)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetMaxNumBarvar() (maxnumbarvar int32,err error) {
  if _tmp535 := C.MSK_getmaxnumbarvar(self.ptr(),(*C.int32_t)(&maxnumbarvar)); _tmp535 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp535)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetMaxNumCon() (maxnumcon int32,err error) {
  if _tmp536 := C.MSK_getmaxnumcon(self.ptr(),(*C.int32_t)(&maxnumcon)); _tmp536 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp536)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetMaxNumCone() (maxnumcone int32,err error) {
  if _tmp537 := C.MSK_getmaxnumcone(self.ptr(),(*C.int32_t)(&maxnumcone)); _tmp537 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp537)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetMaxNumQNz() (maxnumqnz int64,err error) {
  if _tmp538 := C.MSK_getmaxnumqnz64(self.ptr(),(*C.int64_t)(&maxnumqnz)); _tmp538 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp538)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetMaxNumVar() (maxnumvar int32,err error) {
  if _tmp539 := C.MSK_getmaxnumvar(self.ptr(),(*C.int32_t)(&maxnumvar)); _tmp539 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp539)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetMemUsage() (meminuse int64,maxmemuse int64,err error) {
  if _tmp540 := C.MSK_getmemusagetask(self.ptr(),(*C.int64_t)(&meminuse),(*C.int64_t)(&maxmemuse)); _tmp540 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp540)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumAcc() (num int64,err error) {
  if _tmp541 := C.MSK_getnumacc(self.ptr(),(*C.int64_t)(&num)); _tmp541 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp541)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumAfe() (numafe int64,err error) {
  if _tmp542 := C.MSK_getnumafe(self.ptr(),(*C.int64_t)(&numafe)); _tmp542 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp542)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumANz() (numanz int32,err error) {
  if _tmp543 := C.MSK_getnumanz(self.ptr(),(*C.int32_t)(&numanz)); _tmp543 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp543)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumANz64() (numanz int64,err error) {
  if _tmp544 := C.MSK_getnumanz64(self.ptr(),(*C.int64_t)(&numanz)); _tmp544 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp544)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumBaraBlockTriplets() (num int64,err error) {
  if _tmp545 := C.MSK_getnumbarablocktriplets(self.ptr(),(*C.int64_t)(&num)); _tmp545 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp545)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumBaraNz() (nz int64,err error) {
  if _tmp546 := C.MSK_getnumbaranz(self.ptr(),(*C.int64_t)(&nz)); _tmp546 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp546)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumBarcBlockTriplets() (num int64,err error) {
  if _tmp547 := C.MSK_getnumbarcblocktriplets(self.ptr(),(*C.int64_t)(&num)); _tmp547 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp547)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumBarcNz() (nz int64,err error) {
  if _tmp548 := C.MSK_getnumbarcnz(self.ptr(),(*C.int64_t)(&nz)); _tmp548 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp548)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumBarvar() (numbarvar int32,err error) {
  if _tmp549 := C.MSK_getnumbarvar(self.ptr(),(*C.int32_t)(&numbarvar)); _tmp549 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp549)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumCon() (numcon int32,err error) {
  if _tmp550 := C.MSK_getnumcon(self.ptr(),(*C.int32_t)(&numcon)); _tmp550 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp550)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumCone() (numcone int32,err error) {
  if _tmp551 := C.MSK_getnumcone(self.ptr(),(*C.int32_t)(&numcone)); _tmp551 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp551)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumConeMem(k int32) (nummem int32,err error) {
  if _tmp552 := C.MSK_getnumconemem(self.ptr(),C.int32_t(k),(*C.int32_t)(&nummem)); _tmp552 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp552)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumDjc() (num int64,err error) {
  if _tmp553 := C.MSK_getnumdjc(self.ptr(),(*C.int64_t)(&num)); _tmp553 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp553)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumDomain() (numdomain int64,err error) {
  if _tmp554 := C.MSK_getnumdomain(self.ptr(),(*C.int64_t)(&numdomain)); _tmp554 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp554)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumIntVar() (numintvar int32,err error) {
  if _tmp555 := C.MSK_getnumintvar(self.ptr(),(*C.int32_t)(&numintvar)); _tmp555 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp555)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumParam(partype Parametertype) (numparam int32,err error) {
  if _tmp556 := C.MSK_getnumparam(self.ptr(),C.int32_t(partype),(*C.int32_t)(&numparam)); _tmp556 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp556)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumQConKNz(k int32) (numqcnz int64,err error) {
  if _tmp557 := C.MSK_getnumqconknz64(self.ptr(),C.int32_t(k),(*C.int64_t)(&numqcnz)); _tmp557 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp557)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumQObjNz() (numqonz int64,err error) {
  if _tmp558 := C.MSK_getnumqobjnz64(self.ptr(),(*C.int64_t)(&numqonz)); _tmp558 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp558)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumSymMat() (num int64,err error) {
  if _tmp559 := C.MSK_getnumsymmat(self.ptr(),(*C.int64_t)(&num)); _tmp559 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp559)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumVar() (numvar int32,err error) {
  if _tmp560 := C.MSK_getnumvar(self.ptr(),(*C.int32_t)(&numvar)); _tmp560 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp560)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetObjName() (objname string,err error) {
  var _tmp561 C.int32_t
  if _tmp562 := C.MSK_getobjnamelen(self.ptr(),(&_tmp561)); _tmp562 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp562)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizeobjname int32 = int32((1 + _tmp561))
  _tmp563 := make([]byte,sizeobjname)
  if _tmp564 := C.MSK_getobjname(self.ptr(),C.int32_t(sizeobjname),(*C.uchar)(&_tmp563[0])); _tmp564 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp564)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp563,byte(0)); p < 0 {
    objname = string(_tmp563)
  } else {
    objname = string(_tmp563[:p])
  }
  return
}
func (self *Task) GetObjNameLen() (len int32,err error) {
  if _tmp565 := C.MSK_getobjnamelen(self.ptr(),(*C.int32_t)(&len)); _tmp565 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp565)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetObjSense() (sense Objsense,err error) {
  var _tmp566 int32
  if _tmp567 := C.MSK_getobjsense(self.ptr(),(*C.int32_t)(&_tmp566)); _tmp567 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp567)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  sense = Objsense(_tmp566)
  return
}
func (self *Task) GetParamMax(partype Parametertype) (parammax int32,err error) {
  if _tmp568 := C.MSK_getparammax(self.ptr(),C.int32_t(partype),(*C.int32_t)(&parammax)); _tmp568 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp568)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetParamName(partype Parametertype,param int32) (parname string,err error) {
  _tmp569 := make([]byte,MSK_MAX_STR_LEN)
  if _tmp570 := C.MSK_getparamname(self.ptr(),C.int32_t(partype),C.int32_t(param),(*C.uchar)(&_tmp569[0])); _tmp570 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp570)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp569,byte(0)); p < 0 {
    parname = string(_tmp569)
  } else {
    parname = string(_tmp569[:p])
  }
  return
}
func (self *Task) GetPowerDomainAlpha(domidx int64) (alpha []float64,err error) {
  var _tmp571 C.int64_t
  var _tmp572 C.int64_t
  if _tmp573 := C.MSK_getpowerdomaininfo(self.ptr(),(C.int64_t)(domidx),(&_tmp571),(&_tmp572)); _tmp573 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp573)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp574 *float64
  alpha = make([]float64,_tmp572)
  if len(alpha) > 0 { _tmp574 = (*float64)(&alpha[0]) }
  if _tmp575 := C.MSK_getpowerdomainalpha(self.ptr(),C.int64_t(domidx),(*C.double)(_tmp574)); _tmp575 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp575)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetPowerDomainInfo(domidx int64) (n int64,nleft int64,err error) {
  if _tmp576 := C.MSK_getpowerdomaininfo(self.ptr(),C.int64_t(domidx),(*C.int64_t)(&n),(*C.int64_t)(&nleft)); _tmp576 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp576)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetPrimalObj(whichsol Soltype) (primalobj float64,err error) {
  if _tmp577 := C.MSK_getprimalobj(self.ptr(),C.int32_t(whichsol),(*C.double)(&primalobj)); _tmp577 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp577)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetPrimalSolutionNorms(whichsol Soltype) (nrmxc float64,nrmxx float64,nrmbarx float64,err error) {
  if _tmp578 := C.MSK_getprimalsolutionnorms(self.ptr(),C.int32_t(whichsol),(*C.double)(&nrmxc),(*C.double)(&nrmxx),(*C.double)(&nrmbarx)); _tmp578 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp578)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetProbType() (probtype Problemtype,err error) {
  var _tmp579 int32
  if _tmp580 := C.MSK_getprobtype(self.ptr(),(*C.int32_t)(&_tmp579)); _tmp580 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp580)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  probtype = Problemtype(_tmp579)
  return
}
func (self *Task) GetProSta(whichsol Soltype) (problemsta Prosta,err error) {
  var _tmp581 int32
  if _tmp582 := C.MSK_getprosta(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(&_tmp581)); _tmp582 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp582)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  problemsta = Prosta(_tmp581)
  return
}
func (self *Task) GetPviolAcc(whichsol Soltype,accidxlist []int64) (viol []float64,err error) {
  _tmp583 := len(accidxlist)
  var numaccidx int64 = int64(_tmp583)
  var _tmp584 *int64
  if accidxlist != nil { _tmp584 = (*int64)(&accidxlist[0]) }
  var _tmp585 *float64
  viol = make([]float64,numaccidx)
  if len(viol) > 0 { _tmp585 = (*float64)(&viol[0]) }
  if _tmp586 := C.MSK_getpviolacc(self.ptr(),C.int32_t(whichsol),C.int64_t(numaccidx),(*C.int64_t)(_tmp584),(*C.double)(_tmp585)); _tmp586 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp586)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetPviolBarvar(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp587 := len(sub)
  var num int32 = int32(_tmp587)
  var _tmp588 *int32
  if sub != nil { _tmp588 = (*int32)(&sub[0]) }
  var _tmp589 *float64
  viol = make([]float64,num)
  if len(viol) > 0 { _tmp589 = (*float64)(&viol[0]) }
  if _tmp590 := C.MSK_getpviolbarvar(self.ptr(),C.int32_t(whichsol),C.int32_t(num),(*C.int32_t)(_tmp588),(*C.double)(_tmp589)); _tmp590 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp590)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetPviolCon(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp591 := len(sub)
  var num int32 = int32(_tmp591)
  var _tmp592 *int32
  if sub != nil { _tmp592 = (*int32)(&sub[0]) }
  var _tmp593 *float64
  viol = make([]float64,num)
  if len(viol) > 0 { _tmp593 = (*float64)(&viol[0]) }
  if _tmp594 := C.MSK_getpviolcon(self.ptr(),C.int32_t(whichsol),C.int32_t(num),(*C.int32_t)(_tmp592),(*C.double)(_tmp593)); _tmp594 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp594)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetPviolCones(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp595 := len(sub)
  var num int32 = int32(_tmp595)
  var _tmp596 *int32
  if sub != nil { _tmp596 = (*int32)(&sub[0]) }
  var _tmp597 *float64
  viol = make([]float64,num)
  if len(viol) > 0 { _tmp597 = (*float64)(&viol[0]) }
  if _tmp598 := C.MSK_getpviolcones(self.ptr(),C.int32_t(whichsol),C.int32_t(num),(*C.int32_t)(_tmp596),(*C.double)(_tmp597)); _tmp598 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp598)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetPviolDjc(whichsol Soltype,djcidxlist []int64) (viol []float64,err error) {
  _tmp599 := len(djcidxlist)
  var numdjcidx int64 = int64(_tmp599)
  var _tmp600 *int64
  if djcidxlist != nil { _tmp600 = (*int64)(&djcidxlist[0]) }
  var _tmp601 *float64
  viol = make([]float64,numdjcidx)
  if len(viol) > 0 { _tmp601 = (*float64)(&viol[0]) }
  if _tmp602 := C.MSK_getpvioldjc(self.ptr(),C.int32_t(whichsol),C.int64_t(numdjcidx),(*C.int64_t)(_tmp600),(*C.double)(_tmp601)); _tmp602 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp602)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetPviolVar(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp603 := len(sub)
  var num int32 = int32(_tmp603)
  var _tmp604 *int32
  if sub != nil { _tmp604 = (*int32)(&sub[0]) }
  var _tmp605 *float64
  viol = make([]float64,num)
  if len(viol) > 0 { _tmp605 = (*float64)(&viol[0]) }
  if _tmp606 := C.MSK_getpviolvar(self.ptr(),C.int32_t(whichsol),C.int32_t(num),(*C.int32_t)(_tmp604),(*C.double)(_tmp605)); _tmp606 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp606)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetQConK(k int32) (numqcnz int64,qcsubi []int32,qcsubj []int32,qcval []float64,err error) {
  var _tmp607 C.int64_t
  if _tmp608 := C.MSK_getnumqconknz64(self.ptr(),(C.int32_t)(k),(&_tmp607)); _tmp608 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp608)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumqcnz int64 = int64(_tmp607)
  var _tmp609 C.int64_t
  if _tmp610 := C.MSK_getnumqconknz64(self.ptr(),(C.int32_t)(k),(&_tmp609)); _tmp610 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp610)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp611 *int32
  qcsubi = make([]int32,_tmp609)
  if len(qcsubi) > 0 { _tmp611 = (*int32)(&qcsubi[0]) }
  var _tmp612 C.int64_t
  if _tmp613 := C.MSK_getnumqconknz64(self.ptr(),(C.int32_t)(k),(&_tmp612)); _tmp613 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp613)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp614 *int32
  qcsubj = make([]int32,_tmp612)
  if len(qcsubj) > 0 { _tmp614 = (*int32)(&qcsubj[0]) }
  var _tmp615 C.int64_t
  if _tmp616 := C.MSK_getnumqconknz64(self.ptr(),(C.int32_t)(k),(&_tmp615)); _tmp616 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp616)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp617 *float64
  qcval = make([]float64,_tmp615)
  if len(qcval) > 0 { _tmp617 = (*float64)(&qcval[0]) }
  if _tmp618 := C.MSK_getqconk64(self.ptr(),C.int32_t(k),C.int64_t(maxnumqcnz),(*C.int64_t)(&numqcnz),(*C.int32_t)(_tmp611),(*C.int32_t)(_tmp614),(*C.double)(_tmp617)); _tmp618 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp618)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetQObj() (numqonz int64,qosubi []int32,qosubj []int32,qoval []float64,err error) {
  var _tmp619 C.int64_t
  if _tmp620 := C.MSK_getnumqobjnz64(self.ptr(),(&_tmp619)); _tmp620 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp620)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxnumqonz int64 = int64(_tmp619)
  var _tmp621 *int32
  qosubi = make([]int32,maxnumqonz)
  if len(qosubi) > 0 { _tmp621 = (*int32)(&qosubi[0]) }
  var _tmp622 *int32
  qosubj = make([]int32,maxnumqonz)
  if len(qosubj) > 0 { _tmp622 = (*int32)(&qosubj[0]) }
  var _tmp623 *float64
  qoval = make([]float64,maxnumqonz)
  if len(qoval) > 0 { _tmp623 = (*float64)(&qoval[0]) }
  if _tmp624 := C.MSK_getqobj64(self.ptr(),C.int64_t(maxnumqonz),(*C.int64_t)(&numqonz),(*C.int32_t)(_tmp621),(*C.int32_t)(_tmp622),(*C.double)(_tmp623)); _tmp624 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp624)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetQObjIJ(i int32,j int32) (qoij float64,err error) {
  if _tmp625 := C.MSK_getqobjij(self.ptr(),C.int32_t(i),C.int32_t(j),(*C.double)(&qoij)); _tmp625 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp625)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetReducedCosts(whichsol Soltype,first int32,last int32) (redcosts []float64,err error) {
  var _tmp626 *float64
  redcosts = make([]float64,(last - first))
  if len(redcosts) > 0 { _tmp626 = (*float64)(&redcosts[0]) }
  if _tmp627 := C.MSK_getreducedcosts(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp626)); _tmp627 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp627)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSkc(whichsol Soltype) (skc []Stakey,err error) {
  var _tmp628 C.int32_t
  if _tmp629 := C.MSK_getnumcon(self.ptr(),(&_tmp628)); _tmp629 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp629)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp630 *Stakey
  skc = make([]Stakey,_tmp628)
  if len(skc) > 0 { _tmp630 = (*Stakey)(&skc[0]) }
  if _tmp631 := C.MSK_getskc(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(_tmp630)); _tmp631 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp631)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSkcSlice(whichsol Soltype,first int32,last int32) (skc []Stakey,err error) {
  var _tmp632 *Stakey
  skc = make([]Stakey,(last - first))
  if len(skc) > 0 { _tmp632 = (*Stakey)(&skc[0]) }
  if _tmp633 := C.MSK_getskcslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.int32_t)(_tmp632)); _tmp633 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp633)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSkn(whichsol Soltype) (skn []Stakey,err error) {
  var _tmp634 C.int32_t
  if _tmp635 := C.MSK_getnumcone(self.ptr(),(&_tmp634)); _tmp635 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp635)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp636 *Stakey
  skn = make([]Stakey,_tmp634)
  if len(skn) > 0 { _tmp636 = (*Stakey)(&skn[0]) }
  if _tmp637 := C.MSK_getskn(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(_tmp636)); _tmp637 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp637)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSkx(whichsol Soltype) (skx []Stakey,err error) {
  var _tmp638 C.int32_t
  if _tmp639 := C.MSK_getnumvar(self.ptr(),(&_tmp638)); _tmp639 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp639)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp640 *Stakey
  skx = make([]Stakey,_tmp638)
  if len(skx) > 0 { _tmp640 = (*Stakey)(&skx[0]) }
  if _tmp641 := C.MSK_getskx(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(_tmp640)); _tmp641 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp641)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSkxSlice(whichsol Soltype,first int32,last int32) (skx []Stakey,err error) {
  var _tmp642 *Stakey
  skx = make([]Stakey,(last - first))
  if len(skx) > 0 { _tmp642 = (*Stakey)(&skx[0]) }
  if _tmp643 := C.MSK_getskxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.int32_t)(_tmp642)); _tmp643 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp643)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSlc(whichsol Soltype) (slc []float64,err error) {
  var _tmp644 C.int32_t
  if _tmp645 := C.MSK_getnumcon(self.ptr(),(&_tmp644)); _tmp645 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp645)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp646 *float64
  slc = make([]float64,_tmp644)
  if len(slc) > 0 { _tmp646 = (*float64)(&slc[0]) }
  if _tmp647 := C.MSK_getslc(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp646)); _tmp647 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp647)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSlcSlice(whichsol Soltype,first int32,last int32) (slc []float64,err error) {
  var _tmp648 *float64
  slc = make([]float64,(last - first))
  if len(slc) > 0 { _tmp648 = (*float64)(&slc[0]) }
  if _tmp649 := C.MSK_getslcslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp648)); _tmp649 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp649)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSlx(whichsol Soltype) (slx []float64,err error) {
  var _tmp650 C.int32_t
  if _tmp651 := C.MSK_getnumvar(self.ptr(),(&_tmp650)); _tmp651 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp651)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp652 *float64
  slx = make([]float64,_tmp650)
  if len(slx) > 0 { _tmp652 = (*float64)(&slx[0]) }
  if _tmp653 := C.MSK_getslx(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp652)); _tmp653 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp653)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSlxSlice(whichsol Soltype,first int32,last int32) (slx []float64,err error) {
  var _tmp654 *float64
  slx = make([]float64,(last - first))
  if len(slx) > 0 { _tmp654 = (*float64)(&slx[0]) }
  if _tmp655 := C.MSK_getslxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp654)); _tmp655 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp655)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSnx(whichsol Soltype) (snx []float64,err error) {
  var _tmp656 C.int32_t
  if _tmp657 := C.MSK_getnumvar(self.ptr(),(&_tmp656)); _tmp657 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp657)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp658 *float64
  snx = make([]float64,_tmp656)
  if len(snx) > 0 { _tmp658 = (*float64)(&snx[0]) }
  if _tmp659 := C.MSK_getsnx(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp658)); _tmp659 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp659)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSnxSlice(whichsol Soltype,first int32,last int32) (snx []float64,err error) {
  var _tmp660 *float64
  snx = make([]float64,(last - first))
  if len(snx) > 0 { _tmp660 = (*float64)(&snx[0]) }
  if _tmp661 := C.MSK_getsnxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp660)); _tmp661 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp661)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSolSta(whichsol Soltype) (solutionsta Solsta,err error) {
  var _tmp662 int32
  if _tmp663 := C.MSK_getsolsta(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(&_tmp662)); _tmp663 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp663)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  solutionsta = Solsta(_tmp662)
  return
}
func (self *Task) GetSolution(whichsol Soltype) (problemsta Prosta,solutionsta Solsta,skc []Stakey,skx []Stakey,skn []Stakey,xc []float64,xx []float64,y []float64,slc []float64,suc []float64,slx []float64,sux []float64,snx []float64,err error) {
  var _tmp664 int32
  var _tmp665 int32
  var _tmp666 C.int32_t
  if _tmp667 := C.MSK_getnumcon(self.ptr(),(&_tmp666)); _tmp667 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp667)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp668 *Stakey
  skc = make([]Stakey,_tmp666)
  if len(skc) > 0 { _tmp668 = (*Stakey)(&skc[0]) }
  var _tmp669 C.int32_t
  if _tmp670 := C.MSK_getnumvar(self.ptr(),(&_tmp669)); _tmp670 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp670)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp671 *Stakey
  skx = make([]Stakey,_tmp669)
  if len(skx) > 0 { _tmp671 = (*Stakey)(&skx[0]) }
  var _tmp672 C.int32_t
  if _tmp673 := C.MSK_getnumcone(self.ptr(),(&_tmp672)); _tmp673 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp673)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp674 *Stakey
  skn = make([]Stakey,_tmp672)
  if len(skn) > 0 { _tmp674 = (*Stakey)(&skn[0]) }
  var _tmp675 C.int32_t
  if _tmp676 := C.MSK_getnumcon(self.ptr(),(&_tmp675)); _tmp676 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp676)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp677 *float64
  xc = make([]float64,_tmp675)
  if len(xc) > 0 { _tmp677 = (*float64)(&xc[0]) }
  var _tmp678 C.int32_t
  if _tmp679 := C.MSK_getnumvar(self.ptr(),(&_tmp678)); _tmp679 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp679)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp680 *float64
  xx = make([]float64,_tmp678)
  if len(xx) > 0 { _tmp680 = (*float64)(&xx[0]) }
  var _tmp681 C.int32_t
  if _tmp682 := C.MSK_getnumcon(self.ptr(),(&_tmp681)); _tmp682 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp682)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp683 *float64
  y = make([]float64,_tmp681)
  if len(y) > 0 { _tmp683 = (*float64)(&y[0]) }
  var _tmp684 C.int32_t
  if _tmp685 := C.MSK_getnumcon(self.ptr(),(&_tmp684)); _tmp685 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp685)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp686 *float64
  slc = make([]float64,_tmp684)
  if len(slc) > 0 { _tmp686 = (*float64)(&slc[0]) }
  var _tmp687 C.int32_t
  if _tmp688 := C.MSK_getnumcon(self.ptr(),(&_tmp687)); _tmp688 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp688)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp689 *float64
  suc = make([]float64,_tmp687)
  if len(suc) > 0 { _tmp689 = (*float64)(&suc[0]) }
  var _tmp690 C.int32_t
  if _tmp691 := C.MSK_getnumvar(self.ptr(),(&_tmp690)); _tmp691 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp691)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp692 *float64
  slx = make([]float64,_tmp690)
  if len(slx) > 0 { _tmp692 = (*float64)(&slx[0]) }
  var _tmp693 C.int32_t
  if _tmp694 := C.MSK_getnumvar(self.ptr(),(&_tmp693)); _tmp694 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp694)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp695 *float64
  sux = make([]float64,_tmp693)
  if len(sux) > 0 { _tmp695 = (*float64)(&sux[0]) }
  var _tmp696 C.int32_t
  if _tmp697 := C.MSK_getnumvar(self.ptr(),(&_tmp696)); _tmp697 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp697)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp698 *float64
  snx = make([]float64,_tmp696)
  if len(snx) > 0 { _tmp698 = (*float64)(&snx[0]) }
  if _tmp699 := C.MSK_getsolution(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(&_tmp664),(*C.int32_t)(&_tmp665),(*C.int32_t)(_tmp668),(*C.int32_t)(_tmp671),(*C.int32_t)(_tmp674),(*C.double)(_tmp677),(*C.double)(_tmp680),(*C.double)(_tmp683),(*C.double)(_tmp686),(*C.double)(_tmp689),(*C.double)(_tmp692),(*C.double)(_tmp695),(*C.double)(_tmp698)); _tmp699 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp699)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  problemsta = Prosta(_tmp664)
  solutionsta = Solsta(_tmp665)
  return
}
func (self *Task) GetSolutionInfo(whichsol Soltype) (pobj float64,pviolcon float64,pviolvar float64,pviolbarvar float64,pviolcone float64,pviolitg float64,dobj float64,dviolcon float64,dviolvar float64,dviolbarvar float64,dviolcone float64,err error) {
  if _tmp700 := C.MSK_getsolutioninfo(self.ptr(),C.int32_t(whichsol),(*C.double)(&pobj),(*C.double)(&pviolcon),(*C.double)(&pviolvar),(*C.double)(&pviolbarvar),(*C.double)(&pviolcone),(*C.double)(&pviolitg),(*C.double)(&dobj),(*C.double)(&dviolcon),(*C.double)(&dviolvar),(*C.double)(&dviolbarvar),(*C.double)(&dviolcone)); _tmp700 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp700)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSolutionInfoNew(whichsol Soltype) (pobj float64,pviolcon float64,pviolvar float64,pviolbarvar float64,pviolcone float64,pviolacc float64,pvioldjc float64,pviolitg float64,dobj float64,dviolcon float64,dviolvar float64,dviolbarvar float64,dviolcone float64,dviolacc float64,err error) {
  if _tmp701 := C.MSK_getsolutioninfonew(self.ptr(),C.int32_t(whichsol),(*C.double)(&pobj),(*C.double)(&pviolcon),(*C.double)(&pviolvar),(*C.double)(&pviolbarvar),(*C.double)(&pviolcone),(*C.double)(&pviolacc),(*C.double)(&pvioldjc),(*C.double)(&pviolitg),(*C.double)(&dobj),(*C.double)(&dviolcon),(*C.double)(&dviolvar),(*C.double)(&dviolbarvar),(*C.double)(&dviolcone),(*C.double)(&dviolacc)); _tmp701 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp701)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSolutionNew(whichsol Soltype) (problemsta Prosta,solutionsta Solsta,skc []Stakey,skx []Stakey,skn []Stakey,xc []float64,xx []float64,y []float64,slc []float64,suc []float64,slx []float64,sux []float64,snx []float64,doty []float64,err error) {
  var _tmp702 int32
  var _tmp703 int32
  var _tmp704 C.int32_t
  if _tmp705 := C.MSK_getnumcon(self.ptr(),(&_tmp704)); _tmp705 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp705)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp706 *Stakey
  skc = make([]Stakey,_tmp704)
  if len(skc) > 0 { _tmp706 = (*Stakey)(&skc[0]) }
  var _tmp707 C.int32_t
  if _tmp708 := C.MSK_getnumvar(self.ptr(),(&_tmp707)); _tmp708 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp708)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp709 *Stakey
  skx = make([]Stakey,_tmp707)
  if len(skx) > 0 { _tmp709 = (*Stakey)(&skx[0]) }
  var _tmp710 C.int32_t
  if _tmp711 := C.MSK_getnumcone(self.ptr(),(&_tmp710)); _tmp711 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp711)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp712 *Stakey
  skn = make([]Stakey,_tmp710)
  if len(skn) > 0 { _tmp712 = (*Stakey)(&skn[0]) }
  var _tmp713 C.int32_t
  if _tmp714 := C.MSK_getnumcon(self.ptr(),(&_tmp713)); _tmp714 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp714)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp715 *float64
  xc = make([]float64,_tmp713)
  if len(xc) > 0 { _tmp715 = (*float64)(&xc[0]) }
  var _tmp716 C.int32_t
  if _tmp717 := C.MSK_getnumvar(self.ptr(),(&_tmp716)); _tmp717 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp717)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp718 *float64
  xx = make([]float64,_tmp716)
  if len(xx) > 0 { _tmp718 = (*float64)(&xx[0]) }
  var _tmp719 C.int32_t
  if _tmp720 := C.MSK_getnumcon(self.ptr(),(&_tmp719)); _tmp720 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp720)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp721 *float64
  y = make([]float64,_tmp719)
  if len(y) > 0 { _tmp721 = (*float64)(&y[0]) }
  var _tmp722 C.int32_t
  if _tmp723 := C.MSK_getnumcon(self.ptr(),(&_tmp722)); _tmp723 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp723)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp724 *float64
  slc = make([]float64,_tmp722)
  if len(slc) > 0 { _tmp724 = (*float64)(&slc[0]) }
  var _tmp725 C.int32_t
  if _tmp726 := C.MSK_getnumcon(self.ptr(),(&_tmp725)); _tmp726 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp726)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp727 *float64
  suc = make([]float64,_tmp725)
  if len(suc) > 0 { _tmp727 = (*float64)(&suc[0]) }
  var _tmp728 C.int32_t
  if _tmp729 := C.MSK_getnumvar(self.ptr(),(&_tmp728)); _tmp729 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp729)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp730 *float64
  slx = make([]float64,_tmp728)
  if len(slx) > 0 { _tmp730 = (*float64)(&slx[0]) }
  var _tmp731 C.int32_t
  if _tmp732 := C.MSK_getnumvar(self.ptr(),(&_tmp731)); _tmp732 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp732)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp733 *float64
  sux = make([]float64,_tmp731)
  if len(sux) > 0 { _tmp733 = (*float64)(&sux[0]) }
  var _tmp734 C.int32_t
  if _tmp735 := C.MSK_getnumvar(self.ptr(),(&_tmp734)); _tmp735 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp735)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp736 *float64
  snx = make([]float64,_tmp734)
  if len(snx) > 0 { _tmp736 = (*float64)(&snx[0]) }
  var _tmp737 C.int64_t
  if _tmp738 := C.MSK_getaccntot(self.ptr(),(&_tmp737)); _tmp738 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp738)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp739 *float64
  doty = make([]float64,_tmp737)
  if len(doty) > 0 { _tmp739 = (*float64)(&doty[0]) }
  if _tmp740 := C.MSK_getsolutionnew(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(&_tmp702),(*C.int32_t)(&_tmp703),(*C.int32_t)(_tmp706),(*C.int32_t)(_tmp709),(*C.int32_t)(_tmp712),(*C.double)(_tmp715),(*C.double)(_tmp718),(*C.double)(_tmp721),(*C.double)(_tmp724),(*C.double)(_tmp727),(*C.double)(_tmp730),(*C.double)(_tmp733),(*C.double)(_tmp736),(*C.double)(_tmp739)); _tmp740 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp740)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  problemsta = Prosta(_tmp702)
  solutionsta = Solsta(_tmp703)
  return
}
func (self *Task) GetSolutionSlice(whichsol Soltype,solitem Solitem,first int32,last int32) (values []float64,err error) {
  var _tmp741 *float64
  values = make([]float64,(last - first))
  if len(values) > 0 { _tmp741 = (*float64)(&values[0]) }
  if _tmp742 := C.MSK_getsolutionslice(self.ptr(),C.int32_t(whichsol),C.int32_t(solitem),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp741)); _tmp742 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp742)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSparseSymMat(idx int64) (subi []int32,subj []int32,valij []float64,err error) {
  var _tmp743 C.int32_t
  var _tmp744 C.int64_t
  var _tmp745 C.int32_t
  if _tmp746 := C.MSK_getsymmatinfo(self.ptr(),(C.int64_t)(idx),(&_tmp743),(&_tmp744),(&_tmp745)); _tmp746 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp746)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxlen int64 = int64(_tmp744)
  var _tmp747 *int32
  subi = make([]int32,maxlen)
  if len(subi) > 0 { _tmp747 = (*int32)(&subi[0]) }
  var _tmp748 *int32
  subj = make([]int32,maxlen)
  if len(subj) > 0 { _tmp748 = (*int32)(&subj[0]) }
  var _tmp749 *float64
  valij = make([]float64,maxlen)
  if len(valij) > 0 { _tmp749 = (*float64)(&valij[0]) }
  if _tmp750 := C.MSK_getsparsesymmat(self.ptr(),C.int64_t(idx),C.int64_t(maxlen),(*C.int32_t)(_tmp747),(*C.int32_t)(_tmp748),(*C.double)(_tmp749)); _tmp750 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp750)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetStrParam(param Sparam) (len int32,parvalue string,err error) {
  var _tmp751 C.int32_t
  if _tmp752 := C.MSK_getstrparamlen(self.ptr(),(C.int32_t)(param),(&_tmp751)); _tmp752 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp752)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var maxlen int32 = int32((1 + _tmp751))
  _tmp753 := make([]byte,maxlen)
  if _tmp754 := C.MSK_getstrparam(self.ptr(),C.int32_t(param),C.int32_t(maxlen),(*C.int32_t)(&len),(*C.uchar)(&_tmp753[0])); _tmp754 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp754)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp753,byte(0)); p < 0 {
    parvalue = string(_tmp753)
  } else {
    parvalue = string(_tmp753[:p])
  }
  return
}
func (self *Task) GetStrParamLen(param Sparam) (len int32,err error) {
  if _tmp755 := C.MSK_getstrparamlen(self.ptr(),C.int32_t(param),(*C.int32_t)(&len)); _tmp755 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp755)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSuc(whichsol Soltype) (suc []float64,err error) {
  var _tmp756 C.int32_t
  if _tmp757 := C.MSK_getnumcon(self.ptr(),(&_tmp756)); _tmp757 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp757)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp758 *float64
  suc = make([]float64,_tmp756)
  if len(suc) > 0 { _tmp758 = (*float64)(&suc[0]) }
  if _tmp759 := C.MSK_getsuc(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp758)); _tmp759 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp759)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSucSlice(whichsol Soltype,first int32,last int32) (suc []float64,err error) {
  var _tmp760 *float64
  suc = make([]float64,(last - first))
  if len(suc) > 0 { _tmp760 = (*float64)(&suc[0]) }
  if _tmp761 := C.MSK_getsucslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp760)); _tmp761 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp761)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSux(whichsol Soltype) (sux []float64,err error) {
  var _tmp762 C.int32_t
  if _tmp763 := C.MSK_getnumvar(self.ptr(),(&_tmp762)); _tmp763 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp763)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp764 *float64
  sux = make([]float64,_tmp762)
  if len(sux) > 0 { _tmp764 = (*float64)(&sux[0]) }
  if _tmp765 := C.MSK_getsux(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp764)); _tmp765 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp765)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSuxSlice(whichsol Soltype,first int32,last int32) (sux []float64,err error) {
  var _tmp766 *float64
  sux = make([]float64,(last - first))
  if len(sux) > 0 { _tmp766 = (*float64)(&sux[0]) }
  if _tmp767 := C.MSK_getsuxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp766)); _tmp767 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp767)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSymMatInfo(idx int64) (dim int32,nz int64,mattype Symmattype,err error) {
  var _tmp768 int32
  if _tmp769 := C.MSK_getsymmatinfo(self.ptr(),C.int64_t(idx),(*C.int32_t)(&dim),(*C.int64_t)(&nz),(*C.int32_t)(&_tmp768)); _tmp769 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp769)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  mattype = Symmattype(_tmp768)
  return
}
func (self *Task) GetTaskName() (taskname string,err error) {
  var _tmp770 C.int32_t
  if _tmp771 := C.MSK_gettasknamelen(self.ptr(),(&_tmp770)); _tmp771 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp771)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizetaskname int32 = int32((1 + _tmp770))
  _tmp772 := make([]byte,sizetaskname)
  if _tmp773 := C.MSK_gettaskname(self.ptr(),C.int32_t(sizetaskname),(*C.uchar)(&_tmp772[0])); _tmp773 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp773)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp772,byte(0)); p < 0 {
    taskname = string(_tmp772)
  } else {
    taskname = string(_tmp772[:p])
  }
  return
}
func (self *Task) GetTaskNameLen() (len int32,err error) {
  if _tmp774 := C.MSK_gettasknamelen(self.ptr(),(*C.int32_t)(&len)); _tmp774 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp774)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetVarBound(i int32) (bk Boundkey,bl float64,bu float64,err error) {
  var _tmp775 int32
  if _tmp776 := C.MSK_getvarbound(self.ptr(),C.int32_t(i),(*C.int32_t)(&_tmp775),(*C.double)(&bl),(*C.double)(&bu)); _tmp776 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp776)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  bk = Boundkey(_tmp775)
  return
}
func (self *Task) GetVarBoundSlice(first int32,last int32) (bk []Boundkey,bl []float64,bu []float64,err error) {
  var _tmp777 *Boundkey
  bk = make([]Boundkey,(last - first))
  if len(bk) > 0 { _tmp777 = (*Boundkey)(&bk[0]) }
  var _tmp778 *float64
  bl = make([]float64,(last - first))
  if len(bl) > 0 { _tmp778 = (*float64)(&bl[0]) }
  var _tmp779 *float64
  bu = make([]float64,(last - first))
  if len(bu) > 0 { _tmp779 = (*float64)(&bu[0]) }
  if _tmp780 := C.MSK_getvarboundslice(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.int32_t)(_tmp777),(*C.double)(_tmp778),(*C.double)(_tmp779)); _tmp780 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp780)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetVarName(j int32) (name string,err error) {
  var _tmp781 C.int32_t
  if _tmp782 := C.MSK_getvarnamelen(self.ptr(),(C.int32_t)(j),(&_tmp781)); _tmp782 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp782)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp781))
  _tmp783 := make([]byte,sizename)
  if _tmp784 := C.MSK_getvarname(self.ptr(),C.int32_t(j),C.int32_t(sizename),(*C.uchar)(&_tmp783[0])); _tmp784 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp784)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp783,byte(0)); p < 0 {
    name = string(_tmp783)
  } else {
    name = string(_tmp783[:p])
  }
  return
}
func (self *Task) GetVarNameIndex(somename string) (asgn int32,index int32,err error) {
  _tmp785 := C.CString(somename)
  if _tmp786 := C.MSK_getvarnameindex(self.ptr(),_tmp785,(*C.int32_t)(&asgn),(*C.int32_t)(&index)); _tmp786 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp786)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetVarNameLen(i int32) (len int32,err error) {
  if _tmp787 := C.MSK_getvarnamelen(self.ptr(),C.int32_t(i),(*C.int32_t)(&len)); _tmp787 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp787)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetVarType(j int32) (vartype Variabletype,err error) {
  var _tmp788 int32
  if _tmp789 := C.MSK_getvartype(self.ptr(),C.int32_t(j),(*C.int32_t)(&_tmp788)); _tmp789 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp789)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  vartype = Variabletype(_tmp788)
  return
}
func (self *Task) GetVarTypeList(subj []int32) (vartype []Variabletype,err error) {
  _tmp790 := len(subj)
  var num int32 = int32(_tmp790)
  var _tmp791 *int32
  if subj != nil { _tmp791 = (*int32)(&subj[0]) }
  var _tmp792 *Variabletype
  vartype = make([]Variabletype,num)
  if len(vartype) > 0 { _tmp792 = (*Variabletype)(&vartype[0]) }
  if _tmp793 := C.MSK_getvartypelist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp791),(*C.int32_t)(_tmp792)); _tmp793 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp793)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetXc(whichsol Soltype) (xc []float64,err error) {
  var _tmp794 C.int32_t
  if _tmp795 := C.MSK_getnumcon(self.ptr(),(&_tmp794)); _tmp795 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp795)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp796 *float64
  xc = make([]float64,_tmp794)
  if len(xc) > 0 { _tmp796 = (*float64)(&xc[0]) }
  if _tmp797 := C.MSK_getxc(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp796)); _tmp797 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp797)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetXcSlice(whichsol Soltype,first int32,last int32) (xc []float64,err error) {
  var _tmp798 *float64
  xc = make([]float64,(last - first))
  if len(xc) > 0 { _tmp798 = (*float64)(&xc[0]) }
  if _tmp799 := C.MSK_getxcslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp798)); _tmp799 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp799)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetXx(whichsol Soltype) (xx []float64,err error) {
  var _tmp800 C.int32_t
  if _tmp801 := C.MSK_getnumvar(self.ptr(),(&_tmp800)); _tmp801 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp801)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp802 *float64
  xx = make([]float64,_tmp800)
  if len(xx) > 0 { _tmp802 = (*float64)(&xx[0]) }
  if _tmp803 := C.MSK_getxx(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp802)); _tmp803 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp803)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetXxSlice(whichsol Soltype,first int32,last int32) (xx []float64,err error) {
  var _tmp804 *float64
  xx = make([]float64,(last - first))
  if len(xx) > 0 { _tmp804 = (*float64)(&xx[0]) }
  if _tmp805 := C.MSK_getxxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp804)); _tmp805 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp805)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetY(whichsol Soltype) (y []float64,err error) {
  var _tmp806 C.int32_t
  if _tmp807 := C.MSK_getnumcon(self.ptr(),(&_tmp806)); _tmp807 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp807)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp808 *float64
  y = make([]float64,_tmp806)
  if len(y) > 0 { _tmp808 = (*float64)(&y[0]) }
  if _tmp809 := C.MSK_gety(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp808)); _tmp809 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp809)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetYSlice(whichsol Soltype,first int32,last int32) (y []float64,err error) {
  var _tmp810 *float64
  y = make([]float64,(last - first))
  if len(y) > 0 { _tmp810 = (*float64)(&y[0]) }
  if _tmp811 := C.MSK_getyslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp810)); _tmp811 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp811)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) InfeasibilityReport(whichstream Streamtype,whichsol Soltype) (err error) {
  if _tmp812 := C.MSK_infeasibilityreport(self.ptr(),C.int32_t(whichstream),C.int32_t(whichsol)); _tmp812 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp812)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) InitBasisSolve() (basis []int32,err error) {
  var _tmp813 C.int32_t
  if _tmp814 := C.MSK_getnumcon(self.ptr(),(&_tmp813)); _tmp814 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp814)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp815 *int32
  basis = make([]int32,_tmp813)
  if len(basis) > 0 { _tmp815 = (*int32)(&basis[0]) }
  if _tmp816 := C.MSK_initbasissolve(self.ptr(),(*C.int32_t)(_tmp815)); _tmp816 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp816)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) InputData(maxnumcon int32,maxnumvar int32,c []float64,cfix float64,aptrb []int64,aptre []int64,asub []int32,aval []float64,bkc []Boundkey,blc []float64,buc []float64,bkx []Boundkey,blx []float64,bux []float64) (err error) {
  _tmp817 := len(buc)
  if _tmp817 < len(bkc) { _tmp817 = len(bkc) }
  if _tmp817 < len(blc) { _tmp817 = len(blc) }
  var numcon int32 = int32(_tmp817)
  _tmp818 := len(aptrb)
  if _tmp818 < len(blx) { _tmp818 = len(blx) }
  if _tmp818 < len(bux) { _tmp818 = len(bux) }
  if _tmp818 < len(bkx) { _tmp818 = len(bkx) }
  if _tmp818 < len(aptre) { _tmp818 = len(aptre) }
  if _tmp818 < len(c) { _tmp818 = len(c) }
  var numvar int32 = int32(_tmp818)
  var _tmp819 *float64
  if c != nil { _tmp819 = (*float64)(&c[0]) }
  var _tmp820 *int64
  if aptrb != nil { _tmp820 = (*int64)(&aptrb[0]) }
  var _tmp821 *int64
  if aptre != nil { _tmp821 = (*int64)(&aptre[0]) }
  var _tmp822 *int32
  if asub != nil { _tmp822 = (*int32)(&asub[0]) }
  var _tmp823 *float64
  if aval != nil { _tmp823 = (*float64)(&aval[0]) }
  var _tmp824 *Boundkey
  if bkc != nil { _tmp824 = (*Boundkey)(&bkc[0]) }
  var _tmp825 *float64
  if blc != nil { _tmp825 = (*float64)(&blc[0]) }
  var _tmp826 *float64
  if buc != nil { _tmp826 = (*float64)(&buc[0]) }
  var _tmp827 *Boundkey
  if bkx != nil { _tmp827 = (*Boundkey)(&bkx[0]) }
  var _tmp828 *float64
  if blx != nil { _tmp828 = (*float64)(&blx[0]) }
  var _tmp829 *float64
  if bux != nil { _tmp829 = (*float64)(&bux[0]) }
  if _tmp830 := C.MSK_inputdata64(self.ptr(),C.int32_t(maxnumcon),C.int32_t(maxnumvar),C.int32_t(numcon),C.int32_t(numvar),(*C.double)(_tmp819),C.double(cfix),(*C.int64_t)(_tmp820),(*C.int64_t)(_tmp821),(*C.int32_t)(_tmp822),(*C.double)(_tmp823),(*C.int32_t)(_tmp824),(*C.double)(_tmp825),(*C.double)(_tmp826),(*C.int32_t)(_tmp827),(*C.double)(_tmp828),(*C.double)(_tmp829)); _tmp830 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp830)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) IsDouParName(parname string) (param Dparam,err error) {
  _tmp831 := C.CString(parname)
  var _tmp832 int32
  if _tmp833 := C.MSK_isdouparname(self.ptr(),_tmp831,(*C.int32_t)(&_tmp832)); _tmp833 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp833)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  param = Dparam(_tmp832)
  return
}
func (self *Task) IsIntParName(parname string) (param Iparam,err error) {
  _tmp834 := C.CString(parname)
  var _tmp835 int32
  if _tmp836 := C.MSK_isintparname(self.ptr(),_tmp834,(*C.int32_t)(&_tmp835)); _tmp836 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp836)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  param = Iparam(_tmp835)
  return
}
func (self *Task) IsStrParName(parname string) (param Sparam,err error) {
  _tmp837 := C.CString(parname)
  var _tmp838 int32
  if _tmp839 := C.MSK_isstrparname(self.ptr(),_tmp837,(*C.int32_t)(&_tmp838)); _tmp839 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp839)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  param = Sparam(_tmp838)
  return
}
func (self *Task) LinkFileToStream(whichstream Streamtype,filename string,append int32) (err error) {
  _tmp840 := C.CString(filename)
  if _tmp841 := C.MSK_linkfiletotaskstream(self.ptr(),C.int32_t(whichstream),_tmp840,C.int32_t(append)); _tmp841 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp841)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) OneSolutionSummary(whichstream Streamtype,whichsol Soltype) (err error) {
  if _tmp842 := C.MSK_onesolutionsummary(self.ptr(),C.int32_t(whichstream),C.int32_t(whichsol)); _tmp842 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp842)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) OptimizerSummary(whichstream Streamtype) (err error) {
  if _tmp843 := C.MSK_optimizersummary(self.ptr(),C.int32_t(whichstream)); _tmp843 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp843)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) Optimize() (trmcode Rescode,err error) {
  var _tmp844 int32
  if _tmp845 := C.MSK_optimizetrm(self.ptr(),(*C.int32_t)(&_tmp844)); _tmp845 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp845)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  trmcode = Rescode(_tmp844)
  return
}
func (self *Task) PrimalRepair(wlc []float64,wuc []float64,wlx []float64,wux []float64) (err error) {
  var _tmp846 C.int32_t
  if _tmp847 := C.MSK_getnumcon(self.ptr(),(&_tmp846)); _tmp847 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp847)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(wlc)) < int64(_tmp846) {
    err = &ArrayLengthError{fun:"PrimalRepair",arg:"wlc"}
    return
  }
  var _tmp848 *float64
  if wlc != nil { _tmp848 = (*float64)(&wlc[0]) }
  var _tmp849 C.int32_t
  if _tmp850 := C.MSK_getnumcon(self.ptr(),(&_tmp849)); _tmp850 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp850)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(wuc)) < int64(_tmp849) {
    err = &ArrayLengthError{fun:"PrimalRepair",arg:"wuc"}
    return
  }
  var _tmp851 *float64
  if wuc != nil { _tmp851 = (*float64)(&wuc[0]) }
  var _tmp852 C.int32_t
  if _tmp853 := C.MSK_getnumvar(self.ptr(),(&_tmp852)); _tmp853 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp853)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(wlx)) < int64(_tmp852) {
    err = &ArrayLengthError{fun:"PrimalRepair",arg:"wlx"}
    return
  }
  var _tmp854 *float64
  if wlx != nil { _tmp854 = (*float64)(&wlx[0]) }
  var _tmp855 C.int32_t
  if _tmp856 := C.MSK_getnumvar(self.ptr(),(&_tmp855)); _tmp856 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp856)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(wux)) < int64(_tmp855) {
    err = &ArrayLengthError{fun:"PrimalRepair",arg:"wux"}
    return
  }
  var _tmp857 *float64
  if wux != nil { _tmp857 = (*float64)(&wux[0]) }
  if _tmp858 := C.MSK_primalrepair(self.ptr(),(*C.double)(_tmp848),(*C.double)(_tmp851),(*C.double)(_tmp854),(*C.double)(_tmp857)); _tmp858 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp858)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PrimalSensitivity(subi []int32,marki []Mark,subj []int32,markj []Mark) (leftpricei []float64,rightpricei []float64,leftrangei []float64,rightrangei []float64,leftpricej []float64,rightpricej []float64,leftrangej []float64,rightrangej []float64,err error) {
  _tmp859 := len(subi)
  if _tmp859 < len(marki) { _tmp859 = len(marki) }
  var numi int32 = int32(_tmp859)
  var _tmp860 *int32
  if subi != nil { _tmp860 = (*int32)(&subi[0]) }
  var _tmp861 *Mark
  if marki != nil { _tmp861 = (*Mark)(&marki[0]) }
  _tmp862 := len(markj)
  if _tmp862 < len(subj) { _tmp862 = len(subj) }
  var numj int32 = int32(_tmp862)
  var _tmp863 *int32
  if subj != nil { _tmp863 = (*int32)(&subj[0]) }
  var _tmp864 *Mark
  if markj != nil { _tmp864 = (*Mark)(&markj[0]) }
  var _tmp865 *float64
  leftpricei = make([]float64,numi)
  if len(leftpricei) > 0 { _tmp865 = (*float64)(&leftpricei[0]) }
  var _tmp866 *float64
  rightpricei = make([]float64,numi)
  if len(rightpricei) > 0 { _tmp866 = (*float64)(&rightpricei[0]) }
  var _tmp867 *float64
  leftrangei = make([]float64,numi)
  if len(leftrangei) > 0 { _tmp867 = (*float64)(&leftrangei[0]) }
  var _tmp868 *float64
  rightrangei = make([]float64,numi)
  if len(rightrangei) > 0 { _tmp868 = (*float64)(&rightrangei[0]) }
  var _tmp869 *float64
  leftpricej = make([]float64,numj)
  if len(leftpricej) > 0 { _tmp869 = (*float64)(&leftpricej[0]) }
  var _tmp870 *float64
  rightpricej = make([]float64,numj)
  if len(rightpricej) > 0 { _tmp870 = (*float64)(&rightpricej[0]) }
  var _tmp871 *float64
  leftrangej = make([]float64,numj)
  if len(leftrangej) > 0 { _tmp871 = (*float64)(&leftrangej[0]) }
  var _tmp872 *float64
  rightrangej = make([]float64,numj)
  if len(rightrangej) > 0 { _tmp872 = (*float64)(&rightrangej[0]) }
  if _tmp873 := C.MSK_primalsensitivity(self.ptr(),C.int32_t(numi),(*C.int32_t)(_tmp860),(*C.int32_t)(_tmp861),C.int32_t(numj),(*C.int32_t)(_tmp863),(*C.int32_t)(_tmp864),(*C.double)(_tmp865),(*C.double)(_tmp866),(*C.double)(_tmp867),(*C.double)(_tmp868),(*C.double)(_tmp869),(*C.double)(_tmp870),(*C.double)(_tmp871),(*C.double)(_tmp872)); _tmp873 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp873)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) ProbTypeToStr(probtype Problemtype) (str string,err error) {
  _tmp874 := make([]byte,MSK_MAX_STR_LEN)
  if _tmp875 := C.MSK_probtypetostr(self.ptr(),C.int32_t(probtype),(*C.uchar)(&_tmp874[0])); _tmp875 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp875)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp874,byte(0)); p < 0 {
    str = string(_tmp874)
  } else {
    str = string(_tmp874[:p])
  }
  return
}
func (self *Task) ProStaToStr(problemsta Prosta) (str string,err error) {
  _tmp876 := make([]byte,MSK_MAX_STR_LEN)
  if _tmp877 := C.MSK_prostatostr(self.ptr(),C.int32_t(problemsta),(*C.uchar)(&_tmp876[0])); _tmp877 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp877)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp876,byte(0)); p < 0 {
    str = string(_tmp876)
  } else {
    str = string(_tmp876[:p])
  }
  return
}
func (self *Task) PutAcc(accidx int64,domidx int64,afeidxlist []int64,b []float64) (err error) {
  _tmp878 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp878)
  var _tmp879 *int64
  if afeidxlist != nil { _tmp879 = (*int64)(&afeidxlist[0]) }
  if int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"PutAcc",arg:"b"}
    return
  }
  var _tmp880 *float64
  if b != nil { _tmp880 = (*float64)(&b[0]) }
  if _tmp881 := C.MSK_putacc(self.ptr(),C.int64_t(accidx),C.int64_t(domidx),C.int64_t(numafeidx),(*C.int64_t)(_tmp879),(*C.double)(_tmp880)); _tmp881 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp881)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAccB(accidx int64,b []float64) (err error) {
  _tmp882 := len(b)
  var lengthb int64 = int64(_tmp882)
  var _tmp883 *float64
  if b != nil { _tmp883 = (*float64)(&b[0]) }
  if _tmp884 := C.MSK_putaccb(self.ptr(),C.int64_t(accidx),C.int64_t(lengthb),(*C.double)(_tmp883)); _tmp884 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp884)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAccBJ(accidx int64,j int64,bj float64) (err error) {
  if _tmp885 := C.MSK_putaccbj(self.ptr(),C.int64_t(accidx),C.int64_t(j),C.double(bj)); _tmp885 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp885)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAccDotY(whichsol Soltype,accidx int64) (doty []float64,err error) {
  var _tmp886 C.int64_t
  if _tmp887 := C.MSK_getaccn(self.ptr(),(C.int64_t)(accidx),(&_tmp886)); _tmp887 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp887)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp888 *float64
  doty = make([]float64,_tmp886)
  if len(doty) > 0 { _tmp888 = (*float64)(&doty[0]) }
  if _tmp889 := C.MSK_putaccdoty(self.ptr(),C.int32_t(whichsol),C.int64_t(accidx),(*C.double)(_tmp888)); _tmp889 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp889)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAccList(accidxs []int64,domidxs []int64,afeidxlist []int64,b []float64) (err error) {
  _tmp890 := len(accidxs)
  if _tmp890 < len(domidxs) { _tmp890 = len(domidxs) }
  var numaccs int64 = int64(_tmp890)
  var _tmp891 *int64
  if accidxs != nil { _tmp891 = (*int64)(&accidxs[0]) }
  var _tmp892 *int64
  if domidxs != nil { _tmp892 = (*int64)(&domidxs[0]) }
  _tmp893 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp893)
  var _tmp894 *int64
  if afeidxlist != nil { _tmp894 = (*int64)(&afeidxlist[0]) }
  if int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"PutAccList",arg:"b"}
    return
  }
  var _tmp895 *float64
  if b != nil { _tmp895 = (*float64)(&b[0]) }
  if _tmp896 := C.MSK_putacclist(self.ptr(),C.int64_t(numaccs),(*C.int64_t)(_tmp891),(*C.int64_t)(_tmp892),C.int64_t(numafeidx),(*C.int64_t)(_tmp894),(*C.double)(_tmp895)); _tmp896 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp896)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAccName(accidx int64,name string) (err error) {
  _tmp897 := C.CString(name)
  if _tmp898 := C.MSK_putaccname(self.ptr(),C.int64_t(accidx),_tmp897); _tmp898 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp898)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutACol(j int32,subj []int32,valj []float64) (err error) {
  _tmp899 := len(valj)
  if _tmp899 < len(subj) { _tmp899 = len(subj) }
  var nzj int32 = int32(_tmp899)
  var _tmp900 *int32
  if subj != nil { _tmp900 = (*int32)(&subj[0]) }
  var _tmp901 *float64
  if valj != nil { _tmp901 = (*float64)(&valj[0]) }
  if _tmp902 := C.MSK_putacol(self.ptr(),C.int32_t(j),C.int32_t(nzj),(*C.int32_t)(_tmp900),(*C.double)(_tmp901)); _tmp902 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp902)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAColList(sub []int32,ptrb []int32,ptre []int32,asub []int32,aval []float64) (err error) {
  _tmp903 := len(ptrb)
  if _tmp903 < len(ptre) { _tmp903 = len(ptre) }
  if _tmp903 < len(sub) { _tmp903 = len(sub) }
  var num int32 = int32(_tmp903)
  var _tmp904 *int32
  if sub != nil { _tmp904 = (*int32)(&sub[0]) }
  var _tmp905 *int32
  if ptrb != nil { _tmp905 = (*int32)(&ptrb[0]) }
  var _tmp906 *int32
  if ptre != nil { _tmp906 = (*int32)(&ptre[0]) }
  var _tmp907 *int32
  if asub != nil { _tmp907 = (*int32)(&asub[0]) }
  var _tmp908 *float64
  if aval != nil { _tmp908 = (*float64)(&aval[0]) }
  if _tmp909 := C.MSK_putacollist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp904),(*C.int32_t)(_tmp905),(*C.int32_t)(_tmp906),(*C.int32_t)(_tmp907),(*C.double)(_tmp908)); _tmp909 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp909)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAColSlice(first int32,last int32,ptrb []int64,ptre []int64,asub []int32,aval []float64) (err error) {
  var _tmp910 *int64
  if ptrb != nil { _tmp910 = (*int64)(&ptrb[0]) }
  var _tmp911 *int64
  if ptre != nil { _tmp911 = (*int64)(&ptre[0]) }
  var _tmp912 *int32
  if asub != nil { _tmp912 = (*int32)(&asub[0]) }
  var _tmp913 *float64
  if aval != nil { _tmp913 = (*float64)(&aval[0]) }
  if _tmp914 := C.MSK_putacolslice64(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.int64_t)(_tmp910),(*C.int64_t)(_tmp911),(*C.int32_t)(_tmp912),(*C.double)(_tmp913)); _tmp914 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp914)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeBarfBlockTriplet(afeidx []int64,barvaridx []int32,subk []int32,subl []int32,valkl []float64) (err error) {
  _tmp915 := len(subl)
  if _tmp915 < len(subk) { _tmp915 = len(subk) }
  if _tmp915 < len(barvaridx) { _tmp915 = len(barvaridx) }
  if _tmp915 < len(valkl) { _tmp915 = len(valkl) }
  if _tmp915 < len(afeidx) { _tmp915 = len(afeidx) }
  var numtrip int64 = int64(_tmp915)
  if int64(len(afeidx)) < int64(numtrip) {
    err = &ArrayLengthError{fun:"PutAfeBarfBlockTriplet",arg:"afeidx"}
    return
  }
  var _tmp916 *int64
  if afeidx != nil { _tmp916 = (*int64)(&afeidx[0]) }
  if int64(len(barvaridx)) < int64(numtrip) {
    err = &ArrayLengthError{fun:"PutAfeBarfBlockTriplet",arg:"barvaridx"}
    return
  }
  var _tmp917 *int32
  if barvaridx != nil { _tmp917 = (*int32)(&barvaridx[0]) }
  if int64(len(subk)) < int64(numtrip) {
    err = &ArrayLengthError{fun:"PutAfeBarfBlockTriplet",arg:"subk"}
    return
  }
  var _tmp918 *int32
  if subk != nil { _tmp918 = (*int32)(&subk[0]) }
  if int64(len(subl)) < int64(numtrip) {
    err = &ArrayLengthError{fun:"PutAfeBarfBlockTriplet",arg:"subl"}
    return
  }
  var _tmp919 *int32
  if subl != nil { _tmp919 = (*int32)(&subl[0]) }
  if int64(len(valkl)) < int64(numtrip) {
    err = &ArrayLengthError{fun:"PutAfeBarfBlockTriplet",arg:"valkl"}
    return
  }
  var _tmp920 *float64
  if valkl != nil { _tmp920 = (*float64)(&valkl[0]) }
  if _tmp921 := C.MSK_putafebarfblocktriplet(self.ptr(),C.int64_t(numtrip),(*C.int64_t)(_tmp916),(*C.int32_t)(_tmp917),(*C.int32_t)(_tmp918),(*C.int32_t)(_tmp919),(*C.double)(_tmp920)); _tmp921 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp921)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeBarfEntry(afeidx int64,barvaridx int32,termidx []int64,termweight []float64) (err error) {
  _tmp922 := len(termidx)
  if _tmp922 < len(termweight) { _tmp922 = len(termweight) }
  var numterm int64 = int64(_tmp922)
  var _tmp923 *int64
  if termidx != nil { _tmp923 = (*int64)(&termidx[0]) }
  var _tmp924 *float64
  if termweight != nil { _tmp924 = (*float64)(&termweight[0]) }
  if _tmp925 := C.MSK_putafebarfentry(self.ptr(),C.int64_t(afeidx),C.int32_t(barvaridx),C.int64_t(numterm),(*C.int64_t)(_tmp923),(*C.double)(_tmp924)); _tmp925 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp925)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeBarfEntryList(afeidx []int64,barvaridx []int32,numterm []int64,ptrterm []int64,termidx []int64,termweight []float64) (err error) {
  _tmp926 := len(barvaridx)
  if _tmp926 < len(ptrterm) { _tmp926 = len(ptrterm) }
  if _tmp926 < len(numterm) { _tmp926 = len(numterm) }
  if _tmp926 < len(afeidx) { _tmp926 = len(afeidx) }
  var numafeidx int64 = int64(_tmp926)
  var _tmp927 *int64
  if afeidx != nil { _tmp927 = (*int64)(&afeidx[0]) }
  var _tmp928 *int32
  if barvaridx != nil { _tmp928 = (*int32)(&barvaridx[0]) }
  var _tmp929 *int64
  if numterm != nil { _tmp929 = (*int64)(&numterm[0]) }
  var _tmp930 *int64
  if ptrterm != nil { _tmp930 = (*int64)(&ptrterm[0]) }
  _tmp931 := len(termidx)
  if _tmp931 < len(termweight) { _tmp931 = len(termweight) }
  var lenterm int64 = int64(_tmp931)
  var _tmp932 *int64
  if termidx != nil { _tmp932 = (*int64)(&termidx[0]) }
  var _tmp933 *float64
  if termweight != nil { _tmp933 = (*float64)(&termweight[0]) }
  if _tmp934 := C.MSK_putafebarfentrylist(self.ptr(),C.int64_t(numafeidx),(*C.int64_t)(_tmp927),(*C.int32_t)(_tmp928),(*C.int64_t)(_tmp929),(*C.int64_t)(_tmp930),C.int64_t(lenterm),(*C.int64_t)(_tmp932),(*C.double)(_tmp933)); _tmp934 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp934)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeBarfRow(afeidx int64,barvaridx []int32,numterm []int64,ptrterm []int64,termidx []int64,termweight []float64) (err error) {
  _tmp935 := len(barvaridx)
  if _tmp935 < len(ptrterm) { _tmp935 = len(ptrterm) }
  if _tmp935 < len(numterm) { _tmp935 = len(numterm) }
  var numentr int32 = int32(_tmp935)
  var _tmp936 *int32
  if barvaridx != nil { _tmp936 = (*int32)(&barvaridx[0]) }
  var _tmp937 *int64
  if numterm != nil { _tmp937 = (*int64)(&numterm[0]) }
  var _tmp938 *int64
  if ptrterm != nil { _tmp938 = (*int64)(&ptrterm[0]) }
  _tmp939 := len(termidx)
  if _tmp939 < len(termweight) { _tmp939 = len(termweight) }
  var lenterm int64 = int64(_tmp939)
  var _tmp940 *int64
  if termidx != nil { _tmp940 = (*int64)(&termidx[0]) }
  var _tmp941 *float64
  if termweight != nil { _tmp941 = (*float64)(&termweight[0]) }
  if _tmp942 := C.MSK_putafebarfrow(self.ptr(),C.int64_t(afeidx),C.int32_t(numentr),(*C.int32_t)(_tmp936),(*C.int64_t)(_tmp937),(*C.int64_t)(_tmp938),C.int64_t(lenterm),(*C.int64_t)(_tmp940),(*C.double)(_tmp941)); _tmp942 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp942)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeFCol(varidx int32,afeidx []int64,val []float64) (err error) {
  _tmp943 := len(val)
  if _tmp943 < len(afeidx) { _tmp943 = len(afeidx) }
  var numnz int64 = int64(_tmp943)
  var _tmp944 *int64
  if afeidx != nil { _tmp944 = (*int64)(&afeidx[0]) }
  var _tmp945 *float64
  if val != nil { _tmp945 = (*float64)(&val[0]) }
  if _tmp946 := C.MSK_putafefcol(self.ptr(),C.int32_t(varidx),C.int64_t(numnz),(*C.int64_t)(_tmp944),(*C.double)(_tmp945)); _tmp946 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp946)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeFEntry(afeidx int64,varidx int32,value float64) (err error) {
  if _tmp947 := C.MSK_putafefentry(self.ptr(),C.int64_t(afeidx),C.int32_t(varidx),C.double(value)); _tmp947 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp947)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeFEntryList(afeidx []int64,varidx []int32,val []float64) (err error) {
  _tmp948 := len(val)
  if _tmp948 < len(varidx) { _tmp948 = len(varidx) }
  if _tmp948 < len(afeidx) { _tmp948 = len(afeidx) }
  var numentr int64 = int64(_tmp948)
  var _tmp949 *int64
  if afeidx != nil { _tmp949 = (*int64)(&afeidx[0]) }
  var _tmp950 *int32
  if varidx != nil { _tmp950 = (*int32)(&varidx[0]) }
  var _tmp951 *float64
  if val != nil { _tmp951 = (*float64)(&val[0]) }
  if _tmp952 := C.MSK_putafefentrylist(self.ptr(),C.int64_t(numentr),(*C.int64_t)(_tmp949),(*C.int32_t)(_tmp950),(*C.double)(_tmp951)); _tmp952 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp952)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeFRow(afeidx int64,varidx []int32,val []float64) (err error) {
  _tmp953 := len(varidx)
  if _tmp953 < len(val) { _tmp953 = len(val) }
  var numnz int32 = int32(_tmp953)
  var _tmp954 *int32
  if varidx != nil { _tmp954 = (*int32)(&varidx[0]) }
  var _tmp955 *float64
  if val != nil { _tmp955 = (*float64)(&val[0]) }
  if _tmp956 := C.MSK_putafefrow(self.ptr(),C.int64_t(afeidx),C.int32_t(numnz),(*C.int32_t)(_tmp954),(*C.double)(_tmp955)); _tmp956 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp956)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeFRowList(afeidx []int64,numnzrow []int32,ptrrow []int64,varidx []int32,val []float64) (err error) {
  _tmp957 := len(ptrrow)
  if _tmp957 < len(numnzrow) { _tmp957 = len(numnzrow) }
  if _tmp957 < len(afeidx) { _tmp957 = len(afeidx) }
  var numafeidx int64 = int64(_tmp957)
  var _tmp958 *int64
  if afeidx != nil { _tmp958 = (*int64)(&afeidx[0]) }
  var _tmp959 *int32
  if numnzrow != nil { _tmp959 = (*int32)(&numnzrow[0]) }
  var _tmp960 *int64
  if ptrrow != nil { _tmp960 = (*int64)(&ptrrow[0]) }
  _tmp961 := len(varidx)
  if _tmp961 < len(val) { _tmp961 = len(val) }
  var lenidxval int64 = int64(_tmp961)
  var _tmp962 *int32
  if varidx != nil { _tmp962 = (*int32)(&varidx[0]) }
  var _tmp963 *float64
  if val != nil { _tmp963 = (*float64)(&val[0]) }
  if _tmp964 := C.MSK_putafefrowlist(self.ptr(),C.int64_t(numafeidx),(*C.int64_t)(_tmp958),(*C.int32_t)(_tmp959),(*C.int64_t)(_tmp960),C.int64_t(lenidxval),(*C.int32_t)(_tmp962),(*C.double)(_tmp963)); _tmp964 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp964)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeG(afeidx int64,g float64) (err error) {
  if _tmp965 := C.MSK_putafeg(self.ptr(),C.int64_t(afeidx),C.double(g)); _tmp965 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp965)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeGList(afeidx []int64,g []float64) (err error) {
  _tmp966 := len(g)
  if _tmp966 < len(afeidx) { _tmp966 = len(afeidx) }
  var numafeidx int64 = int64(_tmp966)
  var _tmp967 *int64
  if afeidx != nil { _tmp967 = (*int64)(&afeidx[0]) }
  var _tmp968 *float64
  if g != nil { _tmp968 = (*float64)(&g[0]) }
  if _tmp969 := C.MSK_putafeglist(self.ptr(),C.int64_t(numafeidx),(*C.int64_t)(_tmp967),(*C.double)(_tmp968)); _tmp969 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp969)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeGSlice(first int64,last int64,slice []float64) (err error) {
  if int64(len(slice)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutAfeGSlice",arg:"slice"}
    return
  }
  var _tmp970 *float64
  if slice != nil { _tmp970 = (*float64)(&slice[0]) }
  if _tmp971 := C.MSK_putafegslice(self.ptr(),C.int64_t(first),C.int64_t(last),(*C.double)(_tmp970)); _tmp971 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp971)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAij(i int32,j int32,aij float64) (err error) {
  if _tmp972 := C.MSK_putaij(self.ptr(),C.int32_t(i),C.int32_t(j),C.double(aij)); _tmp972 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp972)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAijList(subi []int32,subj []int32,valij []float64) (err error) {
  _tmp973 := len(subi)
  if _tmp973 < len(valij) { _tmp973 = len(valij) }
  if _tmp973 < len(subj) { _tmp973 = len(subj) }
  var num int64 = int64(_tmp973)
  var _tmp974 *int32
  if subi != nil { _tmp974 = (*int32)(&subi[0]) }
  var _tmp975 *int32
  if subj != nil { _tmp975 = (*int32)(&subj[0]) }
  var _tmp976 *float64
  if valij != nil { _tmp976 = (*float64)(&valij[0]) }
  if _tmp977 := C.MSK_putaijlist64(self.ptr(),C.int64_t(num),(*C.int32_t)(_tmp974),(*C.int32_t)(_tmp975),(*C.double)(_tmp976)); _tmp977 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp977)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutARow(i int32,subi []int32,vali []float64) (err error) {
  _tmp978 := len(vali)
  if _tmp978 < len(subi) { _tmp978 = len(subi) }
  var nzi int32 = int32(_tmp978)
  var _tmp979 *int32
  if subi != nil { _tmp979 = (*int32)(&subi[0]) }
  var _tmp980 *float64
  if vali != nil { _tmp980 = (*float64)(&vali[0]) }
  if _tmp981 := C.MSK_putarow(self.ptr(),C.int32_t(i),C.int32_t(nzi),(*C.int32_t)(_tmp979),(*C.double)(_tmp980)); _tmp981 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp981)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutARowList(sub []int32,ptrb []int64,ptre []int64,asub []int32,aval []float64) (err error) {
  _tmp982 := len(ptrb)
  if _tmp982 < len(ptre) { _tmp982 = len(ptre) }
  if _tmp982 < len(sub) { _tmp982 = len(sub) }
  var num int32 = int32(_tmp982)
  var _tmp983 *int32
  if sub != nil { _tmp983 = (*int32)(&sub[0]) }
  var _tmp984 *int64
  if ptrb != nil { _tmp984 = (*int64)(&ptrb[0]) }
  var _tmp985 *int64
  if ptre != nil { _tmp985 = (*int64)(&ptre[0]) }
  var _tmp986 *int32
  if asub != nil { _tmp986 = (*int32)(&asub[0]) }
  var _tmp987 *float64
  if aval != nil { _tmp987 = (*float64)(&aval[0]) }
  if _tmp988 := C.MSK_putarowlist64(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp983),(*C.int64_t)(_tmp984),(*C.int64_t)(_tmp985),(*C.int32_t)(_tmp986),(*C.double)(_tmp987)); _tmp988 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp988)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutARowSlice(first int32,last int32,ptrb []int64,ptre []int64,asub []int32,aval []float64) (err error) {
  if int64(len(ptrb)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutARowSlice",arg:"ptrb"}
    return
  }
  var _tmp989 *int64
  if ptrb != nil { _tmp989 = (*int64)(&ptrb[0]) }
  if int64(len(ptre)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutARowSlice",arg:"ptre"}
    return
  }
  var _tmp990 *int64
  if ptre != nil { _tmp990 = (*int64)(&ptre[0]) }
  var _tmp991 *int32
  if asub != nil { _tmp991 = (*int32)(&asub[0]) }
  var _tmp992 *float64
  if aval != nil { _tmp992 = (*float64)(&aval[0]) }
  if _tmp993 := C.MSK_putarowslice64(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.int64_t)(_tmp989),(*C.int64_t)(_tmp990),(*C.int32_t)(_tmp991),(*C.double)(_tmp992)); _tmp993 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp993)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutATruncateTol(tolzero float64) (err error) {
  if _tmp994 := C.MSK_putatruncatetol(self.ptr(),C.double(tolzero)); _tmp994 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp994)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBaraBlockTriplet(subi []int32,subj []int32,subk []int32,subl []int32,valijkl []float64) (err error) {
  _tmp995 := len(subl)
  if _tmp995 < len(valijkl) { _tmp995 = len(valijkl) }
  if _tmp995 < len(subk) { _tmp995 = len(subk) }
  if _tmp995 < len(subj) { _tmp995 = len(subj) }
  var num int64 = int64(_tmp995)
  if int64(len(subi)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBaraBlockTriplet",arg:"subi"}
    return
  }
  var _tmp996 *int32
  if subi != nil { _tmp996 = (*int32)(&subi[0]) }
  if int64(len(subj)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBaraBlockTriplet",arg:"subj"}
    return
  }
  var _tmp997 *int32
  if subj != nil { _tmp997 = (*int32)(&subj[0]) }
  if int64(len(subk)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBaraBlockTriplet",arg:"subk"}
    return
  }
  var _tmp998 *int32
  if subk != nil { _tmp998 = (*int32)(&subk[0]) }
  if int64(len(subl)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBaraBlockTriplet",arg:"subl"}
    return
  }
  var _tmp999 *int32
  if subl != nil { _tmp999 = (*int32)(&subl[0]) }
  if int64(len(valijkl)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBaraBlockTriplet",arg:"valijkl"}
    return
  }
  var _tmp1000 *float64
  if valijkl != nil { _tmp1000 = (*float64)(&valijkl[0]) }
  if _tmp1001 := C.MSK_putbarablocktriplet(self.ptr(),C.int64_t(num),(*C.int32_t)(_tmp996),(*C.int32_t)(_tmp997),(*C.int32_t)(_tmp998),(*C.int32_t)(_tmp999),(*C.double)(_tmp1000)); _tmp1001 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1001)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBaraIj(i int32,j int32,sub []int64,weights []float64) (err error) {
  _tmp1002 := len(weights)
  if _tmp1002 < len(sub) { _tmp1002 = len(sub) }
  var num int64 = int64(_tmp1002)
  var _tmp1003 *int64
  if sub != nil { _tmp1003 = (*int64)(&sub[0]) }
  var _tmp1004 *float64
  if weights != nil { _tmp1004 = (*float64)(&weights[0]) }
  if _tmp1005 := C.MSK_putbaraij(self.ptr(),C.int32_t(i),C.int32_t(j),C.int64_t(num),(*C.int64_t)(_tmp1003),(*C.double)(_tmp1004)); _tmp1005 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1005)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBaraIjList(subi []int32,subj []int32,alphaptrb []int64,alphaptre []int64,matidx []int64,weights []float64) (err error) {
  _tmp1006 := len(alphaptrb)
  if _tmp1006 < len(subi) { _tmp1006 = len(subi) }
  if _tmp1006 < len(alphaptre) { _tmp1006 = len(alphaptre) }
  if _tmp1006 < len(subj) { _tmp1006 = len(subj) }
  var num int32 = int32(_tmp1006)
  var _tmp1007 *int32
  if subi != nil { _tmp1007 = (*int32)(&subi[0]) }
  var _tmp1008 *int32
  if subj != nil { _tmp1008 = (*int32)(&subj[0]) }
  var _tmp1009 *int64
  if alphaptrb != nil { _tmp1009 = (*int64)(&alphaptrb[0]) }
  var _tmp1010 *int64
  if alphaptre != nil { _tmp1010 = (*int64)(&alphaptre[0]) }
  var _tmp1011 *int64
  if matidx != nil { _tmp1011 = (*int64)(&matidx[0]) }
  var _tmp1012 *float64
  if weights != nil { _tmp1012 = (*float64)(&weights[0]) }
  if _tmp1013 := C.MSK_putbaraijlist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1007),(*C.int32_t)(_tmp1008),(*C.int64_t)(_tmp1009),(*C.int64_t)(_tmp1010),(*C.int64_t)(_tmp1011),(*C.double)(_tmp1012)); _tmp1013 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1013)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBaraRowList(subi []int32,ptrb []int64,ptre []int64,subj []int32,nummat []int64,matidx []int64,weights []float64) (err error) {
  _tmp1014 := len(ptrb)
  if _tmp1014 < len(ptre) { _tmp1014 = len(ptre) }
  if _tmp1014 < len(subi) { _tmp1014 = len(subi) }
  var num int32 = int32(_tmp1014)
  var _tmp1015 *int32
  if subi != nil { _tmp1015 = (*int32)(&subi[0]) }
  var _tmp1016 *int64
  if ptrb != nil { _tmp1016 = (*int64)(&ptrb[0]) }
  var _tmp1017 *int64
  if ptre != nil { _tmp1017 = (*int64)(&ptre[0]) }
  var _tmp1018 *int32
  if subj != nil { _tmp1018 = (*int32)(&subj[0]) }
  if int64(len(nummat)) < int64(len(subj)) {
    err = &ArrayLengthError{fun:"PutBaraRowList",arg:"nummat"}
    return
  }
  var _tmp1019 *int64
  if nummat != nil { _tmp1019 = (*int64)(&nummat[0]) }
  if int64(len(matidx)) < int64(sum(nummat)) {
    err = &ArrayLengthError{fun:"PutBaraRowList",arg:"matidx"}
    return
  }
  var _tmp1020 *int64
  if matidx != nil { _tmp1020 = (*int64)(&matidx[0]) }
  if int64(len(weights)) < int64(sum(nummat)) {
    err = &ArrayLengthError{fun:"PutBaraRowList",arg:"weights"}
    return
  }
  var _tmp1021 *float64
  if weights != nil { _tmp1021 = (*float64)(&weights[0]) }
  if _tmp1022 := C.MSK_putbararowlist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1015),(*C.int64_t)(_tmp1016),(*C.int64_t)(_tmp1017),(*C.int32_t)(_tmp1018),(*C.int64_t)(_tmp1019),(*C.int64_t)(_tmp1020),(*C.double)(_tmp1021)); _tmp1022 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1022)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBarcBlockTriplet(subj []int32,subk []int32,subl []int32,valjkl []float64) (err error) {
  _tmp1023 := len(valjkl)
  if _tmp1023 < len(subl) { _tmp1023 = len(subl) }
  if _tmp1023 < len(subk) { _tmp1023 = len(subk) }
  if _tmp1023 < len(subj) { _tmp1023 = len(subj) }
  var num int64 = int64(_tmp1023)
  if int64(len(subj)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBarcBlockTriplet",arg:"subj"}
    return
  }
  var _tmp1024 *int32
  if subj != nil { _tmp1024 = (*int32)(&subj[0]) }
  if int64(len(subk)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBarcBlockTriplet",arg:"subk"}
    return
  }
  var _tmp1025 *int32
  if subk != nil { _tmp1025 = (*int32)(&subk[0]) }
  if int64(len(subl)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBarcBlockTriplet",arg:"subl"}
    return
  }
  var _tmp1026 *int32
  if subl != nil { _tmp1026 = (*int32)(&subl[0]) }
  if int64(len(valjkl)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBarcBlockTriplet",arg:"valjkl"}
    return
  }
  var _tmp1027 *float64
  if valjkl != nil { _tmp1027 = (*float64)(&valjkl[0]) }
  if _tmp1028 := C.MSK_putbarcblocktriplet(self.ptr(),C.int64_t(num),(*C.int32_t)(_tmp1024),(*C.int32_t)(_tmp1025),(*C.int32_t)(_tmp1026),(*C.double)(_tmp1027)); _tmp1028 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1028)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBarcJ(j int32,sub []int64,weights []float64) (err error) {
  _tmp1029 := len(weights)
  if _tmp1029 < len(sub) { _tmp1029 = len(sub) }
  var num int64 = int64(_tmp1029)
  var _tmp1030 *int64
  if sub != nil { _tmp1030 = (*int64)(&sub[0]) }
  var _tmp1031 *float64
  if weights != nil { _tmp1031 = (*float64)(&weights[0]) }
  if _tmp1032 := C.MSK_putbarcj(self.ptr(),C.int32_t(j),C.int64_t(num),(*C.int64_t)(_tmp1030),(*C.double)(_tmp1031)); _tmp1032 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1032)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBarsJ(whichsol Soltype,j int32,barsj []float64) (err error) {
  var _tmp1033 C.int64_t
  if _tmp1034 := C.MSK_getlenbarvarj(self.ptr(),(C.int32_t)(j),(&_tmp1033)); _tmp1034 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1034)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(barsj)) < int64(_tmp1033) {
    err = &ArrayLengthError{fun:"PutBarsJ",arg:"barsj"}
    return
  }
  var _tmp1035 *float64
  if barsj != nil { _tmp1035 = (*float64)(&barsj[0]) }
  if _tmp1036 := C.MSK_putbarsj(self.ptr(),C.int32_t(whichsol),C.int32_t(j),(*C.double)(_tmp1035)); _tmp1036 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1036)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBarvarName(j int32,name string) (err error) {
  _tmp1037 := C.CString(name)
  if _tmp1038 := C.MSK_putbarvarname(self.ptr(),C.int32_t(j),_tmp1037); _tmp1038 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1038)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBarxJ(whichsol Soltype,j int32,barxj []float64) (err error) {
  var _tmp1039 C.int64_t
  if _tmp1040 := C.MSK_getlenbarvarj(self.ptr(),(C.int32_t)(j),(&_tmp1039)); _tmp1040 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1040)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(barxj)) < int64(_tmp1039) {
    err = &ArrayLengthError{fun:"PutBarxJ",arg:"barxj"}
    return
  }
  var _tmp1041 *float64
  if barxj != nil { _tmp1041 = (*float64)(&barxj[0]) }
  if _tmp1042 := C.MSK_putbarxj(self.ptr(),C.int32_t(whichsol),C.int32_t(j),(*C.double)(_tmp1041)); _tmp1042 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1042)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutCfix(cfix float64) (err error) {
  if _tmp1043 := C.MSK_putcfix(self.ptr(),C.double(cfix)); _tmp1043 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1043)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutCJ(j int32,cj float64) (err error) {
  if _tmp1044 := C.MSK_putcj(self.ptr(),C.int32_t(j),C.double(cj)); _tmp1044 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1044)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutCList(subj []int32,val []float64) (err error) {
  _tmp1045 := len(val)
  if _tmp1045 < len(subj) { _tmp1045 = len(subj) }
  var num int32 = int32(_tmp1045)
  var _tmp1046 *int32
  if subj != nil { _tmp1046 = (*int32)(&subj[0]) }
  var _tmp1047 *float64
  if val != nil { _tmp1047 = (*float64)(&val[0]) }
  if _tmp1048 := C.MSK_putclist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1046),(*C.double)(_tmp1047)); _tmp1048 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1048)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutConBound(i int32,bkc Boundkey,blc float64,buc float64) (err error) {
  if _tmp1049 := C.MSK_putconbound(self.ptr(),C.int32_t(i),C.int32_t(bkc),C.double(blc),C.double(buc)); _tmp1049 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1049)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutConBoundList(sub []int32,bkc []Boundkey,blc []float64,buc []float64) (err error) {
  _tmp1050 := len(bkc)
  if _tmp1050 < len(blc) { _tmp1050 = len(blc) }
  if _tmp1050 < len(buc) { _tmp1050 = len(buc) }
  if _tmp1050 < len(sub) { _tmp1050 = len(sub) }
  var num int32 = int32(_tmp1050)
  var _tmp1051 *int32
  if sub != nil { _tmp1051 = (*int32)(&sub[0]) }
  var _tmp1052 *Boundkey
  if bkc != nil { _tmp1052 = (*Boundkey)(&bkc[0]) }
  var _tmp1053 *float64
  if blc != nil { _tmp1053 = (*float64)(&blc[0]) }
  var _tmp1054 *float64
  if buc != nil { _tmp1054 = (*float64)(&buc[0]) }
  if _tmp1055 := C.MSK_putconboundlist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1051),(*C.int32_t)(_tmp1052),(*C.double)(_tmp1053),(*C.double)(_tmp1054)); _tmp1055 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1055)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutConBoundListConst(sub []int32,bkc Boundkey,blc float64,buc float64) (err error) {
  _tmp1056 := len(sub)
  var num int32 = int32(_tmp1056)
  var _tmp1057 *int32
  if sub != nil { _tmp1057 = (*int32)(&sub[0]) }
  if _tmp1058 := C.MSK_putconboundlistconst(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1057),C.int32_t(bkc),C.double(blc),C.double(buc)); _tmp1058 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1058)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutConBoundSlice(first int32,last int32,bkc []Boundkey,blc []float64,buc []float64) (err error) {
  if int64(len(bkc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutConBoundSlice",arg:"bkc"}
    return
  }
  var _tmp1059 *Boundkey
  if bkc != nil { _tmp1059 = (*Boundkey)(&bkc[0]) }
  if int64(len(blc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutConBoundSlice",arg:"blc"}
    return
  }
  var _tmp1060 *float64
  if blc != nil { _tmp1060 = (*float64)(&blc[0]) }
  if int64(len(buc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutConBoundSlice",arg:"buc"}
    return
  }
  var _tmp1061 *float64
  if buc != nil { _tmp1061 = (*float64)(&buc[0]) }
  if _tmp1062 := C.MSK_putconboundslice(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.int32_t)(_tmp1059),(*C.double)(_tmp1060),(*C.double)(_tmp1061)); _tmp1062 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1062)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutConBoundSliceConst(first int32,last int32,bkc Boundkey,blc float64,buc float64) (err error) {
  if _tmp1063 := C.MSK_putconboundsliceconst(self.ptr(),C.int32_t(first),C.int32_t(last),C.int32_t(bkc),C.double(blc),C.double(buc)); _tmp1063 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1063)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutCone(k int32,ct Conetype,conepar float64,submem []int32) (err error) {
  _tmp1064 := len(submem)
  var nummem int32 = int32(_tmp1064)
  var _tmp1065 *int32
  if submem != nil { _tmp1065 = (*int32)(&submem[0]) }
  if _tmp1066 := C.MSK_putcone(self.ptr(),C.int32_t(k),C.int32_t(ct),C.double(conepar),C.int32_t(nummem),(*C.int32_t)(_tmp1065)); _tmp1066 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1066)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutConeName(j int32,name string) (err error) {
  _tmp1067 := C.CString(name)
  if _tmp1068 := C.MSK_putconename(self.ptr(),C.int32_t(j),_tmp1067); _tmp1068 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1068)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutConName(i int32,name string) (err error) {
  _tmp1069 := C.CString(name)
  if _tmp1070 := C.MSK_putconname(self.ptr(),C.int32_t(i),_tmp1069); _tmp1070 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1070)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutConSolutionI(i int32,whichsol Soltype,sk Stakey,x float64,sl float64,su float64) (err error) {
  if _tmp1071 := C.MSK_putconsolutioni(self.ptr(),C.int32_t(i),C.int32_t(whichsol),C.int32_t(sk),C.double(x),C.double(sl),C.double(su)); _tmp1071 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1071)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutCSlice(first int32,last int32,slice []float64) (err error) {
  if int64(len(slice)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutCSlice",arg:"slice"}
    return
  }
  var _tmp1072 *float64
  if slice != nil { _tmp1072 = (*float64)(&slice[0]) }
  if _tmp1073 := C.MSK_putcslice(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1072)); _tmp1073 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1073)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutDjc(djcidx int64,domidxlist []int64,afeidxlist []int64,b []float64,termsizelist []int64) (err error) {
  _tmp1074 := len(domidxlist)
  var numdomidx int64 = int64(_tmp1074)
  var _tmp1075 *int64
  if domidxlist != nil { _tmp1075 = (*int64)(&domidxlist[0]) }
  _tmp1076 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp1076)
  var _tmp1077 *int64
  if afeidxlist != nil { _tmp1077 = (*int64)(&afeidxlist[0]) }
  if int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"PutDjc",arg:"b"}
    return
  }
  var _tmp1078 *float64
  if b != nil { _tmp1078 = (*float64)(&b[0]) }
  _tmp1079 := len(termsizelist)
  var numterms int64 = int64(_tmp1079)
  var _tmp1080 *int64
  if termsizelist != nil { _tmp1080 = (*int64)(&termsizelist[0]) }
  if _tmp1081 := C.MSK_putdjc(self.ptr(),C.int64_t(djcidx),C.int64_t(numdomidx),(*C.int64_t)(_tmp1075),C.int64_t(numafeidx),(*C.int64_t)(_tmp1077),(*C.double)(_tmp1078),C.int64_t(numterms),(*C.int64_t)(_tmp1080)); _tmp1081 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1081)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutDjcName(djcidx int64,name string) (err error) {
  _tmp1082 := C.CString(name)
  if _tmp1083 := C.MSK_putdjcname(self.ptr(),C.int64_t(djcidx),_tmp1082); _tmp1083 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1083)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutDjcSlice(idxfirst int64,idxlast int64,domidxlist []int64,afeidxlist []int64,b []float64,termsizelist []int64,termsindjc []int64) (err error) {
  _tmp1084 := len(domidxlist)
  var numdomidx int64 = int64(_tmp1084)
  var _tmp1085 *int64
  if domidxlist != nil { _tmp1085 = (*int64)(&domidxlist[0]) }
  _tmp1086 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp1086)
  var _tmp1087 *int64
  if afeidxlist != nil { _tmp1087 = (*int64)(&afeidxlist[0]) }
  if int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"PutDjcSlice",arg:"b"}
    return
  }
  var _tmp1088 *float64
  if b != nil { _tmp1088 = (*float64)(&b[0]) }
  _tmp1089 := len(termsizelist)
  var numterms int64 = int64(_tmp1089)
  var _tmp1090 *int64
  if termsizelist != nil { _tmp1090 = (*int64)(&termsizelist[0]) }
  if int64(len(termsindjc)) < int64((idxlast - idxfirst)) {
    err = &ArrayLengthError{fun:"PutDjcSlice",arg:"termsindjc"}
    return
  }
  var _tmp1091 *int64
  if termsindjc != nil { _tmp1091 = (*int64)(&termsindjc[0]) }
  if _tmp1092 := C.MSK_putdjcslice(self.ptr(),C.int64_t(idxfirst),C.int64_t(idxlast),C.int64_t(numdomidx),(*C.int64_t)(_tmp1085),C.int64_t(numafeidx),(*C.int64_t)(_tmp1087),(*C.double)(_tmp1088),C.int64_t(numterms),(*C.int64_t)(_tmp1090),(*C.int64_t)(_tmp1091)); _tmp1092 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1092)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutDomainName(domidx int64,name string) (err error) {
  _tmp1093 := C.CString(name)
  if _tmp1094 := C.MSK_putdomainname(self.ptr(),C.int64_t(domidx),_tmp1093); _tmp1094 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1094)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutDouParam(param Dparam,parvalue float64) (err error) {
  if _tmp1095 := C.MSK_putdouparam(self.ptr(),C.int32_t(param),C.double(parvalue)); _tmp1095 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1095)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutIntParam(param Iparam,parvalue int32) (err error) {
  if _tmp1096 := C.MSK_putintparam(self.ptr(),C.int32_t(param),C.int32_t(parvalue)); _tmp1096 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1096)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumAcc(maxnumacc int64) (err error) {
  if _tmp1097 := C.MSK_putmaxnumacc(self.ptr(),C.int64_t(maxnumacc)); _tmp1097 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1097)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumAfe(maxnumafe int64) (err error) {
  if _tmp1098 := C.MSK_putmaxnumafe(self.ptr(),C.int64_t(maxnumafe)); _tmp1098 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1098)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumANz(maxnumanz int64) (err error) {
  if _tmp1099 := C.MSK_putmaxnumanz(self.ptr(),C.int64_t(maxnumanz)); _tmp1099 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1099)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumBarvar(maxnumbarvar int32) (err error) {
  if _tmp1100 := C.MSK_putmaxnumbarvar(self.ptr(),C.int32_t(maxnumbarvar)); _tmp1100 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1100)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumCon(maxnumcon int32) (err error) {
  if _tmp1101 := C.MSK_putmaxnumcon(self.ptr(),C.int32_t(maxnumcon)); _tmp1101 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1101)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumCone(maxnumcone int32) (err error) {
  if _tmp1102 := C.MSK_putmaxnumcone(self.ptr(),C.int32_t(maxnumcone)); _tmp1102 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1102)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumDjc(maxnumdjc int64) (err error) {
  if _tmp1103 := C.MSK_putmaxnumdjc(self.ptr(),C.int64_t(maxnumdjc)); _tmp1103 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1103)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumDomain(maxnumdomain int64) (err error) {
  if _tmp1104 := C.MSK_putmaxnumdomain(self.ptr(),C.int64_t(maxnumdomain)); _tmp1104 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1104)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumQNz(maxnumqnz int64) (err error) {
  if _tmp1105 := C.MSK_putmaxnumqnz(self.ptr(),C.int64_t(maxnumqnz)); _tmp1105 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1105)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumVar(maxnumvar int32) (err error) {
  if _tmp1106 := C.MSK_putmaxnumvar(self.ptr(),C.int32_t(maxnumvar)); _tmp1106 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1106)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutNaDouParam(paramname string,parvalue float64) (err error) {
  _tmp1107 := C.CString(paramname)
  if _tmp1108 := C.MSK_putnadouparam(self.ptr(),_tmp1107,C.double(parvalue)); _tmp1108 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1108)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutNaIntParam(paramname string,parvalue int32) (err error) {
  _tmp1109 := C.CString(paramname)
  if _tmp1110 := C.MSK_putnaintparam(self.ptr(),_tmp1109,C.int32_t(parvalue)); _tmp1110 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1110)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutNaStrParam(paramname string,parvalue string) (err error) {
  _tmp1111 := C.CString(paramname)
  _tmp1112 := C.CString(parvalue)
  if _tmp1113 := C.MSK_putnastrparam(self.ptr(),_tmp1111,_tmp1112); _tmp1113 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1113)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutObjName(objname string) (err error) {
  _tmp1114 := C.CString(objname)
  if _tmp1115 := C.MSK_putobjname(self.ptr(),_tmp1114); _tmp1115 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1115)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutObjSense(sense Objsense) (err error) {
  if _tmp1116 := C.MSK_putobjsense(self.ptr(),C.int32_t(sense)); _tmp1116 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1116)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutOptserverHost(host string) (err error) {
  _tmp1117 := C.CString(host)
  if _tmp1118 := C.MSK_putoptserverhost(self.ptr(),_tmp1117); _tmp1118 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1118)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutParam(parname string,parvalue string) (err error) {
  _tmp1119 := C.CString(parname)
  _tmp1120 := C.CString(parvalue)
  if _tmp1121 := C.MSK_putparam(self.ptr(),_tmp1119,_tmp1120); _tmp1121 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1121)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutQCon(qcsubk []int32,qcsubi []int32,qcsubj []int32,qcval []float64) (err error) {
  _tmp1122 := len(qcval)
  if _tmp1122 < len(qcsubi) { _tmp1122 = len(qcsubi) }
  if _tmp1122 < len(qcsubj) { _tmp1122 = len(qcsubj) }
  var numqcnz int32 = int32(_tmp1122)
  var _tmp1123 *int32
  if qcsubk != nil { _tmp1123 = (*int32)(&qcsubk[0]) }
  var _tmp1124 *int32
  if qcsubi != nil { _tmp1124 = (*int32)(&qcsubi[0]) }
  var _tmp1125 *int32
  if qcsubj != nil { _tmp1125 = (*int32)(&qcsubj[0]) }
  var _tmp1126 *float64
  if qcval != nil { _tmp1126 = (*float64)(&qcval[0]) }
  if _tmp1127 := C.MSK_putqcon(self.ptr(),C.int32_t(numqcnz),(*C.int32_t)(_tmp1123),(*C.int32_t)(_tmp1124),(*C.int32_t)(_tmp1125),(*C.double)(_tmp1126)); _tmp1127 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1127)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutQConK(k int32,qcsubi []int32,qcsubj []int32,qcval []float64) (err error) {
  _tmp1128 := len(qcval)
  if _tmp1128 < len(qcsubi) { _tmp1128 = len(qcsubi) }
  if _tmp1128 < len(qcsubj) { _tmp1128 = len(qcsubj) }
  var numqcnz int32 = int32(_tmp1128)
  var _tmp1129 *int32
  if qcsubi != nil { _tmp1129 = (*int32)(&qcsubi[0]) }
  var _tmp1130 *int32
  if qcsubj != nil { _tmp1130 = (*int32)(&qcsubj[0]) }
  var _tmp1131 *float64
  if qcval != nil { _tmp1131 = (*float64)(&qcval[0]) }
  if _tmp1132 := C.MSK_putqconk(self.ptr(),C.int32_t(k),C.int32_t(numqcnz),(*C.int32_t)(_tmp1129),(*C.int32_t)(_tmp1130),(*C.double)(_tmp1131)); _tmp1132 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1132)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutQObj(qosubi []int32,qosubj []int32,qoval []float64) (err error) {
  _tmp1133 := len(qosubi)
  if _tmp1133 < len(qoval) { _tmp1133 = len(qoval) }
  if _tmp1133 < len(qosubj) { _tmp1133 = len(qosubj) }
  var numqonz int32 = int32(_tmp1133)
  var _tmp1134 *int32
  if qosubi != nil { _tmp1134 = (*int32)(&qosubi[0]) }
  var _tmp1135 *int32
  if qosubj != nil { _tmp1135 = (*int32)(&qosubj[0]) }
  var _tmp1136 *float64
  if qoval != nil { _tmp1136 = (*float64)(&qoval[0]) }
  if _tmp1137 := C.MSK_putqobj(self.ptr(),C.int32_t(numqonz),(*C.int32_t)(_tmp1134),(*C.int32_t)(_tmp1135),(*C.double)(_tmp1136)); _tmp1137 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1137)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutQObjIJ(i int32,j int32,qoij float64) (err error) {
  if _tmp1138 := C.MSK_putqobjij(self.ptr(),C.int32_t(i),C.int32_t(j),C.double(qoij)); _tmp1138 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1138)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSkc(whichsol Soltype,skc []Stakey) (err error) {
  var _tmp1139 C.int32_t
  if _tmp1140 := C.MSK_getnumcon(self.ptr(),(&_tmp1139)); _tmp1140 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1140)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(skc)) < int64(_tmp1139) {
    err = &ArrayLengthError{fun:"PutSkc",arg:"skc"}
    return
  }
  var _tmp1141 *Stakey
  if skc != nil { _tmp1141 = (*Stakey)(&skc[0]) }
  if _tmp1142 := C.MSK_putskc(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(_tmp1141)); _tmp1142 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1142)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSkcSlice(whichsol Soltype,first int32,last int32,skc []Stakey) (err error) {
  if int64(len(skc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSkcSlice",arg:"skc"}
    return
  }
  var _tmp1143 *Stakey
  if skc != nil { _tmp1143 = (*Stakey)(&skc[0]) }
  if _tmp1144 := C.MSK_putskcslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.int32_t)(_tmp1143)); _tmp1144 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1144)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSkx(whichsol Soltype,skx []Stakey) (err error) {
  var _tmp1145 C.int32_t
  if _tmp1146 := C.MSK_getnumvar(self.ptr(),(&_tmp1145)); _tmp1146 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1146)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(skx)) < int64(_tmp1145) {
    err = &ArrayLengthError{fun:"PutSkx",arg:"skx"}
    return
  }
  var _tmp1147 *Stakey
  if skx != nil { _tmp1147 = (*Stakey)(&skx[0]) }
  if _tmp1148 := C.MSK_putskx(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(_tmp1147)); _tmp1148 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1148)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSkxSlice(whichsol Soltype,first int32,last int32,skx []Stakey) (err error) {
  if int64(len(skx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSkxSlice",arg:"skx"}
    return
  }
  var _tmp1149 *Stakey
  if skx != nil { _tmp1149 = (*Stakey)(&skx[0]) }
  if _tmp1150 := C.MSK_putskxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.int32_t)(_tmp1149)); _tmp1150 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1150)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSlc(whichsol Soltype,slc []float64) (err error) {
  var _tmp1151 C.int32_t
  if _tmp1152 := C.MSK_getnumcon(self.ptr(),(&_tmp1151)); _tmp1152 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1152)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(slc)) < int64(_tmp1151) {
    err = &ArrayLengthError{fun:"PutSlc",arg:"slc"}
    return
  }
  var _tmp1153 *float64
  if slc != nil { _tmp1153 = (*float64)(&slc[0]) }
  if _tmp1154 := C.MSK_putslc(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp1153)); _tmp1154 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1154)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSlcSlice(whichsol Soltype,first int32,last int32,slc []float64) (err error) {
  if int64(len(slc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSlcSlice",arg:"slc"}
    return
  }
  var _tmp1155 *float64
  if slc != nil { _tmp1155 = (*float64)(&slc[0]) }
  if _tmp1156 := C.MSK_putslcslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1155)); _tmp1156 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1156)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSlx(whichsol Soltype,slx []float64) (err error) {
  var _tmp1157 C.int32_t
  if _tmp1158 := C.MSK_getnumvar(self.ptr(),(&_tmp1157)); _tmp1158 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1158)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(slx)) < int64(_tmp1157) {
    err = &ArrayLengthError{fun:"PutSlx",arg:"slx"}
    return
  }
  var _tmp1159 *float64
  if slx != nil { _tmp1159 = (*float64)(&slx[0]) }
  if _tmp1160 := C.MSK_putslx(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp1159)); _tmp1160 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1160)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSlxSlice(whichsol Soltype,first int32,last int32,slx []float64) (err error) {
  if int64(len(slx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSlxSlice",arg:"slx"}
    return
  }
  var _tmp1161 *float64
  if slx != nil { _tmp1161 = (*float64)(&slx[0]) }
  if _tmp1162 := C.MSK_putslxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1161)); _tmp1162 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1162)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSnx(whichsol Soltype,sux []float64) (err error) {
  var _tmp1163 C.int32_t
  if _tmp1164 := C.MSK_getnumvar(self.ptr(),(&_tmp1163)); _tmp1164 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1164)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(sux)) < int64(_tmp1163) {
    err = &ArrayLengthError{fun:"PutSnx",arg:"sux"}
    return
  }
  var _tmp1165 *float64
  if sux != nil { _tmp1165 = (*float64)(&sux[0]) }
  if _tmp1166 := C.MSK_putsnx(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp1165)); _tmp1166 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1166)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSnxSlice(whichsol Soltype,first int32,last int32,snx []float64) (err error) {
  if int64(len(snx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSnxSlice",arg:"snx"}
    return
  }
  var _tmp1167 *float64
  if snx != nil { _tmp1167 = (*float64)(&snx[0]) }
  if _tmp1168 := C.MSK_putsnxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1167)); _tmp1168 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1168)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSolution(whichsol Soltype,skc []Stakey,skx []Stakey,skn []Stakey,xc []float64,xx []float64,y []float64,slc []float64,suc []float64,slx []float64,sux []float64,snx []float64) (err error) {
  var _tmp1169 *Stakey
  if skc != nil { _tmp1169 = (*Stakey)(&skc[0]) }
  var _tmp1170 *Stakey
  if skx != nil { _tmp1170 = (*Stakey)(&skx[0]) }
  var _tmp1171 *Stakey
  if skn != nil { _tmp1171 = (*Stakey)(&skn[0]) }
  var _tmp1172 *float64
  if xc != nil { _tmp1172 = (*float64)(&xc[0]) }
  var _tmp1173 *float64
  if xx != nil { _tmp1173 = (*float64)(&xx[0]) }
  var _tmp1174 *float64
  if y != nil { _tmp1174 = (*float64)(&y[0]) }
  var _tmp1175 *float64
  if slc != nil { _tmp1175 = (*float64)(&slc[0]) }
  var _tmp1176 *float64
  if suc != nil { _tmp1176 = (*float64)(&suc[0]) }
  var _tmp1177 *float64
  if slx != nil { _tmp1177 = (*float64)(&slx[0]) }
  var _tmp1178 *float64
  if sux != nil { _tmp1178 = (*float64)(&sux[0]) }
  var _tmp1179 *float64
  if snx != nil { _tmp1179 = (*float64)(&snx[0]) }
  if _tmp1180 := C.MSK_putsolution(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(_tmp1169),(*C.int32_t)(_tmp1170),(*C.int32_t)(_tmp1171),(*C.double)(_tmp1172),(*C.double)(_tmp1173),(*C.double)(_tmp1174),(*C.double)(_tmp1175),(*C.double)(_tmp1176),(*C.double)(_tmp1177),(*C.double)(_tmp1178),(*C.double)(_tmp1179)); _tmp1180 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1180)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSolutionNew(whichsol Soltype,skc []Stakey,skx []Stakey,skn []Stakey,xc []float64,xx []float64,y []float64,slc []float64,suc []float64,slx []float64,sux []float64,snx []float64,doty []float64) (err error) {
  var _tmp1181 *Stakey
  if skc != nil { _tmp1181 = (*Stakey)(&skc[0]) }
  var _tmp1182 *Stakey
  if skx != nil { _tmp1182 = (*Stakey)(&skx[0]) }
  var _tmp1183 *Stakey
  if skn != nil { _tmp1183 = (*Stakey)(&skn[0]) }
  var _tmp1184 *float64
  if xc != nil { _tmp1184 = (*float64)(&xc[0]) }
  var _tmp1185 *float64
  if xx != nil { _tmp1185 = (*float64)(&xx[0]) }
  var _tmp1186 *float64
  if y != nil { _tmp1186 = (*float64)(&y[0]) }
  var _tmp1187 *float64
  if slc != nil { _tmp1187 = (*float64)(&slc[0]) }
  var _tmp1188 *float64
  if suc != nil { _tmp1188 = (*float64)(&suc[0]) }
  var _tmp1189 *float64
  if slx != nil { _tmp1189 = (*float64)(&slx[0]) }
  var _tmp1190 *float64
  if sux != nil { _tmp1190 = (*float64)(&sux[0]) }
  var _tmp1191 *float64
  if snx != nil { _tmp1191 = (*float64)(&snx[0]) }
  var _tmp1192 *float64
  if doty != nil { _tmp1192 = (*float64)(&doty[0]) }
  if _tmp1193 := C.MSK_putsolutionnew(self.ptr(),C.int32_t(whichsol),(*C.int32_t)(_tmp1181),(*C.int32_t)(_tmp1182),(*C.int32_t)(_tmp1183),(*C.double)(_tmp1184),(*C.double)(_tmp1185),(*C.double)(_tmp1186),(*C.double)(_tmp1187),(*C.double)(_tmp1188),(*C.double)(_tmp1189),(*C.double)(_tmp1190),(*C.double)(_tmp1191),(*C.double)(_tmp1192)); _tmp1193 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1193)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSolutionYI(i int32,whichsol Soltype,y float64) (err error) {
  if _tmp1194 := C.MSK_putsolutionyi(self.ptr(),C.int32_t(i),C.int32_t(whichsol),C.double(y)); _tmp1194 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1194)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutStrParam(param Sparam,parvalue string) (err error) {
  _tmp1195 := C.CString(parvalue)
  if _tmp1196 := C.MSK_putstrparam(self.ptr(),C.int32_t(param),_tmp1195); _tmp1196 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1196)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSuc(whichsol Soltype,suc []float64) (err error) {
  var _tmp1197 C.int32_t
  if _tmp1198 := C.MSK_getnumcon(self.ptr(),(&_tmp1197)); _tmp1198 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1198)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(suc)) < int64(_tmp1197) {
    err = &ArrayLengthError{fun:"PutSuc",arg:"suc"}
    return
  }
  var _tmp1199 *float64
  if suc != nil { _tmp1199 = (*float64)(&suc[0]) }
  if _tmp1200 := C.MSK_putsuc(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp1199)); _tmp1200 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1200)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSucSlice(whichsol Soltype,first int32,last int32,suc []float64) (err error) {
  if int64(len(suc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSucSlice",arg:"suc"}
    return
  }
  var _tmp1201 *float64
  if suc != nil { _tmp1201 = (*float64)(&suc[0]) }
  if _tmp1202 := C.MSK_putsucslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1201)); _tmp1202 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1202)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSux(whichsol Soltype,sux []float64) (err error) {
  var _tmp1203 C.int32_t
  if _tmp1204 := C.MSK_getnumvar(self.ptr(),(&_tmp1203)); _tmp1204 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1204)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(sux)) < int64(_tmp1203) {
    err = &ArrayLengthError{fun:"PutSux",arg:"sux"}
    return
  }
  var _tmp1205 *float64
  if sux != nil { _tmp1205 = (*float64)(&sux[0]) }
  if _tmp1206 := C.MSK_putsux(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp1205)); _tmp1206 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1206)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSuxSlice(whichsol Soltype,first int32,last int32,sux []float64) (err error) {
  if int64(len(sux)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSuxSlice",arg:"sux"}
    return
  }
  var _tmp1207 *float64
  if sux != nil { _tmp1207 = (*float64)(&sux[0]) }
  if _tmp1208 := C.MSK_putsuxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1207)); _tmp1208 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1208)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutTaskName(taskname string) (err error) {
  _tmp1209 := C.CString(taskname)
  if _tmp1210 := C.MSK_puttaskname(self.ptr(),_tmp1209); _tmp1210 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1210)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarBound(j int32,bkx Boundkey,blx float64,bux float64) (err error) {
  if _tmp1211 := C.MSK_putvarbound(self.ptr(),C.int32_t(j),C.int32_t(bkx),C.double(blx),C.double(bux)); _tmp1211 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1211)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarBoundList(sub []int32,bkx []Boundkey,blx []float64,bux []float64) (err error) {
  _tmp1212 := len(bkx)
  if _tmp1212 < len(blx) { _tmp1212 = len(blx) }
  if _tmp1212 < len(bux) { _tmp1212 = len(bux) }
  if _tmp1212 < len(sub) { _tmp1212 = len(sub) }
  var num int32 = int32(_tmp1212)
  var _tmp1213 *int32
  if sub != nil { _tmp1213 = (*int32)(&sub[0]) }
  var _tmp1214 *Boundkey
  if bkx != nil { _tmp1214 = (*Boundkey)(&bkx[0]) }
  var _tmp1215 *float64
  if blx != nil { _tmp1215 = (*float64)(&blx[0]) }
  var _tmp1216 *float64
  if bux != nil { _tmp1216 = (*float64)(&bux[0]) }
  if _tmp1217 := C.MSK_putvarboundlist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1213),(*C.int32_t)(_tmp1214),(*C.double)(_tmp1215),(*C.double)(_tmp1216)); _tmp1217 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1217)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarBoundListConst(sub []int32,bkx Boundkey,blx float64,bux float64) (err error) {
  _tmp1218 := len(sub)
  var num int32 = int32(_tmp1218)
  var _tmp1219 *int32
  if sub != nil { _tmp1219 = (*int32)(&sub[0]) }
  if _tmp1220 := C.MSK_putvarboundlistconst(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1219),C.int32_t(bkx),C.double(blx),C.double(bux)); _tmp1220 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1220)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarBoundSlice(first int32,last int32,bkx []Boundkey,blx []float64,bux []float64) (err error) {
  if int64(len(bkx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutVarBoundSlice",arg:"bkx"}
    return
  }
  var _tmp1221 *Boundkey
  if bkx != nil { _tmp1221 = (*Boundkey)(&bkx[0]) }
  if int64(len(blx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutVarBoundSlice",arg:"blx"}
    return
  }
  var _tmp1222 *float64
  if blx != nil { _tmp1222 = (*float64)(&blx[0]) }
  if int64(len(bux)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutVarBoundSlice",arg:"bux"}
    return
  }
  var _tmp1223 *float64
  if bux != nil { _tmp1223 = (*float64)(&bux[0]) }
  if _tmp1224 := C.MSK_putvarboundslice(self.ptr(),C.int32_t(first),C.int32_t(last),(*C.int32_t)(_tmp1221),(*C.double)(_tmp1222),(*C.double)(_tmp1223)); _tmp1224 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1224)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarBoundSliceConst(first int32,last int32,bkx Boundkey,blx float64,bux float64) (err error) {
  if _tmp1225 := C.MSK_putvarboundsliceconst(self.ptr(),C.int32_t(first),C.int32_t(last),C.int32_t(bkx),C.double(blx),C.double(bux)); _tmp1225 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1225)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarName(j int32,name string) (err error) {
  _tmp1226 := C.CString(name)
  if _tmp1227 := C.MSK_putvarname(self.ptr(),C.int32_t(j),_tmp1226); _tmp1227 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1227)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarSolutionJ(j int32,whichsol Soltype,sk Stakey,x float64,sl float64,su float64,sn float64) (err error) {
  if _tmp1228 := C.MSK_putvarsolutionj(self.ptr(),C.int32_t(j),C.int32_t(whichsol),C.int32_t(sk),C.double(x),C.double(sl),C.double(su),C.double(sn)); _tmp1228 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1228)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarType(j int32,vartype Variabletype) (err error) {
  if _tmp1229 := C.MSK_putvartype(self.ptr(),C.int32_t(j),C.int32_t(vartype)); _tmp1229 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1229)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarTypeList(subj []int32,vartype []Variabletype) (err error) {
  _tmp1230 := len(vartype)
  if _tmp1230 < len(subj) { _tmp1230 = len(subj) }
  var num int32 = int32(_tmp1230)
  var _tmp1231 *int32
  if subj != nil { _tmp1231 = (*int32)(&subj[0]) }
  var _tmp1232 *Variabletype
  if vartype != nil { _tmp1232 = (*Variabletype)(&vartype[0]) }
  if _tmp1233 := C.MSK_putvartypelist(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1231),(*C.int32_t)(_tmp1232)); _tmp1233 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1233)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutXc(whichsol Soltype) (xc []float64,err error) {
  var _tmp1234 C.int32_t
  if _tmp1235 := C.MSK_getnumcon(self.ptr(),(&_tmp1234)); _tmp1235 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1235)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  var _tmp1236 *float64
  xc = make([]float64,_tmp1234)
  if len(xc) > 0 { _tmp1236 = (*float64)(&xc[0]) }
  if _tmp1237 := C.MSK_putxc(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp1236)); _tmp1237 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1237)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutXcSlice(whichsol Soltype,first int32,last int32,xc []float64) (err error) {
  if int64(len(xc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutXcSlice",arg:"xc"}
    return
  }
  var _tmp1238 *float64
  if xc != nil { _tmp1238 = (*float64)(&xc[0]) }
  if _tmp1239 := C.MSK_putxcslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1238)); _tmp1239 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1239)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutXx(whichsol Soltype,xx []float64) (err error) {
  var _tmp1240 C.int32_t
  if _tmp1241 := C.MSK_getnumvar(self.ptr(),(&_tmp1240)); _tmp1241 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1241)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(xx)) < int64(_tmp1240) {
    err = &ArrayLengthError{fun:"PutXx",arg:"xx"}
    return
  }
  var _tmp1242 *float64
  if xx != nil { _tmp1242 = (*float64)(&xx[0]) }
  if _tmp1243 := C.MSK_putxx(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp1242)); _tmp1243 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1243)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutXxSlice(whichsol Soltype,first int32,last int32,xx []float64) (err error) {
  if int64(len(xx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutXxSlice",arg:"xx"}
    return
  }
  var _tmp1244 *float64
  if xx != nil { _tmp1244 = (*float64)(&xx[0]) }
  if _tmp1245 := C.MSK_putxxslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1244)); _tmp1245 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1245)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutY(whichsol Soltype,y []float64) (err error) {
  var _tmp1246 C.int32_t
  if _tmp1247 := C.MSK_getnumcon(self.ptr(),(&_tmp1246)); _tmp1247 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1247)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(y)) < int64(_tmp1246) {
    err = &ArrayLengthError{fun:"PutY",arg:"y"}
    return
  }
  var _tmp1248 *float64
  if y != nil { _tmp1248 = (*float64)(&y[0]) }
  if _tmp1249 := C.MSK_puty(self.ptr(),C.int32_t(whichsol),(*C.double)(_tmp1248)); _tmp1249 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1249)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutYSlice(whichsol Soltype,first int32,last int32,y []float64) (err error) {
  if int64(len(y)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutYSlice",arg:"y"}
    return
  }
  var _tmp1250 *float64
  if y != nil { _tmp1250 = (*float64)(&y[0]) }
  if _tmp1251 := C.MSK_putyslice(self.ptr(),C.int32_t(whichsol),C.int32_t(first),C.int32_t(last),(*C.double)(_tmp1250)); _tmp1251 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1251)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadBSolution(filename string,compress Compresstype) (err error) {
  _tmp1252 := C.CString(filename)
  if _tmp1253 := C.MSK_readbsolution(self.ptr(),_tmp1252,C.int32_t(compress)); _tmp1253 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1253)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadData(filename string) (err error) {
  _tmp1254 := C.CString(filename)
  if _tmp1255 := C.MSK_readdataautoformat(self.ptr(),_tmp1254); _tmp1255 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1255)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadDataFormat(filename string,format Dataformat,compress Compresstype) (err error) {
  _tmp1256 := C.CString(filename)
  if _tmp1257 := C.MSK_readdataformat(self.ptr(),_tmp1256,C.int32_t(format),C.int32_t(compress)); _tmp1257 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1257)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadJsonSol(filename string) (err error) {
  _tmp1258 := C.CString(filename)
  if _tmp1259 := C.MSK_readjsonsol(self.ptr(),_tmp1258); _tmp1259 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1259)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadJsonString(data string) (err error) {
  _tmp1260 := C.CString(data)
  if _tmp1261 := C.MSK_readjsonstring(self.ptr(),_tmp1260); _tmp1261 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1261)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadLpString(data string) (err error) {
  _tmp1262 := C.CString(data)
  if _tmp1263 := C.MSK_readlpstring(self.ptr(),_tmp1262); _tmp1263 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1263)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadOpfString(data string) (err error) {
  _tmp1264 := C.CString(data)
  if _tmp1265 := C.MSK_readopfstring(self.ptr(),_tmp1264); _tmp1265 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1265)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadParamFile(filename string) (err error) {
  _tmp1266 := C.CString(filename)
  if _tmp1267 := C.MSK_readparamfile(self.ptr(),_tmp1266); _tmp1267 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1267)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadPtfString(data string) (err error) {
  _tmp1268 := C.CString(data)
  if _tmp1269 := C.MSK_readptfstring(self.ptr(),_tmp1268); _tmp1269 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1269)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadSolution(whichsol Soltype,filename string) (err error) {
  _tmp1270 := C.CString(filename)
  if _tmp1271 := C.MSK_readsolution(self.ptr(),C.int32_t(whichsol),_tmp1270); _tmp1271 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1271)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadSolutionFile(filename string) (err error) {
  _tmp1272 := C.CString(filename)
  if _tmp1273 := C.MSK_readsolutionfile(self.ptr(),_tmp1272); _tmp1273 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1273)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadSummary(whichstream Streamtype) (err error) {
  if _tmp1274 := C.MSK_readsummary(self.ptr(),C.int32_t(whichstream)); _tmp1274 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1274)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadTask(filename string) (err error) {
  _tmp1275 := C.CString(filename)
  if _tmp1276 := C.MSK_readtask(self.ptr(),_tmp1275); _tmp1276 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1276)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) RemoveBarvars(subset []int32) (err error) {
  _tmp1277 := len(subset)
  var num int32 = int32(_tmp1277)
  var _tmp1278 *int32
  if subset != nil { _tmp1278 = (*int32)(&subset[0]) }
  if _tmp1279 := C.MSK_removebarvars(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1278)); _tmp1279 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1279)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) RemoveCones(subset []int32) (err error) {
  _tmp1280 := len(subset)
  var num int32 = int32(_tmp1280)
  var _tmp1281 *int32
  if subset != nil { _tmp1281 = (*int32)(&subset[0]) }
  if _tmp1282 := C.MSK_removecones(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1281)); _tmp1282 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1282)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) RemoveCons(subset []int32) (err error) {
  _tmp1283 := len(subset)
  var num int32 = int32(_tmp1283)
  var _tmp1284 *int32
  if subset != nil { _tmp1284 = (*int32)(&subset[0]) }
  if _tmp1285 := C.MSK_removecons(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1284)); _tmp1285 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1285)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) RemoveVars(subset []int32) (err error) {
  _tmp1286 := len(subset)
  var num int32 = int32(_tmp1286)
  var _tmp1287 *int32
  if subset != nil { _tmp1287 = (*int32)(&subset[0]) }
  if _tmp1288 := C.MSK_removevars(self.ptr(),C.int32_t(num),(*C.int32_t)(_tmp1287)); _tmp1288 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1288)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) ResizeTask(maxnumcon int32,maxnumvar int32,maxnumcone int32,maxnumanz int64,maxnumqnz int64) (err error) {
  if _tmp1289 := C.MSK_resizetask(self.ptr(),C.int32_t(maxnumcon),C.int32_t(maxnumvar),C.int32_t(maxnumcone),C.int64_t(maxnumanz),C.int64_t(maxnumqnz)); _tmp1289 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1289)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) SensitivityReport(whichstream Streamtype) (err error) {
  if _tmp1290 := C.MSK_sensitivityreport(self.ptr(),C.int32_t(whichstream)); _tmp1290 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1290)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) SetDefaults() (err error) {
  if _tmp1291 := C.MSK_setdefaults(self.ptr()); _tmp1291 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1291)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) SkToStr(sk Stakey) (str string,err error) {
  _tmp1292 := make([]byte,MSK_MAX_STR_LEN)
  if _tmp1293 := C.MSK_sktostr(self.ptr(),C.int32_t(sk),(*C.uchar)(&_tmp1292[0])); _tmp1293 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1293)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp1292,byte(0)); p < 0 {
    str = string(_tmp1292)
  } else {
    str = string(_tmp1292[:p])
  }
  return
}
func (self *Task) SolStaToStr(solutionsta Solsta) (str string,err error) {
  _tmp1294 := make([]byte,MSK_MAX_STR_LEN)
  if _tmp1295 := C.MSK_solstatostr(self.ptr(),C.int32_t(solutionsta),(*C.uchar)(&_tmp1294[0])); _tmp1295 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1295)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if p := bytes.IndexByte(_tmp1294,byte(0)); p < 0 {
    str = string(_tmp1294)
  } else {
    str = string(_tmp1294[:p])
  }
  return
}
func (self *Task) SolutionDef(whichsol Soltype) (isdef bool,err error) {
  var _tmp1296 C.int
  if _tmp1297 := C.MSK_solutiondef(self.ptr(),C.int32_t(whichsol),(&_tmp1296)); _tmp1297 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1297)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  isdef = _tmp1296 != 0
  return
}
func (self *Task) SolutionSummary(whichstream Streamtype) (err error) {
  if _tmp1298 := C.MSK_solutionsummary(self.ptr(),C.int32_t(whichstream)); _tmp1298 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1298)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) SolveWithBasis(transp bool,numnz int32,sub []int32,val []float64) (numnzout int32,err error) {
  var _tmp1299 C.int; if transp { _tmp1299 = 1; }
  var _tmp1300 C.int32_t
  if _tmp1301 := C.MSK_getnumcon(self.ptr(),(&_tmp1300)); _tmp1301 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1301)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(sub)) < int64(_tmp1300) {
    err = &ArrayLengthError{fun:"SolveWithBasis",arg:"sub"}
    return
  }
  var _tmp1302 *int32
  if sub != nil { _tmp1302 = (*int32)(&sub[0]) }
  var _tmp1303 C.int32_t
  if _tmp1304 := C.MSK_getnumcon(self.ptr(),(&_tmp1303)); _tmp1304 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1304)
     err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}
    return
  }
  if int64(len(val)) < int64(_tmp1303) {
    err = &ArrayLengthError{fun:"SolveWithBasis",arg:"val"}
    return
  }
  var _tmp1305 *float64
  if val != nil { _tmp1305 = (*float64)(&val[0]) }
  if _tmp1306 := C.MSK_solvewithbasis(self.ptr(),_tmp1299,C.int32_t(numnz),(*C.int32_t)(_tmp1302),(*C.double)(_tmp1305),(*C.int32_t)(&numnzout)); _tmp1306 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1306)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) StrToConeType(str string) (conetype Conetype,err error) {
  _tmp1307 := C.CString(str)
  var _tmp1308 int32
  if _tmp1309 := C.MSK_strtoconetype(self.ptr(),_tmp1307,(*C.int32_t)(&_tmp1308)); _tmp1309 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1309)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  conetype = Conetype(_tmp1308)
  return
}
func (self *Task) StrToSk(str string) (sk Stakey,err error) {
  _tmp1310 := C.CString(str)
  var _tmp1311 int32
  if _tmp1312 := C.MSK_strtosk(self.ptr(),_tmp1310,(*C.int32_t)(&_tmp1311)); _tmp1312 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1312)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  sk = Stakey(_tmp1311)
  return
}
func (self *Task) Toconic() (err error) {
  if _tmp1313 := C.MSK_toconic(self.ptr()); _tmp1313 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1313)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) UpdateSolutionInfo(whichsol Soltype) (err error) {
  if _tmp1314 := C.MSK_updatesolutioninfo(self.ptr(),C.int32_t(whichsol)); _tmp1314 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1314)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) WriteBSolution(filename string,compress Compresstype) (err error) {
  _tmp1315 := C.CString(filename)
  if _tmp1316 := C.MSK_writebsolution(self.ptr(),_tmp1315,C.int32_t(compress)); _tmp1316 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1316)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) WriteData(filename string) (err error) {
  _tmp1317 := C.CString(filename)
  if _tmp1318 := C.MSK_writedata(self.ptr(),_tmp1317); _tmp1318 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1318)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) WriteJsonSol(filename string) (err error) {
  _tmp1319 := C.CString(filename)
  if _tmp1320 := C.MSK_writejsonsol(self.ptr(),_tmp1319); _tmp1320 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1320)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) WriteParamFile(filename string) (err error) {
  _tmp1321 := C.CString(filename)
  if _tmp1322 := C.MSK_writeparamfile(self.ptr(),_tmp1321); _tmp1322 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1322)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) WriteSolution(whichsol Soltype,filename string) (err error) {
  _tmp1323 := C.CString(filename)
  if _tmp1324 := C.MSK_writesolution(self.ptr(),C.int32_t(whichsol),_tmp1323); _tmp1324 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1324)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) WriteSolutionFile(filename string) (err error) {
  _tmp1325 := C.CString(filename)
  if _tmp1326 := C.MSK_writesolutionfile(self.ptr(),_tmp1325); _tmp1326 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1326)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Task) WriteTask(filename string) (err error) {
  _tmp1328 := C.CString(filename)
  if _tmp1329 := C.MSK_writetask(self.ptr(),_tmp1328); _tmp1329 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1329)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) Axpy(n int32,alpha float64,x []float64,y []float64) (err error) {
  if int64(len(x)) < int64(n) {
    err = &ArrayLengthError{fun:"Axpy",arg:"x"}
    return
  }
  var _tmp1330 *float64
  if x != nil { _tmp1330 = (*float64)(&x[0]) }
  if int64(len(y)) < int64(n) {
    err = &ArrayLengthError{fun:"Axpy",arg:"y"}
    return
  }
  var _tmp1331 *float64
  if y != nil { _tmp1331 = (*float64)(&y[0]) }
  if _tmp1332 := C.MSK_axpy(self.ptr(),C.int32_t(n),C.double(alpha),(*C.double)(_tmp1330),(*C.double)(_tmp1331)); _tmp1332 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1332)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) CheckInAll() (err error) {
  if _tmp1333 := C.MSK_checkinall(self.ptr()); _tmp1333 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1333)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) CheckInLicense(feature Feature) (err error) {
  if _tmp1334 := C.MSK_checkinlicense(self.ptr(),C.int32_t(feature)); _tmp1334 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1334)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) CheckOutLicense(feature Feature) (err error) {
  if _tmp1335 := C.MSK_checkoutlicense(self.ptr(),C.int32_t(feature)); _tmp1335 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1335)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) Dot(n int32,x []float64,y []float64) (xty float64,err error) {
  if int64(len(x)) < int64(n) {
    err = &ArrayLengthError{fun:"Dot",arg:"x"}
    return
  }
  var _tmp1336 *float64
  if x != nil { _tmp1336 = (*float64)(&x[0]) }
  if int64(len(y)) < int64(n) {
    err = &ArrayLengthError{fun:"Dot",arg:"y"}
    return
  }
  var _tmp1337 *float64
  if y != nil { _tmp1337 = (*float64)(&y[0]) }
  if _tmp1338 := C.MSK_dot(self.ptr(),C.int32_t(n),(*C.double)(_tmp1336),(*C.double)(_tmp1337),(*C.double)(&xty)); _tmp1338 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1338)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) EchoIntro(longver int32) (err error) {
  if _tmp1339 := C.MSK_echointro(self.ptr(),C.int32_t(longver)); _tmp1339 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1339)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) Expirylicenses() (expiry int64,err error) {
  if _tmp1340 := C.MSK_expirylicenses(self.ptr(),(*C.int64_t)(&expiry)); _tmp1340 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1340)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) Gemm(transa Transpose,transb Transpose,m int32,n int32,k int32,alpha float64,a []float64,b []float64,beta float64,c []float64) (err error) {
  if int64(len(a)) < int64((m * k)) {
    err = &ArrayLengthError{fun:"Gemm",arg:"a"}
    return
  }
  var _tmp1341 *float64
  if a != nil { _tmp1341 = (*float64)(&a[0]) }
  if int64(len(b)) < int64((k * n)) {
    err = &ArrayLengthError{fun:"Gemm",arg:"b"}
    return
  }
  var _tmp1342 *float64
  if b != nil { _tmp1342 = (*float64)(&b[0]) }
  if int64(len(c)) < int64((m * n)) {
    err = &ArrayLengthError{fun:"Gemm",arg:"c"}
    return
  }
  var _tmp1343 *float64
  if c != nil { _tmp1343 = (*float64)(&c[0]) }
  if _tmp1344 := C.MSK_gemm(self.ptr(),C.int32_t(transa),C.int32_t(transb),C.int32_t(m),C.int32_t(n),C.int32_t(k),C.double(alpha),(*C.double)(_tmp1341),(*C.double)(_tmp1342),C.double(beta),(*C.double)(_tmp1343)); _tmp1344 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1344)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) Gemv(transa Transpose,m int32,n int32,alpha float64,a []float64,x []float64,beta float64,y []float64) (err error) {
  if int64(len(a)) < int64((n * m)) {
    err = &ArrayLengthError{fun:"Gemv",arg:"a"}
    return
  }
  var _tmp1345 *float64
  if a != nil { _tmp1345 = (*float64)(&a[0]) }
  var _tmp1346 C.int32_t
  if (transa == MSK_TRANSPOSE_NO) {
    _tmp1346 = (C.int32_t)(n)
  } else {
    _tmp1346 = (C.int32_t)(m)
  }
  if int64(len(x)) < int64(_tmp1346) {
    err = &ArrayLengthError{fun:"Gemv",arg:"x"}
    return
  }
  var _tmp1347 *float64
  if x != nil { _tmp1347 = (*float64)(&x[0]) }
  var _tmp1348 C.int32_t
  if (transa == MSK_TRANSPOSE_NO) {
    _tmp1348 = (C.int32_t)(m)
  } else {
    _tmp1348 = (C.int32_t)(n)
  }
  if int64(len(y)) < int64(_tmp1348) {
    err = &ArrayLengthError{fun:"Gemv",arg:"y"}
    return
  }
  var _tmp1349 *float64
  if y != nil { _tmp1349 = (*float64)(&y[0]) }
  if _tmp1350 := C.MSK_gemv(self.ptr(),C.int32_t(transa),C.int32_t(m),C.int32_t(n),C.double(alpha),(*C.double)(_tmp1345),(*C.double)(_tmp1347),C.double(beta),(*C.double)(_tmp1349)); _tmp1350 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1350)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func GetCodeDesc(code Rescode) (symname string,str string,err error) {
  _tmp1351 := make([]byte,MSK_MAX_STR_LEN)
  _tmp1352 := make([]byte,MSK_MAX_STR_LEN)
  if _tmp1353 := C.MSK_getcodedesc(C.int32_t(code),(*C.uchar)(&_tmp1351[0]),(*C.uchar)(&_tmp1352[0])); _tmp1353 != 0 {
    err = &MosekError{ code:Rescode(_tmp1353) }
    return
  }
  if p := bytes.IndexByte(_tmp1351,byte(0)); p < 0 {
    symname = string(_tmp1351)
  } else {
    symname = string(_tmp1351[:p])
  }
  if p := bytes.IndexByte(_tmp1352,byte(0)); p < 0 {
    str = string(_tmp1352)
  } else {
    str = string(_tmp1352[:p])
  }
  return
}
func GetVersion() (major int32,minor int32,revision int32,err error) {
  if _tmp1354 := C.MSK_getversion((*C.int32_t)(&major),(*C.int32_t)(&minor),(*C.int32_t)(&revision)); _tmp1354 != 0 {
    err = &MosekError{ code:Rescode(_tmp1354) }
    return
  }
  return
}
func LicenseCleanup() (err error) {
  if _tmp1355 := C.MSK_licensecleanup(); _tmp1355 != 0 {
    err = &MosekError{ code:Rescode(_tmp1355) }
    return
  }
  return
}
func (self *Env) Linkfiletostream(whichstream Streamtype,filename string,append int32) (err error) {
  _tmp1356 := C.CString(filename)
  if _tmp1357 := C.MSK_linkfiletoenvstream(self.ptr(),C.int32_t(whichstream),_tmp1356,C.int32_t(append)); _tmp1357 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1357)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) Potrf(uplo Uplo,n int32,a []float64) (err error) {
  if int64(len(a)) < int64((n * n)) {
    err = &ArrayLengthError{fun:"Potrf",arg:"a"}
    return
  }
  var _tmp1358 *float64
  if a != nil { _tmp1358 = (*float64)(&a[0]) }
  if _tmp1359 := C.MSK_potrf(self.ptr(),C.int32_t(uplo),C.int32_t(n),(*C.double)(_tmp1358)); _tmp1359 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1359)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) PutLicenseCode(code []int32) (err error) {
  if int64(len(code)) < int64(MSK_LICENSE_BUFFER_LENGTH) {
    err = &ArrayLengthError{fun:"PutLicenseCode",arg:"code"}
    return
  }
  var _tmp1360 *int32
  if code != nil { _tmp1360 = (*int32)(&code[0]) }
  if _tmp1361 := C.MSK_putlicensecode(self.ptr(),(*C.int32_t)(_tmp1360)); _tmp1361 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1361)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) PutLicenseDebug(licdebug int32) (err error) {
  if _tmp1362 := C.MSK_putlicensedebug(self.ptr(),C.int32_t(licdebug)); _tmp1362 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1362)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) PutLicensePath(licensepath string) (err error) {
  _tmp1363 := C.CString(licensepath)
  if _tmp1364 := C.MSK_putlicensepath(self.ptr(),_tmp1363); _tmp1364 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1364)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) PutLicenseWait(licwait int32) (err error) {
  if _tmp1365 := C.MSK_putlicensewait(self.ptr(),C.int32_t(licwait)); _tmp1365 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1365)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) ResetExpiryLicenses() (err error) {
  if _tmp1366 := C.MSK_resetexpirylicenses(self.ptr()); _tmp1366 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1366)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) Syeig(uplo Uplo,n int32,a []float64) (w []float64,err error) {
  if int64(len(a)) < int64((n * n)) {
    err = &ArrayLengthError{fun:"Syeig",arg:"a"}
    return
  }
  var _tmp1367 *float64
  if a != nil { _tmp1367 = (*float64)(&a[0]) }
  var _tmp1368 *float64
  w = make([]float64,n)
  if len(w) > 0 { _tmp1368 = (*float64)(&w[0]) }
  if _tmp1369 := C.MSK_syeig(self.ptr(),C.int32_t(uplo),C.int32_t(n),(*C.double)(_tmp1367),(*C.double)(_tmp1368)); _tmp1369 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1369)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) Syevd(uplo Uplo,n int32,a []float64) (w []float64,err error) {
  if int64(len(a)) < int64((n * n)) {
    err = &ArrayLengthError{fun:"Syevd",arg:"a"}
    return
  }
  var _tmp1370 *float64
  if a != nil { _tmp1370 = (*float64)(&a[0]) }
  var _tmp1371 *float64
  w = make([]float64,n)
  if len(w) > 0 { _tmp1371 = (*float64)(&w[0]) }
  if _tmp1372 := C.MSK_syevd(self.ptr(),C.int32_t(uplo),C.int32_t(n),(*C.double)(_tmp1370),(*C.double)(_tmp1371)); _tmp1372 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1372)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}
func (self *Env) Syrk(uplo Uplo,trans Transpose,n int32,k int32,alpha float64,a []float64,beta float64,c []float64) (err error) {
  if int64(len(a)) < int64((n * k)) {
    err = &ArrayLengthError{fun:"Syrk",arg:"a"}
    return
  }
  var _tmp1373 *float64
  if a != nil { _tmp1373 = (*float64)(&a[0]) }
  if int64(len(c)) < int64((n * n)) {
    err = &ArrayLengthError{fun:"Syrk",arg:"c"}
    return
  }
  var _tmp1374 *float64
  if c != nil { _tmp1374 = (*float64)(&c[0]) }
  if _tmp1375 := C.MSK_syrk(self.ptr(),C.int32_t(uplo),C.int32_t(trans),C.int32_t(n),C.int32_t(k),C.double(alpha),(*C.double)(_tmp1373),C.double(beta),(*C.double)(_tmp1374)); _tmp1375 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1375)
    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    return
  }
  return
}

