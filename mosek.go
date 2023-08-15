package mosek

// /*<comment>*/

// #include <stdlib.h>
// #include <stdint.h>
// #cgo LDFLAGS: -lmosek64
//
// typedef const char * string_t;
//
// extern void streamfunc_log(void *, char *);
// extern void streamfunc_wrn(void *, char *);
// extern void streamfunc_msg(void *, char *);
// extern void streamfunc_err(void *, char *);
// extern int callbackfunc(void *, void *, int, double*,int32_t*, int64_t*);
// extern int MSK_analyzenames(void *,int,int);
// extern int MSK_analyzeproblem(void *,int);
// extern int MSK_analyzesolution(void *,int,int);
// extern int MSK_appendacc(void *,int64_t,int64_t *,int64_t *,double *);
// extern int MSK_appendaccs(void *,int64_t *,int64_t *,int64_t *,int64_t *,double *);
// extern int MSK_appendaccseq(void *,int64_t,int64_t *,int64_t,double *);
// extern int MSK_appendaccsseq(void *,int64_t *,int64_t *,int64_t,int64_t,double *);
// extern int MSK_appendafes(void *,int64_t);
// extern int MSK_appendbarvars(void *,int32_t *,int32_t *);
// extern int MSK_appendcone(void *,int,double,int32_t *,int32_t *);
// extern int MSK_appendconeseq(void *,int,double,int32_t,int32_t);
// extern int MSK_appendconesseq(void *,int32_t *,int *,double *,int32_t *,int32_t);
// extern int MSK_appendcons(void *,int32_t);
// extern int MSK_appenddjcs(void *,int64_t);
// extern int MSK_appenddualexpconedomain(void *,int64_t *);
// extern int MSK_appenddualgeomeanconedomain(void *,int64_t,int64_t *);
// extern int MSK_appenddualpowerconedomain(void *,int64_t,int64_t *,double *,int64_t *);
// extern int MSK_appendprimalexpconedomain(void *,int64_t *);
// extern int MSK_appendprimalgeomeanconedomain(void *,int64_t,int64_t *);
// extern int MSK_appendprimalpowerconedomain(void *,int64_t,int64_t *,double *,int64_t *);
// extern int MSK_appendquadraticconedomain(void *,int64_t,int64_t *);
// extern int MSK_appendrdomain(void *,int64_t,int64_t *);
// extern int MSK_appendrminusdomain(void *,int64_t,int64_t *);
// extern int MSK_appendrplusdomain(void *,int64_t,int64_t *);
// extern int MSK_appendrquadraticconedomain(void *,int64_t,int64_t *);
// extern int MSK_appendrzerodomain(void *,int64_t,int64_t *);
// extern int MSK_appendsparsesymmat(void *,int32_t,int64_t *,int32_t *,int32_t *,double *,int64_t *);
// extern int MSK_appendsparsesymmatlist(void *,int32_t *,int32_t *,int64_t *,int32_t *,int32_t *,double *,int64_t *);
// extern int MSK_appendsvecpsdconedomain(void *,int64_t,int64_t *);
// extern int MSK_appendvars(void *,int32_t);
// extern int MSK_basiscond(void *,double *,double *);
// extern int MSK_checkmemtask(void *,const char *,int32_t);
// extern int MSK_chgconbound(void *,int32_t,int32_t,int32_t,double);
// extern int MSK_chgvarbound(void *,int32_t,int32_t,int32_t,double);
// extern int MSK_commitchanges(void *);
// extern int MSK_deletesolution(void *,int);
// extern int MSK_dualsensitivity(void *,int32_t *,int32_t *,double *,double *,double *,double *);
// extern int MSK_emptyafebarfrow(void *,int64_t);
// extern int MSK_emptyafebarfrowlist(void *,int64_t *,int64_t *);
// extern int MSK_emptyafefcol(void *,int32_t);
// extern int MSK_emptyafefcollist(void *,int64_t *,int32_t *);
// extern int MSK_emptyafefrow(void *,int64_t);
// extern int MSK_emptyafefrowlist(void *,int64_t *,int64_t *);
// extern int MSK_evaluateacc(void *,int,int64_t,double *);
// extern int MSK_evaluateaccs(void *,int,double *);
// extern int MSK_generateaccnames(void *,int64_t *,int64_t *,const char *,int32_t *,int32_t *,int64_t *,int32_t *,int32_t *,int64_t *,string_t *);
// extern int MSK_generatebarvarnames(void *,int32_t *,int32_t *,const char *,int32_t *,int32_t *,int64_t *,int32_t *,int32_t *,int64_t *,string_t *);
// extern int MSK_generateconenames(void *,int32_t *,int32_t *,const char *,int32_t *,int32_t *,int64_t *,int32_t *,int32_t *,int64_t *,string_t *);
// extern int MSK_generateconnames(void *,int32_t *,int32_t *,const char *,int32_t *,int32_t *,int64_t *,int32_t *,int32_t *,int64_t *,string_t *);
// extern int MSK_generatedjcnames(void *,int64_t *,int64_t *,const char *,int32_t *,int32_t *,int64_t *,int32_t *,int32_t *,int64_t *,string_t *);
// extern int MSK_generatevarnames(void *,int32_t *,int32_t *,const char *,int32_t *,int32_t *,int64_t *,int32_t *,int32_t *,int64_t *,string_t *);
// extern int MSK_getaccafeidxlist(void *,int64_t,int64_t *);
// extern int MSK_getaccb(void *,int64_t,double *);
// extern int MSK_getaccbarfblocktriplet(void *,int64_t *,int64_t *,int64_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_getaccbarfnumblocktriplets(void *,int64_t *);
// extern int MSK_getaccdomain(void *,int64_t,int64_t *);
// extern int MSK_getaccdoty(void *,int,int64_t,double *);
// extern int MSK_getaccdotys(void *,int,double *);
// extern int MSK_getaccfnumnz(void *,int64_t *);
// extern int MSK_getaccftrip(void *,int64_t *,int32_t *,double *);
// extern int MSK_getaccgvector(void *,double *);
// extern int MSK_getaccn(void *,int64_t,int64_t *);
// extern int MSK_getaccname(void *,int64_t,int32_t *,char *);
// extern int MSK_getaccnamelen(void *,int64_t,int32_t *);
// extern int MSK_getaccntot(void *,int64_t *);
// extern int MSK_getaccs(void *,int64_t *,int64_t *,double *);
// extern int MSK_getacol(void *,int32_t,int32_t *,int32_t *,double *);
// extern int MSK_getacolnumnz(void *,int32_t,int32_t *);
// extern int MSK_getacolslice64(void *,int32_t,int32_t,int64_t *,int64_t *,int64_t *,int32_t *,double *);
// extern int MSK_getacolslicenumnz64(void *,int32_t,int32_t,int64_t *);
// extern int MSK_getacolslicetrip(void *,int32_t,int32_t,int64_t *,int32_t *,int32_t *,double *);
// extern int MSK_getafebarfblocktriplet(void *,int64_t *,int64_t *,int64_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_getafebarfnumblocktriplets(void *,int64_t *);
// extern int MSK_getafebarfnumrowentries(void *,int64_t,int32_t *);
// extern int MSK_getafebarfrow(void *,int64_t,int32_t *,int64_t *,int64_t *,int64_t *,double *);
// extern int MSK_getafebarfrowinfo(void *,int64_t,int32_t *,int64_t *);
// extern int MSK_getafefnumnz(void *,int64_t *);
// extern int MSK_getafefrow(void *,int64_t,int32_t *,int32_t *,double *);
// extern int MSK_getafefrownumnz(void *,int64_t,int32_t *);
// extern int MSK_getafeftrip(void *,int64_t *,int32_t *,double *);
// extern int MSK_getafeg(void *,int64_t,double *);
// extern int MSK_getafegslice(void *,int64_t,int64_t,double *);
// extern int MSK_getaij(void *,int32_t,int32_t,double *);
// extern int MSK_getapiecenumnz(void *,int32_t,int32_t,int32_t,int32_t,int32_t *);
// extern int MSK_getarow(void *,int32_t,int32_t *,int32_t *,double *);
// extern int MSK_getarownumnz(void *,int32_t,int32_t *);
// extern int MSK_getarowslice64(void *,int32_t,int32_t,int64_t *,int64_t *,int64_t *,int32_t *,double *);
// extern int MSK_getarowslicenumnz64(void *,int32_t,int32_t,int64_t *);
// extern int MSK_getarowslicetrip(void *,int32_t,int32_t,int64_t *,int32_t *,int32_t *,double *);
// extern int MSK_getatrip(void *,int64_t *,int32_t *,int32_t *,double *);
// extern int MSK_getatruncatetol(void *,double *);
// extern int MSK_getbarablocktriplet(void *,int64_t *,int64_t *,int32_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_getbaraidx(void *,int64_t,int64_t *,int32_t *,int32_t *,int64_t *,int64_t *,double *);
// extern int MSK_getbaraidxij(void *,int64_t,int32_t *,int32_t *);
// extern int MSK_getbaraidxinfo(void *,int64_t,int64_t *);
// extern int MSK_getbarasparsity(void *,int64_t *,int64_t *,int64_t *);
// extern int MSK_getbarcblocktriplet(void *,int64_t *,int64_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_getbarcidx(void *,int64_t,int64_t *,int32_t *,int64_t *,int64_t *,double *);
// extern int MSK_getbarcidxinfo(void *,int64_t,int64_t *);
// extern int MSK_getbarcidxj(void *,int64_t,int32_t *);
// extern int MSK_getbarcsparsity(void *,int64_t *,int64_t *,int64_t *);
// extern int MSK_getbarsj(void *,int,int32_t,double *);
// extern int MSK_getbarsslice(void *,int,int32_t,int32_t,int64_t,double *);
// extern int MSK_getbarvarname(void *,int32_t,int32_t *,char *);
// extern int MSK_getbarvarnameindex(void *,const char *,int32_t *,int32_t *);
// extern int MSK_getbarvarnamelen(void *,int32_t,int32_t *);
// extern int MSK_getbarxj(void *,int,int32_t,double *);
// extern int MSK_getbarxslice(void *,int,int32_t,int32_t,int64_t,double *);
// extern int MSK_getc(void *,double *);
// extern int MSK_getcfix(void *,double *);
// extern int MSK_getcj(void *,int32_t,double *);
// extern int MSK_getclist(void *,int32_t *,int32_t *,double *);
// extern int MSK_getconbound(void *,int32_t,int *,double *,double *);
// extern int MSK_getconboundslice(void *,int32_t,int32_t,int *,double *,double *);
// extern int MSK_getcone(void *,int32_t,int *,double *,int32_t *,int32_t *);
// extern int MSK_getconeinfo(void *,int32_t,int *,double *,int32_t *);
// extern int MSK_getconename(void *,int32_t,int32_t *,char *);
// extern int MSK_getconenameindex(void *,const char *,int32_t *,int32_t *);
// extern int MSK_getconenamelen(void *,int32_t,int32_t *);
// extern int MSK_getconname(void *,int32_t,int32_t *,char *);
// extern int MSK_getconnameindex(void *,const char *,int32_t *,int32_t *);
// extern int MSK_getconnamelen(void *,int32_t,int32_t *);
// extern int MSK_getcslice(void *,int32_t,int32_t,double *);
// extern int MSK_getdimbarvarj(void *,int32_t,int32_t *);
// extern int MSK_getdjcafeidxlist(void *,int64_t,int64_t *);
// extern int MSK_getdjcb(void *,int64_t,double *);
// extern int MSK_getdjcdomainidxlist(void *,int64_t,int64_t *);
// extern int MSK_getdjcname(void *,int64_t,int32_t *,char *);
// extern int MSK_getdjcnamelen(void *,int64_t,int32_t *);
// extern int MSK_getdjcnumafe(void *,int64_t,int64_t *);
// extern int MSK_getdjcnumafetot(void *,int64_t *);
// extern int MSK_getdjcnumdomain(void *,int64_t,int64_t *);
// extern int MSK_getdjcnumdomaintot(void *,int64_t *);
// extern int MSK_getdjcnumterm(void *,int64_t,int64_t *);
// extern int MSK_getdjcnumtermtot(void *,int64_t *);
// extern int MSK_getdjcs(void *,int64_t *,int64_t *,double *,int64_t *,int64_t *);
// extern int MSK_getdjctermsizelist(void *,int64_t,int64_t *);
// extern int MSK_getdomainn(void *,int64_t,int64_t *);
// extern int MSK_getdomainname(void *,int64_t,int32_t *,char *);
// extern int MSK_getdomainnamelen(void *,int64_t,int32_t *);
// extern int MSK_getdomaintype(void *,int64_t,int *);
// extern int MSK_getdouinf(void *,int,double *);
// extern int MSK_getdouparam(void *,int,double *);
// extern int MSK_getdualobj(void *,int,double *);
// extern int MSK_getdualsolutionnorms(void *,int,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_getdviolacc(void *,int,int64_t *,int64_t *,double *);
// extern int MSK_getdviolbarvar(void *,int,int32_t *,int32_t *,double *);
// extern int MSK_getdviolcon(void *,int,int32_t *,int32_t *,double *);
// extern int MSK_getdviolcones(void *,int,int32_t *,int32_t *,double *);
// extern int MSK_getdviolvar(void *,int,int32_t *,int32_t *,double *);
// extern int MSK_getinfindex(void *,int,const char *,int32_t *);
// extern int MSK_getinfmax(void *,int,int32_t *);
// extern int MSK_getinfname(void *,int,int32_t,char *);
// extern int MSK_getintinf(void *,int,int32_t *);
// extern int MSK_getintparam(void *,int,int32_t *);
// extern int MSK_getlenbarvarj(void *,int32_t,int64_t *);
// extern int MSK_getlintinf(void *,int,int64_t *);
// extern int MSK_getmaxnumanz64(void *,int64_t *);
// extern int MSK_getmaxnumbarvar(void *,int32_t *);
// extern int MSK_getmaxnumcon(void *,int32_t *);
// extern int MSK_getmaxnumcone(void *,int32_t *);
// extern int MSK_getmaxnumqnz64(void *,int64_t *);
// extern int MSK_getmaxnumvar(void *,int32_t *);
// extern int MSK_getmemusagetask(void *,int64_t *,int64_t *);
// extern int MSK_getnumacc(void *,int64_t *);
// extern int MSK_getnumafe(void *,int64_t *);
// extern int MSK_getnumanz(void *,int32_t *);
// extern int MSK_getnumanz64(void *,int64_t *);
// extern int MSK_getnumbarablocktriplets(void *,int64_t *);
// extern int MSK_getnumbaranz(void *,int64_t *);
// extern int MSK_getnumbarcblocktriplets(void *,int64_t *);
// extern int MSK_getnumbarcnz(void *,int64_t *);
// extern int MSK_getnumbarvar(void *,int32_t *);
// extern int MSK_getnumcon(void *,int32_t *);
// extern int MSK_getnumcone(void *,int32_t *);
// extern int MSK_getnumconemem(void *,int32_t,int32_t *);
// extern int MSK_getnumdjc(void *,int64_t *);
// extern int MSK_getnumdomain(void *,int64_t *);
// extern int MSK_getnumintvar(void *,int32_t *);
// extern int MSK_getnumparam(void *,int,int32_t *);
// extern int MSK_getnumqconknz64(void *,int32_t,int64_t *);
// extern int MSK_getnumqobjnz64(void *,int64_t *);
// extern int MSK_getnumsymmat(void *,int64_t *);
// extern int MSK_getnumvar(void *,int32_t *);
// extern int MSK_getobjname(void *,int32_t *,char *);
// extern int MSK_getobjnamelen(void *,int32_t *);
// extern int MSK_getobjsense(void *,int *);
// extern int MSK_getparammax(void *,int,int32_t *);
// extern int MSK_getparamname(void *,int,int32_t,char *);
// extern int MSK_getpowerdomainalpha(void *,int64_t,double *);
// extern int MSK_getpowerdomaininfo(void *,int64_t,int64_t *,int64_t *);
// extern int MSK_getprimalobj(void *,int,double *);
// extern int MSK_getprimalsolutionnorms(void *,int,double *,double *,double *);
// extern int MSK_getprobtype(void *,int *);
// extern int MSK_getprosta(void *,int,int *);
// extern int MSK_getpviolacc(void *,int,int64_t *,int64_t *,double *);
// extern int MSK_getpviolbarvar(void *,int,int32_t *,int32_t *,double *);
// extern int MSK_getpviolcon(void *,int,int32_t *,int32_t *,double *);
// extern int MSK_getpviolcones(void *,int,int32_t *,int32_t *,double *);
// extern int MSK_getpvioldjc(void *,int,int64_t *,int64_t *,double *);
// extern int MSK_getpviolvar(void *,int,int32_t *,int32_t *,double *);
// extern int MSK_getqconk64(void *,int32_t,int64_t *,int64_t *,int32_t *,int32_t *,double *);
// extern int MSK_getqobj64(void *,int64_t *,int64_t *,int32_t *,int32_t *,double *);
// extern int MSK_getqobjij(void *,int32_t,int32_t,double *);
// extern int MSK_getreducedcosts(void *,int,int32_t,int32_t,double *);
// extern int MSK_getskc(void *,int,int *);
// extern int MSK_getskcslice(void *,int,int32_t,int32_t,int *);
// extern int MSK_getskn(void *,int,int *);
// extern int MSK_getskx(void *,int,int *);
// extern int MSK_getskxslice(void *,int,int32_t,int32_t,int *);
// extern int MSK_getslc(void *,int,double *);
// extern int MSK_getslcslice(void *,int,int32_t,int32_t,double *);
// extern int MSK_getslx(void *,int,double *);
// extern int MSK_getslxslice(void *,int,int32_t,int32_t,double *);
// extern int MSK_getsnx(void *,int,double *);
// extern int MSK_getsnxslice(void *,int,int32_t,int32_t,double *);
// extern int MSK_getsolsta(void *,int,int *);
// extern int MSK_getsolution(void *,int,int *,int *,int *,int *,int *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_getsolutioninfo(void *,int,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_getsolutioninfonew(void *,int,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_getsolutionnew(void *,int,int *,int *,int *,int *,int *,double *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_getsolutionslice(void *,int,int,int32_t,int32_t,double *);
// extern int MSK_getsparsesymmat(void *,int64_t,int64_t *,int32_t *,int32_t *,double *);
// extern int MSK_getstrparam(void *,int,int32_t *,int32_t *,char *);
// extern int MSK_getstrparamlen(void *,int,int32_t *);
// extern int MSK_getsuc(void *,int,double *);
// extern int MSK_getsucslice(void *,int,int32_t,int32_t,double *);
// extern int MSK_getsux(void *,int,double *);
// extern int MSK_getsuxslice(void *,int,int32_t,int32_t,double *);
// extern int MSK_getsymmatinfo(void *,int64_t,int32_t *,int64_t *,int *);
// extern int MSK_gettaskname(void *,int32_t *,char *);
// extern int MSK_gettasknamelen(void *,int32_t *);
// extern int MSK_getvarbound(void *,int32_t,int *,double *,double *);
// extern int MSK_getvarboundslice(void *,int32_t,int32_t,int *,double *,double *);
// extern int MSK_getvarname(void *,int32_t,int32_t *,char *);
// extern int MSK_getvarnameindex(void *,const char *,int32_t *,int32_t *);
// extern int MSK_getvarnamelen(void *,int32_t,int32_t *);
// extern int MSK_getvartype(void *,int32_t,int *);
// extern int MSK_getvartypelist(void *,int32_t *,int32_t *,int *);
// extern int MSK_getxc(void *,int,double *);
// extern int MSK_getxcslice(void *,int,int32_t,int32_t,double *);
// extern int MSK_getxx(void *,int,double *);
// extern int MSK_getxxslice(void *,int,int32_t,int32_t,double *);
// extern int MSK_gety(void *,int,double *);
// extern int MSK_getyslice(void *,int,int32_t,int32_t,double *);
// extern int MSK_infeasibilityreport(void *,int,int);
// extern int MSK_initbasissolve(void *,int32_t *);
// extern int MSK_inputdata64(void *,int32_t,int32_t,int32_t *,int32_t *,double *,double,int64_t *,int64_t *,int32_t *,double *,int *,double *,double *,int *,double *,double *);
// extern int MSK_isdouparname(void *,const char *,int *);
// extern int MSK_isintparname(void *,const char *,int *);
// extern int MSK_isstrparname(void *,const char *,int *);
// extern int MSK_linkfiletotaskstream(void *,int,const char *,int32_t);
// extern int MSK_onesolutionsummary(void *,int,int);
// extern int MSK_optimizersummary(void *,int);
// extern int MSK_optimizetrm(void *,int *);
// extern int MSK_primalrepair(void *,double *,double *,double *,double *);
// extern int MSK_primalsensitivity(void *,int32_t *,int32_t *,int *,int32_t *,int32_t *,int *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_probtypetostr(void *,int,char *);
// extern int MSK_prostatostr(void *,int,char *);
// extern int MSK_putacc(void *,int64_t,int64_t,int64_t *,int64_t *,double *);
// extern int MSK_putaccb(void *,int64_t,int64_t *,double *);
// extern int MSK_putaccbj(void *,int64_t,int64_t,double);
// extern int MSK_putaccdoty(void *,int,int64_t,double *);
// extern int MSK_putacclist(void *,int64_t *,int64_t *,int64_t *,int64_t *,int64_t *,double *);
// extern int MSK_putaccname(void *,int64_t,const char *);
// extern int MSK_putacol(void *,int32_t,int32_t *,int32_t *,double *);
// extern int MSK_putacollist(void *,int32_t *,int32_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_putacolslice64(void *,int32_t,int32_t,int64_t *,int64_t *,int32_t *,double *);
// extern int MSK_putafebarfblocktriplet(void *,int64_t *,int64_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_putafebarfentry(void *,int64_t,int32_t,int64_t *,int64_t *,double *);
// extern int MSK_putafebarfentrylist(void *,int64_t *,int64_t *,int32_t *,int64_t *,int64_t *,int64_t *,int64_t *,double *);
// extern int MSK_putafebarfrow(void *,int64_t,int32_t *,int32_t *,int64_t *,int64_t *,int64_t *,int64_t *,double *);
// extern int MSK_putafefcol(void *,int32_t,int64_t *,int64_t *,double *);
// extern int MSK_putafefentry(void *,int64_t,int32_t,double);
// extern int MSK_putafefentrylist(void *,int64_t *,int64_t *,int32_t *,double *);
// extern int MSK_putafefrow(void *,int64_t,int32_t *,int32_t *,double *);
// extern int MSK_putafefrowlist(void *,int64_t *,int64_t *,int32_t *,int64_t *,int64_t *,int32_t *,double *);
// extern int MSK_putafeg(void *,int64_t,double);
// extern int MSK_putafeglist(void *,int64_t *,int64_t *,double *);
// extern int MSK_putafegslice(void *,int64_t,int64_t,double *);
// extern int MSK_putaij(void *,int32_t,int32_t,double);
// extern int MSK_putaijlist64(void *,int64_t *,int32_t *,int32_t *,double *);
// extern int MSK_putarow(void *,int32_t,int32_t *,int32_t *,double *);
// extern int MSK_putarowlist64(void *,int32_t *,int32_t *,int64_t *,int64_t *,int32_t *,double *);
// extern int MSK_putarowslice64(void *,int32_t,int32_t,int64_t *,int64_t *,int32_t *,double *);
// extern int MSK_putatruncatetol(void *,double);
// extern int MSK_putbarablocktriplet(void *,int64_t *,int32_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_putbaraij(void *,int32_t,int32_t,int64_t *,int64_t *,double *);
// extern int MSK_putbaraijlist(void *,int32_t *,int32_t *,int32_t *,int64_t *,int64_t *,int64_t *,double *);
// extern int MSK_putbararowlist(void *,int32_t *,int32_t *,int64_t *,int64_t *,int32_t *,int64_t *,int64_t *,double *);
// extern int MSK_putbarcblocktriplet(void *,int64_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_putbarcj(void *,int32_t,int64_t *,int64_t *,double *);
// extern int MSK_putbarsj(void *,int,int32_t,double *);
// extern int MSK_putbarvarname(void *,int32_t,const char *);
// extern int MSK_putbarxj(void *,int,int32_t,double *);
// extern int MSK_putcfix(void *,double);
// extern int MSK_putcj(void *,int32_t,double);
// extern int MSK_putclist(void *,int32_t *,int32_t *,double *);
// extern int MSK_putconbound(void *,int32_t,int,double,double);
// extern int MSK_putconboundlist(void *,int32_t *,int32_t *,int *,double *,double *);
// extern int MSK_putconboundlistconst(void *,int32_t *,int32_t *,int,double,double);
// extern int MSK_putconboundslice(void *,int32_t,int32_t,int *,double *,double *);
// extern int MSK_putconboundsliceconst(void *,int32_t,int32_t,int,double,double);
// extern int MSK_putcone(void *,int32_t,int,double,int32_t *,int32_t *);
// extern int MSK_putconename(void *,int32_t,const char *);
// extern int MSK_putconname(void *,int32_t,const char *);
// extern int MSK_putconsolutioni(void *,int32_t,int,int,double,double,double);
// extern int MSK_putcslice(void *,int32_t,int32_t,double *);
// extern int MSK_putdjc(void *,int64_t,int64_t *,int64_t *,int64_t *,int64_t *,double *,int64_t *,int64_t *);
// extern int MSK_putdjcname(void *,int64_t,const char *);
// extern int MSK_putdjcslice(void *,int64_t,int64_t,int64_t *,int64_t *,int64_t *,int64_t *,double *,int64_t *,int64_t *,int64_t *);
// extern int MSK_putdomainname(void *,int64_t,const char *);
// extern int MSK_putdouparam(void *,int,double);
// extern int MSK_putintparam(void *,int,int32_t);
// extern int MSK_putmaxnumacc(void *,int64_t);
// extern int MSK_putmaxnumafe(void *,int64_t);
// extern int MSK_putmaxnumanz(void *,int64_t);
// extern int MSK_putmaxnumbarvar(void *,int32_t);
// extern int MSK_putmaxnumcon(void *,int32_t);
// extern int MSK_putmaxnumcone(void *,int32_t);
// extern int MSK_putmaxnumdjc(void *,int64_t);
// extern int MSK_putmaxnumdomain(void *,int64_t);
// extern int MSK_putmaxnumqnz(void *,int64_t);
// extern int MSK_putmaxnumvar(void *,int32_t);
// extern int MSK_putnadouparam(void *,const char *,double);
// extern int MSK_putnaintparam(void *,const char *,int32_t);
// extern int MSK_putnastrparam(void *,const char *,const char *);
// extern int MSK_putobjname(void *,const char *);
// extern int MSK_putobjsense(void *,int);
// extern int MSK_putoptserverhost(void *,const char *);
// extern int MSK_putparam(void *,const char *,const char *);
// extern int MSK_putqcon(void *,int32_t *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_putqconk(void *,int32_t,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_putqobj(void *,int32_t *,int32_t *,int32_t *,double *);
// extern int MSK_putqobjij(void *,int32_t,int32_t,double);
// extern int MSK_putskc(void *,int,int *);
// extern int MSK_putskcslice(void *,int,int32_t,int32_t,int *);
// extern int MSK_putskx(void *,int,int *);
// extern int MSK_putskxslice(void *,int,int32_t,int32_t,int *);
// extern int MSK_putslc(void *,int,double *);
// extern int MSK_putslcslice(void *,int,int32_t,int32_t,double *);
// extern int MSK_putslx(void *,int,double *);
// extern int MSK_putslxslice(void *,int,int32_t,int32_t,double *);
// extern int MSK_putsnx(void *,int,double *);
// extern int MSK_putsnxslice(void *,int,int32_t,int32_t,double *);
// extern int MSK_putsolution(void *,int,int *,int *,int *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_putsolutionnew(void *,int,int *,int *,int *,double *,double *,double *,double *,double *,double *,double *,double *,double *);
// extern int MSK_putsolutionyi(void *,int32_t,int,double);
// extern int MSK_putstrparam(void *,int,const char *);
// extern int MSK_putsuc(void *,int,double *);
// extern int MSK_putsucslice(void *,int,int32_t,int32_t,double *);
// extern int MSK_putsux(void *,int,double *);
// extern int MSK_putsuxslice(void *,int,int32_t,int32_t,double *);
// extern int MSK_puttaskname(void *,const char *);
// extern int MSK_putvarbound(void *,int32_t,int,double,double);
// extern int MSK_putvarboundlist(void *,int32_t *,int32_t *,int *,double *,double *);
// extern int MSK_putvarboundlistconst(void *,int32_t *,int32_t *,int,double,double);
// extern int MSK_putvarboundslice(void *,int32_t,int32_t,int *,double *,double *);
// extern int MSK_putvarboundsliceconst(void *,int32_t,int32_t,int,double,double);
// extern int MSK_putvarname(void *,int32_t,const char *);
// extern int MSK_putvarsolutionj(void *,int32_t,int,int,double,double,double,double);
// extern int MSK_putvartype(void *,int32_t,int);
// extern int MSK_putvartypelist(void *,int32_t *,int32_t *,int *);
// extern int MSK_putxc(void *,int,double *);
// extern int MSK_putxcslice(void *,int,int32_t,int32_t,double *);
// extern int MSK_putxx(void *,int,double *);
// extern int MSK_putxxslice(void *,int,int32_t,int32_t,double *);
// extern int MSK_puty(void *,int,double *);
// extern int MSK_putyslice(void *,int,int32_t,int32_t,double *);
// extern int MSK_readbsolution(void *,const char *,int);
// extern int MSK_readdataautoformat(void *,const char *);
// extern int MSK_readdataformat(void *,const char *,int,int);
// extern int MSK_readjsonsol(void *,const char *);
// extern int MSK_readjsonstring(void *,const char *);
// extern int MSK_readlpstring(void *,const char *);
// extern int MSK_readopfstring(void *,const char *);
// extern int MSK_readparamfile(void *,const char *);
// extern int MSK_readptfstring(void *,const char *);
// extern int MSK_readsolution(void *,int,const char *);
// extern int MSK_readsolutionfile(void *,const char *);
// extern int MSK_readsummary(void *,int);
// extern int MSK_readtask(void *,const char *);
// extern int MSK_removebarvars(void *,int32_t *,int32_t *);
// extern int MSK_removecones(void *,int32_t *,int32_t *);
// extern int MSK_removecons(void *,int32_t *,int32_t *);
// extern int MSK_removevars(void *,int32_t *,int32_t *);
// extern int MSK_resizetask(void *,int32_t,int32_t,int32_t,int64_t,int64_t);
// extern int MSK_sensitivityreport(void *,int);
// extern int MSK_setdefaults(void *);
// extern int MSK_sktostr(void *,int,char *);
// extern int MSK_solstatostr(void *,int,char *);
// extern int MSK_solutiondef(void *,int,int *);
// extern int MSK_solutionsummary(void *,int);
// extern int MSK_solvewithbasis(void *,int,int32_t,int32_t *,double *,int32_t *);
// extern int MSK_strtoconetype(void *,const char *,int *);
// extern int MSK_strtosk(void *,const char *,int *);
// extern int MSK_toconic(void *);
// extern int MSK_updatesolutioninfo(void *,int);
// extern int MSK_writebsolution(void *,const char *,int);
// extern int MSK_writedata(void *,const char *);
// extern int MSK_writejsonsol(void *,const char *);
// extern int MSK_writeparamfile(void *,const char *);
// extern int MSK_writesolution(void *,int,const char *);
// extern int MSK_writesolutionfile(void *,const char *);
// extern int MSK_writetask(void *,const char *);
// extern int MSK_axpy(void *,int32_t,double,double *,double *);
// extern int MSK_checkinall(void *);
// extern int MSK_checkinlicense(void *,int);
// extern int MSK_checkoutlicense(void *,int);
// extern int MSK_dot(void *,int32_t,double *,double *,double *);
// extern int MSK_echointro(void *,int32_t);
// extern int MSK_expirylicenses(void *,int64_t *);
// extern int MSK_gemm(void *,int,int,int32_t,int32_t,int32_t,double,double *,double *,double,double *);
// extern int MSK_gemv(void *,int,int32_t,int32_t,double,double *,double *,double,double *);
// extern int MSK_getcodedesc(int,char *,char *);
// extern int MSK_getversion(int32_t *,int32_t *,int32_t *);
// extern int MSK_licensecleanup();
// extern int MSK_linkfiletoenvstream(void *,int,const char *,int32_t);
// extern int MSK_potrf(void *,int,int32_t,double *);
// extern int MSK_putlicensecode(void *,int32_t *);
// extern int MSK_putlicensedebug(void *,int32_t);
// extern int MSK_putlicensepath(void *,const char *);
// extern int MSK_putlicensewait(void *,int32_t);
// extern int MSK_resetexpirylicenses(void *);
// extern int MSK_syeig(void *,int,int32_t,double *,double *);
// extern int MSK_syevd(void *,int,int32_t,double *,double *);
// extern int MSK_syrk(void *,int,int,int32_t,int32_t,double,double *,double,double *);
import "C"

