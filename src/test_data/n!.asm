
n!.out:     file format elf64-littleriscv


Disassembly of section .text:

00000000000100b0 <register_fini>:
   100b0:	ffff0797          	auipc	a5,0xffff0
   100b4:	f5078793          	addi	a5,a5,-176 # 0 <register_fini-0x100b0>
   100b8:	cb89                	beqz	a5,100ca <register_fini+0x1a>
   100ba:	00000517          	auipc	a0,0x0
   100be:	1be50513          	addi	a0,a0,446 # 10278 <__libc_fini_array>
   100c2:	00000317          	auipc	t1,0x0
   100c6:	17a30067          	jr	378(t1) # 1023c <atexit>
   100ca:	8082                	ret

00000000000100cc <_start>:
   100cc:	00002197          	auipc	gp,0x2
   100d0:	c9418193          	addi	gp,gp,-876 # 11d60 <__global_pointer$>
   100d4:	00002517          	auipc	a0,0x2
   100d8:	bec50513          	addi	a0,a0,-1044 # 11cc0 <_edata>
   100dc:	00002617          	auipc	a2,0x2
   100e0:	c2460613          	addi	a2,a2,-988 # 11d00 <__BSS_END__>
   100e4:	8e09                	sub	a2,a2,a0
   100e6:	4581                	li	a1,0
   100e8:	00000097          	auipc	ra,0x0
   100ec:	232080e7          	jalr	562(ra) # 1031a <memset>
   100f0:	00000517          	auipc	a0,0x0
   100f4:	18850513          	addi	a0,a0,392 # 10278 <__libc_fini_array>
   100f8:	00000097          	auipc	ra,0x0
   100fc:	144080e7          	jalr	324(ra) # 1023c <atexit>
   10100:	00000097          	auipc	ra,0x0
   10104:	1b0080e7          	jalr	432(ra) # 102b0 <__libc_init_array>
   10108:	4502                	lw	a0,0(sp)
   1010a:	002c                	addi	a1,sp,8
   1010c:	4601                	li	a2,0
   1010e:	00000097          	auipc	ra,0x0
   10112:	0ea080e7          	jalr	234(ra) # 101f8 <main>
   10116:	00000317          	auipc	t1,0x0
   1011a:	13630067          	jr	310(t1) # 1024c <exit>

000000000001011e <__do_global_dtors_aux>:
   1011e:	00002797          	auipc	a5,0x2
   10122:	baa7c783          	lbu	a5,-1110(a5) # 11cc8 <completed.5470>
   10126:	ef95                	bnez	a5,10162 <__do_global_dtors_aux+0x44>
   10128:	ffff0797          	auipc	a5,0xffff0
   1012c:	ed878793          	addi	a5,a5,-296 # 0 <register_fini-0x100b0>
   10130:	c39d                	beqz	a5,10156 <__do_global_dtors_aux+0x38>
   10132:	1141                	addi	sp,sp,-16
   10134:	00001517          	auipc	a0,0x1
   10138:	41050513          	addi	a0,a0,1040 # 11544 <__FRAME_END__>
   1013c:	e406                	sd	ra,8(sp)
   1013e:	00000097          	auipc	ra,0x0
   10142:	000000e7          	jalr	zero # 0 <register_fini-0x100b0>
   10146:	60a2                	ld	ra,8(sp)
   10148:	4785                	li	a5,1
   1014a:	00002717          	auipc	a4,0x2
   1014e:	b6f70f23          	sb	a5,-1154(a4) # 11cc8 <completed.5470>
   10152:	0141                	addi	sp,sp,16
   10154:	8082                	ret
   10156:	4785                	li	a5,1
   10158:	00002717          	auipc	a4,0x2
   1015c:	b6f70823          	sb	a5,-1168(a4) # 11cc8 <completed.5470>
   10160:	8082                	ret
   10162:	8082                	ret

