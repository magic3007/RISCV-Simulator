
gemm.out:     file format elf64-littleriscv


Disassembly of section .text:

00000000000100b0 <register_fini>:
   100b0:	ffff0797          	auipc	a5,0xffff0
   100b4:	f5078793          	addi	a5,a5,-176 # 0 <register_fini-0x100b0>
   100b8:	cb89                	beqz	a5,100ca <register_fini+0x1a>
   100ba:	00000517          	auipc	a0,0x0
   100be:	37a50513          	addi	a0,a0,890 # 10434 <__libc_fini_array>
   100c2:	00000317          	auipc	t1,0x0
   100c6:	33630067          	jr	822(t1) # 103f8 <atexit>
   100ca:	8082                	ret

00000000000100cc <_start>:
   100cc:	00002197          	auipc	gp,0x2
   100d0:	e5418193          	addi	gp,gp,-428 # 11f20 <__global_pointer$>
   100d4:	00002517          	auipc	a0,0x2
   100d8:	dac50513          	addi	a0,a0,-596 # 11e80 <_edata>
   100dc:	00002617          	auipc	a2,0x2
   100e0:	e5c60613          	addi	a2,a2,-420 # 11f38 <__BSS_END__>
   100e4:	8e09                	sub	a2,a2,a0
   100e6:	4581                	li	a1,0
   100e8:	00000097          	auipc	ra,0x0
   100ec:	3ee080e7          	jalr	1006(ra) # 104d6 <memset>
   100f0:	00000517          	auipc	a0,0x0
   100f4:	34450513          	addi	a0,a0,836 # 10434 <__libc_fini_array>
   100f8:	00000097          	auipc	ra,0x0
   100fc:	300080e7          	jalr	768(ra) # 103f8 <atexit>
   10100:	00000097          	auipc	ra,0x0
   10104:	36c080e7          	jalr	876(ra) # 1046c <__libc_init_array>
   10108:	4502                	lw	a0,0(sp)
   1010a:	002c                	addi	a1,sp,8
   1010c:	4601                	li	a2,0
   1010e:	00000097          	auipc	ra,0x0
   10112:	07a080e7          	jalr	122(ra) # 10188 <main>
   10116:	00000317          	auipc	t1,0x0
   1011a:	2f230067          	jr	754(t1) # 10408 <exit>

000000000001011e <__do_global_dtors_aux>:
   1011e:	00002797          	auipc	a5,0x2
   10122:	d627c783          	lbu	a5,-670(a5) # 11e80 <_edata>
   10126:	ef95                	bnez	a5,10162 <__do_global_dtors_aux+0x44>
   10128:	ffff0797          	auipc	a5,0xffff0
   1012c:	ed878793          	addi	a5,a5,-296 # 0 <register_fini-0x100b0>
   10130:	c39d                	beqz	a5,10156 <__do_global_dtors_aux+0x38>
   10132:	1141                	addi	sp,sp,-16
   10134:	00001517          	auipc	a0,0x1
   10138:	5cc50513          	addi	a0,a0,1484 # 11700 <__FRAME_END__>
   1013c:	e406                	sd	ra,8(sp)
   1013e:	00000097          	auipc	ra,0x0
   10142:	000000e7          	jalr	zero # 0 <register_fini-0x100b0>
   10146:	60a2                	ld	ra,8(sp)
   10148:	4785                	li	a5,1
   1014a:	00002717          	auipc	a4,0x2
   1014e:	d2f70b23          	sb	a5,-714(a4) # 11e80 <_edata>
   10152:	0141                	addi	sp,sp,16
   10154:	8082                	ret
   10156:	4785                	li	a5,1
   10158:	00002717          	auipc	a4,0x2
   1015c:	d2f70423          	sb	a5,-728(a4) # 11e80 <_edata>
   10160:	8082                	ret
   10162:	8082                	ret

