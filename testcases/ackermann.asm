
ackermann.out:     file format elf64-littleriscv


Disassembly of section .text:

00000000000100b0 <register_fini>:
   100b0:	ffff0797          	auipc	a5,0xffff0
   100b4:	f5078793          	addi	a5,a5,-176 # 0 <register_fini-0x100b0>
   100b8:	cb89                	beqz	a5,100ca <register_fini+0x1a>
   100ba:	00000517          	auipc	a0,0x0
   100be:	22e50513          	addi	a0,a0,558 # 102e8 <__libc_fini_array>
   100c2:	00000317          	auipc	t1,0x0
   100c6:	1ea30067          	jr	490(t1) # 102ac <atexit>
   100ca:	8082                	ret

00000000000100cc <_start>:
   100cc:	00002197          	auipc	gp,0x2
   100d0:	d0418193          	addi	gp,gp,-764 # 11dd0 <__global_pointer$>
   100d4:	00002517          	auipc	a0,0x2
   100d8:	c5c50513          	addi	a0,a0,-932 # 11d30 <_edata>
   100dc:	00002617          	auipc	a2,0x2
   100e0:	c9460613          	addi	a2,a2,-876 # 11d70 <__BSS_END__>
   100e4:	8e09                	sub	a2,a2,a0
   100e6:	4581                	li	a1,0
   100e8:	00000097          	auipc	ra,0x0
   100ec:	2a2080e7          	jalr	674(ra) # 1038a <memset>
   100f0:	00000517          	auipc	a0,0x0
   100f4:	1f850513          	addi	a0,a0,504 # 102e8 <__libc_fini_array>
   100f8:	00000097          	auipc	ra,0x0
   100fc:	1b4080e7          	jalr	436(ra) # 102ac <atexit>
   10100:	00000097          	auipc	ra,0x0
   10104:	220080e7          	jalr	544(ra) # 10320 <__libc_init_array>
   10108:	4502                	lw	a0,0(sp)
   1010a:	002c                	addi	a1,sp,8
   1010c:	4601                	li	a2,0
   1010e:	00000097          	auipc	ra,0x0
   10112:	14a080e7          	jalr	330(ra) # 10258 <main>
   10116:	00000317          	auipc	t1,0x0
   1011a:	1a630067          	jr	422(t1) # 102bc <exit>

000000000001011e <__do_global_dtors_aux>:
   1011e:	00002797          	auipc	a5,0x2
   10122:	c127c783          	lbu	a5,-1006(a5) # 11d30 <_edata>
   10126:	ef95                	bnez	a5,10162 <__do_global_dtors_aux+0x44>
   10128:	ffff0797          	auipc	a5,0xffff0
   1012c:	ed878793          	addi	a5,a5,-296 # 0 <register_fini-0x100b0>
   10130:	c39d                	beqz	a5,10156 <__do_global_dtors_aux+0x38>
   10132:	1141                	addi	sp,sp,-16
   10134:	00001517          	auipc	a0,0x1
   10138:	48050513          	addi	a0,a0,1152 # 115b4 <__FRAME_END__>
   1013c:	e406                	sd	ra,8(sp)
   1013e:	00000097          	auipc	ra,0x0
   10142:	000000e7          	jalr	zero # 0 <register_fini-0x100b0>
   10146:	60a2                	ld	ra,8(sp)
   10148:	4785                	li	a5,1
   1014a:	00002717          	auipc	a4,0x2
   1014e:	bef70323          	sb	a5,-1050(a4) # 11d30 <_edata>
   10152:	0141                	addi	sp,sp,16
   10154:	8082                	ret
   10156:	4785                	li	a5,1
   10158:	00002717          	auipc	a4,0x2
   1015c:	bcf70c23          	sb	a5,-1064(a4) # 11d30 <_edata>
   10160:	8082                	ret
   10162:	8082                	ret

0000000000010164 <frame_dummy>:
   10164:	ffff0797          	auipc	a5,0xffff0
   10168:	e9c78793          	addi	a5,a5,-356 # 0 <register_fini-0x100b0>
   1016c:	cf89                	beqz	a5,10186 <frame_dummy+0x22>
   1016e:	00002597          	auipc	a1,0x2
   10172:	bca58593          	addi	a1,a1,-1078 # 11d38 <object.5475>
   10176:	00001517          	auipc	a0,0x1
   1017a:	43e50513          	addi	a0,a0,1086 # 115b4 <__FRAME_END__>
   1017e:	00000317          	auipc	t1,0x0
   10182:	00000067          	jr	zero # 0 <register_fini-0x100b0>
   10186:	8082                	ret