0000000000010164 <frame_dummy>:
   10164:	ffff0797          	auipc	a5,0xffff0
   10168:	e9c78793          	addi	a5,a5,-356 # 0 <register_fini-0x100b0>
   1016c:	cf89                	beqz	a5,10186 <frame_dummy+0x22>
   1016e:	00002597          	auipc	a1,0x2
   10172:	b6258593          	addi	a1,a1,-1182 # 11cd0 <object.5475>
   10176:	00001517          	auipc	a0,0x1
   1017a:	3ce50513          	addi	a0,a0,974 # 11544 <__FRAME_END__>
   1017e:	00000317          	auipc	t1,0x0
   10182:	00000067          	jr	zero # 0 <register_fini-0x100b0>
   10186:	8082                	ret

0000000000010188 <cal_n>:
#include <stdio.h>

int result=0;
//result 3628800
int cal_n(int i)
{
   10188:	fe010113          	addi	sp,sp,-32
   1018c:	00113c23          	sd	ra,24(sp)
   10190:	00813823          	sd	s0,16(sp)
   10194:	02010413          	addi	s0,sp,32
   10198:	00050793          	mv	a5,a0
   1019c:	fef42623          	sw	a5,-20(s0)
	if(i==1)
   101a0:	fec42783          	lw	a5,-20(s0)
   101a4:	0007871b          	sext.w	a4,a5
   101a8:	00100793          	li	a5,1
   101ac:	00f71663          	bne	a4,a5,101b8 <cal_n+0x30>
		return i;
   101b0:	fec42783          	lw	a5,-20(s0)
   101b4:	0300006f          	j	101e4 <cal_n+0x5c>
	else
		return i*cal_n(i-1);
   101b8:	fec42783          	lw	a5,-20(s0)
   101bc:	fff7879b          	addiw	a5,a5,-1
   101c0:	0007879b          	sext.w	a5,a5
   101c4:	00078513          	mv	a0,a5
   101c8:	00000097          	auipc	ra,0x0
   101cc:	fc0080e7          	jalr	-64(ra) # 10188 <cal_n>
   101d0:	00050793          	mv	a5,a0
   101d4:	00078713          	mv	a4,a5
   101d8:	fec42783          	lw	a5,-20(s0)
   101dc:	02e787bb          	mulw	a5,a5,a4
   101e0:	0007879b          	sext.w	a5,a5
}
   101e4:	00078513          	mv	a0,a5
   101e8:	01813083          	ld	ra,24(sp)
   101ec:	01013403          	ld	s0,16(sp)
   101f0:	02010113          	addi	sp,sp,32
   101f4:	00008067          	ret

00000000000101f8 <main>:

int main()  
{
   101f8:	ff010113          	addi	sp,sp,-16
   101fc:	00113423          	sd	ra,8(sp)
   10200:	00813023          	sd	s0,0(sp)
   10204:	01010413          	addi	s0,sp,16
	result=cal_n(10);
   10208:	00a00513          	li	a0,10
   1020c:	00000097          	auipc	ra,0x0
   10210:	f7c080e7          	jalr	-132(ra) # 10188 <cal_n>
   10214:	00050793          	mv	a5,a0
   10218:	00078713          	mv	a4,a5
   1021c:	000127b7          	lui	a5,0x12
   10220:	cce7a023          	sw	a4,-832(a5) # 11cc0 <_edata>
	return 0;
   10224:	00000793          	li	a5,0
}
   10228:	00078513          	mv	a0,a5
   1022c:	00813083          	ld	ra,8(sp)
   10230:	00013403          	ld	s0,0(sp)
   10234:	01010113          	addi	sp,sp,16
   10238:	00008067          	ret

000000000001023c <atexit>:
   1023c:	85aa                	mv	a1,a0
   1023e:	4681                	li	a3,0
   10240:	4601                	li	a2,0
   10242:	4501                	li	a0,0
   10244:	00000317          	auipc	t1,0x0
   10248:	18030067          	jr	384(t1) # 103c4 <__register_exitproc>