0000000000010164 <frame_dummy>:
   10164:	ffff0797          	auipc	a5,0xffff0
   10168:	e9c78793          	addi	a5,a5,-356 # 0 <register_fini-0x100b0>
   1016c:	cf89                	beqz	a5,10186 <frame_dummy+0x22>
   1016e:	00002597          	auipc	a1,0x2
   10172:	d1a58593          	addi	a1,a1,-742 # 11e88 <object.5475>
   10176:	00001517          	auipc	a0,0x1
   1017a:	58a50513          	addi	a0,a0,1418 # 11700 <__FRAME_END__>
   1017e:	00000317          	auipc	t1,0x0
   10182:	00000067          	jr	zero # 0 <register_fini-0x100b0>
   10186:	8082                	ret

0000000000010188 <main>:
#include <stdlib.h>

int N, K, M;
int A[3][3], B[3][3], C[3][3];

int main(){
   10188:	fd010113          	addi	sp,sp,-48
   1018c:	02813423          	sd	s0,40(sp)
   10190:	03010413          	addi	s0,sp,48
	N = 3;
   10194:	000127b7          	lui	a5,0x12
   10198:	00300713          	li	a4,3
   1019c:	eae7ac23          	sw	a4,-328(a5) # 11eb8 <N>
    K = 3;
   101a0:	000127b7          	lui	a5,0x12
   101a4:	00300713          	li	a4,3
   101a8:	eee7a223          	sw	a4,-284(a5) # 11ee4 <K>
    M = 3;
   101ac:	000127b7          	lui	a5,0x12
   101b0:	00300713          	li	a4,3
   101b4:	f2e7aa23          	sw	a4,-204(a5) # 11f34 <M>

	for(int i = 0; i < N; i++)
   101b8:	fe042623          	sw	zero,-20(s0)
   101bc:	0dc0006f          	j	10298 <main+0x110>
		for(int j = 0; j < K; j++)
   101c0:	fe042423          	sw	zero,-24(s0)
   101c4:	0b40006f          	j	10278 <main+0xf0>
			A[i][j] = B[i][j]= i * N + j;
   101c8:	000127b7          	lui	a5,0x12
   101cc:	eb87a783          	lw	a5,-328(a5) # 11eb8 <N>
   101d0:	fec42703          	lw	a4,-20(s0)
   101d4:	02f707bb          	mulw	a5,a4,a5
   101d8:	0007879b          	sext.w	a5,a5
   101dc:	fe842703          	lw	a4,-24(s0)
   101e0:	00f707bb          	addw	a5,a4,a5
   101e4:	0007869b          	sext.w	a3,a5
   101e8:	000127b7          	lui	a5,0x12
   101ec:	ee878613          	addi	a2,a5,-280 # 11ee8 <B>
   101f0:	fe842583          	lw	a1,-24(s0)
   101f4:	fec42703          	lw	a4,-20(s0)
   101f8:	00070793          	mv	a5,a4
   101fc:	00179793          	slli	a5,a5,0x1
   10200:	00e787b3          	add	a5,a5,a4
   10204:	00b787b3          	add	a5,a5,a1
   10208:	00279793          	slli	a5,a5,0x2
   1020c:	00f607b3          	add	a5,a2,a5
   10210:	00d7a023          	sw	a3,0(a5)
   10214:	000127b7          	lui	a5,0x12
   10218:	ee878693          	addi	a3,a5,-280 # 11ee8 <B>
   1021c:	fe842603          	lw	a2,-24(s0)
   10220:	fec42703          	lw	a4,-20(s0)
   10224:	00070793          	mv	a5,a4
   10228:	00179793          	slli	a5,a5,0x1
   1022c:	00e787b3          	add	a5,a5,a4
   10230:	00c787b3          	add	a5,a5,a2
   10234:	00279793          	slli	a5,a5,0x2
   10238:	00f687b3          	add	a5,a3,a5
   1023c:	0007a683          	lw	a3,0(a5)
   10240:	000127b7          	lui	a5,0x12
   10244:	ec078613          	addi	a2,a5,-320 # 11ec0 <A>
   10248:	fe842583          	lw	a1,-24(s0)
   1024c:	fec42703          	lw	a4,-20(s0)
   10250:	00070793          	mv	a5,a4
   10254:	00179793          	slli	a5,a5,0x1
   10258:	00e787b3          	add	a5,a5,a4
   1025c:	00b787b3          	add	a5,a5,a1
   10260:	00279793          	slli	a5,a5,0x2
   10264:	00f607b3          	add	a5,a2,a5
   10268:	00d7a023          	sw	a3,0(a5)
		for(int j = 0; j < K; j++)
   1026c:	fe842783          	lw	a5,-24(s0)
   10270:	0017879b          	addiw	a5,a5,1
   10274:	fef42423          	sw	a5,-24(s0)
   10278:	000127b7          	lui	a5,0x12
   1027c:	ee47a703          	lw	a4,-284(a5) # 11ee4 <K>
   10280:	fe842783          	lw	a5,-24(s0)
   10284:	0007879b          	sext.w	a5,a5
   10288:	f4e7c0e3          	blt	a5,a4,101c8 <main+0x40>
	for(int i = 0; i < N; i++)
   1028c:	fec42783          	lw	a5,-20(s0)
   10290:	0017879b          	addiw	a5,a5,1
   10294:	fef42623          	sw	a5,-20(s0)
   10298:	000127b7          	lui	a5,0x12
   1029c:	eb87a703          	lw	a4,-328(a5) # 11eb8 <N>
   102a0:	fec42783          	lw	a5,-20(s0)
   102a4:	0007879b          	sext.w	a5,a5
   102a8:	f0e7cce3          	blt	a5,a4,101c0 <main+0x38>

	for(int i = 0; i < N; i++)
   102ac:	fe042223          	sw	zero,-28(s0)
   102b0:	1200006f          	j	103d0 <main+0x248>
        for(int k = 0; k < K; k++)
   102b4:	fe042023          	sw	zero,-32(s0)
   102b8:	0f80006f          	j	103b0 <main+0x228>
            for(int j = 0; j < M; j++)
   102bc:	fc042e23          	sw	zero,-36(s0)
   102c0:	0d00006f          	j	10390 <main+0x208>
                C[i][j] += A[i][k] * B[k][j];
   102c4:	000127b7          	lui	a5,0x12
   102c8:	f1078693          	addi	a3,a5,-240 # 11f10 <C>
   102cc:	fdc42603          	lw	a2,-36(s0)
   102d0:	fe442703          	lw	a4,-28(s0)
   102d4:	00070793          	mv	a5,a4
   102d8:	00179793          	slli	a5,a5,0x1
   102dc:	00e787b3          	add	a5,a5,a4
   102e0:	00c787b3          	add	a5,a5,a2
   102e4:	00279793          	slli	a5,a5,0x2
   102e8:	00f687b3          	add	a5,a3,a5
   102ec:	0007a683          	lw	a3,0(a5)
   102f0:	000127b7          	lui	a5,0x12
   102f4:	ec078613          	addi	a2,a5,-320 # 11ec0 <A>
   102f8:	fe042583          	lw	a1,-32(s0)
   102fc:	fe442703          	lw	a4,-28(s0)
   10300:	00070793          	mv	a5,a4
   10304:	00179793          	slli	a5,a5,0x1
   10308:	00e787b3          	add	a5,a5,a4
   1030c:	00b787b3          	add	a5,a5,a1
   10310:	00279793          	slli	a5,a5,0x2
   10314:	00f607b3          	add	a5,a2,a5
   10318:	0007a603          	lw	a2,0(a5)
   1031c:	000127b7          	lui	a5,0x12
   10320:	ee878593          	addi	a1,a5,-280 # 11ee8 <B>
   10324:	fdc42503          	lw	a0,-36(s0)
   10328:	fe042703          	lw	a4,-32(s0)
   1032c:	00070793          	mv	a5,a4
   10330:	00179793          	slli	a5,a5,0x1
   10334:	00e787b3          	add	a5,a5,a4
   10338:	00a787b3          	add	a5,a5,a0
   1033c:	00279793          	slli	a5,a5,0x2
   10340:	00f587b3          	add	a5,a1,a5
   10344:	0007a783          	lw	a5,0(a5)
   10348:	02f607bb          	mulw	a5,a2,a5
   1034c:	0007879b          	sext.w	a5,a5
   10350:	00f687bb          	addw	a5,a3,a5
   10354:	0007869b          	sext.w	a3,a5
   10358:	000127b7          	lui	a5,0x12
   1035c:	f1078613          	addi	a2,a5,-240 # 11f10 <C>
   10360:	fdc42583          	lw	a1,-36(s0)
   10364:	fe442703          	lw	a4,-28(s0)
   10368:	00070793          	mv	a5,a4
   1036c:	00179793          	slli	a5,a5,0x1
   10370:	00e787b3          	add	a5,a5,a4
   10374:	00b787b3          	add	a5,a5,a1
   10378:	00279793          	slli	a5,a5,0x2
   1037c:	00f607b3          	add	a5,a2,a5
   10380:	00d7a023          	sw	a3,0(a5)
            for(int j = 0; j < M; j++)
   10384:	fdc42783          	lw	a5,-36(s0)
   10388:	0017879b          	addiw	a5,a5,1
   1038c:	fcf42e23          	sw	a5,-36(s0)
   10390:	000127b7          	lui	a5,0x12
   10394:	f347a703          	lw	a4,-204(a5) # 11f34 <M>
   10398:	fdc42783          	lw	a5,-36(s0)
   1039c:	0007879b          	sext.w	a5,a5
   103a0:	f2e7c2e3          	blt	a5,a4,102c4 <main+0x13c>
        for(int k = 0; k < K; k++)
   103a4:	fe042783          	lw	a5,-32(s0)
   103a8:	0017879b          	addiw	a5,a5,1
   103ac:	fef42023          	sw	a5,-32(s0)
   103b0:	000127b7          	lui	a5,0x12
   103b4:	ee47a703          	lw	a4,-284(a5) # 11ee4 <K>
   103b8:	fe042783          	lw	a5,-32(s0)
   103bc:	0007879b          	sext.w	a5,a5
   103c0:	eee7cee3          	blt	a5,a4,102bc <main+0x134>
	for(int i = 0; i < N; i++)
   103c4:	fe442783          	lw	a5,-28(s0)
   103c8:	0017879b          	addiw	a5,a5,1
   103cc:	fef42223          	sw	a5,-28(s0)
   103d0:	000127b7          	lui	a5,0x12
   103d4:	eb87a703          	lw	a4,-328(a5) # 11eb8 <N>
   103d8:	fe442783          	lw	a5,-28(s0)
   103dc:	0007879b          	sext.w	a5,a5
   103e0:	ece7cae3          	blt	a5,a4,102b4 <main+0x12c>


    return 0;
   103e4:	00000793          	li	a5,0
}
   103e8:	00078513          	mv	a0,a5
   103ec:	02813403          	ld	s0,40(sp)
   103f0:	03010113          	addi	sp,sp,48
   103f4:	00008067          	ret