0000000000010188 <ackermann>:
#include <stdlib.h>

int ackermann(int m, int n){
   10188:	fd010113          	addi	sp,sp,-48
   1018c:	02113423          	sd	ra,40(sp)
   10190:	02813023          	sd	s0,32(sp)
   10194:	00913c23          	sd	s1,24(sp)
   10198:	03010413          	addi	s0,sp,48
   1019c:	00050793          	mv	a5,a0
   101a0:	00058713          	mv	a4,a1
   101a4:	fcf42e23          	sw	a5,-36(s0)
   101a8:	00070793          	mv	a5,a4
   101ac:	fcf42c23          	sw	a5,-40(s0)
  if (m == 0) return n+1;
   101b0:	fdc42783          	lw	a5,-36(s0)
   101b4:	0007879b          	sext.w	a5,a5
   101b8:	00079a63          	bnez	a5,101cc <ackermann+0x44>
   101bc:	fd842783          	lw	a5,-40(s0)
   101c0:	0017879b          	addiw	a5,a5,1
   101c4:	0007879b          	sext.w	a5,a5
   101c8:	0780006f          	j	10240 <ackermann+0xb8>
  if (n == 0) return ackermann( m - 1, 1 );
   101cc:	fd842783          	lw	a5,-40(s0)
   101d0:	0007879b          	sext.w	a5,a5
   101d4:	02079463          	bnez	a5,101fc <ackermann+0x74>
   101d8:	fdc42783          	lw	a5,-36(s0)
   101dc:	fff7879b          	addiw	a5,a5,-1
   101e0:	0007879b          	sext.w	a5,a5
   101e4:	00100593          	li	a1,1
   101e8:	00078513          	mv	a0,a5
   101ec:	00000097          	auipc	ra,0x0
   101f0:	f9c080e7          	jalr	-100(ra) # 10188 <ackermann>
   101f4:	00050793          	mv	a5,a0
   101f8:	0480006f          	j	10240 <ackermann+0xb8>
  return ackermann( m - 1, ackermann( m, n - 1 ) );
   101fc:	fdc42783          	lw	a5,-36(s0)
   10200:	fff7879b          	addiw	a5,a5,-1
   10204:	0007849b          	sext.w	s1,a5
   10208:	fd842783          	lw	a5,-40(s0)
   1020c:	fff7879b          	addiw	a5,a5,-1
   10210:	0007871b          	sext.w	a4,a5
   10214:	fdc42783          	lw	a5,-36(s0)
   10218:	00070593          	mv	a1,a4
   1021c:	00078513          	mv	a0,a5
   10220:	00000097          	auipc	ra,0x0
   10224:	f68080e7          	jalr	-152(ra) # 10188 <ackermann>
   10228:	00050793          	mv	a5,a0
   1022c:	00078593          	mv	a1,a5
   10230:	00048513          	mv	a0,s1
   10234:	00000097          	auipc	ra,0x0
   10238:	f54080e7          	jalr	-172(ra) # 10188 <ackermann>
   1023c:	00050793          	mv	a5,a0
}
   10240:	00078513          	mv	a0,a5
   10244:	02813083          	ld	ra,40(sp)
   10248:	02013403          	ld	s0,32(sp)
   1024c:	01813483          	ld	s1,24(sp)
   10250:	03010113          	addi	sp,sp,48
   10254:	00008067          	ret

0000000000010258 <main>:

int result;

int main(int argc, const char** argv) {
   10258:	fe010113          	addi	sp,sp,-32
   1025c:	00113c23          	sd	ra,24(sp)
   10260:	00813823          	sd	s0,16(sp)
   10264:	02010413          	addi	s0,sp,32
   10268:	00050793          	mv	a5,a0
   1026c:	feb43023          	sd	a1,-32(s0)
   10270:	fef42623          	sw	a5,-20(s0)

	result = ackermann(3, 2);
   10274:	00200593          	li	a1,2
   10278:	00300513          	li	a0,3
   1027c:	00000097          	auipc	ra,0x0
   10280:	f0c080e7          	jalr	-244(ra) # 10188 <ackermann>
   10284:	00050793          	mv	a5,a0
   10288:	00078713          	mv	a4,a5
   1028c:	000127b7          	lui	a5,0x12
   10290:	d6e7a423          	sw	a4,-664(a5) # 11d68 <result>

    return 0;
   10294:	00000793          	li	a5,0
}
   10298:	00078513          	mv	a0,a5
   1029c:	01813083          	ld	ra,24(sp)
   102a0:	01013403          	ld	s0,16(sp)
   102a4:	02010113          	addi	sp,sp,32
   102a8:	00008067          	ret