000000000001024c <exit>:
   1024c:	1141                	addi	sp,sp,-16
   1024e:	4581                	li	a1,0
   10250:	e022                	sd	s0,0(sp)
   10252:	e406                	sd	ra,8(sp)
   10254:	842a                	mv	s0,a0
   10256:	00000097          	auipc	ra,0x0
   1025a:	1ea080e7          	jalr	490(ra) # 10440 <__call_exitprocs>
   1025e:	00002797          	auipc	a5,0x2
   10262:	a4a78793          	addi	a5,a5,-1462 # 11ca8 <_global_impure_ptr>
   10266:	6388                	ld	a0,0(a5)
   10268:	6d3c                	ld	a5,88(a0)
   1026a:	c391                	beqz	a5,1026e <exit+0x22>
   1026c:	9782                	jalr	a5
   1026e:	8522                	mv	a0,s0
   10270:	00000097          	auipc	ra,0x0
   10274:	296080e7          	jalr	662(ra) # 10506 <_exit>

0000000000010278 <__libc_fini_array>:
   10278:	1101                	addi	sp,sp,-32
   1027a:	e822                	sd	s0,16(sp)
   1027c:	00001797          	auipc	a5,0x1
   10280:	2e478793          	addi	a5,a5,740 # 11560 <__fini_array_end>
   10284:	00001417          	auipc	s0,0x1
   10288:	2d440413          	addi	s0,s0,724 # 11558 <__init_array_end>
   1028c:	8f81                	sub	a5,a5,s0
   1028e:	e426                	sd	s1,8(sp)
   10290:	ec06                	sd	ra,24(sp)
   10292:	4037d493          	srai	s1,a5,0x3
   10296:	c881                	beqz	s1,102a6 <__libc_fini_array+0x2e>
   10298:	17e1                	addi	a5,a5,-8
   1029a:	943e                	add	s0,s0,a5
   1029c:	601c                	ld	a5,0(s0)
   1029e:	14fd                	addi	s1,s1,-1
   102a0:	1461                	addi	s0,s0,-8
   102a2:	9782                	jalr	a5
   102a4:	fce5                	bnez	s1,1029c <__libc_fini_array+0x24>
   102a6:	60e2                	ld	ra,24(sp)
   102a8:	6442                	ld	s0,16(sp)
   102aa:	64a2                	ld	s1,8(sp)
   102ac:	6105                	addi	sp,sp,32
   102ae:	8082                	ret

00000000000102b0 <__libc_init_array>:
   102b0:	1101                	addi	sp,sp,-32
   102b2:	e822                	sd	s0,16(sp)
   102b4:	e04a                	sd	s2,0(sp)
   102b6:	00001417          	auipc	s0,0x1
   102ba:	29240413          	addi	s0,s0,658 # 11548 <__init_array_start>
   102be:	00001917          	auipc	s2,0x1
   102c2:	28a90913          	addi	s2,s2,650 # 11548 <__init_array_start>
   102c6:	40890933          	sub	s2,s2,s0
   102ca:	ec06                	sd	ra,24(sp)
   102cc:	e426                	sd	s1,8(sp)
   102ce:	40395913          	srai	s2,s2,0x3
   102d2:	00090963          	beqz	s2,102e4 <__libc_init_array+0x34>
   102d6:	4481                	li	s1,0
   102d8:	601c                	ld	a5,0(s0)
   102da:	0485                	addi	s1,s1,1
   102dc:	0421                	addi	s0,s0,8
   102de:	9782                	jalr	a5
   102e0:	fe991ce3          	bne	s2,s1,102d8 <__libc_init_array+0x28>
   102e4:	00001417          	auipc	s0,0x1
   102e8:	26440413          	addi	s0,s0,612 # 11548 <__init_array_start>
   102ec:	00001917          	auipc	s2,0x1
   102f0:	26c90913          	addi	s2,s2,620 # 11558 <__init_array_end>
   102f4:	40890933          	sub	s2,s2,s0
   102f8:	40395913          	srai	s2,s2,0x3
   102fc:	00090963          	beqz	s2,1030e <__libc_init_array+0x5e>
   10300:	4481                	li	s1,0
   10302:	601c                	ld	a5,0(s0)
   10304:	0485                	addi	s1,s1,1
   10306:	0421                	addi	s0,s0,8
   10308:	9782                	jalr	a5
   1030a:	fe991ce3          	bne	s2,s1,10302 <__libc_init_array+0x52>
   1030e:	60e2                	ld	ra,24(sp)
   10310:	6442                	ld	s0,16(sp)
   10312:	64a2                	ld	s1,8(sp)
   10314:	6902                	ld	s2,0(sp)
   10316:	6105                	addi	sp,sp,32
   10318:	8082                	ret