00000000000103f8 <atexit>:
   103f8:	85aa                	mv	a1,a0
   103fa:	4681                	li	a3,0
   103fc:	4601                	li	a2,0
   103fe:	4501                	li	a0,0
   10400:	00000317          	auipc	t1,0x0
   10404:	18030067          	jr	384(t1) # 10580 <__register_exitproc>

0000000000010408 <exit>:
   10408:	1141                	addi	sp,sp,-16
   1040a:	4581                	li	a1,0
   1040c:	e022                	sd	s0,0(sp)
   1040e:	e406                	sd	ra,8(sp)
   10410:	842a                	mv	s0,a0
   10412:	00000097          	auipc	ra,0x0
   10416:	1ea080e7          	jalr	490(ra) # 105fc <__call_exitprocs>
   1041a:	00002797          	auipc	a5,0x2
   1041e:	a4e78793          	addi	a5,a5,-1458 # 11e68 <_global_impure_ptr>
   10422:	6388                	ld	a0,0(a5)
   10424:	6d3c                	ld	a5,88(a0)
   10426:	c391                	beqz	a5,1042a <exit+0x22>
   10428:	9782                	jalr	a5
   1042a:	8522                	mv	a0,s0
   1042c:	00000097          	auipc	ra,0x0
   10430:	296080e7          	jalr	662(ra) # 106c2 <_exit>