00000000000102ac <atexit>:
   102ac:	85aa                	mv	a1,a0
   102ae:	4681                	li	a3,0
   102b0:	4601                	li	a2,0
   102b2:	4501                	li	a0,0
   102b4:	00000317          	auipc	t1,0x0
   102b8:	18030067          	jr	384(t1) # 10434 <__register_exitproc>

00000000000102bc <exit>:
   102bc:	1141                	addi	sp,sp,-16
   102be:	4581                	li	a1,0
   102c0:	e022                	sd	s0,0(sp)
   102c2:	e406                	sd	ra,8(sp)
   102c4:	842a                	mv	s0,a0
   102c6:	00000097          	auipc	ra,0x0
   102ca:	1ea080e7          	jalr	490(ra) # 104b0 <__call_exitprocs>
   102ce:	00002797          	auipc	a5,0x2
   102d2:	a4a78793          	addi	a5,a5,-1462 # 11d18 <_global_impure_ptr>
   102d6:	6388                	ld	a0,0(a5)
   102d8:	6d3c                	ld	a5,88(a0)
   102da:	c391                	beqz	a5,102de <exit+0x22>
   102dc:	9782                	jalr	a5
   102de:	8522                	mv	a0,s0
   102e0:	00000097          	auipc	ra,0x0
   102e4:	296080e7          	jalr	662(ra) # 10576 <_exit>

00000000000102e8 <__libc_fini_array>:
   102e8:	1101                	addi	sp,sp,-32
   102ea:	e822                	sd	s0,16(sp)
   102ec:	00001797          	auipc	a5,0x1
   102f0:	2e478793          	addi	a5,a5,740 # 115d0 <__fini_array_end>
   102f4:	00001417          	auipc	s0,0x1
   102f8:	2d440413          	addi	s0,s0,724 # 115c8 <__init_array_end>
   102fc:	8f81                	sub	a5,a5,s0
   102fe:	e426                	sd	s1,8(sp)
   10300:	ec06                	sd	ra,24(sp)
   10302:	4037d493          	srai	s1,a5,0x3
   10306:	c881                	beqz	s1,10316 <__libc_fini_array+0x2e>
   10308:	17e1                	addi	a5,a5,-8
   1030a:	943e                	add	s0,s0,a5
   1030c:	601c                	ld	a5,0(s0)
   1030e:	14fd                	addi	s1,s1,-1
   10310:	1461                	addi	s0,s0,-8
   10312:	9782                	jalr	a5
   10314:	fce5                	bnez	s1,1030c <__libc_fini_array+0x24>
   10316:	60e2                	ld	ra,24(sp)
   10318:	6442                	ld	s0,16(sp)
   1031a:	64a2                	ld	s1,8(sp)
   1031c:	6105                	addi	sp,sp,32
   1031e:	8082                	ret