000000000001031a <memset>:
   1031a:	433d                	li	t1,15
   1031c:	872a                	mv	a4,a0
   1031e:	02c37163          	bgeu	t1,a2,10340 <memset+0x26>
   10322:	00f77793          	andi	a5,a4,15
   10326:	e3c1                	bnez	a5,103a6 <memset+0x8c>
   10328:	e1bd                	bnez	a1,1038e <memset+0x74>
   1032a:	ff067693          	andi	a3,a2,-16
   1032e:	8a3d                	andi	a2,a2,15
   10330:	96ba                	add	a3,a3,a4
   10332:	e30c                	sd	a1,0(a4)
   10334:	e70c                	sd	a1,8(a4)
   10336:	0741                	addi	a4,a4,16
   10338:	fed76de3          	bltu	a4,a3,10332 <memset+0x18>
   1033c:	e211                	bnez	a2,10340 <memset+0x26>
   1033e:	8082                	ret
   10340:	40c306b3          	sub	a3,t1,a2
   10344:	068a                	slli	a3,a3,0x2
   10346:	00000297          	auipc	t0,0x0
   1034a:	9696                	add	a3,a3,t0
   1034c:	00a68067          	jr	10(a3)
   10350:	00b70723          	sb	a1,14(a4)
   10354:	00b706a3          	sb	a1,13(a4)
   10358:	00b70623          	sb	a1,12(a4)
   1035c:	00b705a3          	sb	a1,11(a4)
   10360:	00b70523          	sb	a1,10(a4)
   10364:	00b704a3          	sb	a1,9(a4)
   10368:	00b70423          	sb	a1,8(a4)
   1036c:	00b703a3          	sb	a1,7(a4)
   10370:	00b70323          	sb	a1,6(a4)
   10374:	00b702a3          	sb	a1,5(a4)
   10378:	00b70223          	sb	a1,4(a4)
   1037c:	00b701a3          	sb	a1,3(a4)
   10380:	00b70123          	sb	a1,2(a4)
   10384:	00b700a3          	sb	a1,1(a4)
   10388:	00b70023          	sb	a1,0(a4)
   1038c:	8082                	ret
   1038e:	0ff5f593          	andi	a1,a1,255
   10392:	00859693          	slli	a3,a1,0x8
   10396:	8dd5                	or	a1,a1,a3
   10398:	01059693          	slli	a3,a1,0x10
   1039c:	8dd5                	or	a1,a1,a3
   1039e:	02059693          	slli	a3,a1,0x20
   103a2:	8dd5                	or	a1,a1,a3
   103a4:	b759                	j	1032a <memset+0x10>
   103a6:	00279693          	slli	a3,a5,0x2
   103aa:	00000297          	auipc	t0,0x0
   103ae:	9696                	add	a3,a3,t0
   103b0:	8286                	mv	t0,ra
   103b2:	fa2680e7          	jalr	-94(a3)
   103b6:	8096                	mv	ra,t0
   103b8:	17c1                	addi	a5,a5,-16
   103ba:	8f1d                	sub	a4,a4,a5
   103bc:	963e                	add	a2,a2,a5
   103be:	f8c371e3          	bgeu	t1,a2,10340 <memset+0x26>
   103c2:	b79d                	j	10328 <memset+0xe>