0000000000010434 <__libc_fini_array>:
   10434:	1101                	addi	sp,sp,-32
   10436:	e822                	sd	s0,16(sp)
   10438:	00001797          	auipc	a5,0x1
   1043c:	2e878793          	addi	a5,a5,744 # 11720 <__fini_array_end>
   10440:	00001417          	auipc	s0,0x1
   10444:	2d840413          	addi	s0,s0,728 # 11718 <__init_array_end>
   10448:	8f81                	sub	a5,a5,s0
   1044a:	e426                	sd	s1,8(sp)
   1044c:	ec06                	sd	ra,24(sp)
   1044e:	4037d493          	srai	s1,a5,0x3
   10452:	c881                	beqz	s1,10462 <__libc_fini_array+0x2e>
   10454:	17e1                	addi	a5,a5,-8
   10456:	943e                	add	s0,s0,a5
   10458:	601c                	ld	a5,0(s0)
   1045a:	14fd                	addi	s1,s1,-1
   1045c:	1461                	addi	s0,s0,-8
   1045e:	9782                	jalr	a5
   10460:	fce5                	bnez	s1,10458 <__libc_fini_array+0x24>
   10462:	60e2                	ld	ra,24(sp)
   10464:	6442                	ld	s0,16(sp)
   10466:	64a2                	ld	s1,8(sp)
   10468:	6105                	addi	sp,sp,32
   1046a:	8082                	ret