0000000000010320 <__libc_init_array>:
   10320:	1101                	addi	sp,sp,-32
   10322:	e822                	sd	s0,16(sp)
   10324:	e04a                	sd	s2,0(sp)
   10326:	00001417          	auipc	s0,0x1
   1032a:	29240413          	addi	s0,s0,658 # 115b8 <__init_array_start>
   1032e:	00001917          	auipc	s2,0x1
   10332:	28a90913          	addi	s2,s2,650 # 115b8 <__init_array_start>
   10336:	40890933          	sub	s2,s2,s0
   1033a:	ec06                	sd	ra,24(sp)
   1033c:	e426                	sd	s1,8(sp)
   1033e:	40395913          	srai	s2,s2,0x3
   10342:	00090963          	beqz	s2,10354 <__libc_init_array+0x34>
   10346:	4481                	li	s1,0
   10348:	601c                	ld	a5,0(s0)
   1034a:	0485                	addi	s1,s1,1
   1034c:	0421                	addi	s0,s0,8
   1034e:	9782                	jalr	a5
   10350:	fe991ce3          	bne	s2,s1,10348 <__libc_init_array+0x28>
   10354:	00001417          	auipc	s0,0x1
   10358:	26440413          	addi	s0,s0,612 # 115b8 <__init_array_start>
   1035c:	00001917          	auipc	s2,0x1
   10360:	26c90913          	addi	s2,s2,620 # 115c8 <__init_array_end>
   10364:	40890933          	sub	s2,s2,s0
   10368:	40395913          	srai	s2,s2,0x3
   1036c:	00090963          	beqz	s2,1037e <__libc_init_array+0x5e>
   10370:	4481                	li	s1,0
   10372:	601c                	ld	a5,0(s0)
   10374:	0485                	addi	s1,s1,1
   10376:	0421                	addi	s0,s0,8
   10378:	9782                	jalr	a5
   1037a:	fe991ce3          	bne	s2,s1,10372 <__libc_init_array+0x52>
   1037e:	60e2                	ld	ra,24(sp)
   10380:	6442                	ld	s0,16(sp)
   10382:	64a2                	ld	s1,8(sp)
   10384:	6902                	ld	s2,0(sp)
   10386:	6105                	addi	sp,sp,32
   10388:	8082                	ret

000000000001038a <memset>:
   1038a:	433d                	li	t1,15
   1038c:	872a                	mv	a4,a0
   1038e:	02c37163          	bgeu	t1,a2,103b0 <memset+0x26>
   10392:	00f77793          	andi	a5,a4,15
   10396:	e3c1                	bnez	a5,10416 <memset+0x8c>
   10398:	e1bd                	bnez	a1,103fe <memset+0x74>
   1039a:	ff067693          	andi	a3,a2,-16
   1039e:	8a3d                	andi	a2,a2,15
   103a0:	96ba                	add	a3,a3,a4
   103a2:	e30c                	sd	a1,0(a4)
   103a4:	e70c                	sd	a1,8(a4)
   103a6:	0741                	addi	a4,a4,16
   103a8:	fed76de3          	bltu	a4,a3,103a2 <memset+0x18>
   103ac:	e211                	bnez	a2,103b0 <memset+0x26>
   103ae:	8082                	ret
   103b0:	40c306b3          	sub	a3,t1,a2
   103b4:	068a                	slli	a3,a3,0x2
   103b6:	00000297          	auipc	t0,0x0
   103ba:	9696                	add	a3,a3,t0
   103bc:	00a68067          	jr	10(a3)
   103c0:	00b70723          	sb	a1,14(a4)
   103c4:	00b706a3          	sb	a1,13(a4)
   103c8:	00b70623          	sb	a1,12(a4)
   103cc:	00b705a3          	sb	a1,11(a4)
   103d0:	00b70523          	sb	a1,10(a4)
   103d4:	00b704a3          	sb	a1,9(a4)
   103d8:	00b70423          	sb	a1,8(a4)
   103dc:	00b703a3          	sb	a1,7(a4)
   103e0:	00b70323          	sb	a1,6(a4)
   103e4:	00b702a3          	sb	a1,5(a4)
   103e8:	00b70223          	sb	a1,4(a4)
   103ec:	00b701a3          	sb	a1,3(a4)
   103f0:	00b70123          	sb	a1,2(a4)
   103f4:	00b700a3          	sb	a1,1(a4)
   103f8:	00b70023          	sb	a1,0(a4)
   103fc:	8082                	ret
   103fe:	0ff5f593          	andi	a1,a1,255
   10402:	00859693          	slli	a3,a1,0x8
   10406:	8dd5                	or	a1,a1,a3
   10408:	01059693          	slli	a3,a1,0x10
   1040c:	8dd5                	or	a1,a1,a3
   1040e:	02059693          	slli	a3,a1,0x20
   10412:	8dd5                	or	a1,a1,a3
   10414:	b759                	j	1039a <memset+0x10>
   10416:	00279693          	slli	a3,a5,0x2
   1041a:	00000297          	auipc	t0,0x0
   1041e:	9696                	add	a3,a3,t0
   10420:	8286                	mv	t0,ra
   10422:	fa2680e7          	jalr	-94(a3)
   10426:	8096                	mv	ra,t0
   10428:	17c1                	addi	a5,a5,-16
   1042a:	8f1d                	sub	a4,a4,a5
   1042c:	963e                	add	a2,a2,a5
   1042e:	f8c371e3          	bgeu	t1,a2,103b0 <memset+0x26>
   10432:	b79d                	j	10398 <memset+0xe>