00000000000103c4 <__register_exitproc>:
   103c4:	00002797          	auipc	a5,0x2
   103c8:	8e478793          	addi	a5,a5,-1820 # 11ca8 <_global_impure_ptr>
   103cc:	6398                	ld	a4,0(a5)
   103ce:	1f873783          	ld	a5,504(a4)
   103d2:	c3b1                	beqz	a5,10416 <__register_exitproc+0x52>
   103d4:	4798                	lw	a4,8(a5)
   103d6:	487d                	li	a6,31
   103d8:	06e84263          	blt	a6,a4,1043c <__register_exitproc+0x78>
   103dc:	c505                	beqz	a0,10404 <__register_exitproc+0x40>
   103de:	00371813          	slli	a6,a4,0x3
   103e2:	983e                	add	a6,a6,a5
   103e4:	10c83823          	sd	a2,272(a6)
   103e8:	3107a883          	lw	a7,784(a5)
   103ec:	4605                	li	a2,1
   103ee:	00e6163b          	sllw	a2,a2,a4
   103f2:	00c8e8b3          	or	a7,a7,a2
   103f6:	3117a823          	sw	a7,784(a5)
   103fa:	20d83823          	sd	a3,528(a6)
   103fe:	4689                	li	a3,2
   10400:	02d50063          	beq	a0,a3,10420 <__register_exitproc+0x5c>
   10404:	00270693          	addi	a3,a4,2
   10408:	068e                	slli	a3,a3,0x3
   1040a:	2705                	addiw	a4,a4,1
   1040c:	c798                	sw	a4,8(a5)
   1040e:	97b6                	add	a5,a5,a3
   10410:	e38c                	sd	a1,0(a5)
   10412:	4501                	li	a0,0
   10414:	8082                	ret
   10416:	20070793          	addi	a5,a4,512
   1041a:	1ef73c23          	sd	a5,504(a4)
   1041e:	bf5d                	j	103d4 <__register_exitproc+0x10>
   10420:	3147a683          	lw	a3,788(a5)
   10424:	4501                	li	a0,0
   10426:	8e55                	or	a2,a2,a3
   10428:	00270693          	addi	a3,a4,2
   1042c:	068e                	slli	a3,a3,0x3
   1042e:	2705                	addiw	a4,a4,1
   10430:	30c7aa23          	sw	a2,788(a5)
   10434:	c798                	sw	a4,8(a5)
   10436:	97b6                	add	a5,a5,a3
   10438:	e38c                	sd	a1,0(a5)
   1043a:	8082                	ret
   1043c:	557d                	li	a0,-1
   1043e:	8082                	ret