000000000001046c <__libc_init_array>:
   1046c:	1101                	addi	sp,sp,-32
   1046e:	e822                	sd	s0,16(sp)
   10470:	e04a                	sd	s2,0(sp)
   10472:	00001417          	auipc	s0,0x1
   10476:	29240413          	addi	s0,s0,658 # 11704 <__preinit_array_end>
   1047a:	00001917          	auipc	s2,0x1
   1047e:	28a90913          	addi	s2,s2,650 # 11704 <__preinit_array_end>
   10482:	40890933          	sub	s2,s2,s0
   10486:	ec06                	sd	ra,24(sp)
   10488:	e426                	sd	s1,8(sp)
   1048a:	40395913          	srai	s2,s2,0x3
   1048e:	00090963          	beqz	s2,104a0 <__libc_init_array+0x34>
   10492:	4481                	li	s1,0
   10494:	601c                	ld	a5,0(s0)
   10496:	0485                	addi	s1,s1,1
   10498:	0421                	addi	s0,s0,8
   1049a:	9782                	jalr	a5
   1049c:	fe991ce3          	bne	s2,s1,10494 <__libc_init_array+0x28>
   104a0:	00001417          	auipc	s0,0x1
   104a4:	26840413          	addi	s0,s0,616 # 11708 <__init_array_start>
   104a8:	00001917          	auipc	s2,0x1
   104ac:	27090913          	addi	s2,s2,624 # 11718 <__init_array_end>
   104b0:	40890933          	sub	s2,s2,s0
   104b4:	40395913          	srai	s2,s2,0x3
   104b8:	00090963          	beqz	s2,104ca <__libc_init_array+0x5e>
   104bc:	4481                	li	s1,0
   104be:	601c                	ld	a5,0(s0)
   104c0:	0485                	addi	s1,s1,1
   104c2:	0421                	addi	s0,s0,8
   104c4:	9782                	jalr	a5
   104c6:	fe991ce3          	bne	s2,s1,104be <__libc_init_array+0x52>
   104ca:	60e2                	ld	ra,24(sp)
   104cc:	6442                	ld	s0,16(sp)
   104ce:	64a2                	ld	s1,8(sp)
   104d0:	6902                	ld	s2,0(sp)
   104d2:	6105                	addi	sp,sp,32
   104d4:	8082                	ret