import (
    "unsafe"
    "fmt"
    "golang.org/x/exp/constraints"
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
        cptr unsafe.Pointer
}

type Task struct {
        r               Rescode
        cptr            unsafe.Pointer
	streamfunc      [4]func(string)
	callbackfunc    func(int32)int
	infcallbackfunc func(int32,[]float64,[]int32,[]int64)int
}

func (t * Task) ptr() C.MSKtask_t { return C.MSKtask_t(t.cptr) }
func (e * Env)  ptr() C.MSKenv_t  { return C.MSKenv_t(e.cptr) }

func (self * Env) getlasterror(res C.MSKrescodee) (Rescode,string) {
    return Rescode(res),""
}
func (self * Task) getlasterror(res C.MSKrescodee) (Rescode,string) {
    var lastcode C.MSKrescodee = res
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
func streamfunc_log(handle unsafe.Pointer, msg * C.char) {
	task := (*Task)(handle)
	if task.streamfunc[MSK_STREAM_LOG] != nil { task.streamfunc[MSK_STREAM_LOG](C.GoString(msg)) }
}

//export streamfunc_msg
func streamfunc_msg(handle unsafe.Pointer, msg * C.char) {
	task := (*Task)(handle)
	if task.streamfunc[MSK_STREAM_MSG] != nil { task.streamfunc[MSK_STREAM_MSG](C.GoString(msg)) }
}

//export streamfunc_wrn
func streamfunc_wrn(handle unsafe.Pointer, msg * C.char) {
	task := (*Task)(handle)
	if task.streamfunc[MSK_STREAM_WRN] != nil { task.streamfunc[MSK_STREAM_WRN](C.GoString(msg)) }
}

//export streamfunc_err
func streamfunc_err(handle unsafe.Pointer, msg * C.char) {
	task := (*Task)(handle)
	if task.streamfunc[MSK_STREAM_ERR] != nil { task.streamfunc[MSK_STREAM_ERR](C.GoString(msg)) }
}

//export callbackfunc
func callbackfunc(
	nativetask unsafe.Pointer,
	handle  unsafe.Pointer,
	code    C.int,
	dinf  * C.MSKrealt,
	iinf  * C.MSKint32t,
	liinf * C.MSKint64t) (C.int) {

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


func MakeEnv() (env Env, res int32) {
        var envptr C.MSKenv_t
        res = int32(C.MSK_makeenv(&envptr,nil))
        if res == 0 {
                env.cptr = unsafe.Pointer(envptr)
        }
        return
}

func (env *Env) MakeTask() (task Task, res int32) {
        var taskptr C.MSKtask_t
        res = int32(C.MSK_makeemptytask(env.ptr(), &taskptr))
        if res != 0 { return }
	task.cptr            = unsafe.Pointer(taskptr)
	task.streamfunc[0]   = nil
	task.streamfunc[1]   = nil
	task.streamfunc[2]   = nil
	task.streamfunc[3]   = nil
	task.callbackfunc    = nil
	task.infcallbackfunc = nil

        return
}

func (e *Env) Delete() {
        envptr := e.ptr()
        _ = C.MSK_deleteenv(&envptr)
        e.cptr = nil
}

func (t *Task) Delete() {
        taskptr := t.ptr()
        _ = C.MSK_deletetask(&taskptr)
        t.cptr = nil
}

func (t *Task) PutStreamFunc(whichstream Streamtype, fun func(string)) {
	t.streamfunc[whichstream] = fun

	if fun == nil {
		C.MSK_linkfunctotaskstream(
			t.ptr(),
			C.MSKstreamtypee(whichstream),
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
			t.ptr(),
			C.MSKstreamtypee(whichstream),
			C.MSKuserhandle_t(unsafe.Pointer(t)),
			strmfun) // ?!?
	}
}

func (t *Task) PutCallbackFunc(fun func(int32) int) {
	t.callbackfunc = fun
	if fun == nil {
		C.MSK_putcallbackfunc(t.ptr(), nil, nil)
	} else {
		C.MSK_putcallbackfunc(t.ptr(), (*[0]byte)(C.callbackfunc), C.MSKuserhandle_t(unsafe.Pointer(t)))
	}
}

func (t *Task) PutInfoCallbackFunc(fun func(int32,[]float64,[]int32,[]int64) int) {
	t.infcallbackfunc = fun
	if fun == nil {
		C.MSK_putcallbackfunc(t.ptr(), nil, nil)
	} else {
		C.MSK_putcallbackfunc(t.ptr(), (*[0]byte)(C.callbackfunc), C.MSKuserhandle_t(unsafe.Pointer(t)))
	}
}

func (e * Env)  ClearError() { e.r = MSK_RES_OK }
func (t * Task) ClearError() { t.r = MSK_RES_OK }

//func (e * Env)  GetRes() int32 { return e.r }
//func (t * Task) GetRes() int32 { return t.r }

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
  if _tmp0 := C.MSK_analyzenames(self.ptr(),C.MSKstreamtypee(whichstream),C.MSKnametypee(nametype)); _tmp0 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp0)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AnalyzeProblem(whichstream Streamtype) (err error) {
  if _tmp1 := C.MSK_analyzeproblem(self.ptr(),C.MSKstreamtypee(whichstream)); _tmp1 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AnalyzeSolution(whichstream Streamtype,whichsol Soltype) (err error) {
  if _tmp2 := C.MSK_analyzesolution(self.ptr(),C.MSKstreamtypee(whichstream),C.MSKsoltypee(whichsol)); _tmp2 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp2)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendAcc(domidx int64,afeidxlist []int64,b []float64) (err error) {
  _tmp3 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp3)
  var _tmp4 *int64
  if afeidxlist != nil { _tmp4 = (*int64)(&afeidxlist[0]) }
  var _tmp5 *float64
  if int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"AppendAcc",arg:"b"}
    return
  }
  if b != nil { _tmp5 = (*float64)(&b[0]) }
  if _tmp6 := C.MSK_appendacc(self.ptr(),C.MSKint32t(domidx),C.MSKint32t(numafeidx),_tmp4,_tmp5); _tmp6 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp6)
    err = &MosekError{code:lastcode,msg:lastmsg}
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
  var _tmp11 *float64
  if int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"AppendAccs",arg:"b"}
    return
  }
  if b != nil { _tmp11 = (*float64)(&b[0]) }
  if _tmp12 := C.MSK_appendaccs(self.ptr(),C.MSKint32t(numaccs),_tmp8,C.MSKint32t(numafeidx),_tmp10,_tmp11); _tmp12 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp12)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendAccSeq(domidx int64,afeidxfirst int64,b []float64) (err error) {
  var _tmp13 int64
  if _tmp14 := C.MSK_getdomainn(self.ptr(),domidx,addr(_tmp13)); _tmp14 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp14)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var numafeidx int64 = int64(_tmp13)
  var _tmp15 *float64
  if int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"AppendAccSeq",arg:"b"}
    return
  }
  if b != nil { _tmp15 = (*float64)(&b[0]) }
  if _tmp16 := C.MSK_appendaccseq(self.ptr(),C.MSKint32t(domidx),numafeidx,C.MSKint32t(afeidxfirst),_tmp15); _tmp16 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp16)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendAccsSeq(domidxs []int64,numafeidx int64,afeidxfirst int64,b []float64) (err error) {
  _tmp17 := len(domidxs)
  var numaccs int64 = int64(_tmp17)
  var _tmp18 *int64
  if domidxs != nil { _tmp18 = (*int64)(&domidxs[0]) }
  var _tmp19 *float64
  if int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"AppendAccsSeq",arg:"b"}
    return
  }
  if b != nil { _tmp19 = (*float64)(&b[0]) }
  if _tmp20 := C.MSK_appendaccsseq(self.ptr(),C.MSKint32t(numaccs),_tmp18,C.MSKint32t(numafeidx),C.MSKint32t(afeidxfirst),_tmp19); _tmp20 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp20)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendAfes(num int64) (err error) {
  if _tmp21 := C.MSK_appendafes(self.ptr(),C.MSKint32t(num)); _tmp21 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp21)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendBarvars(dim []int32) (err error) {
  _tmp22 := len(dim)
  var num int32 = int32(_tmp22)
  var _tmp23 *int32
  if dim != nil { _tmp23 = (*int32)(&dim[0]) }
  if _tmp24 := C.MSK_appendbarvars(self.ptr(),C.MSKint32t(num),_tmp23); _tmp24 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp24)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendCone(ct Conetype,conepar float64,submem []int32) (err error) {
  _tmp25 := len(submem)
  var nummem int32 = int32(_tmp25)
  var _tmp26 *int32
  if submem != nil { _tmp26 = (*int32)(&submem[0]) }
  if _tmp27 := C.MSK_appendcone(self.ptr(),C.MSKconetypee(ct),C.MSKrealt(conepar),C.MSKint32t(nummem),_tmp26); _tmp27 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp27)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendConeSeq(ct Conetype,conepar float64,nummem int32,j int32) (err error) {
  if _tmp28 := C.MSK_appendconeseq(self.ptr(),C.MSKconetypee(ct),C.MSKrealt(conepar),C.MSKint32t(nummem),C.MSKint32t(j)); _tmp28 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp28)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendConesSeq(ct []Conetype,conepar []float64,nummem []int32,j int32) (err error) {
  _tmp29 := len(conepar)
  if _tmp29 < ct { _tmp29 = lof["name"] }
  if _tmp29 < nummem { _tmp29 = lof["name"] }
  var num int32 = int32(_tmp29)
  var _tmp30 *Conetype
  if ct != nil { _tmp30 = (*Conetype)(&ct[0]) }
  var _tmp31 *float64
  if conepar != nil { _tmp31 = (*float64)(&conepar[0]) }
  var _tmp32 *int32
  if nummem != nil { _tmp32 = (*int32)(&nummem[0]) }
  if _tmp33 := C.MSK_appendconesseq(self.ptr(),C.MSKint32t(num),_tmp30,_tmp31,_tmp32,C.MSKint32t(j)); _tmp33 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp33)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendCons(num int32) (err error) {
  if _tmp34 := C.MSK_appendcons(self.ptr(),C.MSKint32t(num)); _tmp34 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp34)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendDjcs(num int64) (err error) {
  if _tmp35 := C.MSK_appenddjcs(self.ptr(),C.MSKint32t(num)); _tmp35 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp35)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendDualExpConeDomain() (domidx int64,err error) {
  if _tmp36 := C.MSK_appenddualexpconedomain(self.ptr(),&domidx); _tmp36 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp36)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendDualGeoMeanConeDomain(n int64) (domidx int64,err error) {
  if _tmp37 := C.MSK_appenddualgeomeanconedomain(self.ptr(),C.MSKint32t(n),&domidx); _tmp37 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp37)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendDualPowerConeDomain(n int64,alpha []float64) (domidx int64,err error) {
  _tmp38 := len(alpha)
  var nleft int64 = int64(_tmp38)
  var _tmp39 *float64
  if alpha != nil { _tmp39 = (*float64)(&alpha[0]) }
  if _tmp40 := C.MSK_appenddualpowerconedomain(self.ptr(),C.MSKint32t(n),C.MSKint32t(nleft),_tmp39,&domidx); _tmp40 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp40)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendPrimalExpConeDomain() (domidx int64,err error) {
  if _tmp41 := C.MSK_appendprimalexpconedomain(self.ptr(),&domidx); _tmp41 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp41)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendPrimalGeoMeanConeDomain(n int64) (domidx int64,err error) {
  if _tmp42 := C.MSK_appendprimalgeomeanconedomain(self.ptr(),C.MSKint32t(n),&domidx); _tmp42 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp42)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendPrimalPowerConeDomain(n int64,alpha []float64) (domidx int64,err error) {
  _tmp43 := len(alpha)
  var nleft int64 = int64(_tmp43)
  var _tmp44 *float64
  if alpha != nil { _tmp44 = (*float64)(&alpha[0]) }
  if _tmp45 := C.MSK_appendprimalpowerconedomain(self.ptr(),C.MSKint32t(n),C.MSKint32t(nleft),_tmp44,&domidx); _tmp45 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp45)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendQuadraticConeDomain(n int64) (domidx int64,err error) {
  if _tmp46 := C.MSK_appendquadraticconedomain(self.ptr(),C.MSKint32t(n),&domidx); _tmp46 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp46)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendRDomain(n int64) (domidx int64,err error) {
  if _tmp47 := C.MSK_appendrdomain(self.ptr(),C.MSKint32t(n),&domidx); _tmp47 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp47)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendRminusDomain(n int64) (domidx int64,err error) {
  if _tmp48 := C.MSK_appendrminusdomain(self.ptr(),C.MSKint32t(n),&domidx); _tmp48 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp48)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendRplusDomain(n int64) (domidx int64,err error) {
  if _tmp49 := C.MSK_appendrplusdomain(self.ptr(),C.MSKint32t(n),&domidx); _tmp49 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp49)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendRQuadraticConeDomain(n int64) (domidx int64,err error) {
  if _tmp50 := C.MSK_appendrquadraticconedomain(self.ptr(),C.MSKint32t(n),&domidx); _tmp50 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp50)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendRzeroDomain(n int64) (domidx int64,err error) {
  if _tmp51 := C.MSK_appendrzerodomain(self.ptr(),C.MSKint32t(n),&domidx); _tmp51 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp51)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendSparseSymMat(dim int32,subi []int32,subj []int32,valij []float64) (idx int64,err error) {
  _tmp52 := len(subi)
  if _tmp52 < valij { _tmp52 = lof["name"] }
  if _tmp52 < subj { _tmp52 = lof["name"] }
  var nz int64 = int64(_tmp52)
  var _tmp53 *int32
  if subi != nil { _tmp53 = (*int32)(&subi[0]) }
  var _tmp54 *int32
  if subj != nil { _tmp54 = (*int32)(&subj[0]) }
  var _tmp55 *float64
  if valij != nil { _tmp55 = (*float64)(&valij[0]) }
  if _tmp56 := C.MSK_appendsparsesymmat(self.ptr(),C.MSKint32t(dim),C.MSKint32t(nz),_tmp53,_tmp54,_tmp55,&idx); _tmp56 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp56)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendSparseSymMatList(dims []int32,nz []int64,subi []int32,subj []int32,valij []float64) (idx []int64,err error) {
  _tmp57 := len(dims)
  if _tmp57 < nz { _tmp57 = lof["name"] }
  var num int32 = int32(_tmp57)
  var _tmp58 *int32
  if dims != nil { _tmp58 = (*int32)(&dims[0]) }
  var _tmp59 *int64
  if nz != nil { _tmp59 = (*int64)(&nz[0]) }
  var _tmp60 *int32
  if int64(len(subi)) < int64(nz.foldl(a+b)) {
    err = &ArrayLengthError{fun:"AppendSparseSymMatList",arg:"subi"}
    return
  }
  if subi != nil { _tmp60 = (*int32)(&subi[0]) }
  var _tmp61 *int32
  if int64(len(subj)) < int64(nz.foldl(a+b)) {
    err = &ArrayLengthError{fun:"AppendSparseSymMatList",arg:"subj"}
    return
  }
  if subj != nil { _tmp61 = (*int32)(&subj[0]) }
  var _tmp62 *float64
  if int64(len(valij)) < int64(nz.foldl(a+b)) {
    err = &ArrayLengthError{fun:"AppendSparseSymMatList",arg:"valij"}
    return
  }
  if valij != nil { _tmp62 = (*float64)(&valij[0]) }
  var _tmp63 *int64
  idx := make([]int64,num)
  if len(idx) > 0 { _tmp63 = (*int64)(&n[0]) }
  if _tmp64 := C.MSK_appendsparsesymmatlist(self.ptr(),C.MSKint32t(num),_tmp58,_tmp59,_tmp60,_tmp61,_tmp62,_tmp63); _tmp64 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp64)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendSvecPsdConeDomain(n int64) (domidx int64,err error) {
  if _tmp65 := C.MSK_appendsvecpsdconedomain(self.ptr(),C.MSKint32t(n),&domidx); _tmp65 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp65)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) AppendVars(num int32) (err error) {
  if _tmp66 := C.MSK_appendvars(self.ptr(),C.MSKint32t(num)); _tmp66 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp66)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) BasisCond() (nrmbasis float64,nrminvbasis float64,err error) {
  if _tmp67 := C.MSK_basiscond(self.ptr(),&nrmbasis,&nrminvbasis); _tmp67 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp67)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) CheckMem(file string,line int32) (err error) {
  _tmp68 := C.CString(file)
  if _tmp69 := C.MSK_checkmemtask(self.ptr(),_tmp68,C.MSKint32t(line)); _tmp69 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp69)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) ChgConBound(i int32,lower int32,finite int32,value float64) (err error) {
  if _tmp70 := C.MSK_chgconbound(self.ptr(),C.MSKint32t(i),C.MSKint32t(lower),C.MSKint32t(finite),C.MSKrealt(value)); _tmp70 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp70)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) ChgVarBound(j int32,lower int32,finite int32,value float64) (err error) {
  if _tmp71 := C.MSK_chgvarbound(self.ptr(),C.MSKint32t(j),C.MSKint32t(lower),C.MSKint32t(finite),C.MSKrealt(value)); _tmp71 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp71)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) CommitChanges() (err error) {
  if _tmp72 := C.MSK_commitchanges(self.ptr()); _tmp72 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp72)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) DeleteSolution(whichsol Soltype) (err error) {
  if _tmp73 := C.MSK_deletesolution(self.ptr(),C.MSKsoltypee(whichsol)); _tmp73 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp73)
    err = &MosekError{code:lastcode,msg:lastmsg}
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
  leftpricej := make([]float64,numj)
  if len(leftpricej) > 0 { _tmp76 = (*float64)(&n[0]) }
  var _tmp77 *float64
  rightpricej := make([]float64,numj)
  if len(rightpricej) > 0 { _tmp77 = (*float64)(&n[0]) }
  var _tmp78 *float64
  leftrangej := make([]float64,numj)
  if len(leftrangej) > 0 { _tmp78 = (*float64)(&n[0]) }
  var _tmp79 *float64
  rightrangej := make([]float64,numj)
  if len(rightrangej) > 0 { _tmp79 = (*float64)(&n[0]) }
  if _tmp80 := C.MSK_dualsensitivity(self.ptr(),C.MSKint32t(numj),_tmp75,_tmp76,_tmp77,_tmp78,_tmp79); _tmp80 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp80)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) EmptyAfeBarfRow(afeidx int64) (err error) {
  if _tmp81 := C.MSK_emptyafebarfrow(self.ptr(),C.MSKint32t(afeidx)); _tmp81 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp81)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) EmptyAfeBarfRowList(afeidxlist []int64) (err error) {
  _tmp82 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp82)
  var _tmp83 *int64
  if afeidxlist != nil { _tmp83 = (*int64)(&afeidxlist[0]) }
  if _tmp84 := C.MSK_emptyafebarfrowlist(self.ptr(),C.MSKint32t(numafeidx),_tmp83); _tmp84 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp84)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) EmptyAfeFCol(varidx int32) (err error) {
  if _tmp85 := C.MSK_emptyafefcol(self.ptr(),C.MSKint32t(varidx)); _tmp85 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp85)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) EmptyAfeFColList(varidx []int32) (err error) {
  _tmp86 := len(varidx)
  var numvaridx int64 = int64(_tmp86)
  var _tmp87 *int32
  if varidx != nil { _tmp87 = (*int32)(&varidx[0]) }
  if _tmp88 := C.MSK_emptyafefcollist(self.ptr(),C.MSKint32t(numvaridx),_tmp87); _tmp88 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp88)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) EmptyAfeFRow(afeidx int64) (err error) {
  if _tmp89 := C.MSK_emptyafefrow(self.ptr(),C.MSKint32t(afeidx)); _tmp89 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp89)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) EmptyAfeFRowList(afeidx []int64) (err error) {
  _tmp90 := len(afeidx)
  var numafeidx int64 = int64(_tmp90)
  var _tmp91 *int64
  if afeidx != nil { _tmp91 = (*int64)(&afeidx[0]) }
  if _tmp92 := C.MSK_emptyafefrowlist(self.ptr(),C.MSKint32t(numafeidx),_tmp91); _tmp92 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp92)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) EvaluateAcc(whichsol Soltype,accidx int64) (activity []float64,err error) {
  var _tmp95 *float64
  var _tmp93 int64
  if _tmp94 := C.MSK_getaccn(self.ptr(),accidx,addr(_tmp93)); _tmp94 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp94)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  activity := make([]float64,_tmp93)
  if len(activity) > 0 { _tmp95 = (*float64)(&n[0]) }
  if _tmp96 := C.MSK_evaluateacc(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(accidx),_tmp95); _tmp96 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp96)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) EvaluateAccs(whichsol Soltype) (activity []float64,err error) {
  var _tmp99 *float64
  var _tmp97 int64
  if _tmp98 := C.MSK_getaccntot(self.ptr(),addr(_tmp97)); _tmp98 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp98)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  activity := make([]float64,_tmp97)
  if len(activity) > 0 { _tmp99 = (*float64)(&n[0]) }
  if _tmp100 := C.MSK_evaluateaccs(self.ptr(),C.MSKsoltypee(whichsol),_tmp99); _tmp100 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp100)
    err = &MosekError{code:lastcode,msg:lastmsg}
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
  var _tmp106 *int64
  if int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateAccNames",arg:"sp"}
    return
  }
  if sp != nil { _tmp106 = (*int64)(&sp[0]) }
  _tmp107 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp107)
  var _tmp108 *int32
  if namedaxisidxs != nil { _tmp108 = (*int32)(&namedaxisidxs[0]) }
  _tmp109 := len(names)
  var numnames int64 = int64(_tmp109)
  var _tmp110 *string
  if names != nil { _tmp110 = (*string)(&names[0]) }
  if _tmp111 := C.MSK_generateaccnames(self.ptr(),C.MSKint32t(num),_tmp102,_tmp103,C.MSKint32t(ndims),_tmp105,_tmp106,C.MSKint32t(numnamedaxis),_tmp108,C.MSKint32t(numnames),_tmp110); _tmp111 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp111)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GenerateBarvarNames(subj []int32,fmt string,dims []int32,sp []int64,namedaxisidxs []int32,names []string) (err error) {
  _tmp112 := len(subj)
  var num int32 = int32(_tmp112)
  var _tmp113 *int32
  if subj != nil { _tmp113 = (*int32)(&subj[0]) }
  _tmp114 := C.CString(fmt)
  _tmp115 := len(dims)
  var ndims int32 = int32(_tmp115)
  var _tmp116 *int32
  if dims != nil { _tmp116 = (*int32)(&dims[0]) }
  var _tmp117 *int64
  if int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateBarvarNames",arg:"sp"}
    return
  }
  if sp != nil { _tmp117 = (*int64)(&sp[0]) }
  _tmp118 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp118)
  var _tmp119 *int32
  if namedaxisidxs != nil { _tmp119 = (*int32)(&namedaxisidxs[0]) }
  _tmp120 := len(names)
  var numnames int64 = int64(_tmp120)
  var _tmp121 *string
  if names != nil { _tmp121 = (*string)(&names[0]) }
  if _tmp122 := C.MSK_generatebarvarnames(self.ptr(),C.MSKint32t(num),_tmp113,_tmp114,C.MSKint32t(ndims),_tmp116,_tmp117,C.MSKint32t(numnamedaxis),_tmp119,C.MSKint32t(numnames),_tmp121); _tmp122 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp122)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GenerateConeNames(subk []int32,fmt string,dims []int32,sp []int64,namedaxisidxs []int32,names []string) (err error) {
  _tmp123 := len(subk)
  var num int32 = int32(_tmp123)
  var _tmp124 *int32
  if subk != nil { _tmp124 = (*int32)(&subk[0]) }
  _tmp125 := C.CString(fmt)
  _tmp126 := len(dims)
  var ndims int32 = int32(_tmp126)
  var _tmp127 *int32
  if dims != nil { _tmp127 = (*int32)(&dims[0]) }
  var _tmp128 *int64
  if int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateConeNames",arg:"sp"}
    return
  }
  if sp != nil { _tmp128 = (*int64)(&sp[0]) }
  _tmp129 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp129)
  var _tmp130 *int32
  if namedaxisidxs != nil { _tmp130 = (*int32)(&namedaxisidxs[0]) }
  _tmp131 := len(names)
  var numnames int64 = int64(_tmp131)
  var _tmp132 *string
  if names != nil { _tmp132 = (*string)(&names[0]) }
  if _tmp133 := C.MSK_generateconenames(self.ptr(),C.MSKint32t(num),_tmp124,_tmp125,C.MSKint32t(ndims),_tmp127,_tmp128,C.MSKint32t(numnamedaxis),_tmp130,C.MSKint32t(numnames),_tmp132); _tmp133 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp133)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GenerateConNames(subi []int32,fmt string,dims []int32,sp []int64,namedaxisidxs []int32,names []string) (err error) {
  _tmp134 := len(subi)
  var num int32 = int32(_tmp134)
  var _tmp135 *int32
  if subi != nil { _tmp135 = (*int32)(&subi[0]) }
  _tmp136 := C.CString(fmt)
  _tmp137 := len(dims)
  var ndims int32 = int32(_tmp137)
  var _tmp138 *int32
  if dims != nil { _tmp138 = (*int32)(&dims[0]) }
  var _tmp139 *int64
  if int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateConNames",arg:"sp"}
    return
  }
  if sp != nil { _tmp139 = (*int64)(&sp[0]) }
  _tmp140 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp140)
  var _tmp141 *int32
  if namedaxisidxs != nil { _tmp141 = (*int32)(&namedaxisidxs[0]) }
  _tmp142 := len(names)
  var numnames int64 = int64(_tmp142)
  var _tmp143 *string
  if names != nil { _tmp143 = (*string)(&names[0]) }
  if _tmp144 := C.MSK_generateconnames(self.ptr(),C.MSKint32t(num),_tmp135,_tmp136,C.MSKint32t(ndims),_tmp138,_tmp139,C.MSKint32t(numnamedaxis),_tmp141,C.MSKint32t(numnames),_tmp143); _tmp144 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp144)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GenerateDjcNames(sub []int64,fmt string,dims []int32,sp []int64,namedaxisidxs []int32,names []string) (err error) {
  _tmp145 := len(sub)
  var num int64 = int64(_tmp145)
  var _tmp146 *int64
  if sub != nil { _tmp146 = (*int64)(&sub[0]) }
  _tmp147 := C.CString(fmt)
  _tmp148 := len(dims)
  var ndims int32 = int32(_tmp148)
  var _tmp149 *int32
  if dims != nil { _tmp149 = (*int32)(&dims[0]) }
  var _tmp150 *int64
  if int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateDjcNames",arg:"sp"}
    return
  }
  if sp != nil { _tmp150 = (*int64)(&sp[0]) }
  _tmp151 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp151)
  var _tmp152 *int32
  if namedaxisidxs != nil { _tmp152 = (*int32)(&namedaxisidxs[0]) }
  _tmp153 := len(names)
  var numnames int64 = int64(_tmp153)
  var _tmp154 *string
  if names != nil { _tmp154 = (*string)(&names[0]) }
  if _tmp155 := C.MSK_generatedjcnames(self.ptr(),C.MSKint32t(num),_tmp146,_tmp147,C.MSKint32t(ndims),_tmp149,_tmp150,C.MSKint32t(numnamedaxis),_tmp152,C.MSKint32t(numnames),_tmp154); _tmp155 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp155)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GenerateVarNames(subj []int32,fmt string,dims []int32,sp []int64,namedaxisidxs []int32,names []string) (err error) {
  _tmp156 := len(subj)
  var num int32 = int32(_tmp156)
  var _tmp157 *int32
  if subj != nil { _tmp157 = (*int32)(&subj[0]) }
  _tmp158 := C.CString(fmt)
  _tmp159 := len(dims)
  var ndims int32 = int32(_tmp159)
  var _tmp160 *int32
  if dims != nil { _tmp160 = (*int32)(&dims[0]) }
  var _tmp161 *int64
  if int64(len(sp)) < int64(num) {
    err = &ArrayLengthError{fun:"GenerateVarNames",arg:"sp"}
    return
  }
  if sp != nil { _tmp161 = (*int64)(&sp[0]) }
  _tmp162 := len(namedaxisidxs)
  var numnamedaxis int32 = int32(_tmp162)
  var _tmp163 *int32
  if namedaxisidxs != nil { _tmp163 = (*int32)(&namedaxisidxs[0]) }
  _tmp164 := len(names)
  var numnames int64 = int64(_tmp164)
  var _tmp165 *string
  if names != nil { _tmp165 = (*string)(&names[0]) }
  if _tmp166 := C.MSK_generatevarnames(self.ptr(),C.MSKint32t(num),_tmp157,_tmp158,C.MSKint32t(ndims),_tmp160,_tmp161,C.MSKint32t(numnamedaxis),_tmp163,C.MSKint32t(numnames),_tmp165); _tmp166 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp166)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccAfeIdxList(accidx int64) (afeidxlist []int64,err error) {
  var _tmp169 *int64
  var _tmp167 int64
  if _tmp168 := C.MSK_getaccn(self.ptr(),accidx,addr(_tmp167)); _tmp168 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp168)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  afeidxlist := make([]int64,_tmp167)
  if len(afeidxlist) > 0 { _tmp169 = (*int64)(&n[0]) }
  if _tmp170 := C.MSK_getaccafeidxlist(self.ptr(),C.MSKint32t(accidx),_tmp169); _tmp170 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp170)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccB(accidx int64) (b []float64,err error) {
  var _tmp173 *float64
  var _tmp171 int64
  if _tmp172 := C.MSK_getaccn(self.ptr(),accidx,addr(_tmp171)); _tmp172 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp172)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  b := make([]float64,_tmp171)
  if len(b) > 0 { _tmp173 = (*float64)(&n[0]) }
  if _tmp174 := C.MSK_getaccb(self.ptr(),C.MSKint32t(accidx),_tmp173); _tmp174 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp174)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccBarfBlockTriplet() (numtrip int64,acc_afe []int64,bar_var []int32,blk_row []int32,blk_col []int32,blk_val []float64,err error) {
  var _tmp175 int64
  if _tmp176 := C.MSK_getaccbarfnumblocktriplets(self.ptr(),addr(_tmp175)); _tmp176 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp176)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var maxnumtrip int64 = int64(_tmp175)
  var _tmp177 *int64
  acc_afe := make([]int64,maxnumtrip)
  if len(acc_afe) > 0 { _tmp177 = (*int64)(&n[0]) }
  var _tmp178 *int32
  bar_var := make([]int32,maxnumtrip)
  if len(bar_var) > 0 { _tmp178 = (*int32)(&n[0]) }
  var _tmp179 *int32
  blk_row := make([]int32,maxnumtrip)
  if len(blk_row) > 0 { _tmp179 = (*int32)(&n[0]) }
  var _tmp180 *int32
  blk_col := make([]int32,maxnumtrip)
  if len(blk_col) > 0 { _tmp180 = (*int32)(&n[0]) }
  var _tmp181 *float64
  blk_val := make([]float64,maxnumtrip)
  if len(blk_val) > 0 { _tmp181 = (*float64)(&n[0]) }
  if _tmp182 := C.MSK_getaccbarfblocktriplet(self.ptr(),maxnumtrip,&numtrip,_tmp177,_tmp178,_tmp179,_tmp180,_tmp181); _tmp182 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp182)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccBarfNumBlockTriplets() (numtrip int64,err error) {
  if _tmp183 := C.MSK_getaccbarfnumblocktriplets(self.ptr(),&numtrip); _tmp183 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp183)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccDomain(accidx int64) (domidx int64,err error) {
  if _tmp184 := C.MSK_getaccdomain(self.ptr(),C.MSKint32t(accidx),&domidx); _tmp184 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp184)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccDotY(whichsol Soltype,accidx int64) (doty []float64,err error) {
  var _tmp187 *float64
  var _tmp185 int64
  if _tmp186 := C.MSK_getaccn(self.ptr(),accidx,addr(_tmp185)); _tmp186 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp186)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  doty := make([]float64,_tmp185)
  if len(doty) > 0 { _tmp187 = (*float64)(&n[0]) }
  if _tmp188 := C.MSK_getaccdoty(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(accidx),_tmp187); _tmp188 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp188)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccDotYS(whichsol Soltype) (doty []float64,err error) {
  var _tmp191 *float64
  var _tmp189 int64
  if _tmp190 := C.MSK_getaccntot(self.ptr(),addr(_tmp189)); _tmp190 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp190)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  doty := make([]float64,_tmp189)
  if len(doty) > 0 { _tmp191 = (*float64)(&n[0]) }
  if _tmp192 := C.MSK_getaccdotys(self.ptr(),C.MSKsoltypee(whichsol),_tmp191); _tmp192 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp192)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccFNumnz() (accfnnz int64,err error) {
  if _tmp193 := C.MSK_getaccfnumnz(self.ptr(),&accfnnz); _tmp193 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp193)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccFTrip() (frow []int64,fcol []int32,fval []float64,err error) {
  var _tmp196 *int64
  var _tmp194 int64
  if _tmp195 := C.MSK_getaccfnumnz(self.ptr(),addr(_tmp194)); _tmp195 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp195)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  frow := make([]int64,_tmp194)
  if len(frow) > 0 { _tmp196 = (*int64)(&n[0]) }
  var _tmp199 *int32
  var _tmp197 int64
  if _tmp198 := C.MSK_getaccfnumnz(self.ptr(),addr(_tmp197)); _tmp198 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp198)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  fcol := make([]int32,_tmp197)
  if len(fcol) > 0 { _tmp199 = (*int32)(&n[0]) }
  var _tmp202 *float64
  var _tmp200 int64
  if _tmp201 := C.MSK_getaccfnumnz(self.ptr(),addr(_tmp200)); _tmp201 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp201)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  fval := make([]float64,_tmp200)
  if len(fval) > 0 { _tmp202 = (*float64)(&n[0]) }
  if _tmp203 := C.MSK_getaccftrip(self.ptr(),_tmp196,_tmp199,_tmp202); _tmp203 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp203)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccGVector() (g []float64,err error) {
  var _tmp206 *float64
  var _tmp204 int64
  if _tmp205 := C.MSK_getaccntot(self.ptr(),addr(_tmp204)); _tmp205 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp205)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  g := make([]float64,_tmp204)
  if len(g) > 0 { _tmp206 = (*float64)(&n[0]) }
  if _tmp207 := C.MSK_getaccgvector(self.ptr(),_tmp206); _tmp207 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp207)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccN(accidx int64) (n int64,err error) {
  if _tmp208 := C.MSK_getaccn(self.ptr(),C.MSKint32t(accidx),&n); _tmp208 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp208)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccName(accidx int64) (name string,err error) {
  var _tmp209 int32
  if _tmp210 := C.MSK_getaccnamelen(self.ptr(),accidx,addr(_tmp209)); _tmp210 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp210)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp209))
  _tmp211 := make([]byte,sizename)
  if _tmp212 := C.MSK_getaccname(self.ptr(),C.MSKint32t(accidx),sizename,C.CString(&tmpvar1[0])); _tmp212 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp212)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  var name string
  if p := strings.IndexByte(_tmp211,byte(0)); p < 0 {
    name = string(_tmp211)
  } else {
    name = string(_tmp211[:p])
  }
  return
}
func (self *Task) GetAccNameLen(accidx int64) (len int32,err error) {
  if _tmp213 := C.MSK_getaccnamelen(self.ptr(),C.MSKint32t(accidx),&len); _tmp213 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp213)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccNTot() (n int64,err error) {
  if _tmp214 := C.MSK_getaccntot(self.ptr(),&n); _tmp214 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp214)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAccs() (domidxlist []int64,afeidxlist []int64,b []float64,err error) {
  var _tmp217 *int64
  var _tmp215 int64
  if _tmp216 := C.MSK_getnumacc(self.ptr(),addr(_tmp215)); _tmp216 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp216)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  domidxlist := make([]int64,_tmp215)
  if len(domidxlist) > 0 { _tmp217 = (*int64)(&n[0]) }
  var _tmp220 *int64
  var _tmp218 int64
  if _tmp219 := C.MSK_getaccntot(self.ptr(),addr(_tmp218)); _tmp219 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp219)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  afeidxlist := make([]int64,_tmp218)
  if len(afeidxlist) > 0 { _tmp220 = (*int64)(&n[0]) }
  var _tmp223 *float64
  var _tmp221 int64
  if _tmp222 := C.MSK_getaccntot(self.ptr(),addr(_tmp221)); _tmp222 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp222)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  b := make([]float64,_tmp221)
  if len(b) > 0 { _tmp223 = (*float64)(&n[0]) }
  if _tmp224 := C.MSK_getaccs(self.ptr(),_tmp217,_tmp220,_tmp223); _tmp224 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp224)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetACol(j int32) (nzj int32,subj []int32,valj []float64,err error) {
  var _tmp227 *int32
  var _tmp225 int32
  if _tmp226 := C.MSK_getacolnumnz(self.ptr(),j,addr(_tmp225)); _tmp226 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp226)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  subj := make([]int32,_tmp225)
  if len(subj) > 0 { _tmp227 = (*int32)(&n[0]) }
  var _tmp230 *float64
  var _tmp228 int32
  if _tmp229 := C.MSK_getacolnumnz(self.ptr(),j,addr(_tmp228)); _tmp229 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp229)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  valj := make([]float64,_tmp228)
  if len(valj) > 0 { _tmp230 = (*float64)(&n[0]) }
  if _tmp231 := C.MSK_getacol(self.ptr(),C.MSKint32t(j),&nzj,_tmp227,_tmp230); _tmp231 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp231)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAColNumNz(i int32) (nzj int32,err error) {
  if _tmp232 := C.MSK_getacolnumnz(self.ptr(),C.MSKint32t(i),&nzj); _tmp232 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp232)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAColSlice(first int32,last int32) (ptrb []int64,ptre []int64,sub []int32,val []float64,err error) {
  var _tmp233 int64
  if _tmp234 := C.MSK_getacolslicenumnz64(self.ptr(),first,last,addr(_tmp233)); _tmp234 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp234)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp233)
  var _tmp235 *int64
  ptrb := make([]int64,(last - first))
  if len(ptrb) > 0 { _tmp235 = (*int64)(&n[0]) }
  var _tmp236 *int64
  ptre := make([]int64,(last - first))
  if len(ptre) > 0 { _tmp236 = (*int64)(&n[0]) }
  var _tmp237 *int32
  sub := make([]int32,maxnumnz)
  if len(sub) > 0 { _tmp237 = (*int32)(&n[0]) }
  var _tmp238 *float64
  val := make([]float64,maxnumnz)
  if len(val) > 0 { _tmp238 = (*float64)(&n[0]) }
  if _tmp239 := C.MSK_getacolslice64(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),maxnumnz,_tmp235,_tmp236,_tmp237,_tmp238); _tmp239 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp239)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAColSliceNumNz(first int32,last int32) (numnz int64,err error) {
  if _tmp240 := C.MSK_getacolslicenumnz64(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),&numnz); _tmp240 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp240)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAColSliceTrip(first int32,last int32) (subi []int32,subj []int32,val []float64,err error) {
  var _tmp241 int64
  if _tmp242 := C.MSK_getacolslicenumnz64(self.ptr(),first,last,addr(_tmp241)); _tmp242 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp242)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp241)
  var _tmp243 *int32
  subi := make([]int32,maxnumnz)
  if len(subi) > 0 { _tmp243 = (*int32)(&n[0]) }
  var _tmp244 *int32
  subj := make([]int32,maxnumnz)
  if len(subj) > 0 { _tmp244 = (*int32)(&n[0]) }
  var _tmp245 *float64
  val := make([]float64,maxnumnz)
  if len(val) > 0 { _tmp245 = (*float64)(&n[0]) }
  if _tmp246 := C.MSK_getacolslicetrip(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),maxnumnz,_tmp243,_tmp244,_tmp245); _tmp246 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp246)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeBarfBlockTriplet() (numtrip int64,afeidx []int64,barvaridx []int32,subk []int32,subl []int32,valkl []float64,err error) {
  var _tmp247 int64
  if _tmp248 := C.MSK_getafebarfnumblocktriplets(self.ptr(),addr(_tmp247)); _tmp248 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp248)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var maxnumtrip int64 = int64(_tmp247)
  var _tmp249 *int64
  afeidx := make([]int64,maxnumtrip)
  if len(afeidx) > 0 { _tmp249 = (*int64)(&n[0]) }
  var _tmp250 *int32
  barvaridx := make([]int32,maxnumtrip)
  if len(barvaridx) > 0 { _tmp250 = (*int32)(&n[0]) }
  var _tmp251 *int32
  subk := make([]int32,maxnumtrip)
  if len(subk) > 0 { _tmp251 = (*int32)(&n[0]) }
  var _tmp252 *int32
  subl := make([]int32,maxnumtrip)
  if len(subl) > 0 { _tmp252 = (*int32)(&n[0]) }
  var _tmp253 *float64
  valkl := make([]float64,maxnumtrip)
  if len(valkl) > 0 { _tmp253 = (*float64)(&n[0]) }
  if _tmp254 := C.MSK_getafebarfblocktriplet(self.ptr(),maxnumtrip,&numtrip,_tmp249,_tmp250,_tmp251,_tmp252,_tmp253); _tmp254 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp254)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeBarfNumBlockTriplets() (numtrip int64,err error) {
  if _tmp255 := C.MSK_getafebarfnumblocktriplets(self.ptr(),&numtrip); _tmp255 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp255)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeBarfNumRowEntries(afeidx int64) (numentr int32,err error) {
  if _tmp256 := C.MSK_getafebarfnumrowentries(self.ptr(),C.MSKint32t(afeidx),&numentr); _tmp256 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp256)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeBarfRow(afeidx int64) (barvaridx []int32,ptrterm []int64,numterm []int64,termidx []int64,termweight []float64,err error) {
  var _tmp260 *int32
  var _tmp257 int32
  var _tmp258 int64
  if _tmp259 := C.MSK_getafebarfrowinfo(self.ptr(),afeidx,addr(_tmp257),addr(_tmp258)); _tmp259 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp259)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  barvaridx := make([]int32,_tmp257)
  if len(barvaridx) > 0 { _tmp260 = (*int32)(&n[0]) }
  var _tmp264 *int64
  var _tmp261 int32
  var _tmp262 int64
  if _tmp263 := C.MSK_getafebarfrowinfo(self.ptr(),afeidx,addr(_tmp261),addr(_tmp262)); _tmp263 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp263)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  ptrterm := make([]int64,_tmp261)
  if len(ptrterm) > 0 { _tmp264 = (*int64)(&n[0]) }
  var _tmp268 *int64
  var _tmp265 int32
  var _tmp266 int64
  if _tmp267 := C.MSK_getafebarfrowinfo(self.ptr(),afeidx,addr(_tmp265),addr(_tmp266)); _tmp267 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp267)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  numterm := make([]int64,_tmp265)
  if len(numterm) > 0 { _tmp268 = (*int64)(&n[0]) }
  var _tmp272 *int64
  var _tmp269 int32
  var _tmp270 int64
  if _tmp271 := C.MSK_getafebarfrowinfo(self.ptr(),afeidx,addr(_tmp269),addr(_tmp270)); _tmp271 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp271)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  termidx := make([]int64,_tmp270)
  if len(termidx) > 0 { _tmp272 = (*int64)(&n[0]) }
  var _tmp276 *float64
  var _tmp273 int32
  var _tmp274 int64
  if _tmp275 := C.MSK_getafebarfrowinfo(self.ptr(),afeidx,addr(_tmp273),addr(_tmp274)); _tmp275 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp275)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  termweight := make([]float64,_tmp274)
  if len(termweight) > 0 { _tmp276 = (*float64)(&n[0]) }
  if _tmp277 := C.MSK_getafebarfrow(self.ptr(),C.MSKint32t(afeidx),_tmp260,_tmp264,_tmp268,_tmp272,_tmp276); _tmp277 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp277)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeBarfRowInfo(afeidx int64) (numentr int32,numterm int64,err error) {
  if _tmp278 := C.MSK_getafebarfrowinfo(self.ptr(),C.MSKint32t(afeidx),&numentr,&numterm); _tmp278 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp278)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeFNumNz() (numnz int64,err error) {
  if _tmp279 := C.MSK_getafefnumnz(self.ptr(),&numnz); _tmp279 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp279)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeFRow(afeidx int64) (numnz int32,varidx []int32,val []float64,err error) {
  var _tmp282 *int32
  var _tmp280 int32
  if _tmp281 := C.MSK_getafefrownumnz(self.ptr(),afeidx,addr(_tmp280)); _tmp281 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp281)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  varidx := make([]int32,_tmp280)
  if len(varidx) > 0 { _tmp282 = (*int32)(&n[0]) }
  var _tmp285 *float64
  var _tmp283 int32
  if _tmp284 := C.MSK_getafefrownumnz(self.ptr(),afeidx,addr(_tmp283)); _tmp284 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp284)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  val := make([]float64,_tmp283)
  if len(val) > 0 { _tmp285 = (*float64)(&n[0]) }
  if _tmp286 := C.MSK_getafefrow(self.ptr(),C.MSKint32t(afeidx),&numnz,_tmp282,_tmp285); _tmp286 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp286)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeFRowNumNz(afeidx int64) (numnz int32,err error) {
  if _tmp287 := C.MSK_getafefrownumnz(self.ptr(),C.MSKint32t(afeidx),&numnz); _tmp287 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp287)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeFTrip() (afeidx []int64,varidx []int32,val []float64,err error) {
  var _tmp290 *int64
  var _tmp288 int64
  if _tmp289 := C.MSK_getafefnumnz(self.ptr(),addr(_tmp288)); _tmp289 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp289)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  afeidx := make([]int64,_tmp288)
  if len(afeidx) > 0 { _tmp290 = (*int64)(&n[0]) }
  var _tmp293 *int32
  var _tmp291 int64
  if _tmp292 := C.MSK_getafefnumnz(self.ptr(),addr(_tmp291)); _tmp292 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp292)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  varidx := make([]int32,_tmp291)
  if len(varidx) > 0 { _tmp293 = (*int32)(&n[0]) }
  var _tmp296 *float64
  var _tmp294 int64
  if _tmp295 := C.MSK_getafefnumnz(self.ptr(),addr(_tmp294)); _tmp295 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp295)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  val := make([]float64,_tmp294)
  if len(val) > 0 { _tmp296 = (*float64)(&n[0]) }
  if _tmp297 := C.MSK_getafeftrip(self.ptr(),_tmp290,_tmp293,_tmp296); _tmp297 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp297)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeG(afeidx int64) (g float64,err error) {
  if _tmp298 := C.MSK_getafeg(self.ptr(),C.MSKint32t(afeidx),&g); _tmp298 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp298)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAfeGSlice(first int64,last int64) (g []float64,err error) {
  var _tmp299 *float64
  g := make([]float64,(last - first))
  if len(g) > 0 { _tmp299 = (*float64)(&n[0]) }
  if _tmp300 := C.MSK_getafegslice(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),_tmp299); _tmp300 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp300)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAij(i int32,j int32) (aij float64,err error) {
  if _tmp301 := C.MSK_getaij(self.ptr(),C.MSKint32t(i),C.MSKint32t(j),&aij); _tmp301 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp301)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetAPieceNumNz(firsti int32,lasti int32,firstj int32,lastj int32) (numnz int32,err error) {
  if _tmp302 := C.MSK_getapiecenumnz(self.ptr(),C.MSKint32t(firsti),C.MSKint32t(lasti),C.MSKint32t(firstj),C.MSKint32t(lastj),&numnz); _tmp302 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp302)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetARow(i int32) (nzi int32,subi []int32,vali []float64,err error) {
  var _tmp305 *int32
  var _tmp303 int32
  if _tmp304 := C.MSK_getarownumnz(self.ptr(),i,addr(_tmp303)); _tmp304 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp304)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  subi := make([]int32,_tmp303)
  if len(subi) > 0 { _tmp305 = (*int32)(&n[0]) }
  var _tmp308 *float64
  var _tmp306 int32
  if _tmp307 := C.MSK_getarownumnz(self.ptr(),i,addr(_tmp306)); _tmp307 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp307)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  vali := make([]float64,_tmp306)
  if len(vali) > 0 { _tmp308 = (*float64)(&n[0]) }
  if _tmp309 := C.MSK_getarow(self.ptr(),C.MSKint32t(i),&nzi,_tmp305,_tmp308); _tmp309 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp309)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetARowNumNz(i int32) (nzi int32,err error) {
  if _tmp310 := C.MSK_getarownumnz(self.ptr(),C.MSKint32t(i),&nzi); _tmp310 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp310)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetARowSlice(first int32,last int32) (ptrb []int64,ptre []int64,sub []int32,val []float64,err error) {
  var _tmp311 int64
  if _tmp312 := C.MSK_getarowslicenumnz64(self.ptr(),first,last,addr(_tmp311)); _tmp312 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp312)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp311)
  var _tmp313 *int64
  ptrb := make([]int64,(last - first))
  if len(ptrb) > 0 { _tmp313 = (*int64)(&n[0]) }
  var _tmp314 *int64
  ptre := make([]int64,(last - first))
  if len(ptre) > 0 { _tmp314 = (*int64)(&n[0]) }
  var _tmp315 *int32
  sub := make([]int32,maxnumnz)
  if len(sub) > 0 { _tmp315 = (*int32)(&n[0]) }
  var _tmp316 *float64
  val := make([]float64,maxnumnz)
  if len(val) > 0 { _tmp316 = (*float64)(&n[0]) }
  if _tmp317 := C.MSK_getarowslice64(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),maxnumnz,_tmp313,_tmp314,_tmp315,_tmp316); _tmp317 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp317)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetARowSliceNumNz(first int32,last int32) (numnz int64,err error) {
  if _tmp318 := C.MSK_getarowslicenumnz64(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),&numnz); _tmp318 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp318)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetARowSliceTrip(first int32,last int32) (subi []int32,subj []int32,val []float64,err error) {
  var _tmp319 int64
  if _tmp320 := C.MSK_getarowslicenumnz64(self.ptr(),first,last,addr(_tmp319)); _tmp320 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp320)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp319)
  var _tmp321 *int32
  subi := make([]int32,maxnumnz)
  if len(subi) > 0 { _tmp321 = (*int32)(&n[0]) }
  var _tmp322 *int32
  subj := make([]int32,maxnumnz)
  if len(subj) > 0 { _tmp322 = (*int32)(&n[0]) }
  var _tmp323 *float64
  val := make([]float64,maxnumnz)
  if len(val) > 0 { _tmp323 = (*float64)(&n[0]) }
  if _tmp324 := C.MSK_getarowslicetrip(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),maxnumnz,_tmp321,_tmp322,_tmp323); _tmp324 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp324)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetATrip() (subi []int32,subj []int32,val []float64,err error) {
  var _tmp325 int64
  if _tmp326 := C.MSK_getnumanz64(self.ptr(),addr(_tmp325)); _tmp326 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp326)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp325)
  var _tmp327 *int32
  subi := make([]int32,maxnumnz)
  if len(subi) > 0 { _tmp327 = (*int32)(&n[0]) }
  var _tmp328 *int32
  subj := make([]int32,maxnumnz)
  if len(subj) > 0 { _tmp328 = (*int32)(&n[0]) }
  var _tmp329 *float64
  val := make([]float64,maxnumnz)
  if len(val) > 0 { _tmp329 = (*float64)(&n[0]) }
  if _tmp330 := C.MSK_getatrip(self.ptr(),maxnumnz,_tmp327,_tmp328,_tmp329); _tmp330 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp330)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetATruncateTol() (tolzero []float64,err error) {
  var _tmp331 *float64
  tolzero := make([]float64,1)
  if len(tolzero) > 0 { _tmp331 = (*float64)(&n[0]) }
  if _tmp332 := C.MSK_getatruncatetol(self.ptr(),_tmp331); _tmp332 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp332)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBaraBlockTriplet() (num int64,subi []int32,subj []int32,subk []int32,subl []int32,valijkl []float64,err error) {
  var _tmp333 int64
  if _tmp334 := C.MSK_getnumbarablocktriplets(self.ptr(),addr(_tmp333)); _tmp334 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp334)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var maxnum int64 = int64(_tmp333)
  var _tmp335 *int32
  subi := make([]int32,maxnum)
  if len(subi) > 0 { _tmp335 = (*int32)(&n[0]) }
  var _tmp336 *int32
  subj := make([]int32,maxnum)
  if len(subj) > 0 { _tmp336 = (*int32)(&n[0]) }
  var _tmp337 *int32
  subk := make([]int32,maxnum)
  if len(subk) > 0 { _tmp337 = (*int32)(&n[0]) }
  var _tmp338 *int32
  subl := make([]int32,maxnum)
  if len(subl) > 0 { _tmp338 = (*int32)(&n[0]) }
  var _tmp339 *float64
  valijkl := make([]float64,maxnum)
  if len(valijkl) > 0 { _tmp339 = (*float64)(&n[0]) }
  if _tmp340 := C.MSK_getbarablocktriplet(self.ptr(),maxnum,&num,_tmp335,_tmp336,_tmp337,_tmp338,_tmp339); _tmp340 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp340)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBaraIdx(idx int64) (i int32,j int32,num int64,sub []int64,weights []float64,err error) {
  var _tmp341 int64
  if _tmp342 := C.MSK_getbaraidxinfo(self.ptr(),idx,addr(_tmp341)); _tmp342 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp342)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var maxnum int64 = int64(_tmp341)
  var _tmp343 *int64
  sub := make([]int64,maxnum)
  if len(sub) > 0 { _tmp343 = (*int64)(&n[0]) }
  var _tmp344 *float64
  weights := make([]float64,maxnum)
  if len(weights) > 0 { _tmp344 = (*float64)(&n[0]) }
  if _tmp345 := C.MSK_getbaraidx(self.ptr(),C.MSKint32t(idx),maxnum,&i,&j,&num,_tmp343,_tmp344); _tmp345 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp345)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBaraIdxIJ(idx int64) (i int32,j int32,err error) {
  if _tmp346 := C.MSK_getbaraidxij(self.ptr(),C.MSKint32t(idx),&i,&j); _tmp346 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp346)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBaraIdxInfo(idx int64) (num int64,err error) {
  if _tmp347 := C.MSK_getbaraidxinfo(self.ptr(),C.MSKint32t(idx),&num); _tmp347 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp347)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBaraSparsity() (numnz int64,idxij []int64,err error) {
  var _tmp348 int64
  if _tmp349 := C.MSK_getnumbaranz(self.ptr(),addr(_tmp348)); _tmp349 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp349)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp348)
  var _tmp350 *int64
  idxij := make([]int64,maxnumnz)
  if len(idxij) > 0 { _tmp350 = (*int64)(&n[0]) }
  if _tmp351 := C.MSK_getbarasparsity(self.ptr(),maxnumnz,&numnz,_tmp350); _tmp351 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp351)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarcBlockTriplet() (num int64,subj []int32,subk []int32,subl []int32,valjkl []float64,err error) {
  var _tmp352 int64
  if _tmp353 := C.MSK_getnumbarcblocktriplets(self.ptr(),addr(_tmp352)); _tmp353 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp353)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var maxnum int64 = int64(_tmp352)
  var _tmp354 *int32
  subj := make([]int32,maxnum)
  if len(subj) > 0 { _tmp354 = (*int32)(&n[0]) }
  var _tmp355 *int32
  subk := make([]int32,maxnum)
  if len(subk) > 0 { _tmp355 = (*int32)(&n[0]) }
  var _tmp356 *int32
  subl := make([]int32,maxnum)
  if len(subl) > 0 { _tmp356 = (*int32)(&n[0]) }
  var _tmp357 *float64
  valjkl := make([]float64,maxnum)
  if len(valjkl) > 0 { _tmp357 = (*float64)(&n[0]) }
  if _tmp358 := C.MSK_getbarcblocktriplet(self.ptr(),maxnum,&num,_tmp354,_tmp355,_tmp356,_tmp357); _tmp358 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp358)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarcIdx(idx int64) (j int32,num int64,sub []int64,weights []float64,err error) {
  var _tmp359 int64
  if _tmp360 := C.MSK_getbarcidxinfo(self.ptr(),idx,addr(_tmp359)); _tmp360 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp360)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var maxnum int64 = int64(_tmp359)
  var _tmp361 *int64
  sub := make([]int64,maxnum)
  if len(sub) > 0 { _tmp361 = (*int64)(&n[0]) }
  var _tmp362 *float64
  weights := make([]float64,maxnum)
  if len(weights) > 0 { _tmp362 = (*float64)(&n[0]) }
  if _tmp363 := C.MSK_getbarcidx(self.ptr(),C.MSKint32t(idx),maxnum,&j,&num,_tmp361,_tmp362); _tmp363 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp363)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarcIdxInfo(idx int64) (num int64,err error) {
  if _tmp364 := C.MSK_getbarcidxinfo(self.ptr(),C.MSKint32t(idx),&num); _tmp364 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp364)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarcIdxJ(idx int64) (j int32,err error) {
  if _tmp365 := C.MSK_getbarcidxj(self.ptr(),C.MSKint32t(idx),&j); _tmp365 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp365)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarcSparsity() (numnz int64,idxj []int64,err error) {
  var _tmp366 int64
  if _tmp367 := C.MSK_getnumbarcnz(self.ptr(),addr(_tmp366)); _tmp367 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp367)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var maxnumnz int64 = int64(_tmp366)
  var _tmp368 *int64
  idxj := make([]int64,maxnumnz)
  if len(idxj) > 0 { _tmp368 = (*int64)(&n[0]) }
  if _tmp369 := C.MSK_getbarcsparsity(self.ptr(),maxnumnz,&numnz,_tmp368); _tmp369 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp369)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarsJ(whichsol Soltype,j int32) (barsj []float64,err error) {
  var _tmp372 *float64
  var _tmp370 int64
  if _tmp371 := C.MSK_getlenbarvarj(self.ptr(),j,addr(_tmp370)); _tmp371 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp371)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  barsj := make([]float64,_tmp370)
  if len(barsj) > 0 { _tmp372 = (*float64)(&n[0]) }
  if _tmp373 := C.MSK_getbarsj(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(j),_tmp372); _tmp373 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp373)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarsSlice(whichsol Soltype,first int32,last int32,slicesize int64) (barsslice []float64,err error) {
  var _tmp374 *float64
  barsslice := make([]float64,slicesize)
  if len(barsslice) > 0 { _tmp374 = (*float64)(&n[0]) }
  if _tmp375 := C.MSK_getbarsslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),C.MSKint32t(slicesize),_tmp374); _tmp375 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp375)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarvarName(i int32) (name string,err error) {
  var _tmp376 int32
  if _tmp377 := C.MSK_getbarvarnamelen(self.ptr(),i,addr(_tmp376)); _tmp377 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp377)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp376))
  _tmp378 := make([]byte,sizename)
  if _tmp379 := C.MSK_getbarvarname(self.ptr(),C.MSKint32t(i),sizename,C.CString(&tmpvar1[0])); _tmp379 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp379)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  var name string
  if p := strings.IndexByte(_tmp378,byte(0)); p < 0 {
    name = string(_tmp378)
  } else {
    name = string(_tmp378[:p])
  }
  return
}
func (self *Task) GetBarvarNameIndex(somename string) (asgn int32,index int32,err error) {
  _tmp380 := C.CString(somename)
  if _tmp381 := C.MSK_getbarvarnameindex(self.ptr(),_tmp380,&asgn,&index); _tmp381 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp381)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarvarNameLen(i int32) (len int32,err error) {
  if _tmp382 := C.MSK_getbarvarnamelen(self.ptr(),C.MSKint32t(i),&len); _tmp382 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp382)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarxJ(whichsol Soltype,j int32) (barxj []float64,err error) {
  var _tmp385 *float64
  var _tmp383 int64
  if _tmp384 := C.MSK_getlenbarvarj(self.ptr(),j,addr(_tmp383)); _tmp384 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp384)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  barxj := make([]float64,_tmp383)
  if len(barxj) > 0 { _tmp385 = (*float64)(&n[0]) }
  if _tmp386 := C.MSK_getbarxj(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(j),_tmp385); _tmp386 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp386)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetBarxSlice(whichsol Soltype,first int32,last int32,slicesize int64) (barxslice []float64,err error) {
  var _tmp387 *float64
  barxslice := make([]float64,slicesize)
  if len(barxslice) > 0 { _tmp387 = (*float64)(&n[0]) }
  if _tmp388 := C.MSK_getbarxslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),C.MSKint32t(slicesize),_tmp387); _tmp388 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp388)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetC() (c []float64,err error) {
  var _tmp391 *float64
  var _tmp389 int32
  if _tmp390 := C.MSK_getnumvar(self.ptr(),addr(_tmp389)); _tmp390 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp390)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  c := make([]float64,_tmp389)
  if len(c) > 0 { _tmp391 = (*float64)(&n[0]) }
  if _tmp392 := C.MSK_getc(self.ptr(),_tmp391); _tmp392 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp392)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetCfix() (cfix float64,err error) {
  if _tmp393 := C.MSK_getcfix(self.ptr(),&cfix); _tmp393 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp393)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetCJ(j int32) (cj float64,err error) {
  if _tmp394 := C.MSK_getcj(self.ptr(),C.MSKint32t(j),&cj); _tmp394 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp394)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetCList(subj []int32) (c []float64,err error) {
  _tmp395 := len(subj)
  var num int32 = int32(_tmp395)
  var _tmp396 *int32
  if subj != nil { _tmp396 = (*int32)(&subj[0]) }
  var _tmp397 *float64
  c := make([]float64,num)
  if len(c) > 0 { _tmp397 = (*float64)(&n[0]) }
  if _tmp398 := C.MSK_getclist(self.ptr(),C.MSKint32t(num),_tmp396,_tmp397); _tmp398 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp398)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetConBound(i int32) (bk Boundkey,bl float64,bu float64,err error) {
  var _tmp399 C.MSKboundkeye
  if _tmp400 := C.MSK_getconbound(self.ptr(),C.MSKint32t(i),&_tmp399,&bl,&bu); _tmp400 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp400)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  bk = Boundkey(_tmp399)
  return
}
func (self *Task) GetConBoundSlice(first int32,last int32) (bk []Boundkey,bl []float64,bu []float64,err error) {
  var _tmp401 *Boundkey
  bk := make([]Boundkey,(last - first))
  if len(bk) > 0 { _tmp401 = (*Boundkey)(&n[0]) }
  var _tmp402 *float64
  bl := make([]float64,(last - first))
  if len(bl) > 0 { _tmp402 = (*float64)(&n[0]) }
  var _tmp403 *float64
  bu := make([]float64,(last - first))
  if len(bu) > 0 { _tmp403 = (*float64)(&n[0]) }
  if _tmp404 := C.MSK_getconboundslice(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),_tmp401,_tmp402,_tmp403); _tmp404 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp404)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetCone(k int32) (ct Conetype,conepar float64,nummem int32,submem []int32,err error) {
  var _tmp405 C.MSKconetypee
  var _tmp410 *int32
  var _tmp406 conetype
  var _tmp407 float64
  var _tmp408 int32
  if _tmp409 := C.MSK_getconeinfo(self.ptr(),k,addr(_tmp406),addr(_tmp407),addr(_tmp408)); _tmp409 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp409)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  submem := make([]int32,_tmp408)
  if len(submem) > 0 { _tmp410 = (*int32)(&n[0]) }
  if _tmp411 := C.MSK_getcone(self.ptr(),C.MSKint32t(k),&_tmp405,&conepar,&nummem,_tmp410); _tmp411 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp411)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  ct = Conetype(_tmp405)
  return
}
func (self *Task) GetConeInfo(k int32) (ct Conetype,conepar float64,nummem int32,err error) {
  var _tmp412 C.MSKconetypee
  if _tmp413 := C.MSK_getconeinfo(self.ptr(),C.MSKint32t(k),&_tmp412,&conepar,&nummem); _tmp413 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp413)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  ct = Conetype(_tmp412)
  return
}
func (self *Task) GetConeName(i int32) (name string,err error) {
  var _tmp414 int32
  if _tmp415 := C.MSK_getconenamelen(self.ptr(),i,addr(_tmp414)); _tmp415 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp415)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp414))
  _tmp416 := make([]byte,sizename)
  if _tmp417 := C.MSK_getconename(self.ptr(),C.MSKint32t(i),sizename,C.CString(&tmpvar1[0])); _tmp417 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp417)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  var name string
  if p := strings.IndexByte(_tmp416,byte(0)); p < 0 {
    name = string(_tmp416)
  } else {
    name = string(_tmp416[:p])
  }
  return
}
func (self *Task) GetConeNameIndex(somename string) (asgn int32,index int32,err error) {
  _tmp418 := C.CString(somename)
  if _tmp419 := C.MSK_getconenameindex(self.ptr(),_tmp418,&asgn,&index); _tmp419 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp419)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetConeNameLen(i int32) (len int32,err error) {
  if _tmp420 := C.MSK_getconenamelen(self.ptr(),C.MSKint32t(i),&len); _tmp420 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp420)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetConName(i int32) (name string,err error) {
  var _tmp421 int32
  if _tmp422 := C.MSK_getconnamelen(self.ptr(),i,addr(_tmp421)); _tmp422 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp422)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp421))
  _tmp423 := make([]byte,sizename)
  if _tmp424 := C.MSK_getconname(self.ptr(),C.MSKint32t(i),sizename,C.CString(&tmpvar1[0])); _tmp424 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp424)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  var name string
  if p := strings.IndexByte(_tmp423,byte(0)); p < 0 {
    name = string(_tmp423)
  } else {
    name = string(_tmp423[:p])
  }
  return
}
func (self *Task) GetConNameIndex(somename string) (asgn int32,index int32,err error) {
  _tmp425 := C.CString(somename)
  if _tmp426 := C.MSK_getconnameindex(self.ptr(),_tmp425,&asgn,&index); _tmp426 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp426)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetConNameLen(i int32) (len int32,err error) {
  if _tmp427 := C.MSK_getconnamelen(self.ptr(),C.MSKint32t(i),&len); _tmp427 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp427)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetCSlice(first int32,last int32) (c []float64,err error) {
  var _tmp428 *float64
  c := make([]float64,(last - first))
  if len(c) > 0 { _tmp428 = (*float64)(&n[0]) }
  if _tmp429 := C.MSK_getcslice(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),_tmp428); _tmp429 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp429)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDimBarvarJ(j int32) (dimbarvarj int32,err error) {
  if _tmp430 := C.MSK_getdimbarvarj(self.ptr(),C.MSKint32t(j),&dimbarvarj); _tmp430 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp430)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcAfeIdxList(djcidx int64) (afeidxlist []int64,err error) {
  var _tmp433 *int64
  var _tmp431 int64
  if _tmp432 := C.MSK_getdjcnumafe(self.ptr(),djcidx,addr(_tmp431)); _tmp432 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp432)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  afeidxlist := make([]int64,_tmp431)
  if len(afeidxlist) > 0 { _tmp433 = (*int64)(&n[0]) }
  if _tmp434 := C.MSK_getdjcafeidxlist(self.ptr(),C.MSKint32t(djcidx),_tmp433); _tmp434 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp434)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcB(djcidx int64) (b []float64,err error) {
  var _tmp437 *float64
  var _tmp435 int64
  if _tmp436 := C.MSK_getdjcnumafe(self.ptr(),djcidx,addr(_tmp435)); _tmp436 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp436)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  b := make([]float64,_tmp435)
  if len(b) > 0 { _tmp437 = (*float64)(&n[0]) }
  if _tmp438 := C.MSK_getdjcb(self.ptr(),C.MSKint32t(djcidx),_tmp437); _tmp438 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp438)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcDomainIdxList(djcidx int64) (domidxlist []int64,err error) {
  var _tmp441 *int64
  var _tmp439 int64
  if _tmp440 := C.MSK_getdjcnumdomain(self.ptr(),djcidx,addr(_tmp439)); _tmp440 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp440)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  domidxlist := make([]int64,_tmp439)
  if len(domidxlist) > 0 { _tmp441 = (*int64)(&n[0]) }
  if _tmp442 := C.MSK_getdjcdomainidxlist(self.ptr(),C.MSKint32t(djcidx),_tmp441); _tmp442 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp442)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcName(djcidx int64) (name string,err error) {
  var _tmp443 int32
  if _tmp444 := C.MSK_getdjcnamelen(self.ptr(),djcidx,addr(_tmp443)); _tmp444 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp444)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp443))
  _tmp445 := make([]byte,sizename)
  if _tmp446 := C.MSK_getdjcname(self.ptr(),C.MSKint32t(djcidx),sizename,C.CString(&tmpvar1[0])); _tmp446 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp446)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  var name string
  if p := strings.IndexByte(_tmp445,byte(0)); p < 0 {
    name = string(_tmp445)
  } else {
    name = string(_tmp445[:p])
  }
  return
}
func (self *Task) GetDjcNameLen(djcidx int64) (len int32,err error) {
  if _tmp447 := C.MSK_getdjcnamelen(self.ptr(),C.MSKint32t(djcidx),&len); _tmp447 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp447)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcNumAfe(djcidx int64) (numafe int64,err error) {
  if _tmp448 := C.MSK_getdjcnumafe(self.ptr(),C.MSKint32t(djcidx),&numafe); _tmp448 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp448)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcNumAfeTot() (numafetot int64,err error) {
  if _tmp449 := C.MSK_getdjcnumafetot(self.ptr(),&numafetot); _tmp449 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp449)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcNumDomain(djcidx int64) (numdomain int64,err error) {
  if _tmp450 := C.MSK_getdjcnumdomain(self.ptr(),C.MSKint32t(djcidx),&numdomain); _tmp450 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp450)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcNumDomainTot() (numdomaintot int64,err error) {
  if _tmp451 := C.MSK_getdjcnumdomaintot(self.ptr(),&numdomaintot); _tmp451 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp451)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcNumTerm(djcidx int64) (numterm int64,err error) {
  if _tmp452 := C.MSK_getdjcnumterm(self.ptr(),C.MSKint32t(djcidx),&numterm); _tmp452 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp452)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcNumTermTot() (numtermtot int64,err error) {
  if _tmp453 := C.MSK_getdjcnumtermtot(self.ptr(),&numtermtot); _tmp453 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp453)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcs() (domidxlist []int64,afeidxlist []int64,b []float64,termsizelist []int64,numterms []int64,err error) {
  var _tmp456 *int64
  var _tmp454 int64
  if _tmp455 := C.MSK_getdjcnumdomaintot(self.ptr(),addr(_tmp454)); _tmp455 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp455)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  domidxlist := make([]int64,_tmp454)
  if len(domidxlist) > 0 { _tmp456 = (*int64)(&n[0]) }
  var _tmp459 *int64
  var _tmp457 int64
  if _tmp458 := C.MSK_getdjcnumafetot(self.ptr(),addr(_tmp457)); _tmp458 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp458)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  afeidxlist := make([]int64,_tmp457)
  if len(afeidxlist) > 0 { _tmp459 = (*int64)(&n[0]) }
  var _tmp462 *float64
  var _tmp460 int64
  if _tmp461 := C.MSK_getdjcnumafetot(self.ptr(),addr(_tmp460)); _tmp461 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp461)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  b := make([]float64,_tmp460)
  if len(b) > 0 { _tmp462 = (*float64)(&n[0]) }
  var _tmp465 *int64
  var _tmp463 int64
  if _tmp464 := C.MSK_getdjcnumtermtot(self.ptr(),addr(_tmp463)); _tmp464 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp464)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  termsizelist := make([]int64,_tmp463)
  if len(termsizelist) > 0 { _tmp465 = (*int64)(&n[0]) }
  var _tmp468 *int64
  var _tmp466 int64
  if _tmp467 := C.MSK_getnumdjc(self.ptr(),addr(_tmp466)); _tmp467 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp467)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  numterms := make([]int64,_tmp466)
  if len(numterms) > 0 { _tmp468 = (*int64)(&n[0]) }
  if _tmp469 := C.MSK_getdjcs(self.ptr(),_tmp456,_tmp459,_tmp462,_tmp465,_tmp468); _tmp469 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp469)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDjcTermSizeList(djcidx int64) (termsizelist []int64,err error) {
  var _tmp472 *int64
  var _tmp470 int64
  if _tmp471 := C.MSK_getdjcnumterm(self.ptr(),djcidx,addr(_tmp470)); _tmp471 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp471)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  termsizelist := make([]int64,_tmp470)
  if len(termsizelist) > 0 { _tmp472 = (*int64)(&n[0]) }
  if _tmp473 := C.MSK_getdjctermsizelist(self.ptr(),C.MSKint32t(djcidx),_tmp472); _tmp473 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp473)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDomainN(domidx int64) (n int64,err error) {
  if _tmp474 := C.MSK_getdomainn(self.ptr(),C.MSKint32t(domidx),&n); _tmp474 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp474)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDomainName(domidx int64) (name string,err error) {
  var _tmp475 int32
  if _tmp476 := C.MSK_getdomainnamelen(self.ptr(),domidx,addr(_tmp475)); _tmp476 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp476)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp475))
  _tmp477 := make([]byte,sizename)
  if _tmp478 := C.MSK_getdomainname(self.ptr(),C.MSKint32t(domidx),sizename,C.CString(&tmpvar1[0])); _tmp478 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp478)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  var name string
  if p := strings.IndexByte(_tmp477,byte(0)); p < 0 {
    name = string(_tmp477)
  } else {
    name = string(_tmp477[:p])
  }
  return
}
func (self *Task) GetDomainNameLen(domidx int64) (len int32,err error) {
  if _tmp479 := C.MSK_getdomainnamelen(self.ptr(),C.MSKint32t(domidx),&len); _tmp479 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp479)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDomainType(domidx int64) (domtype Domaintype,err error) {
  var _tmp480 C.MSKdomaintypee
  if _tmp481 := C.MSK_getdomaintype(self.ptr(),C.MSKint32t(domidx),&_tmp480); _tmp481 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp481)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  domtype = Domaintype(_tmp480)
  return
}
func (self *Task) GetDouInf(whichdinf Dinfitem) (dvalue float64,err error) {
  if _tmp482 := C.MSK_getdouinf(self.ptr(),C.MSKdinfiteme(whichdinf),&dvalue); _tmp482 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp482)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDouParam(param Dparam) (parvalue float64,err error) {
  if _tmp483 := C.MSK_getdouparam(self.ptr(),C.MSKdparame(param),&parvalue); _tmp483 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp483)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDualObj(whichsol Soltype) (dualobj float64,err error) {
  if _tmp484 := C.MSK_getdualobj(self.ptr(),C.MSKsoltypee(whichsol),&dualobj); _tmp484 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp484)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDualSolutionNorms(whichsol Soltype) (nrmy float64,nrmslc float64,nrmsuc float64,nrmslx float64,nrmsux float64,nrmsnx float64,nrmbars float64,err error) {
  if _tmp485 := C.MSK_getdualsolutionnorms(self.ptr(),C.MSKsoltypee(whichsol),&nrmy,&nrmslc,&nrmsuc,&nrmslx,&nrmsux,&nrmsnx,&nrmbars); _tmp485 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp485)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDviolAcc(whichsol Soltype,accidxlist []int64) (viol []float64,err error) {
  _tmp486 := len(accidxlist)
  var numaccidx int64 = int64(_tmp486)
  var _tmp487 *int64
  if accidxlist != nil { _tmp487 = (*int64)(&accidxlist[0]) }
  var _tmp488 *float64
  viol := make([]float64,numaccidx)
  if len(viol) > 0 { _tmp488 = (*float64)(&n[0]) }
  if _tmp489 := C.MSK_getdviolacc(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(numaccidx),_tmp487,_tmp488); _tmp489 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp489)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDviolBarvar(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp490 := len(sub)
  var num int32 = int32(_tmp490)
  var _tmp491 *int32
  if sub != nil { _tmp491 = (*int32)(&sub[0]) }
  var _tmp492 *float64
  viol := make([]float64,num)
  if len(viol) > 0 { _tmp492 = (*float64)(&n[0]) }
  if _tmp493 := C.MSK_getdviolbarvar(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(num),_tmp491,_tmp492); _tmp493 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp493)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDviolCon(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp494 := len(sub)
  var num int32 = int32(_tmp494)
  var _tmp495 *int32
  if sub != nil { _tmp495 = (*int32)(&sub[0]) }
  var _tmp496 *float64
  viol := make([]float64,num)
  if len(viol) > 0 { _tmp496 = (*float64)(&n[0]) }
  if _tmp497 := C.MSK_getdviolcon(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(num),_tmp495,_tmp496); _tmp497 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp497)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDviolCones(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp498 := len(sub)
  var num int32 = int32(_tmp498)
  var _tmp499 *int32
  if sub != nil { _tmp499 = (*int32)(&sub[0]) }
  var _tmp500 *float64
  viol := make([]float64,num)
  if len(viol) > 0 { _tmp500 = (*float64)(&n[0]) }
  if _tmp501 := C.MSK_getdviolcones(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(num),_tmp499,_tmp500); _tmp501 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp501)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetDviolVar(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp502 := len(sub)
  var num int32 = int32(_tmp502)
  var _tmp503 *int32
  if sub != nil { _tmp503 = (*int32)(&sub[0]) }
  var _tmp504 *float64
  viol := make([]float64,num)
  if len(viol) > 0 { _tmp504 = (*float64)(&n[0]) }
  if _tmp505 := C.MSK_getdviolvar(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(num),_tmp503,_tmp504); _tmp505 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp505)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetInfIndex(inftype Inftype,infname string) (infindex int32,err error) {
  _tmp506 := C.CString(infname)
  if _tmp507 := C.MSK_getinfindex(self.ptr(),C.MSKinftypee(inftype),_tmp506,&infindex); _tmp507 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp507)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetInfMax(inftype Inftype) (infmax []int32,err error) {
  var _tmp508 *int32
  infmax := make([]int32,max_str_len)
  if len(infmax) > 0 { _tmp508 = (*int32)(&n[0]) }
  if _tmp509 := C.MSK_getinfmax(self.ptr(),C.MSKinftypee(inftype),_tmp508); _tmp509 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp509)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetInfName(inftype Inftype,whichinf int32) (infname string,err error) {
  _tmp510 := make([]byte,max_str_len)
  if _tmp511 := C.MSK_getinfname(self.ptr(),C.MSKinftypee(inftype),C.MSKint32t(whichinf),C.CString(&tmpvar1[0])); _tmp511 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp511)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  var infname string
  if p := strings.IndexByte(_tmp510,byte(0)); p < 0 {
    infname = string(_tmp510)
  } else {
    infname = string(_tmp510[:p])
  }
  return
}
func (self *Task) GetIntInf(whichiinf Iinfitem) (ivalue int32,err error) {
  if _tmp512 := C.MSK_getintinf(self.ptr(),C.MSKiinfiteme(whichiinf),&ivalue); _tmp512 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp512)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetIntParam(param Iparam) (parvalue int32,err error) {
  if _tmp513 := C.MSK_getintparam(self.ptr(),C.MSKiparame(param),&parvalue); _tmp513 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp513)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetLenBarvarJ(j int32) (lenbarvarj int64,err error) {
  if _tmp514 := C.MSK_getlenbarvarj(self.ptr(),C.MSKint32t(j),&lenbarvarj); _tmp514 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp514)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetLintInf(whichliinf Liinfitem) (ivalue int64,err error) {
  if _tmp515 := C.MSK_getlintinf(self.ptr(),C.MSKliinfiteme(whichliinf),&ivalue); _tmp515 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp515)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetMaxNumANz() (maxnumanz int64,err error) {
  if _tmp516 := C.MSK_getmaxnumanz64(self.ptr(),&maxnumanz); _tmp516 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp516)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetMaxNumBarvar() (maxnumbarvar int32,err error) {
  if _tmp517 := C.MSK_getmaxnumbarvar(self.ptr(),&maxnumbarvar); _tmp517 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp517)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetMaxNumCon() (maxnumcon int32,err error) {
  if _tmp518 := C.MSK_getmaxnumcon(self.ptr(),&maxnumcon); _tmp518 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp518)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetMaxNumCone() (maxnumcone int32,err error) {
  if _tmp519 := C.MSK_getmaxnumcone(self.ptr(),&maxnumcone); _tmp519 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp519)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetMaxNumQNz() (maxnumqnz int64,err error) {
  if _tmp520 := C.MSK_getmaxnumqnz64(self.ptr(),&maxnumqnz); _tmp520 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp520)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetMaxNumVar() (maxnumvar int32,err error) {
  if _tmp521 := C.MSK_getmaxnumvar(self.ptr(),&maxnumvar); _tmp521 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp521)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetMemUsage() (meminuse int64,maxmemuse int64,err error) {
  if _tmp522 := C.MSK_getmemusagetask(self.ptr(),&meminuse,&maxmemuse); _tmp522 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp522)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumAcc() (num int64,err error) {
  if _tmp523 := C.MSK_getnumacc(self.ptr(),&num); _tmp523 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp523)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumAfe() (numafe int64,err error) {
  if _tmp524 := C.MSK_getnumafe(self.ptr(),&numafe); _tmp524 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp524)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumANz() (numanz int32,err error) {
  if _tmp525 := C.MSK_getnumanz(self.ptr(),&numanz); _tmp525 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp525)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumANz64() (numanz int64,err error) {
  if _tmp526 := C.MSK_getnumanz64(self.ptr(),&numanz); _tmp526 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp526)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumBaraBlockTriplets() (num int64,err error) {
  if _tmp527 := C.MSK_getnumbarablocktriplets(self.ptr(),&num); _tmp527 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp527)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumBaraNz() (nz int64,err error) {
  if _tmp528 := C.MSK_getnumbaranz(self.ptr(),&nz); _tmp528 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp528)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumBarcBlockTriplets() (num int64,err error) {
  if _tmp529 := C.MSK_getnumbarcblocktriplets(self.ptr(),&num); _tmp529 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp529)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumBarcNz() (nz int64,err error) {
  if _tmp530 := C.MSK_getnumbarcnz(self.ptr(),&nz); _tmp530 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp530)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumBarvar() (numbarvar int32,err error) {
  if _tmp531 := C.MSK_getnumbarvar(self.ptr(),&numbarvar); _tmp531 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp531)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumCon() (numcon int32,err error) {
  if _tmp532 := C.MSK_getnumcon(self.ptr(),&numcon); _tmp532 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp532)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumCone() (numcone int32,err error) {
  if _tmp533 := C.MSK_getnumcone(self.ptr(),&numcone); _tmp533 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp533)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumConeMem(k int32) (nummem int32,err error) {
  if _tmp534 := C.MSK_getnumconemem(self.ptr(),C.MSKint32t(k),&nummem); _tmp534 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp534)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumDjc() (num int64,err error) {
  if _tmp535 := C.MSK_getnumdjc(self.ptr(),&num); _tmp535 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp535)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumDomain() (numdomain int64,err error) {
  if _tmp536 := C.MSK_getnumdomain(self.ptr(),&numdomain); _tmp536 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp536)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumIntVar() (numintvar int32,err error) {
  if _tmp537 := C.MSK_getnumintvar(self.ptr(),&numintvar); _tmp537 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp537)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumParam(partype Parametertype) (numparam int32,err error) {
  if _tmp538 := C.MSK_getnumparam(self.ptr(),C.MSKparametertypee(partype),&numparam); _tmp538 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp538)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumQConKNz(k int32) (numqcnz int64,err error) {
  if _tmp539 := C.MSK_getnumqconknz64(self.ptr(),C.MSKint32t(k),&numqcnz); _tmp539 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp539)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumQObjNz() (numqonz int64,err error) {
  if _tmp540 := C.MSK_getnumqobjnz64(self.ptr(),&numqonz); _tmp540 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp540)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumSymMat() (num int64,err error) {
  if _tmp541 := C.MSK_getnumsymmat(self.ptr(),&num); _tmp541 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp541)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetNumVar() (numvar int32,err error) {
  if _tmp542 := C.MSK_getnumvar(self.ptr(),&numvar); _tmp542 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp542)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetObjName() (objname string,err error) {
  var _tmp543 int32
  if _tmp544 := C.MSK_getobjnamelen(self.ptr(),addr(_tmp543)); _tmp544 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp544)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var sizeobjname int32 = int32((1 + _tmp543))
  _tmp545 := make([]byte,sizeobjname)
  if _tmp546 := C.MSK_getobjname(self.ptr(),sizeobjname,C.CString(&tmpvar1[0])); _tmp546 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp546)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  var objname string
  if p := strings.IndexByte(_tmp545,byte(0)); p < 0 {
    objname = string(_tmp545)
  } else {
    objname = string(_tmp545[:p])
  }
  return
}
func (self *Task) GetObjNameLen() (len int32,err error) {
  if _tmp547 := C.MSK_getobjnamelen(self.ptr(),&len); _tmp547 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp547)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetObjSense() (sense Objsense,err error) {
  var _tmp548 C.MSKobjsensee
  if _tmp549 := C.MSK_getobjsense(self.ptr(),&_tmp548); _tmp549 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp549)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  sense = Objsense(_tmp548)
  return
}
func (self *Task) GetParamMax(partype Parametertype) (parammax int32,err error) {
  if _tmp550 := C.MSK_getparammax(self.ptr(),C.MSKparametertypee(partype),&parammax); _tmp550 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp550)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetParamName(partype Parametertype,param int32) (parname string,err error) {
  _tmp551 := make([]byte,max_str_len)
  if _tmp552 := C.MSK_getparamname(self.ptr(),C.MSKparametertypee(partype),C.MSKint32t(param),C.CString(&tmpvar1[0])); _tmp552 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp552)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  var parname string
  if p := strings.IndexByte(_tmp551,byte(0)); p < 0 {
    parname = string(_tmp551)
  } else {
    parname = string(_tmp551[:p])
  }
  return
}
func (self *Task) GetPowerDomainAlpha(domidx int64) (alpha []float64,err error) {
  var _tmp556 *float64
  var _tmp553 int64
  var _tmp554 int64
  if _tmp555 := C.MSK_getpowerdomaininfo(self.ptr(),domidx,addr(_tmp553),addr(_tmp554)); _tmp555 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp555)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  alpha := make([]float64,_tmp554)
  if len(alpha) > 0 { _tmp556 = (*float64)(&n[0]) }
  if _tmp557 := C.MSK_getpowerdomainalpha(self.ptr(),C.MSKint32t(domidx),_tmp556); _tmp557 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp557)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetPowerDomainInfo(domidx int64) (n int64,nleft int64,err error) {
  if _tmp558 := C.MSK_getpowerdomaininfo(self.ptr(),C.MSKint32t(domidx),&n,&nleft); _tmp558 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp558)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetPrimalObj(whichsol Soltype) (primalobj float64,err error) {
  if _tmp559 := C.MSK_getprimalobj(self.ptr(),C.MSKsoltypee(whichsol),&primalobj); _tmp559 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp559)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetPrimalSolutionNorms(whichsol Soltype) (nrmxc float64,nrmxx float64,nrmbarx float64,err error) {
  if _tmp560 := C.MSK_getprimalsolutionnorms(self.ptr(),C.MSKsoltypee(whichsol),&nrmxc,&nrmxx,&nrmbarx); _tmp560 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp560)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetProbType() (probtype Problemtype,err error) {
  var _tmp561 C.MSKproblemtypee
  if _tmp562 := C.MSK_getprobtype(self.ptr(),&_tmp561); _tmp562 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp562)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  probtype = Problemtype(_tmp561)
  return
}
func (self *Task) GetProSta(whichsol Soltype) (problemsta Prosta,err error) {
  var _tmp563 C.MSKprostae
  if _tmp564 := C.MSK_getprosta(self.ptr(),C.MSKsoltypee(whichsol),&_tmp563); _tmp564 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp564)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  problemsta = Prosta(_tmp563)
  return
}
func (self *Task) GetPviolAcc(whichsol Soltype,accidxlist []int64) (viol []float64,err error) {
  _tmp565 := len(accidxlist)
  var numaccidx int64 = int64(_tmp565)
  var _tmp566 *int64
  if accidxlist != nil { _tmp566 = (*int64)(&accidxlist[0]) }
  var _tmp567 *float64
  viol := make([]float64,numaccidx)
  if len(viol) > 0 { _tmp567 = (*float64)(&n[0]) }
  if _tmp568 := C.MSK_getpviolacc(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(numaccidx),_tmp566,_tmp567); _tmp568 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp568)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetPviolBarvar(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp569 := len(sub)
  var num int32 = int32(_tmp569)
  var _tmp570 *int32
  if sub != nil { _tmp570 = (*int32)(&sub[0]) }
  var _tmp571 *float64
  viol := make([]float64,num)
  if len(viol) > 0 { _tmp571 = (*float64)(&n[0]) }
  if _tmp572 := C.MSK_getpviolbarvar(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(num),_tmp570,_tmp571); _tmp572 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp572)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetPviolCon(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp573 := len(sub)
  var num int32 = int32(_tmp573)
  var _tmp574 *int32
  if sub != nil { _tmp574 = (*int32)(&sub[0]) }
  var _tmp575 *float64
  viol := make([]float64,num)
  if len(viol) > 0 { _tmp575 = (*float64)(&n[0]) }
  if _tmp576 := C.MSK_getpviolcon(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(num),_tmp574,_tmp575); _tmp576 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp576)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetPviolCones(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp577 := len(sub)
  var num int32 = int32(_tmp577)
  var _tmp578 *int32
  if sub != nil { _tmp578 = (*int32)(&sub[0]) }
  var _tmp579 *float64
  viol := make([]float64,num)
  if len(viol) > 0 { _tmp579 = (*float64)(&n[0]) }
  if _tmp580 := C.MSK_getpviolcones(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(num),_tmp578,_tmp579); _tmp580 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp580)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetPviolDjc(whichsol Soltype,djcidxlist []int64) (viol []float64,err error) {
  _tmp581 := len(djcidxlist)
  var numdjcidx int64 = int64(_tmp581)
  var _tmp582 *int64
  if djcidxlist != nil { _tmp582 = (*int64)(&djcidxlist[0]) }
  var _tmp583 *float64
  viol := make([]float64,numdjcidx)
  if len(viol) > 0 { _tmp583 = (*float64)(&n[0]) }
  if _tmp584 := C.MSK_getpvioldjc(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(numdjcidx),_tmp582,_tmp583); _tmp584 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp584)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetPviolVar(whichsol Soltype,sub []int32) (viol []float64,err error) {
  _tmp585 := len(sub)
  var num int32 = int32(_tmp585)
  var _tmp586 *int32
  if sub != nil { _tmp586 = (*int32)(&sub[0]) }
  var _tmp587 *float64
  viol := make([]float64,num)
  if len(viol) > 0 { _tmp587 = (*float64)(&n[0]) }
  if _tmp588 := C.MSK_getpviolvar(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(num),_tmp586,_tmp587); _tmp588 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp588)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetQConK(k int32) (numqcnz int64,qcsubi []int32,qcsubj []int32,qcval []float64,err error) {
  var _tmp589 int64
  if _tmp590 := C.MSK_getnumqconknz64(self.ptr(),k,addr(_tmp589)); _tmp590 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp590)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var maxnumqcnz int64 = int64(_tmp589)
  var _tmp593 *int32
  var _tmp591 int64
  if _tmp592 := C.MSK_getnumqconknz64(self.ptr(),k,addr(_tmp591)); _tmp592 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp592)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  qcsubi := make([]int32,_tmp591)
  if len(qcsubi) > 0 { _tmp593 = (*int32)(&n[0]) }
  var _tmp596 *int32
  var _tmp594 int64
  if _tmp595 := C.MSK_getnumqconknz64(self.ptr(),k,addr(_tmp594)); _tmp595 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp595)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  qcsubj := make([]int32,_tmp594)
  if len(qcsubj) > 0 { _tmp596 = (*int32)(&n[0]) }
  var _tmp599 *float64
  var _tmp597 int64
  if _tmp598 := C.MSK_getnumqconknz64(self.ptr(),k,addr(_tmp597)); _tmp598 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp598)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  qcval := make([]float64,_tmp597)
  if len(qcval) > 0 { _tmp599 = (*float64)(&n[0]) }
  if _tmp600 := C.MSK_getqconk64(self.ptr(),C.MSKint32t(k),maxnumqcnz,&numqcnz,_tmp593,_tmp596,_tmp599); _tmp600 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp600)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetQObj() (numqonz int64,qosubi []int32,qosubj []int32,qoval []float64,err error) {
  var _tmp601 int64
  if _tmp602 := C.MSK_getnumqobjnz64(self.ptr(),addr(_tmp601)); _tmp602 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp602)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var maxnumqonz int64 = int64(_tmp601)
  var _tmp603 *int32
  qosubi := make([]int32,maxnumqonz)
  if len(qosubi) > 0 { _tmp603 = (*int32)(&n[0]) }
  var _tmp604 *int32
  qosubj := make([]int32,maxnumqonz)
  if len(qosubj) > 0 { _tmp604 = (*int32)(&n[0]) }
  var _tmp605 *float64
  qoval := make([]float64,maxnumqonz)
  if len(qoval) > 0 { _tmp605 = (*float64)(&n[0]) }
  if _tmp606 := C.MSK_getqobj64(self.ptr(),maxnumqonz,&numqonz,_tmp603,_tmp604,_tmp605); _tmp606 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp606)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetQObjIJ(i int32,j int32) (qoij float64,err error) {
  if _tmp607 := C.MSK_getqobjij(self.ptr(),C.MSKint32t(i),C.MSKint32t(j),&qoij); _tmp607 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp607)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetReducedCosts(whichsol Soltype,first int32,last int32) (redcosts []float64,err error) {
  var _tmp608 *float64
  redcosts := make([]float64,(last - first))
  if len(redcosts) > 0 { _tmp608 = (*float64)(&n[0]) }
  if _tmp609 := C.MSK_getreducedcosts(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp608); _tmp609 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp609)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSkc(whichsol Soltype) (skc []Stakey,err error) {
  var _tmp612 *Stakey
  var _tmp610 int32
  if _tmp611 := C.MSK_getnumcon(self.ptr(),addr(_tmp610)); _tmp611 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp611)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  skc := make([]Stakey,_tmp610)
  if len(skc) > 0 { _tmp612 = (*Stakey)(&n[0]) }
  if _tmp613 := C.MSK_getskc(self.ptr(),C.MSKsoltypee(whichsol),_tmp612); _tmp613 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp613)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSkcSlice(whichsol Soltype,first int32,last int32) (skc []Stakey,err error) {
  var _tmp614 *Stakey
  skc := make([]Stakey,(last - first))
  if len(skc) > 0 { _tmp614 = (*Stakey)(&n[0]) }
  if _tmp615 := C.MSK_getskcslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp614); _tmp615 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp615)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSkn(whichsol Soltype) (skn []Stakey,err error) {
  var _tmp618 *Stakey
  var _tmp616 int32
  if _tmp617 := C.MSK_getnumcone(self.ptr(),addr(_tmp616)); _tmp617 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp617)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  skn := make([]Stakey,_tmp616)
  if len(skn) > 0 { _tmp618 = (*Stakey)(&n[0]) }
  if _tmp619 := C.MSK_getskn(self.ptr(),C.MSKsoltypee(whichsol),_tmp618); _tmp619 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp619)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSkx(whichsol Soltype) (skx []Stakey,err error) {
  var _tmp622 *Stakey
  var _tmp620 int32
  if _tmp621 := C.MSK_getnumvar(self.ptr(),addr(_tmp620)); _tmp621 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp621)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  skx := make([]Stakey,_tmp620)
  if len(skx) > 0 { _tmp622 = (*Stakey)(&n[0]) }
  if _tmp623 := C.MSK_getskx(self.ptr(),C.MSKsoltypee(whichsol),_tmp622); _tmp623 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp623)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSkxSlice(whichsol Soltype,first int32,last int32) (skx []Stakey,err error) {
  var _tmp624 *Stakey
  skx := make([]Stakey,(last - first))
  if len(skx) > 0 { _tmp624 = (*Stakey)(&n[0]) }
  if _tmp625 := C.MSK_getskxslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp624); _tmp625 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp625)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSlc(whichsol Soltype) (slc []float64,err error) {
  var _tmp628 *float64
  var _tmp626 int32
  if _tmp627 := C.MSK_getnumcon(self.ptr(),addr(_tmp626)); _tmp627 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp627)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  slc := make([]float64,_tmp626)
  if len(slc) > 0 { _tmp628 = (*float64)(&n[0]) }
  if _tmp629 := C.MSK_getslc(self.ptr(),C.MSKsoltypee(whichsol),_tmp628); _tmp629 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp629)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSlcSlice(whichsol Soltype,first int32,last int32) (slc []float64,err error) {
  var _tmp630 *float64
  slc := make([]float64,(last - first))
  if len(slc) > 0 { _tmp630 = (*float64)(&n[0]) }
  if _tmp631 := C.MSK_getslcslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp630); _tmp631 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp631)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSlx(whichsol Soltype) (slx []float64,err error) {
  var _tmp634 *float64
  var _tmp632 int32
  if _tmp633 := C.MSK_getnumvar(self.ptr(),addr(_tmp632)); _tmp633 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp633)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  slx := make([]float64,_tmp632)
  if len(slx) > 0 { _tmp634 = (*float64)(&n[0]) }
  if _tmp635 := C.MSK_getslx(self.ptr(),C.MSKsoltypee(whichsol),_tmp634); _tmp635 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp635)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSlxSlice(whichsol Soltype,first int32,last int32) (slx []float64,err error) {
  var _tmp636 *float64
  slx := make([]float64,(last - first))
  if len(slx) > 0 { _tmp636 = (*float64)(&n[0]) }
  if _tmp637 := C.MSK_getslxslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp636); _tmp637 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp637)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSnx(whichsol Soltype) (snx []float64,err error) {
  var _tmp640 *float64
  var _tmp638 int32
  if _tmp639 := C.MSK_getnumvar(self.ptr(),addr(_tmp638)); _tmp639 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp639)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  snx := make([]float64,_tmp638)
  if len(snx) > 0 { _tmp640 = (*float64)(&n[0]) }
  if _tmp641 := C.MSK_getsnx(self.ptr(),C.MSKsoltypee(whichsol),_tmp640); _tmp641 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp641)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSnxSlice(whichsol Soltype,first int32,last int32) (snx []float64,err error) {
  var _tmp642 *float64
  snx := make([]float64,(last - first))
  if len(snx) > 0 { _tmp642 = (*float64)(&n[0]) }
  if _tmp643 := C.MSK_getsnxslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp642); _tmp643 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp643)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSolSta(whichsol Soltype) (solutionsta Solsta,err error) {
  var _tmp644 C.MSKsolstae
  if _tmp645 := C.MSK_getsolsta(self.ptr(),C.MSKsoltypee(whichsol),&_tmp644); _tmp645 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp645)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  solutionsta = Solsta(_tmp644)
  return
}
func (self *Task) GetSolution(whichsol Soltype) (problemsta Prosta,solutionsta Solsta,skc []Stakey,skx []Stakey,skn []Stakey,xc []float64,xx []float64,y []float64,slc []float64,suc []float64,slx []float64,sux []float64,snx []float64,err error) {
  var _tmp646 C.MSKprostae
  var _tmp647 C.MSKsolstae
  var _tmp650 *Stakey
  var _tmp648 int32
  if _tmp649 := C.MSK_getnumcon(self.ptr(),addr(_tmp648)); _tmp649 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp649)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  skc := make([]Stakey,_tmp648)
  if len(skc) > 0 { _tmp650 = (*Stakey)(&n[0]) }
  var _tmp653 *Stakey
  var _tmp651 int32
  if _tmp652 := C.MSK_getnumvar(self.ptr(),addr(_tmp651)); _tmp652 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp652)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  skx := make([]Stakey,_tmp651)
  if len(skx) > 0 { _tmp653 = (*Stakey)(&n[0]) }
  var _tmp656 *Stakey
  var _tmp654 int32
  if _tmp655 := C.MSK_getnumcone(self.ptr(),addr(_tmp654)); _tmp655 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp655)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  skn := make([]Stakey,_tmp654)
  if len(skn) > 0 { _tmp656 = (*Stakey)(&n[0]) }
  var _tmp659 *float64
  var _tmp657 int32
  if _tmp658 := C.MSK_getnumcon(self.ptr(),addr(_tmp657)); _tmp658 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp658)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  xc := make([]float64,_tmp657)
  if len(xc) > 0 { _tmp659 = (*float64)(&n[0]) }
  var _tmp662 *float64
  var _tmp660 int32
  if _tmp661 := C.MSK_getnumvar(self.ptr(),addr(_tmp660)); _tmp661 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp661)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  xx := make([]float64,_tmp660)
  if len(xx) > 0 { _tmp662 = (*float64)(&n[0]) }
  var _tmp665 *float64
  var _tmp663 int32
  if _tmp664 := C.MSK_getnumcon(self.ptr(),addr(_tmp663)); _tmp664 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp664)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  y := make([]float64,_tmp663)
  if len(y) > 0 { _tmp665 = (*float64)(&n[0]) }
  var _tmp668 *float64
  var _tmp666 int32
  if _tmp667 := C.MSK_getnumcon(self.ptr(),addr(_tmp666)); _tmp667 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp667)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  slc := make([]float64,_tmp666)
  if len(slc) > 0 { _tmp668 = (*float64)(&n[0]) }
  var _tmp671 *float64
  var _tmp669 int32
  if _tmp670 := C.MSK_getnumcon(self.ptr(),addr(_tmp669)); _tmp670 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp670)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  suc := make([]float64,_tmp669)
  if len(suc) > 0 { _tmp671 = (*float64)(&n[0]) }
  var _tmp674 *float64
  var _tmp672 int32
  if _tmp673 := C.MSK_getnumvar(self.ptr(),addr(_tmp672)); _tmp673 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp673)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  slx := make([]float64,_tmp672)
  if len(slx) > 0 { _tmp674 = (*float64)(&n[0]) }
  var _tmp677 *float64
  var _tmp675 int32
  if _tmp676 := C.MSK_getnumvar(self.ptr(),addr(_tmp675)); _tmp676 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp676)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  sux := make([]float64,_tmp675)
  if len(sux) > 0 { _tmp677 = (*float64)(&n[0]) }
  var _tmp680 *float64
  var _tmp678 int32
  if _tmp679 := C.MSK_getnumvar(self.ptr(),addr(_tmp678)); _tmp679 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp679)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  snx := make([]float64,_tmp678)
  if len(snx) > 0 { _tmp680 = (*float64)(&n[0]) }
  if _tmp681 := C.MSK_getsolution(self.ptr(),C.MSKsoltypee(whichsol),&_tmp646,&_tmp647,_tmp650,_tmp653,_tmp656,_tmp659,_tmp662,_tmp665,_tmp668,_tmp671,_tmp674,_tmp677,_tmp680); _tmp681 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp681)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  problemsta = Prosta(_tmp646)
  solutionsta = Solsta(_tmp647)
  return
}
func (self *Task) GetSolutionInfo(whichsol Soltype) (pobj float64,pviolcon float64,pviolvar float64,pviolbarvar float64,pviolcone float64,pviolitg float64,dobj float64,dviolcon float64,dviolvar float64,dviolbarvar float64,dviolcone float64,err error) {
  if _tmp682 := C.MSK_getsolutioninfo(self.ptr(),C.MSKsoltypee(whichsol),&pobj,&pviolcon,&pviolvar,&pviolbarvar,&pviolcone,&pviolitg,&dobj,&dviolcon,&dviolvar,&dviolbarvar,&dviolcone); _tmp682 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp682)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSolutionInfoNew(whichsol Soltype) (pobj float64,pviolcon float64,pviolvar float64,pviolbarvar float64,pviolcone float64,pviolacc float64,pvioldjc float64,pviolitg float64,dobj float64,dviolcon float64,dviolvar float64,dviolbarvar float64,dviolcone float64,dviolacc float64,err error) {
  if _tmp683 := C.MSK_getsolutioninfonew(self.ptr(),C.MSKsoltypee(whichsol),&pobj,&pviolcon,&pviolvar,&pviolbarvar,&pviolcone,&pviolacc,&pvioldjc,&pviolitg,&dobj,&dviolcon,&dviolvar,&dviolbarvar,&dviolcone,&dviolacc); _tmp683 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp683)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSolutionNew(whichsol Soltype) (problemsta Prosta,solutionsta Solsta,skc []Stakey,skx []Stakey,skn []Stakey,xc []float64,xx []float64,y []float64,slc []float64,suc []float64,slx []float64,sux []float64,snx []float64,doty []float64,err error) {
  var _tmp684 C.MSKprostae
  var _tmp685 C.MSKsolstae
  var _tmp688 *Stakey
  var _tmp686 int32
  if _tmp687 := C.MSK_getnumcon(self.ptr(),addr(_tmp686)); _tmp687 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp687)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  skc := make([]Stakey,_tmp686)
  if len(skc) > 0 { _tmp688 = (*Stakey)(&n[0]) }
  var _tmp691 *Stakey
  var _tmp689 int32
  if _tmp690 := C.MSK_getnumvar(self.ptr(),addr(_tmp689)); _tmp690 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp690)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  skx := make([]Stakey,_tmp689)
  if len(skx) > 0 { _tmp691 = (*Stakey)(&n[0]) }
  var _tmp694 *Stakey
  var _tmp692 int32
  if _tmp693 := C.MSK_getnumcone(self.ptr(),addr(_tmp692)); _tmp693 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp693)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  skn := make([]Stakey,_tmp692)
  if len(skn) > 0 { _tmp694 = (*Stakey)(&n[0]) }
  var _tmp697 *float64
  var _tmp695 int32
  if _tmp696 := C.MSK_getnumcon(self.ptr(),addr(_tmp695)); _tmp696 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp696)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  xc := make([]float64,_tmp695)
  if len(xc) > 0 { _tmp697 = (*float64)(&n[0]) }
  var _tmp700 *float64
  var _tmp698 int32
  if _tmp699 := C.MSK_getnumvar(self.ptr(),addr(_tmp698)); _tmp699 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp699)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  xx := make([]float64,_tmp698)
  if len(xx) > 0 { _tmp700 = (*float64)(&n[0]) }
  var _tmp703 *float64
  var _tmp701 int32
  if _tmp702 := C.MSK_getnumcon(self.ptr(),addr(_tmp701)); _tmp702 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp702)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  y := make([]float64,_tmp701)
  if len(y) > 0 { _tmp703 = (*float64)(&n[0]) }
  var _tmp706 *float64
  var _tmp704 int32
  if _tmp705 := C.MSK_getnumcon(self.ptr(),addr(_tmp704)); _tmp705 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp705)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  slc := make([]float64,_tmp704)
  if len(slc) > 0 { _tmp706 = (*float64)(&n[0]) }
  var _tmp709 *float64
  var _tmp707 int32
  if _tmp708 := C.MSK_getnumcon(self.ptr(),addr(_tmp707)); _tmp708 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp708)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  suc := make([]float64,_tmp707)
  if len(suc) > 0 { _tmp709 = (*float64)(&n[0]) }
  var _tmp712 *float64
  var _tmp710 int32
  if _tmp711 := C.MSK_getnumvar(self.ptr(),addr(_tmp710)); _tmp711 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp711)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  slx := make([]float64,_tmp710)
  if len(slx) > 0 { _tmp712 = (*float64)(&n[0]) }
  var _tmp715 *float64
  var _tmp713 int32
  if _tmp714 := C.MSK_getnumvar(self.ptr(),addr(_tmp713)); _tmp714 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp714)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  sux := make([]float64,_tmp713)
  if len(sux) > 0 { _tmp715 = (*float64)(&n[0]) }
  var _tmp718 *float64
  var _tmp716 int32
  if _tmp717 := C.MSK_getnumvar(self.ptr(),addr(_tmp716)); _tmp717 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp717)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  snx := make([]float64,_tmp716)
  if len(snx) > 0 { _tmp718 = (*float64)(&n[0]) }
  var _tmp721 *float64
  var _tmp719 int64
  if _tmp720 := C.MSK_getaccntot(self.ptr(),addr(_tmp719)); _tmp720 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp720)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  doty := make([]float64,_tmp719)
  if len(doty) > 0 { _tmp721 = (*float64)(&n[0]) }
  if _tmp722 := C.MSK_getsolutionnew(self.ptr(),C.MSKsoltypee(whichsol),&_tmp684,&_tmp685,_tmp688,_tmp691,_tmp694,_tmp697,_tmp700,_tmp703,_tmp706,_tmp709,_tmp712,_tmp715,_tmp718,_tmp721); _tmp722 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp722)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  problemsta = Prosta(_tmp684)
  solutionsta = Solsta(_tmp685)
  return
}
func (self *Task) GetSolutionSlice(whichsol Soltype,solitem Solitem,first int32,last int32) (values []float64,err error) {
  var _tmp723 *float64
  values := make([]float64,(last - first))
  if len(values) > 0 { _tmp723 = (*float64)(&n[0]) }
  if _tmp724 := C.MSK_getsolutionslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKsoliteme(solitem),C.MSKint32t(first),C.MSKint32t(last),_tmp723); _tmp724 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp724)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSparseSymMat(idx int64) (subi []int32,subj []int32,valij []float64,err error) {
  var _tmp725 int32
  var _tmp726 int64
  var _tmp727 symmattype
  if _tmp728 := C.MSK_getsymmatinfo(self.ptr(),idx,addr(_tmp725),addr(_tmp726),addr(_tmp727)); _tmp728 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp728)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var maxlen int64 = int64(_tmp726)
  var _tmp729 *int32
  subi := make([]int32,maxlen)
  if len(subi) > 0 { _tmp729 = (*int32)(&n[0]) }
  var _tmp730 *int32
  subj := make([]int32,maxlen)
  if len(subj) > 0 { _tmp730 = (*int32)(&n[0]) }
  var _tmp731 *float64
  valij := make([]float64,maxlen)
  if len(valij) > 0 { _tmp731 = (*float64)(&n[0]) }
  if _tmp732 := C.MSK_getsparsesymmat(self.ptr(),C.MSKint32t(idx),maxlen,_tmp729,_tmp730,_tmp731); _tmp732 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp732)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetStrParam(param Sparam) (len int32,parvalue string,err error) {
  var _tmp733 int32
  if _tmp734 := C.MSK_getstrparamlen(self.ptr(),param,addr(_tmp733)); _tmp734 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp734)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var maxlen int32 = int32((1 + _tmp733))
  _tmp735 := make([]byte,maxlen)
  if _tmp736 := C.MSK_getstrparam(self.ptr(),C.MSKsparame(param),maxlen,&len,C.CString(&tmpvar1[0])); _tmp736 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp736)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  var parvalue string
  if p := strings.IndexByte(_tmp735,byte(0)); p < 0 {
    parvalue = string(_tmp735)
  } else {
    parvalue = string(_tmp735[:p])
  }
  return
}
func (self *Task) GetStrParamLen(param Sparam) (len int32,err error) {
  if _tmp737 := C.MSK_getstrparamlen(self.ptr(),C.MSKsparame(param),&len); _tmp737 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp737)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSuc(whichsol Soltype) (suc []float64,err error) {
  var _tmp740 *float64
  var _tmp738 int32
  if _tmp739 := C.MSK_getnumcon(self.ptr(),addr(_tmp738)); _tmp739 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp739)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  suc := make([]float64,_tmp738)
  if len(suc) > 0 { _tmp740 = (*float64)(&n[0]) }
  if _tmp741 := C.MSK_getsuc(self.ptr(),C.MSKsoltypee(whichsol),_tmp740); _tmp741 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp741)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSucSlice(whichsol Soltype,first int32,last int32) (suc []float64,err error) {
  var _tmp742 *float64
  suc := make([]float64,(last - first))
  if len(suc) > 0 { _tmp742 = (*float64)(&n[0]) }
  if _tmp743 := C.MSK_getsucslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp742); _tmp743 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp743)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSux(whichsol Soltype) (sux []float64,err error) {
  var _tmp746 *float64
  var _tmp744 int32
  if _tmp745 := C.MSK_getnumvar(self.ptr(),addr(_tmp744)); _tmp745 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp745)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  sux := make([]float64,_tmp744)
  if len(sux) > 0 { _tmp746 = (*float64)(&n[0]) }
  if _tmp747 := C.MSK_getsux(self.ptr(),C.MSKsoltypee(whichsol),_tmp746); _tmp747 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp747)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSuxSlice(whichsol Soltype,first int32,last int32) (sux []float64,err error) {
  var _tmp748 *float64
  sux := make([]float64,(last - first))
  if len(sux) > 0 { _tmp748 = (*float64)(&n[0]) }
  if _tmp749 := C.MSK_getsuxslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp748); _tmp749 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp749)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetSymMatInfo(idx int64) (dim int32,nz int64,mattype Symmattype,err error) {
  var _tmp750 C.MSKsymmattypee
  if _tmp751 := C.MSK_getsymmatinfo(self.ptr(),C.MSKint32t(idx),&dim,&nz,&_tmp750); _tmp751 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp751)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  mattype = Symmattype(_tmp750)
  return
}
func (self *Task) GetTaskName() (taskname string,err error) {
  var _tmp752 int32
  if _tmp753 := C.MSK_gettasknamelen(self.ptr(),addr(_tmp752)); _tmp753 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp753)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var sizetaskname int32 = int32((1 + _tmp752))
  _tmp754 := make([]byte,sizetaskname)
  if _tmp755 := C.MSK_gettaskname(self.ptr(),sizetaskname,C.CString(&tmpvar1[0])); _tmp755 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp755)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  var taskname string
  if p := strings.IndexByte(_tmp754,byte(0)); p < 0 {
    taskname = string(_tmp754)
  } else {
    taskname = string(_tmp754[:p])
  }
  return
}
func (self *Task) GetTaskNameLen() (len int32,err error) {
  if _tmp756 := C.MSK_gettasknamelen(self.ptr(),&len); _tmp756 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp756)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetVarBound(i int32) (bk Boundkey,bl float64,bu float64,err error) {
  var _tmp757 C.MSKboundkeye
  if _tmp758 := C.MSK_getvarbound(self.ptr(),C.MSKint32t(i),&_tmp757,&bl,&bu); _tmp758 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp758)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  bk = Boundkey(_tmp757)
  return
}
func (self *Task) GetVarBoundSlice(first int32,last int32) (bk []Boundkey,bl []float64,bu []float64,err error) {
  var _tmp759 *Boundkey
  bk := make([]Boundkey,(last - first))
  if len(bk) > 0 { _tmp759 = (*Boundkey)(&n[0]) }
  var _tmp760 *float64
  bl := make([]float64,(last - first))
  if len(bl) > 0 { _tmp760 = (*float64)(&n[0]) }
  var _tmp761 *float64
  bu := make([]float64,(last - first))
  if len(bu) > 0 { _tmp761 = (*float64)(&n[0]) }
  if _tmp762 := C.MSK_getvarboundslice(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),_tmp759,_tmp760,_tmp761); _tmp762 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp762)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetVarName(j int32) (name string,err error) {
  var _tmp763 int32
  if _tmp764 := C.MSK_getvarnamelen(self.ptr(),j,addr(_tmp763)); _tmp764 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp764)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  var sizename int32 = int32((1 + _tmp763))
  _tmp765 := make([]byte,sizename)
  if _tmp766 := C.MSK_getvarname(self.ptr(),C.MSKint32t(j),sizename,C.CString(&tmpvar1[0])); _tmp766 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp766)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  var name string
  if p := strings.IndexByte(_tmp765,byte(0)); p < 0 {
    name = string(_tmp765)
  } else {
    name = string(_tmp765[:p])
  }
  return
}
func (self *Task) GetVarNameIndex(somename string) (asgn int32,index int32,err error) {
  _tmp767 := C.CString(somename)
  if _tmp768 := C.MSK_getvarnameindex(self.ptr(),_tmp767,&asgn,&index); _tmp768 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp768)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetVarNameLen(i int32) (len int32,err error) {
  if _tmp769 := C.MSK_getvarnamelen(self.ptr(),C.MSKint32t(i),&len); _tmp769 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp769)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetVarType(j int32) (vartype Variabletype,err error) {
  var _tmp770 C.MSKvariabletypee
  if _tmp771 := C.MSK_getvartype(self.ptr(),C.MSKint32t(j),&_tmp770); _tmp771 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp771)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  vartype = Variabletype(_tmp770)
  return
}
func (self *Task) GetVarTypeList(subj []int32) (vartype []Variabletype,err error) {
  _tmp772 := len(subj)
  var num int32 = int32(_tmp772)
  var _tmp773 *int32
  if subj != nil { _tmp773 = (*int32)(&subj[0]) }
  var _tmp774 *Variabletype
  vartype := make([]Variabletype,num)
  if len(vartype) > 0 { _tmp774 = (*Variabletype)(&n[0]) }
  if _tmp775 := C.MSK_getvartypelist(self.ptr(),C.MSKint32t(num),_tmp773,_tmp774); _tmp775 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp775)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetXc(whichsol Soltype) (xc []float64,err error) {
  var _tmp778 *float64
  var _tmp776 int32
  if _tmp777 := C.MSK_getnumcon(self.ptr(),addr(_tmp776)); _tmp777 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp777)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  xc := make([]float64,_tmp776)
  if len(xc) > 0 { _tmp778 = (*float64)(&n[0]) }
  if _tmp779 := C.MSK_getxc(self.ptr(),C.MSKsoltypee(whichsol),_tmp778); _tmp779 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp779)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetXcSlice(whichsol Soltype,first int32,last int32) (xc []float64,err error) {
  var _tmp780 *float64
  xc := make([]float64,(last - first))
  if len(xc) > 0 { _tmp780 = (*float64)(&n[0]) }
  if _tmp781 := C.MSK_getxcslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp780); _tmp781 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp781)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetXx(whichsol Soltype) (xx []float64,err error) {
  var _tmp784 *float64
  var _tmp782 int32
  if _tmp783 := C.MSK_getnumvar(self.ptr(),addr(_tmp782)); _tmp783 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp783)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  xx := make([]float64,_tmp782)
  if len(xx) > 0 { _tmp784 = (*float64)(&n[0]) }
  if _tmp785 := C.MSK_getxx(self.ptr(),C.MSKsoltypee(whichsol),_tmp784); _tmp785 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp785)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetXxSlice(whichsol Soltype,first int32,last int32) (xx []float64,err error) {
  var _tmp786 *float64
  xx := make([]float64,(last - first))
  if len(xx) > 0 { _tmp786 = (*float64)(&n[0]) }
  if _tmp787 := C.MSK_getxxslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp786); _tmp787 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp787)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetY(whichsol Soltype) (y []float64,err error) {
  var _tmp790 *float64
  var _tmp788 int32
  if _tmp789 := C.MSK_getnumcon(self.ptr(),addr(_tmp788)); _tmp789 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp789)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  y := make([]float64,_tmp788)
  if len(y) > 0 { _tmp790 = (*float64)(&n[0]) }
  if _tmp791 := C.MSK_gety(self.ptr(),C.MSKsoltypee(whichsol),_tmp790); _tmp791 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp791)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) GetYSlice(whichsol Soltype,first int32,last int32) (y []float64,err error) {
  var _tmp792 *float64
  y := make([]float64,(last - first))
  if len(y) > 0 { _tmp792 = (*float64)(&n[0]) }
  if _tmp793 := C.MSK_getyslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp792); _tmp793 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp793)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) InfeasibilityReport(whichstream Streamtype,whichsol Soltype) (err error) {
  if _tmp794 := C.MSK_infeasibilityreport(self.ptr(),C.MSKstreamtypee(whichstream),C.MSKsoltypee(whichsol)); _tmp794 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp794)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) InitBasisSolve() (basis []int32,err error) {
  var _tmp797 *int32
  var _tmp795 int32
  if _tmp796 := C.MSK_getnumcon(self.ptr(),addr(_tmp795)); _tmp796 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp796)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  basis := make([]int32,_tmp795)
  if len(basis) > 0 { _tmp797 = (*int32)(&n[0]) }
  if _tmp798 := C.MSK_initbasissolve(self.ptr(),_tmp797); _tmp798 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp798)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) InputData(maxnumcon int32,maxnumvar int32,c []float64,cfix float64,aptrb []int64,aptre []int64,asub []int32,aval []float64,bkc []Boundkey,blc []float64,buc []float64,bkx []Boundkey,blx []float64,bux []float64) (err error) {
  _tmp799 := len(buc)
  if _tmp799 < bkc { _tmp799 = lof["name"] }
  if _tmp799 < blc { _tmp799 = lof["name"] }
  var numcon int32 = int32(_tmp799)
  _tmp800 := len(aptrb)
  if _tmp800 < blx { _tmp800 = lof["name"] }
  if _tmp800 < bux { _tmp800 = lof["name"] }
  if _tmp800 < bkx { _tmp800 = lof["name"] }
  if _tmp800 < aptre { _tmp800 = lof["name"] }
  if _tmp800 < c { _tmp800 = lof["name"] }
  var numvar int32 = int32(_tmp800)
  var _tmp801 *float64
  if c != nil { _tmp801 = (*float64)(&c[0]) }
  var _tmp802 *int64
  if aptrb != nil { _tmp802 = (*int64)(&aptrb[0]) }
  var _tmp803 *int64
  if aptre != nil { _tmp803 = (*int64)(&aptre[0]) }
  var _tmp804 *int32
  if asub != nil { _tmp804 = (*int32)(&asub[0]) }
  var _tmp805 *float64
  if aval != nil { _tmp805 = (*float64)(&aval[0]) }
  var _tmp806 *Boundkey
  if bkc != nil { _tmp806 = (*Boundkey)(&bkc[0]) }
  var _tmp807 *float64
  if blc != nil { _tmp807 = (*float64)(&blc[0]) }
  var _tmp808 *float64
  if buc != nil { _tmp808 = (*float64)(&buc[0]) }
  var _tmp809 *Boundkey
  if bkx != nil { _tmp809 = (*Boundkey)(&bkx[0]) }
  var _tmp810 *float64
  if blx != nil { _tmp810 = (*float64)(&blx[0]) }
  var _tmp811 *float64
  if bux != nil { _tmp811 = (*float64)(&bux[0]) }
  if _tmp812 := C.MSK_inputdata64(self.ptr(),C.MSKint32t(maxnumcon),C.MSKint32t(maxnumvar),C.MSKint32t(numcon),C.MSKint32t(numvar),_tmp801,C.MSKrealt(cfix),_tmp802,_tmp803,_tmp804,_tmp805,_tmp806,_tmp807,_tmp808,_tmp809,_tmp810,_tmp811); _tmp812 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp812)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) IsDouParName(parname string) (param Dparam,err error) {
  _tmp813 := C.CString(parname)
  var _tmp814 C.MSKdparame
  if _tmp815 := C.MSK_isdouparname(self.ptr(),_tmp813,&_tmp814); _tmp815 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp815)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  param = Dparam(_tmp814)
  return
}
func (self *Task) IsIntParName(parname string) (param Iparam,err error) {
  _tmp816 := C.CString(parname)
  var _tmp817 C.MSKiparame
  if _tmp818 := C.MSK_isintparname(self.ptr(),_tmp816,&_tmp817); _tmp818 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp818)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  param = Iparam(_tmp817)
  return
}
func (self *Task) IsStrParName(parname string) (param Sparam,err error) {
  _tmp819 := C.CString(parname)
  var _tmp820 C.MSKsparame
  if _tmp821 := C.MSK_isstrparname(self.ptr(),_tmp819,&_tmp820); _tmp821 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp821)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  param = Sparam(_tmp820)
  return
}
func (self *Task) LinkFileToStream(whichstream Streamtype,filename string,append int32) (err error) {
  _tmp822 := C.CString(filename)
  if _tmp823 := C.MSK_linkfiletotaskstream(self.ptr(),C.MSKstreamtypee(whichstream),_tmp822,C.MSKint32t(append)); _tmp823 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp823)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) OneSolutionSummary(whichstream Streamtype,whichsol Soltype) (err error) {
  if _tmp824 := C.MSK_onesolutionsummary(self.ptr(),C.MSKstreamtypee(whichstream),C.MSKsoltypee(whichsol)); _tmp824 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp824)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) OptimizerSummary(whichstream Streamtype) (err error) {
  if _tmp825 := C.MSK_optimizersummary(self.ptr(),C.MSKstreamtypee(whichstream)); _tmp825 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp825)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) Optimize() (trmcode Rescode,err error) {
  var _tmp826 C.MSKrescodee
  if _tmp827 := C.MSK_optimizetrm(self.ptr(),&_tmp826); _tmp827 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp827)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  trmcode = Rescode(_tmp826)
  return
}
func (self *Task) PrimalRepair(wlc []float64,wuc []float64,wlx []float64,wux []float64) (err error) {
  var _tmp830 *float64
  var _tmp828 int32
  if _tmp829 := C.MSK_getnumcon(self.ptr(),addr(_tmp828)); _tmp829 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp829)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  if int64(len(wlc)) < int64(_tmp828) {
    err = &ArrayLengthError{fun:"PrimalRepair",arg:"wlc"}
    return
  }
  if wlc != nil { _tmp830 = (*float64)(&wlc[0]) }
  var _tmp833 *float64
  var _tmp831 int32
  if _tmp832 := C.MSK_getnumcon(self.ptr(),addr(_tmp831)); _tmp832 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp832)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  if int64(len(wuc)) < int64(_tmp831) {
    err = &ArrayLengthError{fun:"PrimalRepair",arg:"wuc"}
    return
  }
  if wuc != nil { _tmp833 = (*float64)(&wuc[0]) }
  var _tmp836 *float64
  var _tmp834 int32
  if _tmp835 := C.MSK_getnumvar(self.ptr(),addr(_tmp834)); _tmp835 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp835)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  if int64(len(wlx)) < int64(_tmp834) {
    err = &ArrayLengthError{fun:"PrimalRepair",arg:"wlx"}
    return
  }
  if wlx != nil { _tmp836 = (*float64)(&wlx[0]) }
  var _tmp839 *float64
  var _tmp837 int32
  if _tmp838 := C.MSK_getnumvar(self.ptr(),addr(_tmp837)); _tmp838 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp838)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  if int64(len(wux)) < int64(_tmp837) {
    err = &ArrayLengthError{fun:"PrimalRepair",arg:"wux"}
    return
  }
  if wux != nil { _tmp839 = (*float64)(&wux[0]) }
  if _tmp840 := C.MSK_primalrepair(self.ptr(),_tmp830,_tmp833,_tmp836,_tmp839); _tmp840 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp840)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PrimalSensitivity(subi []int32,marki []Mark,subj []int32,markj []Mark) (leftpricei []float64,rightpricei []float64,leftrangei []float64,rightrangei []float64,leftpricej []float64,rightpricej []float64,leftrangej []float64,rightrangej []float64,err error) {
  _tmp841 := len(subi)
  if _tmp841 < marki { _tmp841 = lof["name"] }
  var numi int32 = int32(_tmp841)
  var _tmp842 *int32
  if subi != nil { _tmp842 = (*int32)(&subi[0]) }
  var _tmp843 *Mark
  if marki != nil { _tmp843 = (*Mark)(&marki[0]) }
  _tmp844 := len(markj)
  if _tmp844 < subj { _tmp844 = lof["name"] }
  var numj int32 = int32(_tmp844)
  var _tmp845 *int32
  if subj != nil { _tmp845 = (*int32)(&subj[0]) }
  var _tmp846 *Mark
  if markj != nil { _tmp846 = (*Mark)(&markj[0]) }
  var _tmp847 *float64
  leftpricei := make([]float64,numi)
  if len(leftpricei) > 0 { _tmp847 = (*float64)(&n[0]) }
  var _tmp848 *float64
  rightpricei := make([]float64,numi)
  if len(rightpricei) > 0 { _tmp848 = (*float64)(&n[0]) }
  var _tmp849 *float64
  leftrangei := make([]float64,numi)
  if len(leftrangei) > 0 { _tmp849 = (*float64)(&n[0]) }
  var _tmp850 *float64
  rightrangei := make([]float64,numi)
  if len(rightrangei) > 0 { _tmp850 = (*float64)(&n[0]) }
  var _tmp851 *float64
  leftpricej := make([]float64,numj)
  if len(leftpricej) > 0 { _tmp851 = (*float64)(&n[0]) }
  var _tmp852 *float64
  rightpricej := make([]float64,numj)
  if len(rightpricej) > 0 { _tmp852 = (*float64)(&n[0]) }
  var _tmp853 *float64
  leftrangej := make([]float64,numj)
  if len(leftrangej) > 0 { _tmp853 = (*float64)(&n[0]) }
  var _tmp854 *float64
  rightrangej := make([]float64,numj)
  if len(rightrangej) > 0 { _tmp854 = (*float64)(&n[0]) }
  if _tmp855 := C.MSK_primalsensitivity(self.ptr(),C.MSKint32t(numi),_tmp842,_tmp843,C.MSKint32t(numj),_tmp845,_tmp846,_tmp847,_tmp848,_tmp849,_tmp850,_tmp851,_tmp852,_tmp853,_tmp854); _tmp855 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp855)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) ProbTypeToStr(probtype Problemtype) (str string,err error) {
  _tmp856 := make([]byte,max_str_len)
  if _tmp857 := C.MSK_probtypetostr(self.ptr(),C.MSKproblemtypee(probtype),C.CString(&tmpvar1[0])); _tmp857 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp857)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  var str string
  if p := strings.IndexByte(_tmp856,byte(0)); p < 0 {
    str = string(_tmp856)
  } else {
    str = string(_tmp856[:p])
  }
  return
}
func (self *Task) ProStaToStr(problemsta Prosta) (str string,err error) {
  _tmp858 := make([]byte,max_str_len)
  if _tmp859 := C.MSK_prostatostr(self.ptr(),C.MSKprostae(problemsta),C.CString(&tmpvar1[0])); _tmp859 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp859)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  var str string
  if p := strings.IndexByte(_tmp858,byte(0)); p < 0 {
    str = string(_tmp858)
  } else {
    str = string(_tmp858[:p])
  }
  return
}
func (self *Task) PutAcc(accidx int64,domidx int64,afeidxlist []int64,b []float64) (err error) {
  _tmp860 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp860)
  var _tmp861 *int64
  if afeidxlist != nil { _tmp861 = (*int64)(&afeidxlist[0]) }
  var _tmp862 *float64
  if int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"PutAcc",arg:"b"}
    return
  }
  if b != nil { _tmp862 = (*float64)(&b[0]) }
  if _tmp863 := C.MSK_putacc(self.ptr(),C.MSKint32t(accidx),C.MSKint32t(domidx),C.MSKint32t(numafeidx),_tmp861,_tmp862); _tmp863 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp863)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAccB(accidx int64,b []float64) (err error) {
  _tmp864 := len(b)
  var lengthb int64 = int64(_tmp864)
  var _tmp865 *float64
  if b != nil { _tmp865 = (*float64)(&b[0]) }
  if _tmp866 := C.MSK_putaccb(self.ptr(),C.MSKint32t(accidx),C.MSKint32t(lengthb),_tmp865); _tmp866 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp866)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAccBJ(accidx int64,j int64,bj float64) (err error) {
  if _tmp867 := C.MSK_putaccbj(self.ptr(),C.MSKint32t(accidx),C.MSKint32t(j),C.MSKrealt(bj)); _tmp867 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp867)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAccDotY(whichsol Soltype,accidx int64) (doty []float64,err error) {
  var _tmp870 *float64
  var _tmp868 int64
  if _tmp869 := C.MSK_getaccn(self.ptr(),accidx,addr(_tmp868)); _tmp869 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp869)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  doty := make([]float64,_tmp868)
  if len(doty) > 0 { _tmp870 = (*float64)(&n[0]) }
  if _tmp871 := C.MSK_putaccdoty(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(accidx),_tmp870); _tmp871 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp871)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAccList(accidxs []int64,domidxs []int64,afeidxlist []int64,b []float64) (err error) {
  _tmp872 := len(accidxs)
  if _tmp872 < domidxs { _tmp872 = lof["name"] }
  var numaccs int64 = int64(_tmp872)
  var _tmp873 *int64
  if accidxs != nil { _tmp873 = (*int64)(&accidxs[0]) }
  var _tmp874 *int64
  if domidxs != nil { _tmp874 = (*int64)(&domidxs[0]) }
  _tmp875 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp875)
  var _tmp876 *int64
  if afeidxlist != nil { _tmp876 = (*int64)(&afeidxlist[0]) }
  var _tmp877 *float64
  if int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"PutAccList",arg:"b"}
    return
  }
  if b != nil { _tmp877 = (*float64)(&b[0]) }
  if _tmp878 := C.MSK_putacclist(self.ptr(),C.MSKint32t(numaccs),_tmp873,_tmp874,C.MSKint32t(numafeidx),_tmp876,_tmp877); _tmp878 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp878)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAccName(accidx int64,name string) (err error) {
  _tmp879 := C.CString(name)
  if _tmp880 := C.MSK_putaccname(self.ptr(),C.MSKint32t(accidx),_tmp879); _tmp880 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp880)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutACol(j int32,subj []int32,valj []float64) (err error) {
  _tmp881 := len(valj)
  if _tmp881 < subj { _tmp881 = lof["name"] }
  var nzj int32 = int32(_tmp881)
  var _tmp882 *int32
  if subj != nil { _tmp882 = (*int32)(&subj[0]) }
  var _tmp883 *float64
  if valj != nil { _tmp883 = (*float64)(&valj[0]) }
  if _tmp884 := C.MSK_putacol(self.ptr(),C.MSKint32t(j),C.MSKint32t(nzj),_tmp882,_tmp883); _tmp884 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp884)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAColList(sub []int32,ptrb []int32,ptre []int32,asub []int32,aval []float64) (err error) {
  _tmp885 := len(ptrb)
  if _tmp885 < ptre { _tmp885 = lof["name"] }
  if _tmp885 < sub { _tmp885 = lof["name"] }
  var num int32 = int32(_tmp885)
  var _tmp886 *int32
  if sub != nil { _tmp886 = (*int32)(&sub[0]) }
  var _tmp887 *int32
  if ptrb != nil { _tmp887 = (*int32)(&ptrb[0]) }
  var _tmp888 *int32
  if ptre != nil { _tmp888 = (*int32)(&ptre[0]) }
  var _tmp889 *int32
  if asub != nil { _tmp889 = (*int32)(&asub[0]) }
  var _tmp890 *float64
  if aval != nil { _tmp890 = (*float64)(&aval[0]) }
  if _tmp891 := C.MSK_putacollist(self.ptr(),C.MSKint32t(num),_tmp886,_tmp887,_tmp888,_tmp889,_tmp890); _tmp891 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp891)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAColSlice(first int32,last int32,ptrb []int64,ptre []int64,asub []int32,aval []float64) (err error) {
  var _tmp892 *int64
  if ptrb != nil { _tmp892 = (*int64)(&ptrb[0]) }
  var _tmp893 *int64
  if ptre != nil { _tmp893 = (*int64)(&ptre[0]) }
  var _tmp894 *int32
  if asub != nil { _tmp894 = (*int32)(&asub[0]) }
  var _tmp895 *float64
  if aval != nil { _tmp895 = (*float64)(&aval[0]) }
  if _tmp896 := C.MSK_putacolslice64(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),_tmp892,_tmp893,_tmp894,_tmp895); _tmp896 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp896)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeBarfBlockTriplet(afeidx []int64,barvaridx []int32,subk []int32,subl []int32,valkl []float64) (err error) {
  _tmp897 := len(subl)
  if _tmp897 < subk { _tmp897 = lof["name"] }
  if _tmp897 < barvaridx { _tmp897 = lof["name"] }
  if _tmp897 < valkl { _tmp897 = lof["name"] }
  if _tmp897 < afeidx { _tmp897 = lof["name"] }
  var numtrip int64 = int64(_tmp897)
  var _tmp898 *int64
  if int64(len(afeidx)) < int64(numtrip) {
    err = &ArrayLengthError{fun:"PutAfeBarfBlockTriplet",arg:"afeidx"}
    return
  }
  if afeidx != nil { _tmp898 = (*int64)(&afeidx[0]) }
  var _tmp899 *int32
  if int64(len(barvaridx)) < int64(numtrip) {
    err = &ArrayLengthError{fun:"PutAfeBarfBlockTriplet",arg:"barvaridx"}
    return
  }
  if barvaridx != nil { _tmp899 = (*int32)(&barvaridx[0]) }
  var _tmp900 *int32
  if int64(len(subk)) < int64(numtrip) {
    err = &ArrayLengthError{fun:"PutAfeBarfBlockTriplet",arg:"subk"}
    return
  }
  if subk != nil { _tmp900 = (*int32)(&subk[0]) }
  var _tmp901 *int32
  if int64(len(subl)) < int64(numtrip) {
    err = &ArrayLengthError{fun:"PutAfeBarfBlockTriplet",arg:"subl"}
    return
  }
  if subl != nil { _tmp901 = (*int32)(&subl[0]) }
  var _tmp902 *float64
  if int64(len(valkl)) < int64(numtrip) {
    err = &ArrayLengthError{fun:"PutAfeBarfBlockTriplet",arg:"valkl"}
    return
  }
  if valkl != nil { _tmp902 = (*float64)(&valkl[0]) }
  if _tmp903 := C.MSK_putafebarfblocktriplet(self.ptr(),C.MSKint32t(numtrip),_tmp898,_tmp899,_tmp900,_tmp901,_tmp902); _tmp903 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp903)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeBarfEntry(afeidx int64,barvaridx int32,termidx []int64,termweight []float64) (err error) {
  _tmp904 := len(termidx)
  if _tmp904 < termweight { _tmp904 = lof["name"] }
  var numterm int64 = int64(_tmp904)
  var _tmp905 *int64
  if termidx != nil { _tmp905 = (*int64)(&termidx[0]) }
  var _tmp906 *float64
  if termweight != nil { _tmp906 = (*float64)(&termweight[0]) }
  if _tmp907 := C.MSK_putafebarfentry(self.ptr(),C.MSKint32t(afeidx),C.MSKint32t(barvaridx),C.MSKint32t(numterm),_tmp905,_tmp906); _tmp907 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp907)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeBarfEntryList(afeidx []int64,barvaridx []int32,numterm []int64,ptrterm []int64,termidx []int64,termweight []float64) (err error) {
  _tmp908 := len(barvaridx)
  if _tmp908 < ptrterm { _tmp908 = lof["name"] }
  if _tmp908 < numterm { _tmp908 = lof["name"] }
  if _tmp908 < afeidx { _tmp908 = lof["name"] }
  var numafeidx int64 = int64(_tmp908)
  var _tmp909 *int64
  if afeidx != nil { _tmp909 = (*int64)(&afeidx[0]) }
  var _tmp910 *int32
  if barvaridx != nil { _tmp910 = (*int32)(&barvaridx[0]) }
  var _tmp911 *int64
  if numterm != nil { _tmp911 = (*int64)(&numterm[0]) }
  var _tmp912 *int64
  if ptrterm != nil { _tmp912 = (*int64)(&ptrterm[0]) }
  _tmp913 := len(termidx)
  if _tmp913 < termweight { _tmp913 = lof["name"] }
  var lenterm int64 = int64(_tmp913)
  var _tmp914 *int64
  if termidx != nil { _tmp914 = (*int64)(&termidx[0]) }
  var _tmp915 *float64
  if termweight != nil { _tmp915 = (*float64)(&termweight[0]) }
  if _tmp916 := C.MSK_putafebarfentrylist(self.ptr(),C.MSKint32t(numafeidx),_tmp909,_tmp910,_tmp911,_tmp912,C.MSKint32t(lenterm),_tmp914,_tmp915); _tmp916 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp916)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeBarfRow(afeidx int64,barvaridx []int32,numterm []int64,ptrterm []int64,termidx []int64,termweight []float64) (err error) {
  _tmp917 := len(barvaridx)
  if _tmp917 < ptrterm { _tmp917 = lof["name"] }
  if _tmp917 < numterm { _tmp917 = lof["name"] }
  var numentr int32 = int32(_tmp917)
  var _tmp918 *int32
  if barvaridx != nil { _tmp918 = (*int32)(&barvaridx[0]) }
  var _tmp919 *int64
  if numterm != nil { _tmp919 = (*int64)(&numterm[0]) }
  var _tmp920 *int64
  if ptrterm != nil { _tmp920 = (*int64)(&ptrterm[0]) }
  _tmp921 := len(termidx)
  if _tmp921 < termweight { _tmp921 = lof["name"] }
  var lenterm int64 = int64(_tmp921)
  var _tmp922 *int64
  if termidx != nil { _tmp922 = (*int64)(&termidx[0]) }
  var _tmp923 *float64
  if termweight != nil { _tmp923 = (*float64)(&termweight[0]) }
  if _tmp924 := C.MSK_putafebarfrow(self.ptr(),C.MSKint32t(afeidx),C.MSKint32t(numentr),_tmp918,_tmp919,_tmp920,C.MSKint32t(lenterm),_tmp922,_tmp923); _tmp924 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp924)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeFCol(varidx int32,afeidx []int64,val []float64) (err error) {
  _tmp925 := len(val)
  if _tmp925 < afeidx { _tmp925 = lof["name"] }
  var numnz int64 = int64(_tmp925)
  var _tmp926 *int64
  if afeidx != nil { _tmp926 = (*int64)(&afeidx[0]) }
  var _tmp927 *float64
  if val != nil { _tmp927 = (*float64)(&val[0]) }
  if _tmp928 := C.MSK_putafefcol(self.ptr(),C.MSKint32t(varidx),C.MSKint32t(numnz),_tmp926,_tmp927); _tmp928 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp928)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeFEntry(afeidx int64,varidx int32,value float64) (err error) {
  if _tmp929 := C.MSK_putafefentry(self.ptr(),C.MSKint32t(afeidx),C.MSKint32t(varidx),C.MSKrealt(value)); _tmp929 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp929)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeFEntryList(afeidx []int64,varidx []int32,val []float64) (err error) {
  _tmp930 := len(val)
  if _tmp930 < varidx { _tmp930 = lof["name"] }
  if _tmp930 < afeidx { _tmp930 = lof["name"] }
  var numentr int64 = int64(_tmp930)
  var _tmp931 *int64
  if afeidx != nil { _tmp931 = (*int64)(&afeidx[0]) }
  var _tmp932 *int32
  if varidx != nil { _tmp932 = (*int32)(&varidx[0]) }
  var _tmp933 *float64
  if val != nil { _tmp933 = (*float64)(&val[0]) }
  if _tmp934 := C.MSK_putafefentrylist(self.ptr(),C.MSKint32t(numentr),_tmp931,_tmp932,_tmp933); _tmp934 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp934)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeFRow(afeidx int64,varidx []int32,val []float64) (err error) {
  _tmp935 := len(varidx)
  if _tmp935 < val { _tmp935 = lof["name"] }
  var numnz int32 = int32(_tmp935)
  var _tmp936 *int32
  if varidx != nil { _tmp936 = (*int32)(&varidx[0]) }
  var _tmp937 *float64
  if val != nil { _tmp937 = (*float64)(&val[0]) }
  if _tmp938 := C.MSK_putafefrow(self.ptr(),C.MSKint32t(afeidx),C.MSKint32t(numnz),_tmp936,_tmp937); _tmp938 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp938)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeFRowList(afeidx []int64,numnzrow []int32,ptrrow []int64,varidx []int32,val []float64) (err error) {
  _tmp939 := len(ptrrow)
  if _tmp939 < numnzrow { _tmp939 = lof["name"] }
  if _tmp939 < afeidx { _tmp939 = lof["name"] }
  var numafeidx int64 = int64(_tmp939)
  var _tmp940 *int64
  if afeidx != nil { _tmp940 = (*int64)(&afeidx[0]) }
  var _tmp941 *int32
  if numnzrow != nil { _tmp941 = (*int32)(&numnzrow[0]) }
  var _tmp942 *int64
  if ptrrow != nil { _tmp942 = (*int64)(&ptrrow[0]) }
  _tmp943 := len(varidx)
  if _tmp943 < val { _tmp943 = lof["name"] }
  var lenidxval int64 = int64(_tmp943)
  var _tmp944 *int32
  if varidx != nil { _tmp944 = (*int32)(&varidx[0]) }
  var _tmp945 *float64
  if val != nil { _tmp945 = (*float64)(&val[0]) }
  if _tmp946 := C.MSK_putafefrowlist(self.ptr(),C.MSKint32t(numafeidx),_tmp940,_tmp941,_tmp942,C.MSKint32t(lenidxval),_tmp944,_tmp945); _tmp946 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp946)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeG(afeidx int64,g float64) (err error) {
  if _tmp947 := C.MSK_putafeg(self.ptr(),C.MSKint32t(afeidx),C.MSKrealt(g)); _tmp947 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp947)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeGList(afeidx []int64,g []float64) (err error) {
  _tmp948 := len(g)
  if _tmp948 < afeidx { _tmp948 = lof["name"] }
  var numafeidx int64 = int64(_tmp948)
  var _tmp949 *int64
  if afeidx != nil { _tmp949 = (*int64)(&afeidx[0]) }
  var _tmp950 *float64
  if g != nil { _tmp950 = (*float64)(&g[0]) }
  if _tmp951 := C.MSK_putafeglist(self.ptr(),C.MSKint32t(numafeidx),_tmp949,_tmp950); _tmp951 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp951)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAfeGSlice(first int64,last int64,slice []float64) (err error) {
  var _tmp952 *float64
  if int64(len(slice)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutAfeGSlice",arg:"slice"}
    return
  }
  if slice != nil { _tmp952 = (*float64)(&slice[0]) }
  if _tmp953 := C.MSK_putafegslice(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),_tmp952); _tmp953 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp953)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAij(i int32,j int32,aij float64) (err error) {
  if _tmp954 := C.MSK_putaij(self.ptr(),C.MSKint32t(i),C.MSKint32t(j),C.MSKrealt(aij)); _tmp954 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp954)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutAijList(subi []int32,subj []int32,valij []float64) (err error) {
  _tmp955 := len(subi)
  if _tmp955 < valij { _tmp955 = lof["name"] }
  if _tmp955 < subj { _tmp955 = lof["name"] }
  var num int64 = int64(_tmp955)
  var _tmp956 *int32
  if subi != nil { _tmp956 = (*int32)(&subi[0]) }
  var _tmp957 *int32
  if subj != nil { _tmp957 = (*int32)(&subj[0]) }
  var _tmp958 *float64
  if valij != nil { _tmp958 = (*float64)(&valij[0]) }
  if _tmp959 := C.MSK_putaijlist64(self.ptr(),C.MSKint32t(num),_tmp956,_tmp957,_tmp958); _tmp959 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp959)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutARow(i int32,subi []int32,vali []float64) (err error) {
  _tmp960 := len(vali)
  if _tmp960 < subi { _tmp960 = lof["name"] }
  var nzi int32 = int32(_tmp960)
  var _tmp961 *int32
  if subi != nil { _tmp961 = (*int32)(&subi[0]) }
  var _tmp962 *float64
  if vali != nil { _tmp962 = (*float64)(&vali[0]) }
  if _tmp963 := C.MSK_putarow(self.ptr(),C.MSKint32t(i),C.MSKint32t(nzi),_tmp961,_tmp962); _tmp963 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp963)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutARowList(sub []int32,ptrb []int64,ptre []int64,asub []int32,aval []float64) (err error) {
  _tmp964 := len(ptrb)
  if _tmp964 < ptre { _tmp964 = lof["name"] }
  if _tmp964 < sub { _tmp964 = lof["name"] }
  var num int32 = int32(_tmp964)
  var _tmp965 *int32
  if sub != nil { _tmp965 = (*int32)(&sub[0]) }
  var _tmp966 *int64
  if ptrb != nil { _tmp966 = (*int64)(&ptrb[0]) }
  var _tmp967 *int64
  if ptre != nil { _tmp967 = (*int64)(&ptre[0]) }
  var _tmp968 *int32
  if asub != nil { _tmp968 = (*int32)(&asub[0]) }
  var _tmp969 *float64
  if aval != nil { _tmp969 = (*float64)(&aval[0]) }
  if _tmp970 := C.MSK_putarowlist64(self.ptr(),C.MSKint32t(num),_tmp965,_tmp966,_tmp967,_tmp968,_tmp969); _tmp970 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp970)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutARowSlice(first int32,last int32,ptrb []int64,ptre []int64,asub []int32,aval []float64) (err error) {
  var _tmp971 *int64
  if int64(len(ptrb)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutARowSlice",arg:"ptrb"}
    return
  }
  if ptrb != nil { _tmp971 = (*int64)(&ptrb[0]) }
  var _tmp972 *int64
  if int64(len(ptre)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutARowSlice",arg:"ptre"}
    return
  }
  if ptre != nil { _tmp972 = (*int64)(&ptre[0]) }
  var _tmp973 *int32
  if asub != nil { _tmp973 = (*int32)(&asub[0]) }
  var _tmp974 *float64
  if aval != nil { _tmp974 = (*float64)(&aval[0]) }
  if _tmp975 := C.MSK_putarowslice64(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),_tmp971,_tmp972,_tmp973,_tmp974); _tmp975 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp975)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutATruncateTol(tolzero float64) (err error) {
  if _tmp976 := C.MSK_putatruncatetol(self.ptr(),C.MSKrealt(tolzero)); _tmp976 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp976)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBaraBlockTriplet(subi []int32,subj []int32,subk []int32,subl []int32,valijkl []float64) (err error) {
  _tmp977 := len(subl)
  if _tmp977 < valijkl { _tmp977 = lof["name"] }
  if _tmp977 < subk { _tmp977 = lof["name"] }
  if _tmp977 < subj { _tmp977 = lof["name"] }
  var num int64 = int64(_tmp977)
  var _tmp978 *int32
  if int64(len(subi)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBaraBlockTriplet",arg:"subi"}
    return
  }
  if subi != nil { _tmp978 = (*int32)(&subi[0]) }
  var _tmp979 *int32
  if int64(len(subj)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBaraBlockTriplet",arg:"subj"}
    return
  }
  if subj != nil { _tmp979 = (*int32)(&subj[0]) }
  var _tmp980 *int32
  if int64(len(subk)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBaraBlockTriplet",arg:"subk"}
    return
  }
  if subk != nil { _tmp980 = (*int32)(&subk[0]) }
  var _tmp981 *int32
  if int64(len(subl)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBaraBlockTriplet",arg:"subl"}
    return
  }
  if subl != nil { _tmp981 = (*int32)(&subl[0]) }
  var _tmp982 *float64
  if int64(len(valijkl)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBaraBlockTriplet",arg:"valijkl"}
    return
  }
  if valijkl != nil { _tmp982 = (*float64)(&valijkl[0]) }
  if _tmp983 := C.MSK_putbarablocktriplet(self.ptr(),C.MSKint32t(num),_tmp978,_tmp979,_tmp980,_tmp981,_tmp982); _tmp983 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp983)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBaraIj(i int32,j int32,sub []int64,weights []float64) (err error) {
  _tmp984 := len(weights)
  if _tmp984 < sub { _tmp984 = lof["name"] }
  var num int64 = int64(_tmp984)
  var _tmp985 *int64
  if sub != nil { _tmp985 = (*int64)(&sub[0]) }
  var _tmp986 *float64
  if weights != nil { _tmp986 = (*float64)(&weights[0]) }
  if _tmp987 := C.MSK_putbaraij(self.ptr(),C.MSKint32t(i),C.MSKint32t(j),C.MSKint32t(num),_tmp985,_tmp986); _tmp987 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp987)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBaraIjList(subi []int32,subj []int32,alphaptrb []int64,alphaptre []int64,matidx []int64,weights []float64) (err error) {
  _tmp988 := len(alphaptrb)
  if _tmp988 < subi { _tmp988 = lof["name"] }
  if _tmp988 < alphaptre { _tmp988 = lof["name"] }
  if _tmp988 < subj { _tmp988 = lof["name"] }
  var num int32 = int32(_tmp988)
  var _tmp989 *int32
  if subi != nil { _tmp989 = (*int32)(&subi[0]) }
  var _tmp990 *int32
  if subj != nil { _tmp990 = (*int32)(&subj[0]) }
  var _tmp991 *int64
  if alphaptrb != nil { _tmp991 = (*int64)(&alphaptrb[0]) }
  var _tmp992 *int64
  if alphaptre != nil { _tmp992 = (*int64)(&alphaptre[0]) }
  var _tmp993 *int64
  if matidx != nil { _tmp993 = (*int64)(&matidx[0]) }
  var _tmp994 *float64
  if weights != nil { _tmp994 = (*float64)(&weights[0]) }
  if _tmp995 := C.MSK_putbaraijlist(self.ptr(),C.MSKint32t(num),_tmp989,_tmp990,_tmp991,_tmp992,_tmp993,_tmp994); _tmp995 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp995)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBaraRowList(subi []int32,ptrb []int64,ptre []int64,subj []int32,nummat []int64,matidx []int64,weights []float64) (err error) {
  _tmp996 := len(ptrb)
  if _tmp996 < ptre { _tmp996 = lof["name"] }
  if _tmp996 < subi { _tmp996 = lof["name"] }
  var num int32 = int32(_tmp996)
  var _tmp997 *int32
  if subi != nil { _tmp997 = (*int32)(&subi[0]) }
  var _tmp998 *int64
  if ptrb != nil { _tmp998 = (*int64)(&ptrb[0]) }
  var _tmp999 *int64
  if ptre != nil { _tmp999 = (*int64)(&ptre[0]) }
  var _tmp1000 *int32
  if subj != nil { _tmp1000 = (*int32)(&subj[0]) }
  var _tmp1001 *int64
  if int64(len(nummat)) < int64(cast[int32](subj.len)) {
    err = &ArrayLengthError{fun:"PutBaraRowList",arg:"nummat"}
    return
  }
  if nummat != nil { _tmp1001 = (*int64)(&nummat[0]) }
  var _tmp1002 *int64
  if int64(len(matidx)) < int64(nummat.foldl(a+b)) {
    err = &ArrayLengthError{fun:"PutBaraRowList",arg:"matidx"}
    return
  }
  if matidx != nil { _tmp1002 = (*int64)(&matidx[0]) }
  var _tmp1003 *float64
  if int64(len(weights)) < int64(nummat.foldl(a+b)) {
    err = &ArrayLengthError{fun:"PutBaraRowList",arg:"weights"}
    return
  }
  if weights != nil { _tmp1003 = (*float64)(&weights[0]) }
  if _tmp1004 := C.MSK_putbararowlist(self.ptr(),C.MSKint32t(num),_tmp997,_tmp998,_tmp999,_tmp1000,_tmp1001,_tmp1002,_tmp1003); _tmp1004 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1004)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBarcBlockTriplet(subj []int32,subk []int32,subl []int32,valjkl []float64) (err error) {
  _tmp1005 := len(valjkl)
  if _tmp1005 < subl { _tmp1005 = lof["name"] }
  if _tmp1005 < subk { _tmp1005 = lof["name"] }
  if _tmp1005 < subj { _tmp1005 = lof["name"] }
  var num int64 = int64(_tmp1005)
  var _tmp1006 *int32
  if int64(len(subj)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBarcBlockTriplet",arg:"subj"}
    return
  }
  if subj != nil { _tmp1006 = (*int32)(&subj[0]) }
  var _tmp1007 *int32
  if int64(len(subk)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBarcBlockTriplet",arg:"subk"}
    return
  }
  if subk != nil { _tmp1007 = (*int32)(&subk[0]) }
  var _tmp1008 *int32
  if int64(len(subl)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBarcBlockTriplet",arg:"subl"}
    return
  }
  if subl != nil { _tmp1008 = (*int32)(&subl[0]) }
  var _tmp1009 *float64
  if int64(len(valjkl)) < int64(num) {
    err = &ArrayLengthError{fun:"PutBarcBlockTriplet",arg:"valjkl"}
    return
  }
  if valjkl != nil { _tmp1009 = (*float64)(&valjkl[0]) }
  if _tmp1010 := C.MSK_putbarcblocktriplet(self.ptr(),C.MSKint32t(num),_tmp1006,_tmp1007,_tmp1008,_tmp1009); _tmp1010 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1010)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBarcJ(j int32,sub []int64,weights []float64) (err error) {
  _tmp1011 := len(weights)
  if _tmp1011 < sub { _tmp1011 = lof["name"] }
  var num int64 = int64(_tmp1011)
  var _tmp1012 *int64
  if sub != nil { _tmp1012 = (*int64)(&sub[0]) }
  var _tmp1013 *float64
  if weights != nil { _tmp1013 = (*float64)(&weights[0]) }
  if _tmp1014 := C.MSK_putbarcj(self.ptr(),C.MSKint32t(j),C.MSKint32t(num),_tmp1012,_tmp1013); _tmp1014 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1014)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBarsJ(whichsol Soltype,j int32,barsj []float64) (err error) {
  var _tmp1017 *float64
  var _tmp1015 int64
  if _tmp1016 := C.MSK_getlenbarvarj(self.ptr(),j,addr(_tmp1015)); _tmp1016 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp1016)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  if int64(len(barsj)) < int64(_tmp1015) {
    err = &ArrayLengthError{fun:"PutBarsJ",arg:"barsj"}
    return
  }
  if barsj != nil { _tmp1017 = (*float64)(&barsj[0]) }
  if _tmp1018 := C.MSK_putbarsj(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(j),_tmp1017); _tmp1018 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1018)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBarvarName(j int32,name string) (err error) {
  _tmp1019 := C.CString(name)
  if _tmp1020 := C.MSK_putbarvarname(self.ptr(),C.MSKint32t(j),_tmp1019); _tmp1020 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1020)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutBarxJ(whichsol Soltype,j int32,barxj []float64) (err error) {
  var _tmp1023 *float64
  var _tmp1021 int64
  if _tmp1022 := C.MSK_getlenbarvarj(self.ptr(),j,addr(_tmp1021)); _tmp1022 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp1022)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  if int64(len(barxj)) < int64(_tmp1021) {
    err = &ArrayLengthError{fun:"PutBarxJ",arg:"barxj"}
    return
  }
  if barxj != nil { _tmp1023 = (*float64)(&barxj[0]) }
  if _tmp1024 := C.MSK_putbarxj(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(j),_tmp1023); _tmp1024 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1024)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutCfix(cfix float64) (err error) {
  if _tmp1025 := C.MSK_putcfix(self.ptr(),C.MSKrealt(cfix)); _tmp1025 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1025)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutCJ(j int32,cj float64) (err error) {
  if _tmp1026 := C.MSK_putcj(self.ptr(),C.MSKint32t(j),C.MSKrealt(cj)); _tmp1026 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1026)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutCList(subj []int32,val []float64) (err error) {
  _tmp1027 := len(val)
  if _tmp1027 < subj { _tmp1027 = lof["name"] }
  var num int32 = int32(_tmp1027)
  var _tmp1028 *int32
  if subj != nil { _tmp1028 = (*int32)(&subj[0]) }
  var _tmp1029 *float64
  if val != nil { _tmp1029 = (*float64)(&val[0]) }
  if _tmp1030 := C.MSK_putclist(self.ptr(),C.MSKint32t(num),_tmp1028,_tmp1029); _tmp1030 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1030)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutConBound(i int32,bkc Boundkey,blc float64,buc float64) (err error) {
  if _tmp1031 := C.MSK_putconbound(self.ptr(),C.MSKint32t(i),C.MSKboundkeye(bkc),C.MSKrealt(blc),C.MSKrealt(buc)); _tmp1031 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1031)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutConBoundList(sub []int32,bkc []Boundkey,blc []float64,buc []float64) (err error) {
  _tmp1032 := len(bkc)
  if _tmp1032 < blc { _tmp1032 = lof["name"] }
  if _tmp1032 < buc { _tmp1032 = lof["name"] }
  if _tmp1032 < sub { _tmp1032 = lof["name"] }
  var num int32 = int32(_tmp1032)
  var _tmp1033 *int32
  if sub != nil { _tmp1033 = (*int32)(&sub[0]) }
  var _tmp1034 *Boundkey
  if bkc != nil { _tmp1034 = (*Boundkey)(&bkc[0]) }
  var _tmp1035 *float64
  if blc != nil { _tmp1035 = (*float64)(&blc[0]) }
  var _tmp1036 *float64
  if buc != nil { _tmp1036 = (*float64)(&buc[0]) }
  if _tmp1037 := C.MSK_putconboundlist(self.ptr(),C.MSKint32t(num),_tmp1033,_tmp1034,_tmp1035,_tmp1036); _tmp1037 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1037)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutConBoundListConst(sub []int32,bkc Boundkey,blc float64,buc float64) (err error) {
  _tmp1038 := len(sub)
  var num int32 = int32(_tmp1038)
  var _tmp1039 *int32
  if sub != nil { _tmp1039 = (*int32)(&sub[0]) }
  if _tmp1040 := C.MSK_putconboundlistconst(self.ptr(),C.MSKint32t(num),_tmp1039,C.MSKboundkeye(bkc),C.MSKrealt(blc),C.MSKrealt(buc)); _tmp1040 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1040)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutConBoundSlice(first int32,last int32,bkc []Boundkey,blc []float64,buc []float64) (err error) {
  var _tmp1041 *Boundkey
  if int64(len(bkc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutConBoundSlice",arg:"bkc"}
    return
  }
  if bkc != nil { _tmp1041 = (*Boundkey)(&bkc[0]) }
  var _tmp1042 *float64
  if int64(len(blc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutConBoundSlice",arg:"blc"}
    return
  }
  if blc != nil { _tmp1042 = (*float64)(&blc[0]) }
  var _tmp1043 *float64
  if int64(len(buc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutConBoundSlice",arg:"buc"}
    return
  }
  if buc != nil { _tmp1043 = (*float64)(&buc[0]) }
  if _tmp1044 := C.MSK_putconboundslice(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),_tmp1041,_tmp1042,_tmp1043); _tmp1044 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1044)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutConBoundSliceConst(first int32,last int32,bkc Boundkey,blc float64,buc float64) (err error) {
  if _tmp1045 := C.MSK_putconboundsliceconst(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),C.MSKboundkeye(bkc),C.MSKrealt(blc),C.MSKrealt(buc)); _tmp1045 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1045)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutCone(k int32,ct Conetype,conepar float64,submem []int32) (err error) {
  _tmp1046 := len(submem)
  var nummem int32 = int32(_tmp1046)
  var _tmp1047 *int32
  if submem != nil { _tmp1047 = (*int32)(&submem[0]) }
  if _tmp1048 := C.MSK_putcone(self.ptr(),C.MSKint32t(k),C.MSKconetypee(ct),C.MSKrealt(conepar),C.MSKint32t(nummem),_tmp1047); _tmp1048 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1048)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutConeName(j int32,name string) (err error) {
  _tmp1049 := C.CString(name)
  if _tmp1050 := C.MSK_putconename(self.ptr(),C.MSKint32t(j),_tmp1049); _tmp1050 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1050)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutConName(i int32,name string) (err error) {
  _tmp1051 := C.CString(name)
  if _tmp1052 := C.MSK_putconname(self.ptr(),C.MSKint32t(i),_tmp1051); _tmp1052 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1052)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutConSolutionI(i int32,whichsol Soltype,sk Stakey,x float64,sl float64,su float64) (err error) {
  if _tmp1053 := C.MSK_putconsolutioni(self.ptr(),C.MSKint32t(i),C.MSKsoltypee(whichsol),C.MSKstakeye(sk),C.MSKrealt(x),C.MSKrealt(sl),C.MSKrealt(su)); _tmp1053 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1053)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutCSlice(first int32,last int32,slice []float64) (err error) {
  var _tmp1054 *float64
  if int64(len(slice)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutCSlice",arg:"slice"}
    return
  }
  if slice != nil { _tmp1054 = (*float64)(&slice[0]) }
  if _tmp1055 := C.MSK_putcslice(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),_tmp1054); _tmp1055 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1055)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutDjc(djcidx int64,domidxlist []int64,afeidxlist []int64,b []float64,termsizelist []int64) (err error) {
  _tmp1056 := len(domidxlist)
  var numdomidx int64 = int64(_tmp1056)
  var _tmp1057 *int64
  if domidxlist != nil { _tmp1057 = (*int64)(&domidxlist[0]) }
  _tmp1058 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp1058)
  var _tmp1059 *int64
  if afeidxlist != nil { _tmp1059 = (*int64)(&afeidxlist[0]) }
  var _tmp1060 *float64
  if int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"PutDjc",arg:"b"}
    return
  }
  if b != nil { _tmp1060 = (*float64)(&b[0]) }
  _tmp1061 := len(termsizelist)
  var numterms int64 = int64(_tmp1061)
  var _tmp1062 *int64
  if termsizelist != nil { _tmp1062 = (*int64)(&termsizelist[0]) }
  if _tmp1063 := C.MSK_putdjc(self.ptr(),C.MSKint32t(djcidx),C.MSKint32t(numdomidx),_tmp1057,C.MSKint32t(numafeidx),_tmp1059,_tmp1060,C.MSKint32t(numterms),_tmp1062); _tmp1063 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1063)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutDjcName(djcidx int64,name string) (err error) {
  _tmp1064 := C.CString(name)
  if _tmp1065 := C.MSK_putdjcname(self.ptr(),C.MSKint32t(djcidx),_tmp1064); _tmp1065 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1065)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutDjcSlice(idxfirst int64,idxlast int64,domidxlist []int64,afeidxlist []int64,b []float64,termsizelist []int64,termsindjc []int64) (err error) {
  _tmp1066 := len(domidxlist)
  var numdomidx int64 = int64(_tmp1066)
  var _tmp1067 *int64
  if domidxlist != nil { _tmp1067 = (*int64)(&domidxlist[0]) }
  _tmp1068 := len(afeidxlist)
  var numafeidx int64 = int64(_tmp1068)
  var _tmp1069 *int64
  if afeidxlist != nil { _tmp1069 = (*int64)(&afeidxlist[0]) }
  var _tmp1070 *float64
  if int64(len(b)) < int64(numafeidx) {
    err = &ArrayLengthError{fun:"PutDjcSlice",arg:"b"}
    return
  }
  if b != nil { _tmp1070 = (*float64)(&b[0]) }
  _tmp1071 := len(termsizelist)
  var numterms int64 = int64(_tmp1071)
  var _tmp1072 *int64
  if termsizelist != nil { _tmp1072 = (*int64)(&termsizelist[0]) }
  var _tmp1073 *int64
  if int64(len(termsindjc)) < int64((idxlast - idxfirst)) {
    err = &ArrayLengthError{fun:"PutDjcSlice",arg:"termsindjc"}
    return
  }
  if termsindjc != nil { _tmp1073 = (*int64)(&termsindjc[0]) }
  if _tmp1074 := C.MSK_putdjcslice(self.ptr(),C.MSKint32t(idxfirst),C.MSKint32t(idxlast),C.MSKint32t(numdomidx),_tmp1067,C.MSKint32t(numafeidx),_tmp1069,_tmp1070,C.MSKint32t(numterms),_tmp1072,_tmp1073); _tmp1074 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1074)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutDomainName(domidx int64,name string) (err error) {
  _tmp1075 := C.CString(name)
  if _tmp1076 := C.MSK_putdomainname(self.ptr(),C.MSKint32t(domidx),_tmp1075); _tmp1076 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1076)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutDouParam(param Dparam,parvalue float64) (err error) {
  if _tmp1077 := C.MSK_putdouparam(self.ptr(),C.MSKdparame(param),C.MSKrealt(parvalue)); _tmp1077 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1077)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutIntParam(param Iparam,parvalue int32) (err error) {
  if _tmp1078 := C.MSK_putintparam(self.ptr(),C.MSKiparame(param),C.MSKint32t(parvalue)); _tmp1078 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1078)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumAcc(maxnumacc int64) (err error) {
  if _tmp1079 := C.MSK_putmaxnumacc(self.ptr(),C.MSKint32t(maxnumacc)); _tmp1079 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1079)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumAfe(maxnumafe int64) (err error) {
  if _tmp1080 := C.MSK_putmaxnumafe(self.ptr(),C.MSKint32t(maxnumafe)); _tmp1080 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1080)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumANz(maxnumanz int64) (err error) {
  if _tmp1081 := C.MSK_putmaxnumanz(self.ptr(),C.MSKint32t(maxnumanz)); _tmp1081 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1081)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumBarvar(maxnumbarvar int32) (err error) {
  if _tmp1082 := C.MSK_putmaxnumbarvar(self.ptr(),C.MSKint32t(maxnumbarvar)); _tmp1082 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1082)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumCon(maxnumcon int32) (err error) {
  if _tmp1083 := C.MSK_putmaxnumcon(self.ptr(),C.MSKint32t(maxnumcon)); _tmp1083 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1083)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumCone(maxnumcone int32) (err error) {
  if _tmp1084 := C.MSK_putmaxnumcone(self.ptr(),C.MSKint32t(maxnumcone)); _tmp1084 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1084)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumDjc(maxnumdjc int64) (err error) {
  if _tmp1085 := C.MSK_putmaxnumdjc(self.ptr(),C.MSKint32t(maxnumdjc)); _tmp1085 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1085)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumDomain(maxnumdomain int64) (err error) {
  if _tmp1086 := C.MSK_putmaxnumdomain(self.ptr(),C.MSKint32t(maxnumdomain)); _tmp1086 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1086)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumQNz(maxnumqnz int64) (err error) {
  if _tmp1087 := C.MSK_putmaxnumqnz(self.ptr(),C.MSKint32t(maxnumqnz)); _tmp1087 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1087)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutMaxNumVar(maxnumvar int32) (err error) {
  if _tmp1088 := C.MSK_putmaxnumvar(self.ptr(),C.MSKint32t(maxnumvar)); _tmp1088 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1088)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutNaDouParam(paramname string,parvalue float64) (err error) {
  _tmp1089 := C.CString(paramname)
  if _tmp1090 := C.MSK_putnadouparam(self.ptr(),_tmp1089,C.MSKrealt(parvalue)); _tmp1090 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1090)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutNaIntParam(paramname string,parvalue int32) (err error) {
  _tmp1091 := C.CString(paramname)
  if _tmp1092 := C.MSK_putnaintparam(self.ptr(),_tmp1091,C.MSKint32t(parvalue)); _tmp1092 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1092)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutNaStrParam(paramname string,parvalue string) (err error) {
  _tmp1093 := C.CString(paramname)
  _tmp1094 := C.CString(parvalue)
  if _tmp1095 := C.MSK_putnastrparam(self.ptr(),_tmp1093,_tmp1094); _tmp1095 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1095)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutObjName(objname string) (err error) {
  _tmp1096 := C.CString(objname)
  if _tmp1097 := C.MSK_putobjname(self.ptr(),_tmp1096); _tmp1097 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1097)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutObjSense(sense Objsense) (err error) {
  if _tmp1098 := C.MSK_putobjsense(self.ptr(),C.MSKobjsensee(sense)); _tmp1098 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1098)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutOptserverHost(host string) (err error) {
  _tmp1099 := C.CString(host)
  if _tmp1100 := C.MSK_putoptserverhost(self.ptr(),_tmp1099); _tmp1100 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1100)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutParam(parname string,parvalue string) (err error) {
  _tmp1101 := C.CString(parname)
  _tmp1102 := C.CString(parvalue)
  if _tmp1103 := C.MSK_putparam(self.ptr(),_tmp1101,_tmp1102); _tmp1103 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1103)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutQCon(qcsubk []int32,qcsubi []int32,qcsubj []int32,qcval []float64) (err error) {
  _tmp1104 := len(qcval)
  if _tmp1104 < qcsubi { _tmp1104 = lof["name"] }
  if _tmp1104 < qcsubj { _tmp1104 = lof["name"] }
  var numqcnz int32 = int32(_tmp1104)
  var _tmp1105 *int32
  if qcsubk != nil { _tmp1105 = (*int32)(&qcsubk[0]) }
  var _tmp1106 *int32
  if qcsubi != nil { _tmp1106 = (*int32)(&qcsubi[0]) }
  var _tmp1107 *int32
  if qcsubj != nil { _tmp1107 = (*int32)(&qcsubj[0]) }
  var _tmp1108 *float64
  if qcval != nil { _tmp1108 = (*float64)(&qcval[0]) }
  if _tmp1109 := C.MSK_putqcon(self.ptr(),C.MSKint32t(numqcnz),_tmp1105,_tmp1106,_tmp1107,_tmp1108); _tmp1109 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1109)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutQConK(k int32,qcsubi []int32,qcsubj []int32,qcval []float64) (err error) {
  _tmp1110 := len(qcval)
  if _tmp1110 < qcsubi { _tmp1110 = lof["name"] }
  if _tmp1110 < qcsubj { _tmp1110 = lof["name"] }
  var numqcnz int32 = int32(_tmp1110)
  var _tmp1111 *int32
  if qcsubi != nil { _tmp1111 = (*int32)(&qcsubi[0]) }
  var _tmp1112 *int32
  if qcsubj != nil { _tmp1112 = (*int32)(&qcsubj[0]) }
  var _tmp1113 *float64
  if qcval != nil { _tmp1113 = (*float64)(&qcval[0]) }
  if _tmp1114 := C.MSK_putqconk(self.ptr(),C.MSKint32t(k),C.MSKint32t(numqcnz),_tmp1111,_tmp1112,_tmp1113); _tmp1114 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1114)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutQObj(qosubi []int32,qosubj []int32,qoval []float64) (err error) {
  _tmp1115 := len(qosubi)
  if _tmp1115 < qoval { _tmp1115 = lof["name"] }
  if _tmp1115 < qosubj { _tmp1115 = lof["name"] }
  var numqonz int32 = int32(_tmp1115)
  var _tmp1116 *int32
  if qosubi != nil { _tmp1116 = (*int32)(&qosubi[0]) }
  var _tmp1117 *int32
  if qosubj != nil { _tmp1117 = (*int32)(&qosubj[0]) }
  var _tmp1118 *float64
  if qoval != nil { _tmp1118 = (*float64)(&qoval[0]) }
  if _tmp1119 := C.MSK_putqobj(self.ptr(),C.MSKint32t(numqonz),_tmp1116,_tmp1117,_tmp1118); _tmp1119 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1119)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutQObjIJ(i int32,j int32,qoij float64) (err error) {
  if _tmp1120 := C.MSK_putqobjij(self.ptr(),C.MSKint32t(i),C.MSKint32t(j),C.MSKrealt(qoij)); _tmp1120 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1120)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSkc(whichsol Soltype,skc []Stakey) (err error) {
  var _tmp1123 *Stakey
  var _tmp1121 int32
  if _tmp1122 := C.MSK_getnumcon(self.ptr(),addr(_tmp1121)); _tmp1122 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp1122)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  if int64(len(skc)) < int64(_tmp1121) {
    err = &ArrayLengthError{fun:"PutSkc",arg:"skc"}
    return
  }
  if skc != nil { _tmp1123 = (*Stakey)(&skc[0]) }
  if _tmp1124 := C.MSK_putskc(self.ptr(),C.MSKsoltypee(whichsol),_tmp1123); _tmp1124 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1124)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSkcSlice(whichsol Soltype,first int32,last int32,skc []Stakey) (err error) {
  var _tmp1125 *Stakey
  if int64(len(skc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSkcSlice",arg:"skc"}
    return
  }
  if skc != nil { _tmp1125 = (*Stakey)(&skc[0]) }
  if _tmp1126 := C.MSK_putskcslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp1125); _tmp1126 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1126)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSkx(whichsol Soltype,skx []Stakey) (err error) {
  var _tmp1129 *Stakey
  var _tmp1127 int32
  if _tmp1128 := C.MSK_getnumvar(self.ptr(),addr(_tmp1127)); _tmp1128 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp1128)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  if int64(len(skx)) < int64(_tmp1127) {
    err = &ArrayLengthError{fun:"PutSkx",arg:"skx"}
    return
  }
  if skx != nil { _tmp1129 = (*Stakey)(&skx[0]) }
  if _tmp1130 := C.MSK_putskx(self.ptr(),C.MSKsoltypee(whichsol),_tmp1129); _tmp1130 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1130)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSkxSlice(whichsol Soltype,first int32,last int32,skx []Stakey) (err error) {
  var _tmp1131 *Stakey
  if int64(len(skx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSkxSlice",arg:"skx"}
    return
  }
  if skx != nil { _tmp1131 = (*Stakey)(&skx[0]) }
  if _tmp1132 := C.MSK_putskxslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp1131); _tmp1132 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1132)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSlc(whichsol Soltype,slc []float64) (err error) {
  var _tmp1135 *float64
  var _tmp1133 int32
  if _tmp1134 := C.MSK_getnumcon(self.ptr(),addr(_tmp1133)); _tmp1134 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp1134)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  if int64(len(slc)) < int64(_tmp1133) {
    err = &ArrayLengthError{fun:"PutSlc",arg:"slc"}
    return
  }
  if slc != nil { _tmp1135 = (*float64)(&slc[0]) }
  if _tmp1136 := C.MSK_putslc(self.ptr(),C.MSKsoltypee(whichsol),_tmp1135); _tmp1136 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1136)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSlcSlice(whichsol Soltype,first int32,last int32,slc []float64) (err error) {
  var _tmp1137 *float64
  if int64(len(slc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSlcSlice",arg:"slc"}
    return
  }
  if slc != nil { _tmp1137 = (*float64)(&slc[0]) }
  if _tmp1138 := C.MSK_putslcslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp1137); _tmp1138 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1138)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSlx(whichsol Soltype,slx []float64) (err error) {
  var _tmp1141 *float64
  var _tmp1139 int32
  if _tmp1140 := C.MSK_getnumvar(self.ptr(),addr(_tmp1139)); _tmp1140 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp1140)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  if int64(len(slx)) < int64(_tmp1139) {
    err = &ArrayLengthError{fun:"PutSlx",arg:"slx"}
    return
  }
  if slx != nil { _tmp1141 = (*float64)(&slx[0]) }
  if _tmp1142 := C.MSK_putslx(self.ptr(),C.MSKsoltypee(whichsol),_tmp1141); _tmp1142 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1142)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSlxSlice(whichsol Soltype,first int32,last int32,slx []float64) (err error) {
  var _tmp1143 *float64
  if int64(len(slx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSlxSlice",arg:"slx"}
    return
  }
  if slx != nil { _tmp1143 = (*float64)(&slx[0]) }
  if _tmp1144 := C.MSK_putslxslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp1143); _tmp1144 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1144)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSnx(whichsol Soltype,sux []float64) (err error) {
  var _tmp1147 *float64
  var _tmp1145 int32
  if _tmp1146 := C.MSK_getnumvar(self.ptr(),addr(_tmp1145)); _tmp1146 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp1146)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  if int64(len(sux)) < int64(_tmp1145) {
    err = &ArrayLengthError{fun:"PutSnx",arg:"sux"}
    return
  }
  if sux != nil { _tmp1147 = (*float64)(&sux[0]) }
  if _tmp1148 := C.MSK_putsnx(self.ptr(),C.MSKsoltypee(whichsol),_tmp1147); _tmp1148 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1148)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSnxSlice(whichsol Soltype,first int32,last int32,snx []float64) (err error) {
  var _tmp1149 *float64
  if int64(len(snx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSnxSlice",arg:"snx"}
    return
  }
  if snx != nil { _tmp1149 = (*float64)(&snx[0]) }
  if _tmp1150 := C.MSK_putsnxslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp1149); _tmp1150 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1150)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSolution(whichsol Soltype,skc []Stakey,skx []Stakey,skn []Stakey,xc []float64,xx []float64,y []float64,slc []float64,suc []float64,slx []float64,sux []float64,snx []float64) (err error) {
  var _tmp1151 *Stakey
  if skc != nil { _tmp1151 = (*Stakey)(&skc[0]) }
  var _tmp1152 *Stakey
  if skx != nil { _tmp1152 = (*Stakey)(&skx[0]) }
  var _tmp1153 *Stakey
  if skn != nil { _tmp1153 = (*Stakey)(&skn[0]) }
  var _tmp1154 *float64
  if xc != nil { _tmp1154 = (*float64)(&xc[0]) }
  var _tmp1155 *float64
  if xx != nil { _tmp1155 = (*float64)(&xx[0]) }
  var _tmp1156 *float64
  if y != nil { _tmp1156 = (*float64)(&y[0]) }
  var _tmp1157 *float64
  if slc != nil { _tmp1157 = (*float64)(&slc[0]) }
  var _tmp1158 *float64
  if suc != nil { _tmp1158 = (*float64)(&suc[0]) }
  var _tmp1159 *float64
  if slx != nil { _tmp1159 = (*float64)(&slx[0]) }
  var _tmp1160 *float64
  if sux != nil { _tmp1160 = (*float64)(&sux[0]) }
  var _tmp1161 *float64
  if snx != nil { _tmp1161 = (*float64)(&snx[0]) }
  if _tmp1162 := C.MSK_putsolution(self.ptr(),C.MSKsoltypee(whichsol),_tmp1151,_tmp1152,_tmp1153,_tmp1154,_tmp1155,_tmp1156,_tmp1157,_tmp1158,_tmp1159,_tmp1160,_tmp1161); _tmp1162 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1162)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSolutionNew(whichsol Soltype,skc []Stakey,skx []Stakey,skn []Stakey,xc []float64,xx []float64,y []float64,slc []float64,suc []float64,slx []float64,sux []float64,snx []float64,doty []float64) (err error) {
  var _tmp1163 *Stakey
  if skc != nil { _tmp1163 = (*Stakey)(&skc[0]) }
  var _tmp1164 *Stakey
  if skx != nil { _tmp1164 = (*Stakey)(&skx[0]) }
  var _tmp1165 *Stakey
  if skn != nil { _tmp1165 = (*Stakey)(&skn[0]) }
  var _tmp1166 *float64
  if xc != nil { _tmp1166 = (*float64)(&xc[0]) }
  var _tmp1167 *float64
  if xx != nil { _tmp1167 = (*float64)(&xx[0]) }
  var _tmp1168 *float64
  if y != nil { _tmp1168 = (*float64)(&y[0]) }
  var _tmp1169 *float64
  if slc != nil { _tmp1169 = (*float64)(&slc[0]) }
  var _tmp1170 *float64
  if suc != nil { _tmp1170 = (*float64)(&suc[0]) }
  var _tmp1171 *float64
  if slx != nil { _tmp1171 = (*float64)(&slx[0]) }
  var _tmp1172 *float64
  if sux != nil { _tmp1172 = (*float64)(&sux[0]) }
  var _tmp1173 *float64
  if snx != nil { _tmp1173 = (*float64)(&snx[0]) }
  var _tmp1174 *float64
  if doty != nil { _tmp1174 = (*float64)(&doty[0]) }
  if _tmp1175 := C.MSK_putsolutionnew(self.ptr(),C.MSKsoltypee(whichsol),_tmp1163,_tmp1164,_tmp1165,_tmp1166,_tmp1167,_tmp1168,_tmp1169,_tmp1170,_tmp1171,_tmp1172,_tmp1173,_tmp1174); _tmp1175 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1175)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSolutionYI(i int32,whichsol Soltype,y float64) (err error) {
  if _tmp1176 := C.MSK_putsolutionyi(self.ptr(),C.MSKint32t(i),C.MSKsoltypee(whichsol),C.MSKrealt(y)); _tmp1176 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1176)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutStrParam(param Sparam,parvalue string) (err error) {
  _tmp1177 := C.CString(parvalue)
  if _tmp1178 := C.MSK_putstrparam(self.ptr(),C.MSKsparame(param),_tmp1177); _tmp1178 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1178)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSuc(whichsol Soltype,suc []float64) (err error) {
  var _tmp1181 *float64
  var _tmp1179 int32
  if _tmp1180 := C.MSK_getnumcon(self.ptr(),addr(_tmp1179)); _tmp1180 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp1180)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  if int64(len(suc)) < int64(_tmp1179) {
    err = &ArrayLengthError{fun:"PutSuc",arg:"suc"}
    return
  }
  if suc != nil { _tmp1181 = (*float64)(&suc[0]) }
  if _tmp1182 := C.MSK_putsuc(self.ptr(),C.MSKsoltypee(whichsol),_tmp1181); _tmp1182 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1182)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSucSlice(whichsol Soltype,first int32,last int32,suc []float64) (err error) {
  var _tmp1183 *float64
  if int64(len(suc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSucSlice",arg:"suc"}
    return
  }
  if suc != nil { _tmp1183 = (*float64)(&suc[0]) }
  if _tmp1184 := C.MSK_putsucslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp1183); _tmp1184 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1184)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSux(whichsol Soltype,sux []float64) (err error) {
  var _tmp1187 *float64
  var _tmp1185 int32
  if _tmp1186 := C.MSK_getnumvar(self.ptr(),addr(_tmp1185)); _tmp1186 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp1186)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  if int64(len(sux)) < int64(_tmp1185) {
    err = &ArrayLengthError{fun:"PutSux",arg:"sux"}
    return
  }
  if sux != nil { _tmp1187 = (*float64)(&sux[0]) }
  if _tmp1188 := C.MSK_putsux(self.ptr(),C.MSKsoltypee(whichsol),_tmp1187); _tmp1188 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1188)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutSuxSlice(whichsol Soltype,first int32,last int32,sux []float64) (err error) {
  var _tmp1189 *float64
  if int64(len(sux)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutSuxSlice",arg:"sux"}
    return
  }
  if sux != nil { _tmp1189 = (*float64)(&sux[0]) }
  if _tmp1190 := C.MSK_putsuxslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp1189); _tmp1190 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1190)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutTaskName(taskname string) (err error) {
  _tmp1191 := C.CString(taskname)
  if _tmp1192 := C.MSK_puttaskname(self.ptr(),_tmp1191); _tmp1192 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1192)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarBound(j int32,bkx Boundkey,blx float64,bux float64) (err error) {
  if _tmp1193 := C.MSK_putvarbound(self.ptr(),C.MSKint32t(j),C.MSKboundkeye(bkx),C.MSKrealt(blx),C.MSKrealt(bux)); _tmp1193 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1193)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarBoundList(sub []int32,bkx []Boundkey,blx []float64,bux []float64) (err error) {
  _tmp1194 := len(bkx)
  if _tmp1194 < blx { _tmp1194 = lof["name"] }
  if _tmp1194 < bux { _tmp1194 = lof["name"] }
  if _tmp1194 < sub { _tmp1194 = lof["name"] }
  var num int32 = int32(_tmp1194)
  var _tmp1195 *int32
  if sub != nil { _tmp1195 = (*int32)(&sub[0]) }
  var _tmp1196 *Boundkey
  if bkx != nil { _tmp1196 = (*Boundkey)(&bkx[0]) }
  var _tmp1197 *float64
  if blx != nil { _tmp1197 = (*float64)(&blx[0]) }
  var _tmp1198 *float64
  if bux != nil { _tmp1198 = (*float64)(&bux[0]) }
  if _tmp1199 := C.MSK_putvarboundlist(self.ptr(),C.MSKint32t(num),_tmp1195,_tmp1196,_tmp1197,_tmp1198); _tmp1199 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1199)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarBoundListConst(sub []int32,bkx Boundkey,blx float64,bux float64) (err error) {
  _tmp1200 := len(sub)
  var num int32 = int32(_tmp1200)
  var _tmp1201 *int32
  if sub != nil { _tmp1201 = (*int32)(&sub[0]) }
  if _tmp1202 := C.MSK_putvarboundlistconst(self.ptr(),C.MSKint32t(num),_tmp1201,C.MSKboundkeye(bkx),C.MSKrealt(blx),C.MSKrealt(bux)); _tmp1202 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1202)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarBoundSlice(first int32,last int32,bkx []Boundkey,blx []float64,bux []float64) (err error) {
  var _tmp1203 *Boundkey
  if int64(len(bkx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutVarBoundSlice",arg:"bkx"}
    return
  }
  if bkx != nil { _tmp1203 = (*Boundkey)(&bkx[0]) }
  var _tmp1204 *float64
  if int64(len(blx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutVarBoundSlice",arg:"blx"}
    return
  }
  if blx != nil { _tmp1204 = (*float64)(&blx[0]) }
  var _tmp1205 *float64
  if int64(len(bux)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutVarBoundSlice",arg:"bux"}
    return
  }
  if bux != nil { _tmp1205 = (*float64)(&bux[0]) }
  if _tmp1206 := C.MSK_putvarboundslice(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),_tmp1203,_tmp1204,_tmp1205); _tmp1206 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1206)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarBoundSliceConst(first int32,last int32,bkx Boundkey,blx float64,bux float64) (err error) {
  if _tmp1207 := C.MSK_putvarboundsliceconst(self.ptr(),C.MSKint32t(first),C.MSKint32t(last),C.MSKboundkeye(bkx),C.MSKrealt(blx),C.MSKrealt(bux)); _tmp1207 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1207)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarName(j int32,name string) (err error) {
  _tmp1208 := C.CString(name)
  if _tmp1209 := C.MSK_putvarname(self.ptr(),C.MSKint32t(j),_tmp1208); _tmp1209 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1209)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarSolutionJ(j int32,whichsol Soltype,sk Stakey,x float64,sl float64,su float64,sn float64) (err error) {
  if _tmp1210 := C.MSK_putvarsolutionj(self.ptr(),C.MSKint32t(j),C.MSKsoltypee(whichsol),C.MSKstakeye(sk),C.MSKrealt(x),C.MSKrealt(sl),C.MSKrealt(su),C.MSKrealt(sn)); _tmp1210 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1210)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarType(j int32,vartype Variabletype) (err error) {
  if _tmp1211 := C.MSK_putvartype(self.ptr(),C.MSKint32t(j),C.MSKvariabletypee(vartype)); _tmp1211 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1211)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutVarTypeList(subj []int32,vartype []Variabletype) (err error) {
  _tmp1212 := len(vartype)
  if _tmp1212 < subj { _tmp1212 = lof["name"] }
  var num int32 = int32(_tmp1212)
  var _tmp1213 *int32
  if subj != nil { _tmp1213 = (*int32)(&subj[0]) }
  var _tmp1214 *Variabletype
  if vartype != nil { _tmp1214 = (*Variabletype)(&vartype[0]) }
  if _tmp1215 := C.MSK_putvartypelist(self.ptr(),C.MSKint32t(num),_tmp1213,_tmp1214); _tmp1215 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1215)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutXc(whichsol Soltype) (xc []float64,err error) {
  var _tmp1218 *float64
  var _tmp1216 int32
  if _tmp1217 := C.MSK_getnumcon(self.ptr(),addr(_tmp1216)); _tmp1217 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp1217)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  xc := make([]float64,_tmp1216)
  if len(xc) > 0 { _tmp1218 = (*float64)(&n[0]) }
  if _tmp1219 := C.MSK_putxc(self.ptr(),C.MSKsoltypee(whichsol),_tmp1218); _tmp1219 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1219)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutXcSlice(whichsol Soltype,first int32,last int32,xc []float64) (err error) {
  var _tmp1220 *float64
  if int64(len(xc)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutXcSlice",arg:"xc"}
    return
  }
  if xc != nil { _tmp1220 = (*float64)(&xc[0]) }
  if _tmp1221 := C.MSK_putxcslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp1220); _tmp1221 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1221)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutXx(whichsol Soltype,xx []float64) (err error) {
  var _tmp1224 *float64
  var _tmp1222 int32
  if _tmp1223 := C.MSK_getnumvar(self.ptr(),addr(_tmp1222)); _tmp1223 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp1223)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  if int64(len(xx)) < int64(_tmp1222) {
    err = &ArrayLengthError{fun:"PutXx",arg:"xx"}
    return
  }
  if xx != nil { _tmp1224 = (*float64)(&xx[0]) }
  if _tmp1225 := C.MSK_putxx(self.ptr(),C.MSKsoltypee(whichsol),_tmp1224); _tmp1225 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1225)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutXxSlice(whichsol Soltype,first int32,last int32,xx []float64) (err error) {
  var _tmp1226 *float64
  if int64(len(xx)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutXxSlice",arg:"xx"}
    return
  }
  if xx != nil { _tmp1226 = (*float64)(&xx[0]) }
  if _tmp1227 := C.MSK_putxxslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp1226); _tmp1227 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1227)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutY(whichsol Soltype,y []float64) (err error) {
  var _tmp1230 *float64
  var _tmp1228 int32
  if _tmp1229 := C.MSK_getnumcon(self.ptr(),addr(_tmp1228)); _tmp1229 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp1229)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  if int64(len(y)) < int64(_tmp1228) {
    err = &ArrayLengthError{fun:"PutY",arg:"y"}
    return
  }
  if y != nil { _tmp1230 = (*float64)(&y[0]) }
  if _tmp1231 := C.MSK_puty(self.ptr(),C.MSKsoltypee(whichsol),_tmp1230); _tmp1231 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1231)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) PutYSlice(whichsol Soltype,first int32,last int32,y []float64) (err error) {
  var _tmp1232 *float64
  if int64(len(y)) < int64((last - first)) {
    err = &ArrayLengthError{fun:"PutYSlice",arg:"y"}
    return
  }
  if y != nil { _tmp1232 = (*float64)(&y[0]) }
  if _tmp1233 := C.MSK_putyslice(self.ptr(),C.MSKsoltypee(whichsol),C.MSKint32t(first),C.MSKint32t(last),_tmp1232); _tmp1233 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1233)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadBSolution(filename string,compress Compresstype) (err error) {
  _tmp1234 := C.CString(filename)
  if _tmp1235 := C.MSK_readbsolution(self.ptr(),_tmp1234,C.MSKcompresstypee(compress)); _tmp1235 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1235)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadData(filename string) (err error) {
  _tmp1236 := C.CString(filename)
  if _tmp1237 := C.MSK_readdataautoformat(self.ptr(),_tmp1236); _tmp1237 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1237)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadDataFormat(filename string,format Dataformat,compress Compresstype) (err error) {
  _tmp1238 := C.CString(filename)
  if _tmp1239 := C.MSK_readdataformat(self.ptr(),_tmp1238,C.MSKdataformate(format),C.MSKcompresstypee(compress)); _tmp1239 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1239)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadJsonSol(filename string) (err error) {
  _tmp1240 := C.CString(filename)
  if _tmp1241 := C.MSK_readjsonsol(self.ptr(),_tmp1240); _tmp1241 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1241)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadJsonString(data string) (err error) {
  _tmp1242 := C.CString(data)
  if _tmp1243 := C.MSK_readjsonstring(self.ptr(),_tmp1242); _tmp1243 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1243)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadLpString(data string) (err error) {
  _tmp1244 := C.CString(data)
  if _tmp1245 := C.MSK_readlpstring(self.ptr(),_tmp1244); _tmp1245 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1245)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadOpfString(data string) (err error) {
  _tmp1246 := C.CString(data)
  if _tmp1247 := C.MSK_readopfstring(self.ptr(),_tmp1246); _tmp1247 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1247)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadParamFile(filename string) (err error) {
  _tmp1248 := C.CString(filename)
  if _tmp1249 := C.MSK_readparamfile(self.ptr(),_tmp1248); _tmp1249 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1249)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadPtfString(data string) (err error) {
  _tmp1250 := C.CString(data)
  if _tmp1251 := C.MSK_readptfstring(self.ptr(),_tmp1250); _tmp1251 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1251)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadSolution(whichsol Soltype,filename string) (err error) {
  _tmp1252 := C.CString(filename)
  if _tmp1253 := C.MSK_readsolution(self.ptr(),C.MSKsoltypee(whichsol),_tmp1252); _tmp1253 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1253)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadSolutionFile(filename string) (err error) {
  _tmp1254 := C.CString(filename)
  if _tmp1255 := C.MSK_readsolutionfile(self.ptr(),_tmp1254); _tmp1255 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1255)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadSummary(whichstream Streamtype) (err error) {
  if _tmp1256 := C.MSK_readsummary(self.ptr(),C.MSKstreamtypee(whichstream)); _tmp1256 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1256)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) ReadTask(filename string) (err error) {
  _tmp1257 := C.CString(filename)
  if _tmp1258 := C.MSK_readtask(self.ptr(),_tmp1257); _tmp1258 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1258)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) RemoveBarvars(subset []int32) (err error) {
  _tmp1259 := len(subset)
  var num int32 = int32(_tmp1259)
  var _tmp1260 *int32
  if subset != nil { _tmp1260 = (*int32)(&subset[0]) }
  if _tmp1261 := C.MSK_removebarvars(self.ptr(),C.MSKint32t(num),_tmp1260); _tmp1261 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1261)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) RemoveCones(subset []int32) (err error) {
  _tmp1262 := len(subset)
  var num int32 = int32(_tmp1262)
  var _tmp1263 *int32
  if subset != nil { _tmp1263 = (*int32)(&subset[0]) }
  if _tmp1264 := C.MSK_removecones(self.ptr(),C.MSKint32t(num),_tmp1263); _tmp1264 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1264)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) RemoveCons(subset []int32) (err error) {
  _tmp1265 := len(subset)
  var num int32 = int32(_tmp1265)
  var _tmp1266 *int32
  if subset != nil { _tmp1266 = (*int32)(&subset[0]) }
  if _tmp1267 := C.MSK_removecons(self.ptr(),C.MSKint32t(num),_tmp1266); _tmp1267 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1267)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) RemoveVars(subset []int32) (err error) {
  _tmp1268 := len(subset)
  var num int32 = int32(_tmp1268)
  var _tmp1269 *int32
  if subset != nil { _tmp1269 = (*int32)(&subset[0]) }
  if _tmp1270 := C.MSK_removevars(self.ptr(),C.MSKint32t(num),_tmp1269); _tmp1270 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1270)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) ResizeTask(maxnumcon int32,maxnumvar int32,maxnumcone int32,maxnumanz int64,maxnumqnz int64) (err error) {
  if _tmp1271 := C.MSK_resizetask(self.ptr(),C.MSKint32t(maxnumcon),C.MSKint32t(maxnumvar),C.MSKint32t(maxnumcone),C.MSKint32t(maxnumanz),C.MSKint32t(maxnumqnz)); _tmp1271 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1271)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) SensitivityReport(whichstream Streamtype) (err error) {
  if _tmp1272 := C.MSK_sensitivityreport(self.ptr(),C.MSKstreamtypee(whichstream)); _tmp1272 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1272)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) SetDefaults() (err error) {
  if _tmp1273 := C.MSK_setdefaults(self.ptr()); _tmp1273 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1273)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) SkToStr(sk Stakey) (str string,err error) {
  _tmp1274 := make([]byte,max_str_len)
  if _tmp1275 := C.MSK_sktostr(self.ptr(),C.MSKstakeye(sk),C.CString(&tmpvar1[0])); _tmp1275 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1275)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  var str string
  if p := strings.IndexByte(_tmp1274,byte(0)); p < 0 {
    str = string(_tmp1274)
  } else {
    str = string(_tmp1274[:p])
  }
  return
}
func (self *Task) SolStaToStr(solutionsta Solsta) (str string,err error) {
  _tmp1276 := make([]byte,max_str_len)
  if _tmp1277 := C.MSK_solstatostr(self.ptr(),C.MSKsolstae(solutionsta),C.CString(&tmpvar1[0])); _tmp1277 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1277)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  var str string
  if p := strings.IndexByte(_tmp1276,byte(0)); p < 0 {
    str = string(_tmp1276)
  } else {
    str = string(_tmp1276[:p])
  }
  return
}
func (self *Task) SolutionDef(whichsol Soltype) (isdef bool,err error) {
  if _tmp1278 := C.MSK_solutiondef(self.ptr(),C.MSKsoltypee(whichsol),&isdef); _tmp1278 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1278)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) SolutionSummary(whichstream Streamtype) (err error) {
  if _tmp1279 := C.MSK_solutionsummary(self.ptr(),C.MSKstreamtypee(whichstream)); _tmp1279 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1279)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) SolveWithBasis(transp bool,numnz int32,sub []int32,val []float64) (numnzout int32,err error) {
  var _tmp1282 *int32
  var _tmp1280 int32
  if _tmp1281 := C.MSK_getnumcon(self.ptr(),addr(_tmp1280)); _tmp1281 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp1281)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  if int64(len(sub)) < int64(_tmp1280) {
    err = &ArrayLengthError{fun:"SolveWithBasis",arg:"sub"}
    return
  }
  if sub != nil { _tmp1282 = (*int32)(&sub[0]) }
  var _tmp1285 *float64
  var _tmp1283 int32
  if _tmp1284 := C.MSK_getnumcon(self.ptr(),addr(_tmp1283)); _tmp1284 != 0 {
    lastcode,lastmsg = self.getlasterror(_tmp1284)
     err = MosekError{ code:lastcode,msg:lastmsg}
    return
  }
  if int64(len(val)) < int64(_tmp1283) {
    err = &ArrayLengthError{fun:"SolveWithBasis",arg:"val"}
    return
  }
  if val != nil { _tmp1285 = (*float64)(&val[0]) }
  if _tmp1286 := C.MSK_solvewithbasis(self.ptr(),C.int(transp),C.MSKint32t(numnz),_tmp1282,_tmp1285,&numnzout); _tmp1286 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1286)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) StrToConeType(str string) (conetype Conetype,err error) {
  _tmp1287 := C.CString(str)
  var _tmp1288 C.MSKconetypee
  if _tmp1289 := C.MSK_strtoconetype(self.ptr(),_tmp1287,&_tmp1288); _tmp1289 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1289)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  conetype = Conetype(_tmp1288)
  return
}
func (self *Task) StrToSk(str string) (sk Stakey,err error) {
  _tmp1290 := C.CString(str)
  var _tmp1291 C.MSKstakeye
  if _tmp1292 := C.MSK_strtosk(self.ptr(),_tmp1290,&_tmp1291); _tmp1292 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1292)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  sk = Stakey(_tmp1291)
  return
}
func (self *Task) Toconic() (err error) {
  if _tmp1293 := C.MSK_toconic(self.ptr()); _tmp1293 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1293)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) UpdateSolutionInfo(whichsol Soltype) (err error) {
  if _tmp1294 := C.MSK_updatesolutioninfo(self.ptr(),C.MSKsoltypee(whichsol)); _tmp1294 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1294)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) WriteBSolution(filename string,compress Compresstype) (err error) {
  _tmp1295 := C.CString(filename)
  if _tmp1296 := C.MSK_writebsolution(self.ptr(),_tmp1295,C.MSKcompresstypee(compress)); _tmp1296 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1296)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) WriteData(filename string) (err error) {
  _tmp1297 := C.CString(filename)
  if _tmp1298 := C.MSK_writedata(self.ptr(),_tmp1297); _tmp1298 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1298)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) WriteJsonSol(filename string) (err error) {
  _tmp1299 := C.CString(filename)
  if _tmp1300 := C.MSK_writejsonsol(self.ptr(),_tmp1299); _tmp1300 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1300)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) WriteParamFile(filename string) (err error) {
  _tmp1301 := C.CString(filename)
  if _tmp1302 := C.MSK_writeparamfile(self.ptr(),_tmp1301); _tmp1302 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1302)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) WriteSolution(whichsol Soltype,filename string) (err error) {
  _tmp1303 := C.CString(filename)
  if _tmp1304 := C.MSK_writesolution(self.ptr(),C.MSKsoltypee(whichsol),_tmp1303); _tmp1304 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1304)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) WriteSolutionFile(filename string) (err error) {
  _tmp1305 := C.CString(filename)
  if _tmp1306 := C.MSK_writesolutionfile(self.ptr(),_tmp1305); _tmp1306 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1306)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Task) WriteTask(filename string) (err error) {
  _tmp1308 := C.CString(filename)
  if _tmp1309 := C.MSK_writetask(self.ptr(),_tmp1308); _tmp1309 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1309)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) Axpy(n int32,alpha float64,x []float64,y []float64) (err error) {
  var _tmp1310 *float64
  if int64(len(x)) < int64(n) {
    err = &ArrayLengthError{fun:"Axpy",arg:"x"}
    return
  }
  if x != nil { _tmp1310 = (*float64)(&x[0]) }
  var _tmp1311 *float64
  if int64(len(y)) < int64(n) {
    err = &ArrayLengthError{fun:"Axpy",arg:"y"}
    return
  }
  if y != nil { _tmp1311 = (*float64)(&y[0]) }
  if _tmp1312 := C.MSK_axpy(self.ptr(),C.MSKint32t(n),C.MSKrealt(alpha),_tmp1310,_tmp1311); _tmp1312 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1312)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) CheckInAll() (err error) {
  if _tmp1313 := C.MSK_checkinall(self.ptr()); _tmp1313 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1313)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) CheckInLicense(feature Feature) (err error) {
  if _tmp1314 := C.MSK_checkinlicense(self.ptr(),C.MSKfeaturee(feature)); _tmp1314 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1314)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) CheckOutLicense(feature Feature) (err error) {
  if _tmp1315 := C.MSK_checkoutlicense(self.ptr(),C.MSKfeaturee(feature)); _tmp1315 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1315)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) Dot(n int32,x []float64,y []float64) (xty float64,err error) {
  var _tmp1316 *float64
  if int64(len(x)) < int64(n) {
    err = &ArrayLengthError{fun:"Dot",arg:"x"}
    return
  }
  if x != nil { _tmp1316 = (*float64)(&x[0]) }
  var _tmp1317 *float64
  if int64(len(y)) < int64(n) {
    err = &ArrayLengthError{fun:"Dot",arg:"y"}
    return
  }
  if y != nil { _tmp1317 = (*float64)(&y[0]) }
  if _tmp1318 := C.MSK_dot(self.ptr(),C.MSKint32t(n),_tmp1316,_tmp1317,&xty); _tmp1318 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1318)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) EchoIntro(longver int32) (err error) {
  if _tmp1319 := C.MSK_echointro(self.ptr(),C.MSKint32t(longver)); _tmp1319 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1319)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) Expirylicenses() (expiry int64,err error) {
  if _tmp1320 := C.MSK_expirylicenses(self.ptr(),&expiry); _tmp1320 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1320)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) Gemm(transa Transpose,transb Transpose,m int32,n int32,k int32,alpha float64,a []float64,b []float64,beta float64,c []float64) (err error) {
  var _tmp1321 *float64
  if int64(len(a)) < int64((m * k)) {
    err = &ArrayLengthError{fun:"Gemm",arg:"a"}
    return
  }
  if a != nil { _tmp1321 = (*float64)(&a[0]) }
  var _tmp1322 *float64
  if int64(len(b)) < int64((k * n)) {
    err = &ArrayLengthError{fun:"Gemm",arg:"b"}
    return
  }
  if b != nil { _tmp1322 = (*float64)(&b[0]) }
  var _tmp1323 *float64
  if int64(len(c)) < int64((m * n)) {
    err = &ArrayLengthError{fun:"Gemm",arg:"c"}
    return
  }
  if c != nil { _tmp1323 = (*float64)(&c[0]) }
  if _tmp1324 := C.MSK_gemm(self.ptr(),C.MSKtransposee(transa),C.MSKtransposee(transb),C.MSKint32t(m),C.MSKint32t(n),C.MSKint32t(k),C.MSKrealt(alpha),_tmp1321,_tmp1322,C.MSKrealt(beta),_tmp1323); _tmp1324 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1324)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) Gemv(transa Transpose,m int32,n int32,alpha float64,a []float64,x []float64,beta float64,y []float64) (err error) {
  var _tmp1325 *float64
  if int64(len(a)) < int64((n * m)) {
    err = &ArrayLengthError{fun:"Gemv",arg:"a"}
    return
  }
  if a != nil { _tmp1325 = (*float64)(&a[0]) }
  var _tmp1327 *float64
  var _tmp1326 int32
  if (transa == transpose_no) {
    _tmp1326 = n
  } else {
    _tmp1326 = m
  }
  if int64(len(x)) < int64(_tmp1326) {
    err = &ArrayLengthError{fun:"Gemv",arg:"x"}
    return
  }
  if x != nil { _tmp1327 = (*float64)(&x[0]) }
  var _tmp1329 *float64
  var _tmp1328 int32
  if (transa == transpose_no) {
    _tmp1328 = m
  } else {
    _tmp1328 = n
  }
  if int64(len(y)) < int64(_tmp1328) {
    err = &ArrayLengthError{fun:"Gemv",arg:"y"}
    return
  }
  if y != nil { _tmp1329 = (*float64)(&y[0]) }
  if _tmp1330 := C.MSK_gemv(self.ptr(),C.MSKtransposee(transa),C.MSKint32t(m),C.MSKint32t(n),C.MSKrealt(alpha),_tmp1325,_tmp1327,C.MSKrealt(beta),_tmp1329); _tmp1330 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1330)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func GetCodeDesc(code Rescode) (symname string,str string,err error) {
  _tmp1331 := make([]byte,max_str_len)
  _tmp1332 := make([]byte,max_str_len)
  if _tmp1333 := C.MSK_getcodedesc(C.MSKrescodee(code),C.CString(&tmpvar1[0]),C.CString(&tmpvar1[0])); _tmp1333 != 0 {
    err = &MosekError{ code:_tmp1333 }
    return
  }
  var symname string
  if p := strings.IndexByte(_tmp1331,byte(0)); p < 0 {
    symname = string(_tmp1331)
  } else {
    symname = string(_tmp1331[:p])
  }
  var str string
  if p := strings.IndexByte(_tmp1332,byte(0)); p < 0 {
    str = string(_tmp1332)
  } else {
    str = string(_tmp1332[:p])
  }
  return
}
func GetVersion() (major int32,minor int32,revision int32,err error) {
  if _tmp1334 := C.MSK_getversion(&major,&minor,&revision); _tmp1334 != 0 {
    err = &MosekError{ code:_tmp1334 }
    return
  }
  return
}
func LicenseCleanup() (err error) {
  if _tmp1335 := C.MSK_licensecleanup(); _tmp1335 != 0 {
    err = &MosekError{ code:_tmp1335 }
    return
  }
  return
}
func (self *Env) Linkfiletostream(whichstream Streamtype,filename string,append int32) (err error) {
  _tmp1336 := C.CString(filename)
  if _tmp1337 := C.MSK_linkfiletoenvstream(self.ptr(),C.MSKstreamtypee(whichstream),_tmp1336,C.MSKint32t(append)); _tmp1337 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1337)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) Potrf(uplo Uplo,n int32,a []float64) (err error) {
  var _tmp1338 *float64
  if int64(len(a)) < int64((n * n)) {
    err = &ArrayLengthError{fun:"Potrf",arg:"a"}
    return
  }
  if a != nil { _tmp1338 = (*float64)(&a[0]) }
  if _tmp1339 := C.MSK_potrf(self.ptr(),C.MSKuploe(uplo),C.MSKint32t(n),_tmp1338); _tmp1339 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1339)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) PutLicenseCode(code []int32) (err error) {
  var _tmp1340 *int32
  if int64(len(code)) < int64(license_buffer_length) {
    err = &ArrayLengthError{fun:"PutLicenseCode",arg:"code"}
    return
  }
  if code != nil { _tmp1340 = (*int32)(&code[0]) }
  if _tmp1341 := C.MSK_putlicensecode(self.ptr(),_tmp1340); _tmp1341 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1341)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) PutLicenseDebug(licdebug int32) (err error) {
  if _tmp1342 := C.MSK_putlicensedebug(self.ptr(),C.MSKint32t(licdebug)); _tmp1342 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1342)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) PutLicensePath(licensepath string) (err error) {
  _tmp1343 := C.CString(licensepath)
  if _tmp1344 := C.MSK_putlicensepath(self.ptr(),_tmp1343); _tmp1344 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1344)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) PutLicenseWait(licwait int32) (err error) {
  if _tmp1345 := C.MSK_putlicensewait(self.ptr(),C.MSKint32t(licwait)); _tmp1345 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1345)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) ResetExpiryLicenses() (err error) {
  if _tmp1346 := C.MSK_resetexpirylicenses(self.ptr()); _tmp1346 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1346)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) Syeig(uplo Uplo,n int32,a []float64) (w []float64,err error) {
  var _tmp1347 *float64
  if int64(len(a)) < int64((n * n)) {
    err = &ArrayLengthError{fun:"Syeig",arg:"a"}
    return
  }
  if a != nil { _tmp1347 = (*float64)(&a[0]) }
  var _tmp1348 *float64
  w := make([]float64,n)
  if len(w) > 0 { _tmp1348 = (*float64)(&n[0]) }
  if _tmp1349 := C.MSK_syeig(self.ptr(),C.MSKuploe(uplo),C.MSKint32t(n),_tmp1347,_tmp1348); _tmp1349 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1349)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) Syevd(uplo Uplo,n int32,a []float64) (w []float64,err error) {
  var _tmp1350 *float64
  if int64(len(a)) < int64((n * n)) {
    err = &ArrayLengthError{fun:"Syevd",arg:"a"}
    return
  }
  if a != nil { _tmp1350 = (*float64)(&a[0]) }
  var _tmp1351 *float64
  w := make([]float64,n)
  if len(w) > 0 { _tmp1351 = (*float64)(&n[0]) }
  if _tmp1352 := C.MSK_syevd(self.ptr(),C.MSKuploe(uplo),C.MSKint32t(n),_tmp1350,_tmp1351); _tmp1352 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1352)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}
func (self *Env) Syrk(uplo Uplo,trans Transpose,n int32,k int32,alpha float64,a []float64,beta float64,c []float64) (err error) {
  var _tmp1353 *float64
  if int64(len(a)) < int64((n * k)) {
    err = &ArrayLengthError{fun:"Syrk",arg:"a"}
    return
  }
  if a != nil { _tmp1353 = (*float64)(&a[0]) }
  var _tmp1354 *float64
  if int64(len(c)) < int64((n * n)) {
    err = &ArrayLengthError{fun:"Syrk",arg:"c"}
    return
  }
  if c != nil { _tmp1354 = (*float64)(&c[0]) }
  if _tmp1355 := C.MSK_syrk(self.ptr(),C.MSKuploe(uplo),C.MSKtransposee(trans),C.MSKint32t(n),C.MSKint32t(k),C.MSKrealt(alpha),_tmp1353,C.MSKrealt(beta),_tmp1354); _tmp1355 != 0 {
    lastcode,lastmsg := self.getlasterror(_tmp1355)
    err = &MosekError{code:lastcode,msg:lastmsg}
    return
  }
  return
}