0000000000010434 <__register_exitproc>:
   10434:	00002797          	auipc	a5,0x2
   10438:	8e478793          	addi	a5,a5,-1820 # 11d18 <_global_impure_ptr>
   1043c:	6398                	ld	a4,0(a5)
   1043e:	1f873783          	ld	a5,504(a4)
   10442:	c3b1                	beqz	a5,10486 <__register_exitproc+0x52>
   10444:	4798                	lw	a4,8(a5)
   10446:	487d                	li	a6,31
   10448:	06e84263          	blt	a6,a4,104ac <__register_exitproc+0x78>
   1044c:	c505                	beqz	a0,10474 <__register_exitproc+0x40>
   1044e:	00371813          	slli	a6,a4,0x3
   10452:	983e                	add	a6,a6,a5
   10454:	10c83823          	sd	a2,272(a6)
   10458:	3107a883          	lw	a7,784(a5)
   1045c:	4605                	li	a2,1
   1045e:	00e6163b          	sllw	a2,a2,a4
   10462:	00c8e8b3          	or	a7,a7,a2
   10466:	3117a823          	sw	a7,784(a5)
   1046a:	20d83823          	sd	a3,528(a6)
   1046e:	4689                	li	a3,2
   10470:	02d50063          	beq	a0,a3,10490 <__register_exitproc+0x5c>
   10474:	00270693          	addi	a3,a4,2
   10478:	068e                	slli	a3,a3,0x3
   1047a:	2705                	addiw	a4,a4,1
   1047c:	c798                	sw	a4,8(a5)
   1047e:	97b6                	add	a5,a5,a3
   10480:	e38c                	sd	a1,0(a5)
   10482:	4501                	li	a0,0
   10484:	8082                	ret
   10486:	20070793          	addi	a5,a4,512
   1048a:	1ef73c23          	sd	a5,504(a4)
   1048e:	bf5d                	j	10444 <__register_exitproc+0x10>
   10490:	3147a683          	lw	a3,788(a5)
   10494:	4501                	li	a0,0
   10496:	8e55                	or	a2,a2,a3
   10498:	00270693          	addi	a3,a4,2
   1049c:	068e                	slli	a3,a3,0x3
   1049e:	2705                	addiw	a4,a4,1
   104a0:	30c7aa23          	sw	a2,788(a5)
   104a4:	c798                	sw	a4,8(a5)
   104a6:	97b6                	add	a5,a5,a3
   104a8:	e38c                	sd	a1,0(a5)
   104aa:	8082                	ret
   104ac:	557d                	li	a0,-1
   104ae:	8082                	ret