00000000000104d6 <memset>:
   104d6:	433d                	li	t1,15
   104d8:	872a                	mv	a4,a0
   104da:	02c37163          	bgeu	t1,a2,104fc <memset+0x26>
   104de:	00f77793          	andi	a5,a4,15
   104e2:	e3c1                	bnez	a5,10562 <memset+0x8c>
   104e4:	e1bd                	bnez	a1,1054a <memset+0x74>
   104e6:	ff067693          	andi	a3,a2,-16
   104ea:	8a3d                	andi	a2,a2,15
   104ec:	96ba                	add	a3,a3,a4
   104ee:	e30c                	sd	a1,0(a4)
   104f0:	e70c                	sd	a1,8(a4)
   104f2:	0741                	addi	a4,a4,16
   104f4:	fed76de3          	bltu	a4,a3,104ee <memset+0x18>
   104f8:	e211                	bnez	a2,104fc <memset+0x26>
   104fa:	8082                	ret
   104fc:	40c306b3          	sub	a3,t1,a2
   10500:	068a                	slli	a3,a3,0x2
   10502:	00000297          	auipc	t0,0x0
   10506:	9696                	add	a3,a3,t0
   10508:	00a68067          	jr	10(a3)
   1050c:	00b70723          	sb	a1,14(a4)
   10510:	00b706a3          	sb	a1,13(a4)
   10514:	00b70623          	sb	a1,12(a4)
   10518:	00b705a3          	sb	a1,11(a4)
   1051c:	00b70523          	sb	a1,10(a4)
   10520:	00b704a3          	sb	a1,9(a4)
   10524:	00b70423          	sb	a1,8(a4)
   10528:	00b703a3          	sb	a1,7(a4)
   1052c:	00b70323          	sb	a1,6(a4)
   10530:	00b702a3          	sb	a1,5(a4)
   10534:	00b70223          	sb	a1,4(a4)
   10538:	00b701a3          	sb	a1,3(a4)
   1053c:	00b70123          	sb	a1,2(a4)
   10540:	00b700a3          	sb	a1,1(a4)
   10544:	00b70023          	sb	a1,0(a4)
   10548:	8082                	ret
   1054a:	0ff5f593          	andi	a1,a1,255
   1054e:	00859693          	slli	a3,a1,0x8
   10552:	8dd5                	or	a1,a1,a3
   10554:	01059693          	slli	a3,a1,0x10
   10558:	8dd5                	or	a1,a1,a3
   1055a:	02059693          	slli	a3,a1,0x20
   1055e:	8dd5                	or	a1,a1,a3
   10560:	b759                	j	104e6 <memset+0x10>
   10562:	00279693          	slli	a3,a5,0x2
   10566:	00000297          	auipc	t0,0x0
   1056a:	9696                	add	a3,a3,t0
   1056c:	8286                	mv	t0,ra
   1056e:	fa2680e7          	jalr	-94(a3)
   10572:	8096                	mv	ra,t0
   10574:	17c1                	addi	a5,a5,-16
   10576:	8f1d                	sub	a4,a4,a5
   10578:	963e                	add	a2,a2,a5
   1057a:	f8c371e3          	bgeu	t1,a2,104fc <memset+0x26>
   1057e:	b79d                	j	104e4 <memset+0xe>