0000000000010440 <__call_exitprocs>:
   10440:	715d                	addi	sp,sp,-80
   10442:	00002797          	auipc	a5,0x2
   10446:	86678793          	addi	a5,a5,-1946 # 11ca8 <_global_impure_ptr>
   1044a:	e062                	sd	s8,0(sp)
   1044c:	0007bc03          	ld	s8,0(a5)
   10450:	f44e                	sd	s3,40(sp)
   10452:	f052                	sd	s4,32(sp)
   10454:	ec56                	sd	s5,24(sp)
   10456:	e85a                	sd	s6,16(sp)
   10458:	e486                	sd	ra,72(sp)
   1045a:	e0a2                	sd	s0,64(sp)
   1045c:	fc26                	sd	s1,56(sp)
   1045e:	f84a                	sd	s2,48(sp)
   10460:	e45e                	sd	s7,8(sp)
   10462:	8aaa                	mv	s5,a0
   10464:	8b2e                	mv	s6,a1
   10466:	4a05                	li	s4,1
   10468:	59fd                	li	s3,-1
   1046a:	1f8c3903          	ld	s2,504(s8)
   1046e:	02090463          	beqz	s2,10496 <__call_exitprocs+0x56>
   10472:	00892483          	lw	s1,8(s2)
   10476:	fff4841b          	addiw	s0,s1,-1
   1047a:	00044e63          	bltz	s0,10496 <__call_exitprocs+0x56>
   1047e:	048e                	slli	s1,s1,0x3
   10480:	94ca                	add	s1,s1,s2
   10482:	020b0663          	beqz	s6,104ae <__call_exitprocs+0x6e>
   10486:	2084b783          	ld	a5,520(s1)
   1048a:	03678263          	beq	a5,s6,104ae <__call_exitprocs+0x6e>
   1048e:	347d                	addiw	s0,s0,-1
   10490:	14e1                	addi	s1,s1,-8
   10492:	ff3418e3          	bne	s0,s3,10482 <__call_exitprocs+0x42>
   10496:	60a6                	ld	ra,72(sp)
   10498:	6406                	ld	s0,64(sp)
   1049a:	74e2                	ld	s1,56(sp)
   1049c:	7942                	ld	s2,48(sp)
   1049e:	79a2                	ld	s3,40(sp)
   104a0:	7a02                	ld	s4,32(sp)
   104a2:	6ae2                	ld	s5,24(sp)
   104a4:	6b42                	ld	s6,16(sp)
   104a6:	6ba2                	ld	s7,8(sp)
   104a8:	6c02                	ld	s8,0(sp)
   104aa:	6161                	addi	sp,sp,80
   104ac:	8082                	ret
   104ae:	00892783          	lw	a5,8(s2)
   104b2:	6498                	ld	a4,8(s1)
   104b4:	37fd                	addiw	a5,a5,-1
   104b6:	04878263          	beq	a5,s0,104fa <__call_exitprocs+0xba>
   104ba:	0004b423          	sd	zero,8(s1)
   104be:	db61                	beqz	a4,1048e <__call_exitprocs+0x4e>
   104c0:	31092783          	lw	a5,784(s2)
   104c4:	008a16bb          	sllw	a3,s4,s0
   104c8:	00892b83          	lw	s7,8(s2)
   104cc:	8ff5                	and	a5,a5,a3
   104ce:	2781                	sext.w	a5,a5
   104d0:	eb99                	bnez	a5,104e6 <__call_exitprocs+0xa6>
   104d2:	9702                	jalr	a4
   104d4:	00892783          	lw	a5,8(s2)
   104d8:	f97799e3          	bne	a5,s7,1046a <__call_exitprocs+0x2a>
   104dc:	1f8c3783          	ld	a5,504(s8)
   104e0:	fb2787e3          	beq	a5,s2,1048e <__call_exitprocs+0x4e>
   104e4:	b759                	j	1046a <__call_exitprocs+0x2a>
   104e6:	31492783          	lw	a5,788(s2)
   104ea:	1084b583          	ld	a1,264(s1)
   104ee:	8ff5                	and	a5,a5,a3
   104f0:	2781                	sext.w	a5,a5
   104f2:	e799                	bnez	a5,10500 <__call_exitprocs+0xc0>
   104f4:	8556                	mv	a0,s5
   104f6:	9702                	jalr	a4
   104f8:	bff1                	j	104d4 <__call_exitprocs+0x94>
   104fa:	00892423          	sw	s0,8(s2)
   104fe:	b7c1                	j	104be <__call_exitprocs+0x7e>
   10500:	852e                	mv	a0,a1
   10502:	9702                	jalr	a4
   10504:	bfc1                	j	104d4 <__call_exitprocs+0x94>

0000000000010506 <_exit>:
   10506:	4581                	li	a1,0
   10508:	4601                	li	a2,0
   1050a:	4681                	li	a3,0
   1050c:	4701                	li	a4,0
   1050e:	4781                	li	a5,0
   10510:	05d00893          	li	a7,93
   10514:	00000073          	ecall
   10518:	00054363          	bltz	a0,1051e <_exit+0x18>
   1051c:	a001                	j	1051c <_exit+0x16>
   1051e:	1141                	addi	sp,sp,-16
   10520:	e022                	sd	s0,0(sp)
   10522:	842a                	mv	s0,a0
   10524:	e406                	sd	ra,8(sp)
   10526:	4080043b          	negw	s0,s0
   1052a:	00000097          	auipc	ra,0x0
   1052e:	00c080e7          	jalr	12(ra) # 10536 <__errno>
   10532:	c100                	sw	s0,0(a0)
   10534:	a001                	j	10534 <_exit+0x2e>

0000000000010536 <__errno>:
   10536:	00001797          	auipc	a5,0x1
   1053a:	78278793          	addi	a5,a5,1922 # 11cb8 <_impure_ptr>
   1053e:	6388                	ld	a0,0(a5)
   10540:	8082                	ret