00000000000104b0 <__call_exitprocs>:
   104b0:	715d                	addi	sp,sp,-80
   104b2:	00002797          	auipc	a5,0x2
   104b6:	86678793          	addi	a5,a5,-1946 # 11d18 <_global_impure_ptr>
   104ba:	e062                	sd	s8,0(sp)
   104bc:	0007bc03          	ld	s8,0(a5)
   104c0:	f44e                	sd	s3,40(sp)
   104c2:	f052                	sd	s4,32(sp)
   104c4:	ec56                	sd	s5,24(sp)
   104c6:	e85a                	sd	s6,16(sp)
   104c8:	e486                	sd	ra,72(sp)
   104ca:	e0a2                	sd	s0,64(sp)
   104cc:	fc26                	sd	s1,56(sp)
   104ce:	f84a                	sd	s2,48(sp)
   104d0:	e45e                	sd	s7,8(sp)
   104d2:	8aaa                	mv	s5,a0
   104d4:	8b2e                	mv	s6,a1
   104d6:	4a05                	li	s4,1
   104d8:	59fd                	li	s3,-1
   104da:	1f8c3903          	ld	s2,504(s8)
   104de:	02090463          	beqz	s2,10506 <__call_exitprocs+0x56>
   104e2:	00892483          	lw	s1,8(s2)
   104e6:	fff4841b          	addiw	s0,s1,-1
   104ea:	00044e63          	bltz	s0,10506 <__call_exitprocs+0x56>
   104ee:	048e                	slli	s1,s1,0x3
   104f0:	94ca                	add	s1,s1,s2
   104f2:	020b0663          	beqz	s6,1051e <__call_exitprocs+0x6e>
   104f6:	2084b783          	ld	a5,520(s1)
   104fa:	03678263          	beq	a5,s6,1051e <__call_exitprocs+0x6e>
   104fe:	347d                	addiw	s0,s0,-1
   10500:	14e1                	addi	s1,s1,-8
   10502:	ff3418e3          	bne	s0,s3,104f2 <__call_exitprocs+0x42>
   10506:	60a6                	ld	ra,72(sp)
   10508:	6406                	ld	s0,64(sp)
   1050a:	74e2                	ld	s1,56(sp)
   1050c:	7942                	ld	s2,48(sp)
   1050e:	79a2                	ld	s3,40(sp)
   10510:	7a02                	ld	s4,32(sp)
   10512:	6ae2                	ld	s5,24(sp)
   10514:	6b42                	ld	s6,16(sp)
   10516:	6ba2                	ld	s7,8(sp)
   10518:	6c02                	ld	s8,0(sp)
   1051a:	6161                	addi	sp,sp,80
   1051c:	8082                	ret
   1051e:	00892783          	lw	a5,8(s2)
   10522:	6498                	ld	a4,8(s1)
   10524:	37fd                	addiw	a5,a5,-1
   10526:	04878263          	beq	a5,s0,1056a <__call_exitprocs+0xba>
   1052a:	0004b423          	sd	zero,8(s1)
   1052e:	db61                	beqz	a4,104fe <__call_exitprocs+0x4e>
   10530:	31092783          	lw	a5,784(s2)
   10534:	008a16bb          	sllw	a3,s4,s0
   10538:	00892b83          	lw	s7,8(s2)
   1053c:	8ff5                	and	a5,a5,a3
   1053e:	2781                	sext.w	a5,a5
   10540:	eb99                	bnez	a5,10556 <__call_exitprocs+0xa6>
   10542:	9702                	jalr	a4
   10544:	00892783          	lw	a5,8(s2)
   10548:	f97799e3          	bne	a5,s7,104da <__call_exitprocs+0x2a>
   1054c:	1f8c3783          	ld	a5,504(s8)
   10550:	fb2787e3          	beq	a5,s2,104fe <__call_exitprocs+0x4e>
   10554:	b759                	j	104da <__call_exitprocs+0x2a>
   10556:	31492783          	lw	a5,788(s2)
   1055a:	1084b583          	ld	a1,264(s1)
   1055e:	8ff5                	and	a5,a5,a3
   10560:	2781                	sext.w	a5,a5
   10562:	e799                	bnez	a5,10570 <__call_exitprocs+0xc0>
   10564:	8556                	mv	a0,s5
   10566:	9702                	jalr	a4
   10568:	bff1                	j	10544 <__call_exitprocs+0x94>
   1056a:	00892423          	sw	s0,8(s2)
   1056e:	b7c1                	j	1052e <__call_exitprocs+0x7e>
   10570:	852e                	mv	a0,a1
   10572:	9702                	jalr	a4
   10574:	bfc1                	j	10544 <__call_exitprocs+0x94>

0000000000010576 <_exit>:
   10576:	4581                	li	a1,0
   10578:	4601                	li	a2,0
   1057a:	4681                	li	a3,0
   1057c:	4701                	li	a4,0
   1057e:	4781                	li	a5,0
   10580:	05d00893          	li	a7,93
   10584:	00000073          	ecall
   10588:	00054363          	bltz	a0,1058e <_exit+0x18>
   1058c:	a001                	j	1058c <_exit+0x16>
   1058e:	1141                	addi	sp,sp,-16
   10590:	e022                	sd	s0,0(sp)
   10592:	842a                	mv	s0,a0
   10594:	e406                	sd	ra,8(sp)
   10596:	4080043b          	negw	s0,s0
   1059a:	00000097          	auipc	ra,0x0
   1059e:	00c080e7          	jalr	12(ra) # 105a6 <__errno>
   105a2:	c100                	sw	s0,0(a0)
   105a4:	a001                	j	105a4 <_exit+0x2e>

00000000000105a6 <__errno>:
   105a6:	00001797          	auipc	a5,0x1
   105aa:	78278793          	addi	a5,a5,1922 # 11d28 <_impure_ptr>
   105ae:	6388                	ld	a0,0(a5)
   105b0:	8082                	ret