0000000000010580 <__register_exitproc>:
   10580:	00002797          	auipc	a5,0x2
   10584:	8e878793          	addi	a5,a5,-1816 # 11e68 <_global_impure_ptr>
   10588:	6398                	ld	a4,0(a5)
   1058a:	1f873783          	ld	a5,504(a4)
   1058e:	c3b1                	beqz	a5,105d2 <__register_exitproc+0x52>
   10590:	4798                	lw	a4,8(a5)
   10592:	487d                	li	a6,31
   10594:	06e84263          	blt	a6,a4,105f8 <__register_exitproc+0x78>
   10598:	c505                	beqz	a0,105c0 <__register_exitproc+0x40>
   1059a:	00371813          	slli	a6,a4,0x3
   1059e:	983e                	add	a6,a6,a5
   105a0:	10c83823          	sd	a2,272(a6)
   105a4:	3107a883          	lw	a7,784(a5)
   105a8:	4605                	li	a2,1
   105aa:	00e6163b          	sllw	a2,a2,a4
   105ae:	00c8e8b3          	or	a7,a7,a2
   105b2:	3117a823          	sw	a7,784(a5)
   105b6:	20d83823          	sd	a3,528(a6)
   105ba:	4689                	li	a3,2
   105bc:	02d50063          	beq	a0,a3,105dc <__register_exitproc+0x5c>
   105c0:	00270693          	addi	a3,a4,2
   105c4:	068e                	slli	a3,a3,0x3
   105c6:	2705                	addiw	a4,a4,1
   105c8:	c798                	sw	a4,8(a5)
   105ca:	97b6                	add	a5,a5,a3
   105cc:	e38c                	sd	a1,0(a5)
   105ce:	4501                	li	a0,0
   105d0:	8082                	ret
   105d2:	20070793          	addi	a5,a4,512
   105d6:	1ef73c23          	sd	a5,504(a4)
   105da:	bf5d                	j	10590 <__register_exitproc+0x10>
   105dc:	3147a683          	lw	a3,788(a5)
   105e0:	4501                	li	a0,0
   105e2:	8e55                	or	a2,a2,a3
   105e4:	00270693          	addi	a3,a4,2
   105e8:	068e                	slli	a3,a3,0x3
   105ea:	2705                	addiw	a4,a4,1
   105ec:	30c7aa23          	sw	a2,788(a5)
   105f0:	c798                	sw	a4,8(a5)
   105f2:	97b6                	add	a5,a5,a3
   105f4:	e38c                	sd	a1,0(a5)
   105f6:	8082                	ret
   105f8:	557d                	li	a0,-1
   105fa:	8082                	ret

00000000000105fc <__call_exitprocs>:
   105fc:	715d                	addi	sp,sp,-80
   105fe:	00002797          	auipc	a5,0x2
   10602:	86a78793          	addi	a5,a5,-1942 # 11e68 <_global_impure_ptr>
   10606:	e062                	sd	s8,0(sp)
   10608:	0007bc03          	ld	s8,0(a5)
   1060c:	f44e                	sd	s3,40(sp)
   1060e:	f052                	sd	s4,32(sp)
   10610:	ec56                	sd	s5,24(sp)
   10612:	e85a                	sd	s6,16(sp)
   10614:	e486                	sd	ra,72(sp)
   10616:	e0a2                	sd	s0,64(sp)
   10618:	fc26                	sd	s1,56(sp)
   1061a:	f84a                	sd	s2,48(sp)
   1061c:	e45e                	sd	s7,8(sp)
   1061e:	8aaa                	mv	s5,a0
   10620:	8b2e                	mv	s6,a1
   10622:	4a05                	li	s4,1
   10624:	59fd                	li	s3,-1
   10626:	1f8c3903          	ld	s2,504(s8)
   1062a:	02090463          	beqz	s2,10652 <__call_exitprocs+0x56>
   1062e:	00892483          	lw	s1,8(s2)
   10632:	fff4841b          	addiw	s0,s1,-1
   10636:	00044e63          	bltz	s0,10652 <__call_exitprocs+0x56>
   1063a:	048e                	slli	s1,s1,0x3
   1063c:	94ca                	add	s1,s1,s2
   1063e:	020b0663          	beqz	s6,1066a <__call_exitprocs+0x6e>
   10642:	2084b783          	ld	a5,520(s1)
   10646:	03678263          	beq	a5,s6,1066a <__call_exitprocs+0x6e>
   1064a:	347d                	addiw	s0,s0,-1
   1064c:	14e1                	addi	s1,s1,-8
   1064e:	ff3418e3          	bne	s0,s3,1063e <__call_exitprocs+0x42>
   10652:	60a6                	ld	ra,72(sp)
   10654:	6406                	ld	s0,64(sp)
   10656:	74e2                	ld	s1,56(sp)
   10658:	7942                	ld	s2,48(sp)
   1065a:	79a2                	ld	s3,40(sp)
   1065c:	7a02                	ld	s4,32(sp)
   1065e:	6ae2                	ld	s5,24(sp)
   10660:	6b42                	ld	s6,16(sp)
   10662:	6ba2                	ld	s7,8(sp)
   10664:	6c02                	ld	s8,0(sp)
   10666:	6161                	addi	sp,sp,80
   10668:	8082                	ret
   1066a:	00892783          	lw	a5,8(s2)
   1066e:	6498                	ld	a4,8(s1)
   10670:	37fd                	addiw	a5,a5,-1
   10672:	04878263          	beq	a5,s0,106b6 <__call_exitprocs+0xba>
   10676:	0004b423          	sd	zero,8(s1)
   1067a:	db61                	beqz	a4,1064a <__call_exitprocs+0x4e>
   1067c:	31092783          	lw	a5,784(s2)
   10680:	008a16bb          	sllw	a3,s4,s0
   10684:	00892b83          	lw	s7,8(s2)
   10688:	8ff5                	and	a5,a5,a3
   1068a:	2781                	sext.w	a5,a5
   1068c:	eb99                	bnez	a5,106a2 <__call_exitprocs+0xa6>
   1068e:	9702                	jalr	a4
   10690:	00892783          	lw	a5,8(s2)
   10694:	f97799e3          	bne	a5,s7,10626 <__call_exitprocs+0x2a>
   10698:	1f8c3783          	ld	a5,504(s8)
   1069c:	fb2787e3          	beq	a5,s2,1064a <__call_exitprocs+0x4e>
   106a0:	b759                	j	10626 <__call_exitprocs+0x2a>
   106a2:	31492783          	lw	a5,788(s2)
   106a6:	1084b583          	ld	a1,264(s1)
   106aa:	8ff5                	and	a5,a5,a3
   106ac:	2781                	sext.w	a5,a5
   106ae:	e799                	bnez	a5,106bc <__call_exitprocs+0xc0>
   106b0:	8556                	mv	a0,s5
   106b2:	9702                	jalr	a4
   106b4:	bff1                	j	10690 <__call_exitprocs+0x94>
   106b6:	00892423          	sw	s0,8(s2)
   106ba:	b7c1                	j	1067a <__call_exitprocs+0x7e>
   106bc:	852e                	mv	a0,a1
   106be:	9702                	jalr	a4
   106c0:	bfc1                	j	10690 <__call_exitprocs+0x94>

00000000000106c2 <_exit>:
   106c2:	4581                	li	a1,0
   106c4:	4601                	li	a2,0
   106c6:	4681                	li	a3,0
   106c8:	4701                	li	a4,0
   106ca:	4781                	li	a5,0
   106cc:	05d00893          	li	a7,93
   106d0:	00000073          	ecall
   106d4:	00054363          	bltz	a0,106da <_exit+0x18>
   106d8:	a001                	j	106d8 <_exit+0x16>
   106da:	1141                	addi	sp,sp,-16
   106dc:	e022                	sd	s0,0(sp)
   106de:	842a                	mv	s0,a0
   106e0:	e406                	sd	ra,8(sp)
   106e2:	4080043b          	negw	s0,s0
   106e6:	00000097          	auipc	ra,0x0
   106ea:	00c080e7          	jalr	12(ra) # 106f2 <__errno>
   106ee:	c100                	sw	s0,0(a0)
   106f0:	a001                	j	106f0 <_exit+0x2e>

00000000000106f2 <__errno>:
   106f2:	00001797          	auipc	a5,0x1
   106f6:	78678793          	addi	a5,a5,1926 # 11e78 <_impure_ptr>
   106fa:	6388                	ld	a0,0(a5)
   106fc:	8082                	ret
